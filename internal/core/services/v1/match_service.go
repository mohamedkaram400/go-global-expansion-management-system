package services

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/core/entities/v1"
	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/ports/v1"
	"github.com/mohamedkaram400/go-global-expansion-management-system/responses/v1"
)

type MatchService struct {
	MatchRepo ports.MatchRepository
	ProjectService *ProjectService
	Notifier ports.Notifier
	ClientService *ClientService
}

func NewMatchService(repo ports.MatchRepository, projectService *ProjectService, notifier ports.Notifier, clientService *ClientService) *MatchService {
	return &MatchService{MatchRepo: repo, ProjectService: projectService, Notifier: notifier, ClientService: clientService}
}

func (svc *MatchService) Rebuild(ctx context.Context, projectID int) ([]responses.MatchResponse, error) {

	// 1. Load project from DB
	project, err := svc.ProjectService.FindProjectByID(ctx, projectID)
	if err != nil {
		return nil, err
	}
 
	// fmt.Println(project)

	// 2. Parse JSON []byte → []string
	var projectServices []string
	if err := json.Unmarshal(project.ServicesNeeded, &projectServices); err != nil {
		return nil, fmt.Errorf("failed to parse services_needed: %w", err)
	}


	// 3. Get all vendors that apply rules
	vendors, err := svc.MatchRepo.GetVendorsForProject(ctx, project)
	if err != nil {
		return nil, err
	}
 
	// fmt.Println(project, vendors)

	// 4. Init object for response with length of vendors matches
	matchResponses := make([]responses.MatchResponse, 0, len(vendors))

	// 5. Loop over vendors Count number of service overlap in project and offered by vendor
	for _, v := range vendors {

		// 6. Parse vendor services []byte → []string
		var vendorServices []string
		if err := json.Unmarshal(v.ServicesOffered, &vendorServices); err != nil {
			return nil, fmt.Errorf("failed to parse vendor services_offered: %w", err)
		}

		// 7. Count how many service overlap with project and vendor
		overlap := countOverlap(projectServices, vendorServices)

		// fmt.Println("overlap:", overlap, "projectServices: ", projectServices, "vendorServices: ", vendorServices)

		if overlap == 0 {
			continue
		}

		// 8. Calculate Score formula: services_overlap * 2 + rating + SLA_weight
		score := float64(overlap*2) + v.Rating + float64(v.ResponseSlaHours)
		match := entities.Match{
			ProjectID: project.ID,
			VendorID:  v.ID,
			Score:     score,
		}

		// 9. Save match result with socre in matches table 
		if err := svc.MatchRepo.UpsertMatch(ctx, &match); err != nil {
			return nil, err
		}


		// 10. Build clean response
		resp := responses.MatchResponse{
			ID:         match.ID,
			ProjectID:  match.ProjectID,
			VendorID:   match.VendorID,
			VendorName: v.Name,
			Score:      match.Score,
			CreatedAt:  match.CreatedAt,
			UpdatedAt:  match.UpdatedAt,
		}

		// 11. Append all match result to matchs
		matchResponses = append(matchResponses, resp)

		// 12. Get client related to current project
		client, err := svc.ClientService.FindClientByID(ctx, project.ClientID)
		if err != nil {
			return nil, err
		}

		// 13. Send email notification
		go func() {
			notif := ports.MatchNotification{
				MatchID:   match.ID,
				ProjectID: match.ProjectID,
				VendorID:  match.VendorID,
				Score:     match.Score,
				To:        []string{client.ContactEmail},
				Subject:   fmt.Sprintf("New Match for Project %d", match.ProjectID),
				Body:      fmt.Sprintf("Vendor %s matched with score %.1f", v.Name, match.Score),
			}

			fmt.Println("notif:", notif)
			
			if err := svc.Notifier.SendMatchNotification(context.Background(), notif); err != nil {
				log.Printf("email send error: %v", err)
			} else {
				log.Println(">>> email sent OK")
			}
		    // Add delay
	        time.Sleep(1 * time.Second)

		}()
	}

	fmt.Println("Matchs: ", matchResponses)

	return matchResponses, nil
}

func countOverlap(projectServices, vendorServices []string) int {
	set := make(map[string]struct{})

	for _, vs := range vendorServices {
		set[vs] = struct{}{}
	}

	count := 0
	for _, ps := range projectServices {
		if _, ok := set[ps]; ok {
			count++
		}
	}
	return count
}


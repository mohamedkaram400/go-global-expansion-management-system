package repositories

import (
	"context"
	"errors"

	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/core/entities/v1"
	"gorm.io/gorm"
)

type ProjectRepo struct {
	DB *gorm.DB
}

func NewProjectRepo(db *gorm.DB) *ProjectRepo {
	return &ProjectRepo{DB: db}
}

func (r *ProjectRepo) GetAllProjects(ctx context.Context, skip int, limit int) ([]*entities.Project, error) {
	var projects []*entities.Project
	if err := r.DB.WithContext(ctx).
		Offset(skip).
		Limit(limit).
		Find(&projects).Error; err != nil {
		return nil, err
	}
	return projects, nil
}

func (r *ProjectRepo) FindProjectByID(ctx context.Context, projectID int) (*entities.Project, error) {
	var project entities.Project
	if err := r.DB.WithContext(ctx).
		Where("id = ?", projectID).
		First(&project).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		return nil, err
	}

	return &project, nil
}

func (r *ProjectRepo) InsertProject(ctx context.Context, project *entities.Project) (*entities.Project, error) {
	if err := r.DB.WithContext(ctx).Create(project).Error; err != nil {
		return nil, err
	}
	return project, nil
}

func (r *ProjectRepo) UpdateProjectByID(ctx context.Context, projectID int, updates map[string]interface{}) (*entities.Project, error) {

	project := &entities.Project{}

	// Update the Project
	if err := r.DB.WithContext(ctx).
		Model(project).
		Where("id = ?", projectID).
		Updates(updates).Error; err != nil {
		return nil, err
	}

	// Fetch the updated record
	if err := r.DB.WithContext(ctx).
		Where("id = ?", projectID).
		First(project).Error; err != nil {
		return nil, err
	}

	return project, nil
}

func (r *ProjectRepo) DeleteProjectByID(ctx context.Context, projectID int) (int, error) {
	if err := r.DB.WithContext(ctx).
		Where("id = ?", projectID).
		Delete(&entities.Project{}).Error; err != nil {
		return 0, err
	}

	return 1, nil
}

func (r *ProjectRepo) GetProjectCountries(ctx context.Context) (map[int]string, error) {
	rows, err := r.DB.WithContext(ctx).Raw("SELECT id, country FROM projects").Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make(map[int]string)
	for rows.Next() {
		var id int
		var country string
		if err := rows.Scan(&id, &country); err != nil {
			return nil, err
		}
		result[id] = country
	}

	return result, nil
}

func (s *ProjectRepo) GetAllActiveProjects(ctx context.Context) ([]*entities.Project, error) {
	var projects []*entities.Project
	if err := s.DB.Where("status = ?", "active").Find(&projects).Error; err != nil {
		return nil, err
	}
	return projects, nil
}

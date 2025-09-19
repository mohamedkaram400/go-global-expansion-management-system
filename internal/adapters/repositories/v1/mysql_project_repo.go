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

func (r *ProjectRepo) GetAllProjects(ctx context.Context, skip int, limit int) ([]entities.Project, error) {
	var projects []entities.Project
	if err := r.DB.WithContext(ctx).
		Offset(skip).
		Limit(limit).
		Find(&projects).Error; err != nil {
		return nil, err
	}
	return projects, nil
}

func (r *ProjectRepo) FindProjectByID(ctx context.Context, projectID uint) (*entities.Project, error) {
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

func (r *ProjectRepo) UpdateProjectByID(ctx context.Context, projectID uint, updates map[string]interface{}) (*entities.Project, error) {

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

func (r *ProjectRepo) DeleteProjectByID(ctx context.Context, projectID uint) (int, error) {
	if err := r.DB.WithContext(ctx).
		Where("id = ?", projectID).
		Delete(&entities.Project{}).Error; err != nil {
		return 0, err
	}

	return 1, nil
}

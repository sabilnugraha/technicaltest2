package repositories

import (
	"technicaltest/models"

	"gorm.io/gorm"
)

type DataRepository interface {
	CreateData(data models.Data) (models.Data, error)
}

func RepositoryData(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateData(data models.Data) (models.Data, error) {
	err := r.db.Create(&data).Error

	return data, err
}

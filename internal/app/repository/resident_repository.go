package repository

import (
	"Filo.Hack/internal/app/model"
	"gorm.io/gorm"
)

type ResidentRepository struct {
	DB *gorm.DB
}

func NewResidentRepository(db *gorm.DB) *ResidentRepository {
	return &ResidentRepository{
		DB: db,
	}
}

func (r ResidentRepository) CreateResident(user *model.Resident) error {
	return r.DB.Create(user).Error
}

func (r ResidentRepository) GetResidentByEmail(email string) (*model.Resident, error) {
	var user model.Resident
	if err := r.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

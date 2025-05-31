package repository

import (
	"errors"
	"fmt"

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

func (s *ResidentRepository) SetInterestsByUser(interests []string, email string) (*model.Resident, error) {
	// Находим пользователя по email
	var resident model.Resident
	if err := s.DB.Where("email = ?", email).Preload("Interests").First(&resident).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("user with email %s not found", email)
		}
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	// Получаем существующие интересы пользователя
	currentInterests := make(map[string]bool)
	for _, interest := range resident.Interests {
		currentInterests[interest.Name] = true
	}

	// Создаем или получаем интересы и формируем новый список
	var newInterests []model.Interest
	for _, interestName := range interests {
		// Пропускаем уже существующие интересы
		if currentInterests[interestName] {
			continue
		}

		// Получаем или создаем интерес
		interest, err := s.GetOrCreateInterest(interestName)
		if err != nil {
			return nil, fmt.Errorf("failed to get or create interest %s: %w", interestName, err)
		}
		newInterests = append(newInterests, *interest)
	}

	// Удаляем старые связи
	if err := s.DB.Model(&resident).Association("Interests").Clear(); err != nil {
		return nil, fmt.Errorf("failed to clear existing interests: %w", err)
	}

	// Добавляем новые связи
	if len(newInterests) > 0 {
		if err := s.DB.Model(&resident).Association("Interests").Append(newInterests); err != nil {
			return nil, fmt.Errorf("failed to set new interests: %w", err)
		}
	}

	// Перезагружаем данные пользователя с интересами
	if err := s.DB.Preload("Interests").First(&resident, resident.ResidentID).Error; err != nil {
		return nil, fmt.Errorf("failed to reload resident: %w", err)
	}

	return &resident, nil
}

func (r *ResidentRepository) GetOrCreateInterest(name string) (*model.Interest, error) {
	var interest model.Interest
	result := r.DB.Where(model.Interest{Name: name}).FirstOrCreate(&interest)

	if result.Error != nil {
		return nil, result.Error
	}

	return &interest, nil
}

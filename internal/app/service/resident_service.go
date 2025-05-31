package service

import (
	"Filo.Hack/internal/app/model"
	"Filo.Hack/internal/app/repository"
)

type ResidentService struct {
	residentRepo *repository.ResidentRepository
}

func (s *ResidentService) SetInterestsByUser(interests []string, email string) (*model.Resident, error) {
	return s.residentRepo.SetInterestsByUser(interests, email)
}

func (s *ResidentService) GetMe(email string) (*model.Resident, error) {
	return s.residentRepo.GetResidentByEmail(email)
}

func NewResidentService(residentRepo *repository.ResidentRepository) *ResidentService {
	return &ResidentService{residentRepo: residentRepo}
}

func (s *ResidentService) CreateResident(resident *model.Resident) error {
	return s.residentRepo.CreateResident(resident)
}

func (s *ResidentService) GetResidentByEmail(email string) (*model.Resident, error) {
	return s.residentRepo.GetResidentByEmail(email)
}

package service

import (
	"errors"
	"time"

	"Filo.Hack/internal/app/model"
	"Filo.Hack/internal/app/repository"
)

type EventService struct {
	eventRepo    *repository.EventRepository
	residentRepo *repository.ResidentRepository
}

func NewEventService(repo *repository.EventRepository, residentRepo *repository.ResidentRepository) *EventService {
	return &EventService{eventRepo: repo, residentRepo: residentRepo}
}

func (s *EventService) CreateEvent(event *model.Event) error {
	if event.EventDate.Before(time.Now()) {
		return errors.New("Date must be in future")
	}

	if _, err := s.residentRepo.GetResidentByEmail(event.Organizer.Email); err != nil {
		return errors.New("User not found")
	} else {
		eventOrganizer, err := s.residentRepo.GetResidentByEmail(event.Organizer.Email)
		if err != nil {
			return err
		}
		event.Organizer = *eventOrganizer
	}

	// Если такого интереса не было добавляем
	s.eventRepo.GetOrCreateInterest(event.EventTypeName)

	return s.eventRepo.Add(event)
}

func (s *EventService) GetEventsByCategory(category string) []model.Event {
	return s.eventRepo.GetAllByCategory(category)
}

func (s *EventService) GetEventsByAddress(address string) []model.Event {
	return s.eventRepo.GetAllByAddress(address)
}

func (s *EventService) GetAllEvents() []model.Event {
	return s.eventRepo.GetAll()
}

func (s *EventService) GetAllInterests() (*[]model.Interest, error) {
	return s.eventRepo.GetAllInterests()
}

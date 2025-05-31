package service

import "Filo.Hack/internal/app/repository"

type EventService struct {
	eventRepo repository.EventRepository
}

func NewEventService(repo repository.EventRepository) *EventService {
	return &EventService{eventRepo: repo}
}

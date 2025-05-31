package endpoint

import "Filo.Hack/internal/app/service"

type eventEndpoint struct {
	eventService *service.EventService
}

func NewEventEndpoint(eventService *service.EventService) *eventEndpoint {
	return &eventEndpoint{eventService: eventService}
}

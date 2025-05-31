package endpoint

import (
	"Filo.Hack/internal/app/model"
	"Filo.Hack/internal/app/service"
	"github.com/labstack/echo/v4"
)

type eventEndpoint struct {
	eventService *service.EventService
}

func NewEventEndpoint(eventService *service.EventService) *eventEndpoint {
	return &eventEndpoint{eventService: eventService}
}

func (e *eventEndpoint) CreateEvent(ctx echo.Context) error {
	var req model.Event

	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(400, map[string]interface{}{
			"status":  "failed",
			"message": err.Error(),
		})
	}

	if err := e.eventService.CreateEvent(&req); err != nil {
		return ctx.JSON(400, map[string]interface{}{
			"status":  "failed",
			"message": err.Error(),
		})
	}

	return ctx.JSON(200, req)
}

func (eventEndpoint) GetAllEvent(ctx echo.Context) error {
	return nil
}

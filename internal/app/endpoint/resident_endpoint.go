package endpoint

import (
	"Filo.Hack/internal/app/middleware"
	"Filo.Hack/internal/app/service"
	"github.com/labstack/echo/v4"
)

type ResidentEndpoint struct {
	residentService *service.ResidentService
}

func NewResidentEndpoint(residentService *service.ResidentService) *ResidentEndpoint {
	return &ResidentEndpoint{
		residentService: residentService,
	}
}

func (r *ResidentEndpoint) GetMe(ctx echo.Context) error {
	userClaims, ok := ctx.Get("user").(*middleware.CustomClaims)
	if !ok {
		return ctx.JSON(400, "invalid token")
	}

	resident, err := r.residentService.GetResidentByEmail(userClaims.Email)
	if err != nil {
		return ctx.JSON(400, map[string]interface{}{
			"status":  "failed",
			"message": err.Error(),
		})
	}

	return ctx.JSON(200, resident)
}

func (r *ResidentEndpoint) SetInterestsByUser(ctx echo.Context) error {
	userClaims, ok := ctx.Get("user").(*middleware.CustomClaims)
	if !ok {
		return ctx.JSON(400, "invalid token")
	}

	var interests struct {
		Interests []string `json:"interests"`
	}

	if err := ctx.Bind(&interests); err != nil {
		return ctx.JSON(400, map[string]interface{}{
			"status":  "failed",
			"message": err.Error(),
		})
	}

	user, err := r.residentService.SetInterestsByUser(interests.Interests, userClaims.Email)
	if err != nil {
		return ctx.JSON(400, map[string]interface{}{
			"status":  "failed",
			"message": err.Error(),
		})
	}

	return ctx.JSON(200, user)
}

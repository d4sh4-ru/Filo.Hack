package endpoint

import (
	"Filo.Hack/internal/app/model"
	"Filo.Hack/internal/app/service"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type AuthEndpoint struct {
	authService     *service.AuthService
	residentService *service.ResidentService
}

func NewAuthEndpoint(residentService *service.ResidentService, authService *service.AuthService) *AuthEndpoint {
	return &AuthEndpoint{
		residentService: residentService,
		authService:     authService,
	}
}

func (e *AuthEndpoint) SignIn(ctx echo.Context) error {
	var req model.Resident

	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(400, map[string]interface{}{
			"status":  "failed",
			"message": err.Error(),
		})
	}

	token, err := e.authService.SignIn(req.Email, req.Password)

	if err != nil {
		return ctx.JSON(400, map[string]interface{}{
			"status":  "failed",
			"message": err.Error(),
		})
	}

	return ctx.JSON(200, map[string]string{"token": token})
}

func (e *AuthEndpoint) SignUp(ctx echo.Context) error {
	var req model.Resident
	var err error
	if err = ctx.Bind(&req); err != nil {
		return ctx.JSON(400, map[string]interface{}{
			"status":  "failed",
			"message": err.Error(),
		})
	}

	req.Password, err = hashPassword(req.Password)
	if err != nil {
		return ctx.JSON(400, map[string]interface{}{
			"status":  "failed",
			"message": err.Error(),
		})
	}

	if err = e.residentService.CreateResident(&req); err != nil {
		return ctx.JSON(400, map[string]interface{}{
			"status":  "failed",
			"message": err.Error(),
		})
	}

	return ctx.JSON(200, req)
}

// hashPassword хеширует пароль используя bcrypt
func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

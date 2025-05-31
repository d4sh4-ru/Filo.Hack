package router

import (
	"Filo.Hack/config"
	"Filo.Hack/internal/app/endpoint"
	"Filo.Hack/internal/app/middleware"
	"Filo.Hack/internal/app/repository"
	"Filo.Hack/internal/app/service"
	"Filo.Hack/internal/lib/storage"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func RegisterRouters(e *echo.Echo, dbClient *storage.DBClient, cfg *config.Config) {
	RegisterPublicRouters(e, dbClient, cfg)

	RegisterPrivateRouters(e, dbClient, cfg)
}

func RegisterPublicRouters(e *echo.Echo, dbClient *storage.DBClient, cfg *config.Config) {
	residentRepo := repository.NewResidentRepository(dbClient.Db)
	authService := service.NewAuthService(residentRepo, cfg.JWTSecret)
	log.Printf("JWTSecret %s", cfg.JWTSecret)
	residentService := service.NewResidentService(residentRepo)
	authController := endpoint.NewAuthEndpoint(residentService, authService)

	// Открытые маршруты
	e.POST("/register", authController.SignUp)
	e.POST("/login", authController.SignIn)
}

func RegisterPrivateRouters(e *echo.Echo, dbClient *storage.DBClient, cfg *config.Config) {
	jwtMiddleware := middleware.JWTMiddlewareConfig{
		SecretKey: []byte(cfg.HTTPServer.JWTSecret),
	}

	// Защищенные роуты
	api := e.Group("/api")
	api.Use(middleware.NewJWTMiddleware(jwtMiddleware))

	// Пользователи
	users := api.Group("/users")
	userRepo := repository.NewResidentRepository(dbClient.Db)
	residentService := service.NewResidentService(userRepo)
	userEndpoint := endpoint.NewResidentEndpoint(residentService)

	users.GET("/me", userEndpoint.GetMe)
	users.POST("/interests", userEndpoint.SetInterestsByUser)

	// События
	events := api.Group("/events")
	eventRepo := repository.NewEventRepository(dbClient.Db)
	eventService := service.NewEventService(eventRepo, userRepo)
	eventEndpoint := endpoint.NewEventEndpoint(eventService)

	events.POST("", eventEndpoint.CreateEvent)
}

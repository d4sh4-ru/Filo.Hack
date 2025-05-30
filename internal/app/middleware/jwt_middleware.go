package middleware

import (
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type JWTMiddlewareConfig struct {
	SecretKey []byte
}

type CustomClaims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

// NewJWTMiddleware - создает новый middleware с поддержкой кастомных claims
func NewJWTMiddleware(cfg JWTMiddlewareConfig) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Извлекаем заголовок Authorization
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return echo.NewHTTPError(http.StatusUnauthorized, "missing token")
			}

			// Убираем префикс "Bearer " из токена
			tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
			if tokenStr == authHeader { // Если префикс не был найден
				return echo.NewHTTPError(http.StatusUnauthorized, "invalid token format")
			}

			// Парсим токен с использованием кастомной структуры CustomClaims
			claims := &CustomClaims{}
			token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
				// Проверяем метод подписи
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, echo.NewHTTPError(http.StatusUnauthorized, "unexpected signing method")
				}
				return cfg.SecretKey, nil
			})

			// Проверяем ошибки парсинга и валидность токена
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "invalid token: "+err.Error())
			}
			if !token.Valid {
				return echo.NewHTTPError(http.StatusUnauthorized, "token is not valid")
			}

			// Сохраняем claims в контекст для использования в обработчиках
			c.Set("user", claims)

			// Передаем управление следующему обработчику
			return next(c)
		}
	}
}

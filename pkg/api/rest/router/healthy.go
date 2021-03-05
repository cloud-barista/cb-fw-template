package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// HealthCheck Method
// @Tags Default
// @Summary Health Check
// @Description for health check
// @ID HealthCheck
// @Accept json
// @Produce json
// @Success 200 {string} string "ok"
// @Router /healthcheck [get]
func HealthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "Cloud-Barista CB-Myfw")
}

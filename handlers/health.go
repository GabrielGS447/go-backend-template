package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// @Summary Show the status of server.
// @Description get the status of server.
// @Tags health
// @Accept */*
// @Produce plain
// @Success 200 "OK"
// @Router /health [get]
func HandleHealthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

package service

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	service Servicer
}

func NewHandler(service Servicer) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) GetTestService(c echo.Context) error {

	res := h.service.GetTestService(c)

	return c.JSON(http.StatusOK, &res)

}

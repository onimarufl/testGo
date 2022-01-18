package service

import "github.com/labstack/echo/v4"

type Service struct {
}

func New() *Service {
	return &Service{}
}

type Servicer interface {
	GetTestService(c echo.Context) string
}

func (s *Service) GetTestService(c echo.Context) string {

	return "test"

}

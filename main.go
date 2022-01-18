package main

import (
	"github.com/labstack/echo/v4"
	"github.com/onimarufl/testGo.git/service"
)

func main() {

	testService := service.NewHandler(
		service.New(),
	)

	e := echo.New()
	e.GET("/test", testService.GetTestService)

	e.Logger.Fatal(e.Start(":1323"))

}

package main

import (
	"botProject/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var e = createMux()

func main() {

	// 呼び出すための関数名は大文字から始めないといけないかも
	e.GET("/", handler.SetWordTest)
	e.POST("/post", handler.ShowWordTest)


	e.Logger.Fatal(e.Start(":8080"))
}

func createMux() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.CSRF())

	e.Static("/js", "src/js")
	

	return e
}

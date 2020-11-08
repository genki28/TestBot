package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// 呼び出すための関数名は大文字から始めないといけないかも
func SetWordTest(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World!!")
}

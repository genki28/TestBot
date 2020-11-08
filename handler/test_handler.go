package handler

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

// 呼び出すための関数名は大文字から始めないといけない
func SetWordTest(c echo.Context) error {

	data := map[string]interface{}{
		"name": "test",
	}

	return render(c, "index.html", data)
}

func ShowWordTest(c echo.Context) error {
	getName := c.FormValue("text")
	if getName == "こんにちは" {
		fmt.Println("こんばんは")
	} else if getName == "こんばんは" {
		fmt.Println("こんにちは")
	} else {
		fmt.Println("データにありません")
	}

	data := map[string]interface{}{
		"name": getName,
	}

	return render(c, "index.html", data)
}

package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/start", func(c echo.Context) error {
		return c.String(http.StatusOK, "staring...")
	})
	e.Logger.Fatal(e.Start(":1323"))
}

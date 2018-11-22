package main

import (
	"io"
	"net/http"
	"os"

	"github.com/labstack/echo"
)

func uploadImage(c echo.Context) error {
	caty := c.FormValue("category")
	image, err := c.FormFile("image")
	if err != nil {
		return err
	}

	// Source
	src, err := image.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Destination
	dst, err := os.Create(image.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"category": caty,
	})
}

func main() {
	e := echo.New()
	e.GET("/start", func(c echo.Context) error {
		return c.String(http.StatusOK, "staring...")
	})

	e.POST("/upload", uploadImage)

	e.Logger.Fatal(e.Start(":1323"))
}

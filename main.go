package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	vision "cloud.google.com/go/vision/apiv1"
	"github.com/labstack/echo"
)

func uploadImage(c echo.Context) error {
	image, err := c.FormFile("image")
	if err != nil {
		return err
	}

	// Source
	file, err := image.Open()
	if err != nil {
		return err
	}
	defer file.Close()

	ctx := context.Background()

	// Creates a client.
	client, err := vision.NewImageAnnotatorClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	img, err := vision.NewImageFromReader(file)
	if err != nil {
		log.Fatalf("Failed to create image: %v", err)
	}

	annotation, err := client.DetectDocumentText(ctx, img, nil)
	if err != nil {
		log.Fatal(err)
	}

	if annotation == nil {
		fmt.Println("No text found.")
	} else {
		fmt.Println("Document Text:")
		fmt.Printf("%q\n", annotation.Text)

		fmt.Println("Pages:")
		for _, page := range annotation.Pages {
			fmt.Printf("\tConfidence: %f, Width: %d, Height: %d\n", page.Confidence, page.Width, page.Height)
			fmt.Println("\tBlocks:")
			for _, block := range page.Blocks {
				fmt.Printf("\t\tConfidence: %f, Block type: %v\n", block.Confidence, block.BlockType)
				fmt.Println("\t\tParagraphs:")
				for _, paragraph := range block.Paragraphs {
					fmt.Printf("\t\t\tConfidence: %f", paragraph.Confidence)
					fmt.Println("\t\t\tWords:")
					for _, word := range paragraph.Words {
						symbols := make([]string, len(word.Symbols))
						for i, s := range word.Symbols {
							symbols[i] = s.Text
						}
						wordText := strings.Join(symbols, "")
						fmt.Printf("\t\t\t\tConfidence: %f, Symbols: %s\n", word.Confidence, wordText)
					}
				}
			}
		}
	}

	return c.JSON(http.StatusOK, annotation)
}

func main() {
	e := echo.New()

	api := e.Group("/api")
	api.GET("/start", func(c echo.Context) error {
		return c.String(http.StatusOK, "staring...")
	})

	api.POST("/upload", uploadImage)

	e.Logger.Fatal(e.Start(":1323"))
}

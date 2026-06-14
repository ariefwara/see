package main

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"path/filepath"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

var supportedExts = map[string]string{
	".png":  "PNG",
	".jpg":  "JPEG",
	".jpeg": "JPEG",
	".gif":  "GIF",
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: see <image-file>")
		os.Exit(1)
	}

	filePath := os.Args[1]

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		fmt.Printf("Error: file '%s' not found\n", filePath)
		os.Exit(1)
	}

	ext := strings.ToLower(filepath.Ext(filePath))
	if _, ok := supportedExts[ext]; !ok {
		fmt.Printf("Error: unsupported format '%s' (supported: png, jpg, jpeg, gif)\n", ext)
		os.Exit(1)
	}

	f, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error: cannot open file '%s'\n", filePath)
		os.Exit(1)
	}
	defer f.Close()

	if _, _, err := image.Decode(f); err != nil {
		fmt.Printf("Error: '%s' is not a valid image file\n", filePath)
		os.Exit(1)
	}
	f.Close()

	a := app.New()
	w := a.NewWindow(filepath.Base(filePath))

	img := canvas.NewImageFromFile(filePath)
	img.FillMode = canvas.ImageFillOriginal

	w.SetContent(container.NewScroll(img))
	w.Resize(fyne.NewSize(900, 700))
	w.CenterOnScreen()
	w.ShowAndRun()
}

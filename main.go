package main

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

type viewer struct {
	w          fyne.Window
	img        *canvas.Image
	scroll     *container.Scroll
	files      []string
	currentIdx int
}

var supportedExts = map[string]bool{
	".png":  true,
	".jpg":  true,
	".jpeg": true,
	".gif":  true,
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: see <image-file | directory>")
		fmt.Println("       see .     # open first image in current directory")
		os.Exit(1)
	}

	arg := os.Args[1]

	if info, err := os.Stat(arg); os.IsNotExist(err) {
		fmt.Printf("Error: '%s' not found\n", arg)
		os.Exit(1)
	} else if info.IsDir() {
		files := listImages(arg)
		path, idx := firstValidImage(files)
		if path == "" {
			fmt.Printf("Error: no valid images found in '%s'\n", arg)
			os.Exit(1)
		}
		openViewer(files, idx)
		return
	}

	// Argument is a file — validate and open
	filePath := arg

	ext := strings.ToLower(filepath.Ext(filePath))
	if !supportedExts[ext] {
		fmt.Printf("Error: unsupported format '%s' (supported: png, jpg, jpeg, gif)\n", ext)
		os.Exit(1)
	}

	f, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error: cannot open file '%s'\n", filePath)
		os.Exit(1)
	}
	if _, _, err := image.Decode(f); err != nil {
		f.Close()
		fmt.Printf("Error: '%s' is not a valid image file\n", filePath)
		os.Exit(1)
	}
	f.Close()

	// List all supported images in the same directory
	dir := filepath.Dir(filePath)
	files := listImages(dir)

	currentIdx := -1
	for i, p := range files {
		if p == filePath {
			currentIdx = i
			break
		}
	}
	if currentIdx == -1 {
		files = append(files, filePath)
		currentIdx = len(files) - 1
	}

	openViewer(files, currentIdx)
}

func openViewer(files []string, startIdx int) {
	path := files[startIdx]

	a := app.New()
	w := a.NewWindow(filepath.Base(path))

	img := canvas.NewImageFromFile(path)
	img.FillMode = canvas.ImageFillOriginal

	scroll := container.NewScroll(img)

	v := &viewer{
		w:          w,
		img:        img,
		scroll:     scroll,
		files:      files,
		currentIdx: startIdx,
	}

	w.Canvas().SetOnTypedKey(func(ev *fyne.KeyEvent) {
		switch ev.Name {
		case fyne.KeyLeft:
			v.prev()
		case fyne.KeyRight:
			v.next()
		case fyne.KeyEscape:
			w.Close()
		}
	})

	w.SetContent(scroll)
	w.Resize(fyne.NewSize(900, 700))
	w.CenterOnScreen()
	w.ShowAndRun()
}

func (v *viewer) loadImage(idx int) {
	if idx < 0 || idx >= len(v.files) {
		return
	}

	path := v.files[idx]
	v.img.File = path
	v.img.Image = nil
	canvas.Refresh(v.img)

	v.currentIdx = idx
	v.w.SetTitle(filepath.Base(path))

	v.scroll.Offset.X = 0
	v.scroll.Offset.Y = 0
	canvas.Refresh(v.scroll)
}

func (v *viewer) next() {
	if v.currentIdx < len(v.files)-1 {
		v.loadImage(v.currentIdx + 1)
	}
}

func (v *viewer) prev() {
	if v.currentIdx > 0 {
		v.loadImage(v.currentIdx - 1)
	}
}

func listImages(dir string) []string {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil
	}

	var files []string
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		ext := strings.ToLower(filepath.Ext(entry.Name()))
		if supportedExts[ext] {
			files = append(files, filepath.Join(dir, entry.Name()))
		}
	}

	sort.Slice(files, func(i, j int) bool {
		return strings.ToLower(filepath.Base(files[i])) < strings.ToLower(filepath.Base(files[j]))
	})

	return files
}

func firstValidImage(files []string) (string, int) {
	for i, path := range files {
		f, err := os.Open(path)
		if err != nil {
			continue
		}
		_, _, err = image.Decode(f)
		f.Close()
		if err == nil {
			return path, i
		}
	}
	return "", -1
}

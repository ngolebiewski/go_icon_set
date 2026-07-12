package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
	"path/filepath"
	"strconv"

	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
)

// GenerateCheatSheet creates a visual directory layout of all your current icons
func GenerateCheatSheet(cfg *Config) {
	padding := 16
	boxSize := 48 // Area allocated for the 16x16 icon + sub-label text

	// Dynamic vertical height calculation row-by-row
	totalRows := (cfg.TotalIcons + cfg.Columns - 1) / cfg.Columns

	sheetW := cfg.Columns*(boxSize+padding) + padding
	sheetH := totalRows*(boxSize+padding+16) + padding

	cheatCanvas := image.NewRGBA(image.Rect(0, 0, sheetW, sheetH))

	// Dark slate theme background color (#0f172a)
	draw.Draw(cheatCanvas, cheatCanvas.Bounds(), &image.Uniform{color.RGBA{15, 23, 42, 255}}, image.Point{}, draw.Src)

	fontDrawer := &font.Drawer{
		Dst:  cheatCanvas,
		Src:  image.NewUniform(color.RGBA{148, 163, 184, 255}), // slate text label color
		Face: basicfont.Face7x13,
	}

	fmt.Printf("🎨 Generating Cheat Sheet canvas (%dx%d pixels)...\n", sheetW, sheetH)

	for i := 0; i < cfg.TotalIcons; i++ {
		col := i % cfg.Columns
		row := i / cfg.Columns

		boxX := padding + col*(boxSize+padding)
		boxY := padding + row*(boxSize+padding+16)

		// Border outline bounding box (#334155)
		drawBorderBox(cheatCanvas, boxX, boxY, boxSize, boxSize, color.RGBA{51, 65, 85, 255})

		name := "blank"
		indexStr := strconv.Itoa(i)
		if mappedName, exists := cfg.SourceList[indexStr]; exists && mappedName != "" {
			name = mappedName
		}

		// Draw icon into the exact dead center of the box grid segment
		if name != "blank" {
			iconX := boxX + (boxSize-cfg.IconWidth)/2
			iconY := boxY + (boxSize-cfg.IconHeight)/2

			srcX := col * cfg.IconWidth
			srcY := row * cfg.IconHeight
			cropRect := image.Rect(srcX, srcY, srcX+cfg.IconWidth, srcY+cfg.IconHeight)

			draw.Draw(cheatCanvas, image.Rect(iconX, iconY, iconX+cfg.IconWidth, iconY+cfg.IconHeight), cfg.LoadedImage, cropRect.Min, draw.Over)
		}

		// Truncate overly long labels so they stay inside the box bounds
		if len(name) > 6 {
			name = name[:5] + ".."
		}

		textX := boxX + (boxSize-len(name)*7)/2
		textY := boxY + boxSize + 12
		fontDrawer.Dot = fixed.Point26_6{X: fixed.I(textX), Y: fixed.I(textY)}
		fontDrawer.DrawString(name)
	}

	outPath := filepath.Join(cfg.OutputDir, "cheatsheet.png")
	out, err := os.Create(outPath)
	if err != nil {
		fmt.Printf("⚠️ Error creating cheatsheet file: %v\n", err)
		return
	}
	defer out.Close()

	png.Encode(out, cheatCanvas)
	fmt.Printf("🎯 Cheat Sheet successfully generated at: %s\n", outPath)
}

func drawBorderBox(img *image.RGBA, x, y, w, h int, c color.RGBA) {
	for i := 0; i < w; i++ {
		img.SetRGBA(x+i, y, c)
		img.SetRGBA(x+i, y+h-1, c)
	}
	for j := 0; j < h; j++ {
		img.SetRGBA(x, y+j, c)
		img.SetRGBA(x+w-1, y+j, c)
	}
}

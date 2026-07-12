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
// GenerateCheatSheet creates a visual directory layout of all your current icons
func GenerateCheatSheet(cfg *Config) {
	padding := 16

	// 🔥 Dynamic layout sizing: If icons are tiny, double them and give them a bigger box
	scaleFactor := 1
	if cfg.IconWidth < 32 && cfg.IconHeight < 32 {
		scaleFactor = 2
	}

	boxSize := 48
	if scaleFactor > 1 {
		boxSize = 64 // Expand container box to accommodate 32x32 upscaled graphics cleanly
	}

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

	fmt.Printf("🎨 Generating Cheat Sheet canvas (%dx%d pixels) [Scale: %dx]...\n", sheetW, sheetH, scaleFactor)

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
			scaledW := cfg.IconWidth * scaleFactor
			scaledH := cfg.IconHeight * scaleFactor

			iconX := boxX + (boxSize-scaledW)/2
			iconY := boxY + (boxSize-scaledH)/2

			srcX := col * cfg.IconWidth
			srcY := row * cfg.IconHeight

			// 🔥 Nearest-Neighbor Upscaling Engine Loop
			// Maps every canvas coordinate target step back to its original source coordinate index
			for dy := 0; dy < scaledH; dy++ {
				for dx := 0; dx < scaledW; dx++ {
					origX := srcX + (dx / scaleFactor)
					origY := srcY + (dy / scaleFactor)

					// Pull original pixel color values safely from source sheet coordinates
					pixelColor := cfg.LoadedImage.At(origX, origY)

					// Paint onto the canvas destination matrix coordinates
					cheatCanvas.Set(iconX+dx, iconY+dy, pixelColor)
				}
			}
		}

		// Truncate overly long labels so they stay inside the box bounds
		if len(name) > (boxSize / 8) {
			name = name[:(boxSize/8)-1] + ".."
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

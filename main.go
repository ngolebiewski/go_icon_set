package main

import (
	"encoding/json"
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

type Config struct {
	IconWidth  int    // icon width ie 16px
	IconHeight int    // icon height
	Columns    int    // number of icons in a row
	SourceFile string // path to spritesheet, defaults to current directory
	OutputDir  string // Directory where to save files, defaults to current directory
	PackageDir string // for the Templ Build

	// JSON Source List mapping a grid slot index string to an icon name
	SourceList map[string]string

	// State fields filled dynamically during verifyConfig
	SheetWidth  int
	SheetHeight int
	TotalIcons  int
	LoadedImage image.Image
}

func initConfig() *Config {
	cfg := Config{
		IconWidth:  16,
		IconHeight: 16,
		Columns:    10,
		SourceFile: "spritesheet.png", // Looks in current directory
		OutputDir:  "icons",           // Subfolder in current directory
		SourceList: make(map[string]string),
		PackageDir: "./temple_icons", // Sub-package folder kept right at top-level
	}
	return &cfg
}

func verifyConfig(cfg *Config) error {
	// 1. Ensure the output preview folders exist
	pngDir := filepath.Join(cfg.OutputDir, "png")
	if err := os.MkdirAll(pngDir, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	svgDir := filepath.Join(cfg.OutputDir, "svg")
	if err := os.MkdirAll(svgDir, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	// 2. Load the JSON configuration list if it exists
	jsonPath := "icon_names.json"
	if jsonFile, err := os.ReadFile(jsonPath); err == nil {
		if err := json.Unmarshal(jsonFile, &cfg.SourceList); err != nil {
			return fmt.Errorf("failed to parse JSON source list: %w", err)
		}
		fmt.Println("📋 Loaded icon name mappings from JSON.")
	}

	// 3. Open and decode the master spritesheet
	file, err := os.Open(cfg.SourceFile)
	if err != nil {
		return fmt.Errorf("cannot open spritesheet file %s: %w", cfg.SourceFile, err)
	}
	defer file.Close()

	img, err := png.Decode(file)
	if err != nil {
		return fmt.Errorf("failed to decode master PNG: %w", err)
	}

	// 4. Extract and verify file structural limits
	bounds := img.Bounds()
	cfg.SheetWidth = bounds.Dx()
	cfg.SheetHeight = bounds.Dy()
	cfg.LoadedImage = img

	if cfg.SheetWidth%cfg.IconWidth != 0 || cfg.SheetHeight%cfg.IconHeight != 0 {
		fmt.Printf("⚠️ Warning: Spritesheet dimensions (%dx%d) do not perfectly match icon sizes (%dx%d).\n",
			cfg.SheetWidth, cfg.SheetHeight, cfg.IconWidth, cfg.IconHeight)
	}

	totalCols := cfg.SheetWidth / cfg.IconWidth
	totalRows := cfg.SheetHeight / cfg.IconHeight
	cfg.TotalIcons = totalCols * totalRows

	return nil
}

// sliceSpritesheet cuts up the image, saves preview files, and returns a map of valid SVG path strings
func sliceSpritesheet(cfg *Config) map[string]string {
	fmt.Printf("✂️  Slicing %d potential icons...\n", cfg.TotalIcons)
	compiledPaths := make(map[string]string)

	for i := 0; i < cfg.TotalIcons; i++ {
		col := i % cfg.Columns
		row := i / cfg.Columns

		x := col * cfg.IconWidth
		y := row * cfg.IconHeight

		if y+cfg.IconHeight > cfg.SheetHeight {
			break
		}

		cropBounds := image.Rect(x, y, x+cfg.IconWidth, y+cfg.IconHeight)
		iconCanvas := image.NewRGBA(image.Rect(0, 0, cfg.IconWidth, cfg.IconHeight))
		draw.Draw(iconCanvas, iconCanvas.Bounds(), cfg.LoadedImage, cropBounds.Min, draw.Src)

		name := fmt.Sprintf("%d", i)
		indexStr := fmt.Sprintf("%d", i)
		if mappedName, exists := cfg.SourceList[indexStr]; exists && mappedName != "" {
			name = mappedName
		}

		outPath := filepath.Join(cfg.OutputDir, "png", name+".png")

		f, err := os.Create(outPath)
		if err != nil {
			fmt.Printf("⚠️ Error creating file %s: %v\n", outPath, err)
			continue
		}

		if err := png.Encode(f, iconCanvas); err != nil {
			f.Close()
			fmt.Printf("⚠️ Error encoding file %s: %v\n", outPath, err)
			continue
		}
		f.Close()

		// Generate SVG path string and create preview file
		pathString := ConvertPNGtoSVG(name, iconCanvas, cfg)

		// If the path isn't empty and isn't explicitly named "blank", track it for the Templ package
		if pathString != "" && name != "blank" {
			compiledPaths[name] = pathString
		}
	}

	return compiledPaths
}

func ConvertPNGtoSVG(name string, img image.Image, cfg *Config) string {
	rgbaImg, ok := img.(*image.RGBA)
	if !ok {
		bounds := img.Bounds()
		rgbaImg = image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
		draw.Draw(rgbaImg, rgbaImg.Bounds(), img, bounds.Min, draw.Src)
	}

	var pathSegments []string

	for y := 0; y < cfg.IconHeight; y++ {
		inSpan := false
		spanStart := 0

		for x := 0; x < cfg.IconWidth; x++ {
			color := rgbaImg.RGBAAt(x, y)
			isSolid := color.A > 0

			if isSolid && !inSpan {
				inSpan = true
				spanStart = x
			} else if !isSolid && inSpan {
				spanLength := x - spanStart
				segment := fmt.Sprintf("M%d,%dh%dv1h-%dz", spanStart, y, spanLength, spanLength)
				pathSegments = append(pathSegments, segment)
				inSpan = false
			}
		}

		if inSpan {
			spanLength := cfg.IconWidth - spanStart
			segment := fmt.Sprintf("M%d,%dh%dv1h-%dz", spanStart, y, spanLength, spanLength)
			pathSegments = append(pathSegments, segment)
		}
	}

	if len(pathSegments) == 0 {
		return ""
	}

	pathData := strings.Join(pathSegments, " ")

	// Save individual preview SVG file locally
	svgPath := filepath.Join(cfg.OutputDir, "svg", name+".svg")
	f, err := os.Create(svgPath)
	if err == nil {
		fmt.Fprintf(f, `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 %d %d" shape-rendering="crispEdges" style="background: transparent;">`+"\n", cfg.IconWidth, cfg.IconHeight)
		fmt.Fprintf(f, `  <path d="%s" fill="currentColor" style="color: white;" />`+"\n", pathData)
		fmt.Fprintln(f, "</svg>")
		f.Close()
	}

	return pathData
}

func main() {
	cfg := initConfig()

	if err := verifyConfig(cfg); err != nil {
		fmt.Printf("❌ Configuration Error: %v\n", err)
		os.Exit(1)
	}

	// 1. Slice image assets and collect path string maps
	pathsMap := sliceSpritesheet(cfg)
	fmt.Println("🎉 Stage 1 complete, SVG and PNG preview files rendered!")

	// 2. Generate the visual cheat sheet layout grid image file (from cheatsheet.go)
	GenerateCheatSheet(cfg)

	// 3. Compile everything together into the temple_icons package distribution folder (from makeTempl.go)
	WriteTemplPackage(cfg.PackageDir, pathsMap)

	fmt.Println("🎉 All systems complete!")
}

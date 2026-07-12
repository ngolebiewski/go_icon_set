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

	// JSON Source List mapping a grid slot index string to an icon name
	// e.g., {"0": "home", "1": "save", "2": "github"}
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
	}
	return &cfg
}

func verifyConfig(cfg *Config) error {
	// 1. Ensure the output folders exist (/png subfolder)
	pngDir := filepath.Join(cfg.OutputDir, "png")
	if err := os.MkdirAll(pngDir, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	svgDir := filepath.Join(cfg.OutputDir, "svg")
	if err := os.MkdirAll(svgDir, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	// 2. Load the JSON configuration list if it exists (optional for numerical defaults)
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

	// Sanity Check: Does the sheet split evenly into your sizes?
	if cfg.SheetWidth%cfg.IconWidth != 0 || cfg.SheetHeight%cfg.IconHeight != 0 {
		fmt.Printf("⚠️ Warning: Spritesheet dimensions (%dx%d) do not perfectly match icon sizes (%dx%d).\n",
			cfg.SheetWidth, cfg.SheetHeight, cfg.IconWidth, cfg.IconHeight)
	}

	totalCols := cfg.SheetWidth / cfg.IconWidth
	totalRows := cfg.SheetHeight / cfg.IconHeight
	cfg.TotalIcons = totalCols * totalRows

	return nil
}

func sliceSpritesheet(cfg *Config) {
	fmt.Printf("✂️  Slicing %d potential icons...\n", cfg.TotalIcons)

	for i := 0; i < cfg.TotalIcons; i++ {
		col := i % cfg.Columns
		row := i / cfg.Columns

		x := col * cfg.IconWidth
		y := row * cfg.IconHeight

		// Break early if our calculation overshoots the hard image boundary
		if y+cfg.IconHeight > cfg.SheetHeight {
			break
		}

		// Define boundaries and cut the layout
		cropBounds := image.Rect(x, y, x+cfg.IconWidth, y+cfg.IconHeight)
		iconCanvas := image.NewRGBA(image.Rect(0, 0, cfg.IconWidth, cfg.IconHeight))
		draw.Draw(iconCanvas, iconCanvas.Bounds(), cfg.LoadedImage, cropBounds.Min, draw.Src)

		// Determine Name: Use JSON registration if available, otherwise fallback to numbers
		name := fmt.Sprintf("%d", i)
		indexStr := fmt.Sprintf("%d", i)
		if mappedName, exists := cfg.SourceList[indexStr]; exists && mappedName != "" {
			name = mappedName
		}

		outPath := filepath.Join(cfg.OutputDir, "png", name+".png")

		// Save the file
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

		// 🔥 CALL THE NEW SVG CONVERTER RIGHT HERE
		ConvertPNGtoSVG(name, iconCanvas, cfg)
	}
}

func ConvertPNGtoSVG(name string, img image.Image, cfg *Config) {
	rgbaImg, ok := img.(*image.RGBA)
	if !ok {
		bounds := img.Bounds()
		rgbaImg = image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
		draw.Draw(rgbaImg, rgbaImg.Bounds(), img, bounds.Min, draw.Src)
	}

	svgPath := filepath.Join(cfg.OutputDir, "svg", name+".svg")
	f, err := os.Create(svgPath)
	if err != nil {
		fmt.Printf("⚠️ Error creating SVG file %s: %v\n", svgPath, err)
		return
	}
	defer f.Close()

	// 1. Setup the wrapper with crispEdges and transparent bounds
	fmt.Fprintf(f, `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 %d %d" shape-rendering="crispEdges" style="background: transparent;">`+"\n", cfg.IconWidth, cfg.IconHeight)

	var pathSegments []string

	for y := 0; y < cfg.IconHeight; y++ {
		inSpan := false
		spanStart := 0

		for x := 0; x < cfg.IconWidth; x++ {
			color := rgbaImg.RGBAAt(x, y)

			// Since your PNG is perfectly sharp, any active pixel has Alpha
			isSolid := color.A > 0

			if isSolid && !inSpan {
				inSpan = true
				spanStart = x
			} else if !isSolid && inSpan {
				spanLength := x - spanStart
				// Standardized absolute/relative instruction path segment
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

	// 2. Join segments with explicit spacing so the SVG string reads cleanly
	if len(pathSegments) > 0 {
		pathData := strings.Join(pathSegments, " ")
		// fill="currentColor" keeps it dynamic, style="color: white;" makes it white by default
		fmt.Fprintf(f, `  <path d="%s" fill="currentColor" style="color: white;" />`+"\n", pathData)
	}

	fmt.Fprintln(f, "</svg>")
}

func main() {
	cfg := initConfig()

	// Verify configuration and exit gracefully if initialization errors hit
	if err := verifyConfig(cfg); err != nil {
		fmt.Printf("❌ Configuration Error: %v\n", err)
		os.Exit(1)
	}

	// Perform the slicing step safely now that our fields are validated
	sliceSpritesheet(cfg)
	fmt.Println("🎉 Stage 1 complete, SVG and PNG files rendered!")

	// 2. 🔥 CALL THE CHEATSHEET ENGINE GENERATOR HERE
	GenerateCheatSheet(cfg)

	fmt.Println("🎉 All systems complete!")
}

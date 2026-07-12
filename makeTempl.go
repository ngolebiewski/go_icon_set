package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func toTitleCase(input string) string {
	input = strings.ReplaceAll(input, "-", " ")
	input = strings.ReplaceAll(input, "_", " ")
	words := strings.Fields(input)
	for i, w := range words {
		words[i] = strings.ToUpper(w[:1]) + w[1:]
	}
	return strings.Join(words, "")
}

// WriteTemplPackage compiles the stored vector paths into a usable sub-package component file
func WriteTemplPackage(outputDir string, compiledPaths map[string]string) {
	// Ensure the target subfolder exists cleanly
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		fmt.Printf("❌ Failed to create package output dir: %v\n", err)
		return
	}

	// Build package icons.go with global defaults injection state manager APIs
	// Inside makeTempl.go -> WriteTemplPackage function:

	code := `package temple_icons

import (
	"context"
	"fmt"
	"io"
	"github.com/a-h/templ"
)

type IconConfig struct {
	Size  string
	Color string
}

var currentDefault = IconConfig{
	Size:  "16px",
	Color: "white",
}

// SetGlobalDefaults alters the rendering token layout definitions application-wide
func SetGlobalDefaults(cfg IconConfig) {
	if cfg.Size != ""  { currentDefault.Size = cfg.Size }
	if cfg.Color != "" { currentDefault.Color = cfg.Color }
}

func buildStyleAttr(cfgs []IconConfig) string {
	cfg := currentDefault // 🔥 FIX: added the colon here so it declares the variable
	if len(cfgs) > 0 {
		if cfgs[0].Size != ""  { cfg.Size = cfgs[0].Size }
		if cfgs[0].Color != "" { cfg.Color = cfgs[0].Color }
	}
	return fmt.Sprintf("style=\"width: %s; height: %s; color: %s; display: inline-block; vertical-align: middle;\"", cfg.Size, cfg.Size, cfg.Color)
}
`

	for name, path := range compiledPaths {
		goName := toTitleCase(name)

		code += fmt.Sprintf(`
// %s renders the sharp vector icon for "%s"
func %s(cfg ...IconConfig) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		style := buildStyleAttr(cfg)
		_, err := io.WriteString(w, fmt.Sprintf("<svg xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 16 16\" shape-rendering=\"crispEdges\" fill=\"currentColor\" %%s><path d=\"%s\" /></svg>", style))
		return err
	})
}
`, goName, name, goName, path)
	}

	err := os.WriteFile(filepath.Join(outputDir, "icons.go"), []byte(code), 0644)
	if err != nil {
		fmt.Printf("❌ Failed to compile icon package: %v\n", err)
		return
	}
	fmt.Println("🚀 Sub-package folder 'temple_icons' compiled successfully!")
}

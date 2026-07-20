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
func WriteTemplPackage(outputDir string, compiledPaths map[string]string, sortedKeys []string) {
	// Ensure the target subfolder exists cleanly
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		fmt.Printf("❌ Failed to create package output dir: %v\n", err)
		return
	}

	// Build package icons.go with global defaults injection state manager APIs
	code := `package templ_icons

import (
	"context"
	"fmt"
	"io"
	"strings"
	"github.com/a-h/templ"
)

type IconConfig struct {
	Size  string
	Color string
	Class string
	Alt   string // Empty = Decorative (Default), Text = Semantic Image
}

var currentDefault = IconConfig{
	Size:  "16px",
	Color: "white",
}

// SetGlobalDefaults alters the rendering token layout definitions application-wide
func SetGlobalDefaults(cfg IconConfig) {
	if cfg.Size != ""  { currentDefault.Size = cfg.Size }
	if cfg.Color != "" { currentDefault.Color = cfg.Color }
	if cfg.Class != "" { currentDefault.Class = cfg.Class }
}

func buildStyleAttr(cfgs []IconConfig) string {
	cfg := currentDefault 
	if len(cfgs) > 0 {
		if cfgs[0].Size != ""  { cfg.Size = cfgs[0].Size }
		if cfgs[0].Color != "" { cfg.Color = cfgs[0].Color }
		if cfgs[0].Class != "" { cfg.Class = cfgs[0].Class }
		if cfgs[0].Alt != ""   { cfg.Alt = cfgs[0].Alt }
	}

	// 1. Build standard inline styling rules
	var styles []string
	if cfg.Size != "" {
		styles = append(styles, fmt.Sprintf("width: %s; height: %s;", cfg.Size, cfg.Size))
	}
	if cfg.Color != "" {
		styles = append(styles, fmt.Sprintf("color: %s;", cfg.Color))
	}
	styles = append(styles, "display: inline-block; vertical-align: middle;")
	styleAttr := fmt.Sprintf("style=\"%s\"", strings.Join(styles, " "))

	// 2. Build explicit CSS class names layout mapping if provided
	classAttr := ""
	if cfg.Class != "" {
		classAttr = fmt.Sprintf("class=%q ", cfg.Class)
	}

	// 3. Build Accessibility (A11y) variables array attributes
	var a11yAttrs []string
	a11yAttrs = append(a11yAttrs, "focusable=\"false\"") // Prevents old/buggy browser navigation traps

	if cfg.Alt != "" {
		// Semantic state: The vector represents a standalone actionable item
		a11yAttrs = append(a11yAttrs, "role=\"img\"")
		a11yAttrs = append(a11yAttrs, fmt.Sprintf("aria-label=%q", cfg.Alt))
	} else {
		// Decorative state (Default): Quietly hidden from structural screen readers
		a11yAttrs = append(a11yAttrs, "aria-hidden=\"true\"")
	}

	// Combine all built attributes cleanly separated by spaces
	return fmt.Sprintf("%s%s %s", classAttr, styleAttr, strings.Join(a11yAttrs, " "))
}
`

	for _, name := range sortedKeys {
		path := compiledPaths[name]
		goName := toTitleCase(name)

		code += fmt.Sprintf(`
// %s renders the sharp vector icon for "%s"
func %s(cfg ...IconConfig) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := buildStyleAttr(cfg)
		_, err := io.WriteString(w, fmt.Sprintf("<svg xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 16 16\" shape-rendering=\"crispEdges\" fill=\"currentColor\" %%s><path d=\"%s\" /></svg>", attrs))
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
	fmt.Println("🚀 Sub-package folder 'templ_icons' compiled successfully with A11y defaults!")
}

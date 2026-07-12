package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// WriteJSPackage compiles your vectors into a lightweight, typed JS library folder
func WriteJSPackage(outputDir string, compiledPaths map[string]string) {
	_ = os.MkdirAll(outputDir, 0755)

	// 1. Write the core JavaScript code
	jsCode := `let globalSize = "16px";
let globalColor = "white";

export function setGlobalDefaults({ size, color } = {}) {
	if (size) globalSize = size;
	if (color) globalColor = color;
}

function renderIcon(path, config = {}) {
	const size = config.size || globalSize;
	const color = config.color || globalColor;
	const style = "width: " + size + "; height: " + size + "; color: " + color + "; display: inline-block; vertical-align: middle;";
	return "<svg xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 16 16\" shape-rendering=\"crispEdges\" fill=\"currentColor\" style=\"" + style + "\"><path d=\"" + path + "\" /></svg>";
}
`

	// 2. Write the matching TypeScript Typings structural layout
	dtsCode := `export interface IconConfig {
	size?: string;
	color?: string;
}

export function setGlobalDefaults(cfg: IconConfig): void;
type IconComponent = (cfg?: IconConfig) => string;
`

	for name, path := range compiledPaths {
		goName := toTitleCase(name)
		jsCode += fmt.Sprintf("export const %s = (cfg) => renderIcon(\"%s\", cfg);\n", goName, path)
		dtsCode += fmt.Sprintf("export const %s: IconComponent;\n", goName)
	}

	_ = os.WriteFile(filepath.Join(outputDir, "icons.js"), []byte(jsCode), 0644)
	_ = os.WriteFile(filepath.Join(outputDir, "icons.d.ts"), []byte(dtsCode), 0644)
	fmt.Println("🚀 JS/TS package folder 'js_icons' compiled successfully!")
}

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// WriteJSPackage compiles your vectors into four distinct frontend flavors
func WriteJSPackage(outputDir string, compiledPaths map[string]string) {
	_ = os.MkdirAll(outputDir, 0755)

	// ==========================================
	// TARGET 1: YOUR ORIGINAL STRING COMPONENTS
	// ==========================================
	stringComponentsCode := `let globalSize = "16px";
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

	// ==========================================
	// TARGET 2: NATIVE WEB COMPONENTS
	// ==========================================
	webComponentCode := `class CrispIcon extends HTMLElement {
	constructor() { super(); this.attachShadow({ mode: 'open' }); }
	connectedCallback() {
		const size = this.getAttribute('size') || '16px';
		const color = this.getAttribute('color') || 'white';
		const svg = document.createElementNS('http://www.w3.org/2000/svg', 'svg');
		svg.setAttribute('viewBox', '0 0 16 16');
		svg.setAttribute('shape-rendering', 'crispEdges');
		svg.setAttribute('fill', 'currentColor');
		svg.style.cssText = "width: " + size + "; height: " + size + "; color: " + color + "; display: inline-block; vertical-align: middle;";
		const path = document.createElementNS('http://www.w3.org/2000/svg', 'path');
		path.setAttribute('d', this.constructor.path || '');
		svg.appendChild(path);
		this.shadowRoot.appendChild(svg);
	}
}
`

	// ==========================================
	// TARGET 3: REACT COMPONENTS
	// ==========================================
	reactCode := `import React from 'react';

function Icon({ path, size = "16px", color = "white", ...props }) {
	const style = { width: size, height: size, color: color, display: "inline-block", verticalAlign: "middle" };
	return (
		<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 16 16" shapeRendering="crispEdges" fill="currentColor" style={style} {...props}>
			<path d={path} />
		</svg>
	);
}
`

	// ==========================================
	// TARGET 4: SSR PURE DICTIONARY
	// ==========================================
	ssrCode := `// Pure raw dictionary of vectors for server-side template resolution
export const IconVectors = {
`

	// Loop once to dynamically compile all targets synchronously
	for name, path := range compiledPaths {
		goName := toTitleCase(name)
		tagName := "icon-" + strings.ToLower(name)

		// 1. Original String Append
		stringComponentsCode += fmt.Sprintf("export const %s = (cfg) => renderIcon(\"%s\", cfg);\n", goName, path)

		// 2. Web Component Append
		webComponentCode += fmt.Sprintf("export class %s extends CrispIcon { static path = %q; }\ncustomElements.define(%q, %s);\n", goName, path, tagName, goName)

		// 3. React Append
		reactCode += fmt.Sprintf("export const %s = (props) => <Icon path=%q {...props} />;\n", goName, path)

		// 4. SSR Dictionary Append (🔥 FIX: Explicitly quote object literal keys to prevent parsing drops)
		ssrCode += fmt.Sprintf("\t%q: %q,\n", name, path)
	}

	ssrCode += "};\n"

	// Commit everything cleanly to disk
	_ = os.WriteFile(filepath.Join(outputDir, "string-components.js"), []byte(stringComponentsCode), 0644)
	_ = os.WriteFile(filepath.Join(outputDir, "web-components.js"), []byte(webComponentCode), 0644)
	_ = os.WriteFile(filepath.Join(outputDir, "react-components.jsx"), []byte(reactCode), 0644)
	_ = os.WriteFile(filepath.Join(outputDir, "ssr-dictionary.js"), []byte(ssrCode), 0644)

	fmt.Println("🚀 All 4 JS delivery targets successfully compiled into 'js_icons/'!")

}

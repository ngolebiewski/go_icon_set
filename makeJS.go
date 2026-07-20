package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// WriteJSPackage compiles your vectors into four distinct frontend flavors with embedded A11y pipelines
func WriteJSPackage(outputDir string, compiledPaths map[string]string, sortedKeys []string) {
	_ = os.MkdirAll(outputDir, 0755)

	// ==========================================
	// TARGET 1: STRING COMPONENTS
	// ==========================================
	stringComponentsCode := `let globalSize = "16px";
let globalColor = "white";
let globalClass = "";

export function setGlobalDefaults({ size, color, className } = {}) {
    if (size) globalSize = size;
    if (color) globalColor = color;
    if (className) globalClass = className;
}

function renderIcon(path, config = {}) {
    const size = config.size || globalSize;
    const color = config.color || globalColor;
    const className = config.className || config.class || globalClass;
    const alt = config.alt || "";

    const style = "width: " + size + "; height: " + size + "; color: " + color + "; display: inline-block; vertical-align: middle;";
    
    let attrs = "style=\"" + style + "\" focusable=\"false\"";
    if (className) {
        attrs += " class=\"" + className + "\"";
    }
    
    if (alt) {
        attrs += " role=\"img\" aria-label=\"" + alt + "\"";
    } else {
        attrs += " aria-hidden=\"true\"";
    }

    return "<svg xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 16 16\" shape-rendering=\"crispEdges\" fill=\"currentColor\" " + attrs + "><path d=\"" + path + "\" /></svg>";
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
        const className = this.getAttribute('class') || '';
        const alt = this.getAttribute('alt') || '';

        const svg = document.createElementNS('http://www.w3.org/2000/svg', 'svg');
        svg.setAttribute('viewBox', '0 0 16 16');
        svg.setAttribute('shape-rendering', 'crispEdges');
        svg.setAttribute('fill', 'currentColor');
        svg.setAttribute('focusable', 'false');
        
        if (className) svg.setAttribute('class', className);
        
        if (alt) {
            svg.setAttribute('role', 'img');
            svg.setAttribute('aria-label', alt);
        } else {
            svg.setAttribute('aria-hidden', 'true');
        }

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

function Icon({ path, size = "16px", color = "white", alt, className, ...props }) {
	const style = { width: size, height: size, color: color, display: "inline-block", verticalAlign: "middle" };
	
	const a11yProps = alt 
		? { role: "img", "aria-label": alt } 
		: { "aria-hidden": "true" };

	return (
		<svg 
			xmlns="http://www.w3.org/2000/svg" 
			viewBox="0 0 16 16" 
			shapeRendering="crispEdges" 
			fill="currentColor" 
			style={style} 
			className={className}
			focusable="false"
			{...a11yProps}
			{...props}
		>
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
	for _, name := range sortedKeys {
		path := compiledPaths[name]

		goName := toTitleCase(name)
		tagName := "icon-" + strings.ToLower(name)

		// 1. Original String Append
		stringComponentsCode += fmt.Sprintf("export const %s = (cfg) => renderIcon(\"%s\", cfg);\n", goName, path)

		// 2. Web Component Append
		webComponentCode += fmt.Sprintf("export class %s extends CrispIcon { static path = %q; }\ncustomElements.define(%q, %s);\n", goName, path, tagName, goName)

		// 3. React Append
		reactCode += fmt.Sprintf("export const %s = (props) => <Icon path=%q {...props} />;\n", goName, path)

		// 4. SSR Dictionary Append
		ssrCode += fmt.Sprintf("\t%q: %q,\n", name, path)
	}

	ssrCode += "};\n"

	// Commit everything cleanly to disk
	_ = os.WriteFile(filepath.Join(outputDir, "string-components.js"), []byte(stringComponentsCode), 0644)
	_ = os.WriteFile(filepath.Join(outputDir, "web-components.js"), []byte(webComponentCode), 0644)
	_ = os.WriteFile(filepath.Join(outputDir, "react-components.jsx"), []byte(reactCode), 0644)
	_ = os.WriteFile(filepath.Join(outputDir, "ssr-dictionary.js"), []byte(ssrCode), 0644)

	fmt.Println("🚀 All 4 JS delivery targets successfully compiled into 'js_icons/' with A11y defaults!")
}

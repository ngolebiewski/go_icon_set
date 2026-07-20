# Go Icon Set
Tool for making a Pixel Art web Icons set from a spritesheet (Asesprite).

*Reference*
Best icon set ever: Shigetaka Kurita
https://www.moma.org/collection/works/196070


## How it works
- Take a spritesheet.png, tell it the icon size 16x16, number of columns and go.
- Optional to make a json sheet with names for the icons. Otherwise defaults to numbered.
```json
	{
		"0": "hamburger",
		"1": "doc",
		"2": "sword",
		...
	}

```
- Optional to change the output folders
- `go run .`
- Auto generated a Templ/Go package and various JavaScript/TypeScript libraries to use in your frontend projects.


## Roadmap
- [x] Make a Templ + Go package
- [x] JavaScript/TypeScript package for React, Next, Nuxt, VanillaJS, etc.
- [ ] Test on Templ and JS options
- [ ] **Finish sample icon set**
- [ ] When finalized, version and add JS version to npm
---

![spritesheet with labels](icons/cheatsheet.png)


# Use in Templ/GO
## Step 1: Configure Your Global Style (Optional)
- If you don't want to specify colors and sizes every single time you call an icon, you can set a global fallback style right when your main web application starts up (usually in your website's main main.go).

```go
package main

import (
	"net/http"
	// Import your icons package
	"github.com/ngolebiewski/go_icon_set/templ_icons"
)

func main() {
	// Set your preferred global tokens ONCE at launch
	templ_icons.SetGlobalDefaults(templ_icons.IconConfig{
		Size:  "24px",          // All icons will render at 24px by default
		Color: "currentColor",  // Force icons to inherit text colors from CSS classes
	})

	// Start your server ...
}

```

--- 
## Step 2: Use the Icons in a .templ File

- Inside any Templ file, import the templ_icons sub-package. Because they are compiled into native templ.Component functions, you can invoke them simply using the @ syntax.

### Accessibility (A11y) by Default
- The icon set automatically manages screen-reader configurations under the hood depending on your UI layout:

1. Decorative Icons (Default): If an icon has accompanying descriptive text, it is automatically hidden from screen readers (aria-hidden="true") to prevent audio clutter.

2.  Semantic Icons: If you use an icon standalone (like inside an interactive menu or a link button), provide an Alt string text description. The generator will automatically swap the markup to expose a clear structural image description (role="img" aria-label="...") to assistive devices.

- Create or open a file like components/nav.templ:

```go
package components

import (
    // Import the sub-package path
    "github.com/ngolebiewski/go_icon_set/templ_icons"
)

templ Navigation() {
    <nav class="flex items-center gap-4 bg-slate-900 p-4 text-slate-300">
        
        <!-- 1. Decorative Default: Hidden from screen readers because "Home" text is next to it -->
        <span class="flex items-center gap-2">
            @templ_icons.Home()
            <span>Home</span>
        </span>

        <!-- 2. Styled Customization: Explicitly sized and colored via inline rules -->
        <span class="flex items-center gap-2">
            @templ_icons.Sword(templ_icons.IconConfig{Size: "2rem", Color: "#f43f5e"})
            <span>Combat</span>
        </span>

        <!-- 3. Semantic Icon Link: Because it has no descriptive text, we supply an 'Alt' tag 
             and append a styling 'Class' framework utility -->
        <a href="https://github.com" class="text-emerald-400 hover:text-emerald-300 transition-colors">
            @templ_icons.Github(templ_icons.IconConfig{
                Size:  "32px",
                Class: "p-1 cursor-pointer",
                Alt:   "Visit our project repository on GitHub",
            })
        </a>

    </nav>
}

```

---
# Use in JavaScript/TypeScript Frontends


-Any JavaScript or TypeScript project can import these components seamlessly.

-If you use innerHTML frameworks (like native web applications, Alpine.js, or vanilla JS components), you call them directly. If you use a framework like Vue, Nuxt, React, or Svelte, you can inject them securely using their respective dynamic HTML utility bindings (such as v-html in Vue/Nuxt or dangerouslySetInnerHTML in React).

```JavaScript

import { setGlobalDefaults, Hamburger, Sword } from './js_icons/icons.js';

// Configure everything once at entry point boot
setGlobalDefaults({
    size: "24px",
    color: "currentColor"
});

// Render dynamic raw string components instantly without network delay overhead:
const navbarHTML = `
  <nav class="nav-bar text-slate-200">
     ${Hamburger()} ${Sword({ color: '#f43f5e' })} </nav>
`;
```


# JS Icon Components Integration Guide

This directory contains your auto-generated cross-platform icon distribution bundles. Every file is completely self-contained and embeds raw vector data directly to eliminate layout shift flicker and network roundtrips.

---

## Accessibility (A11y) Architecture

Every JavaScript flavor manages screen-reader configurations out of the box based on how properties are passed:
- **Decorative Mode (Default):** If an icon has accompanying descriptive text or context labels, omit the `alt` property. The runtime automatically hides the vector from screen readers (`aria-hidden="true" focusable="false"`) to prevent auditory clutter.
- **Semantic Mode:** If an icon is used as a standalone actionable item (e.g., inside an icon-only button, link, or status banner), pass an explicit `alt` text description string. The runtime automatically updates the markup structure to reflect a screen-reader-safe element (`role="img" aria-label="..." focusable="false"`).

---

## 1. String Components ("string-components.js")
**Best For:** Vanilla JavaScript, Alpine.js, HTMX, or lightweight projects. Returns raw HTML strings.

### Usage
```javascript
import { setGlobalDefaults, Home, Sword } from './js_icons/string-components.js';

// Optional: Set global defaults once at application boot
setGlobalDefaults({
    size: "24px",
    color: "currentColor", // Inherits color from parent text classes
    className: "pixel-art"
});

// Render anywhere strings are parsed into HTML
document.getElementById('app').innerHTML = `
    <nav class="nav-bar">
        <!-- Decorative Default: Hidden from screen readers -->
        <span class="item">
            ${Home()} <span>Home</span>
        </span>

        <!-- Semantic Override: Accessible standalone button with layout styling rules -->
        <button class="action-btn">
            ${Sword({ 
                size: "32px", 
                color: "#f43f5e", 
                className: "hover:scale-105", 
                alt: "Equip ruby blade steel weapon" 
            })}
        </button>
    </nav>
`;
```

---

## 2. Native Web Components ("web-components.js")
**Best For:** Modern Single-Page Apps (Vue, Svelte, Vite) or Vanilla apps seeking 100% safe DOM manipulation without raw string injection.

### Usage
Import this file **exactly once** at the absolute root entry point of your frontend architecture (e.g., `main.js` or `index.js`). 

```javascript
import './js_icons/web-components.js'; 
// That's it! Elements register globally with the browser engine.
```

Now, use the custom tags anywhere in your templates without importing them ever again. Pass an `alt="..."` text payload string to trigger accessible semantic mode:

```html
<!-- Decorative Default: Skip screen reader tracking because adjacent label exists -->
<div class="menu-item">
    <icon-home size="24px" color="white"></icon-home>
    <span>Home</span>
</div>

<!-- Semantic Standalone Item: Automatically converts internally to role="img" -->
<button class="icon-only-btn">
    <icon-sword size="32px" color="#f43f5e" alt="Equip combat weapon sword"></icon-sword>
</button>
```

---

## 3. React Components ("react-components.jsx")
**Best For:** React.js, Next.js (Client Components), Vite + React, Remix, or Astro projects.

### Usage
These are compiled as standard functional Virtual DOM components. They are fully compatible with bundler tree-shaking and accept standard JSX property spreads.

```jsx
import React from 'react';
import { Home, Sword } from './js_icons/react-components.jsx';

export function Navigation() {
    return (
        <nav className="flex gap-4 bg-slate-900 p-4">
            {/* 1. Decorative Default: Omit alt because textual context is next to it */}
            <div className="flex items-center gap-2 text-white">
                <Home size="24px" color="white" />
                <span>Home</span>
            </div>

            {/* 2. Semantic Interactive Item: Providing an 'alt' string automatically 
                 strips aria-hidden and applies role="img" + aria-label */}
            <button className="toolbar-btn">
                <Sword 
                    size="2rem" 
                    color="#f43f5e" 
                    className="hover:scale-110 transition-transform"
                    alt="Equip secondary melee combat sword weapon"
                    onClick={() => console.log("Equipped!")} 
                />
            </button>
        </nav>
    );
}
```

---

## 4. SSR Pure Dictionary ("ssr-dictionary.js")
**Best For:** Server-Side Frameworks (Nuxt, SvelteKit, Astro) that need to securely draw vectors on the server before shipping HTML down to the browser.

### Usage
Because this format is just a flat key-value string database, your meta-framework templates map or loop through keys directly to construct server-painted inline SVGs while controlling accessibility flags natively.

#### Example Component in Vue / Nuxt:
```typescript
<script setup>
import { IconVectors } from './js_icons/ssr-dictionary.js';

const props = defineProps({
    name: String, // e.g., 'home' or 'sword'
    size: { type: String, default: '16px' },
    color: { type: String, default: 'white' },
    alt: { type: String, default: '' } // Empty = Decorative, String = Semantic Image
});

const pathData = computed(() => IconVectors[props.name]);
</script>

<template>
  <svg 
    xmlns="http://www.w3.org/2000/svg" 
    viewBox="0 0 16 16" 
    shape-rendering="crispEdges" 
    fill="currentColor"
    focusable="false"
    :aria-hidden="!alt ? 'true' : undefined"
    :role="alt ? 'img' : undefined"
    :aria-label="alt || undefined"
    :style="{ width: size, height: size, color: color, display: 'inline-block', verticalAlign: 'middle' }"
  >
    <path :d="pathData" />
  </svg>
</template>
```

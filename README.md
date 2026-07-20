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

## Step 2: Use the Icons in a .templ File

-Inside any Templ file, import the templ_icons sub-package. Because they are compiled into native templ.Component functions, you can invoke them simply using the @ syntax.

-Create or open a file like components/nav.templ:

```go
package components

import (
	// Import the sub-package path
	"github.com/ngolebiewski/go_icon_set/templ_icons"
)

templ Navigation() {
	<nav class="flex items-center gap-4 bg-slate-900 p-4 text-slate-300">
		
		@templ_icons.Home()
		<span>Home</span>

		@templ_icons.Sword(templ_icons.IconConfig{Size: "2rem", Color: "#f43f5e"})
		<span>Combat</span>

		<div class="text-emerald-400 hover:text-emerald-300 transition-colors">
			@templ_icons.Github(templ_icons.IconConfig{Size: "32px"})
		</div>

	</nav>
}

```

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

## 1. String Components ("string-components.js")
**Best For:** Vanilla JavaScript, Alpine.js, HTMX, or lightweight projects. Returns raw HTML strings.

### Usage
```javascript
import { setGlobalDefaults, Home, Sword } from './js_icons/string-components.js';

// Optional: Set global styles once at application boot
setGlobalDefaults({
	size: "24px",
	color: "currentColor" // Inherits color from parent text classes
});

// Render anywhere strings are parsed into HTML
document.getElementById('app').innerHTML = \`
	<div class="nav-bar">
		\${Home()} \${Sword({ color: '#f43f5e', size: '32px' })} </div>
;
```

---

## 2. Native Web Components ("web-components.js")
**Best For:** Modern Single-Page Apps (Vue, Svelte, Vite) or Vanilla apps seeking 100% safe DOM manipulation without raw string injection.

### Usage
Import this file **exactly once** at the absolute root entry point of your frontend architecture (e.g., "main.js" or "index.js"). 

```javascript
import './js_icons/web-components.js'; 
// That's it! Elements register globally with the browser engine.
```

Now, use the custom tags anywhere in your templates without importing them ever again:

```html
<icon-home size="24px" color="white"></icon-home>
<icon-sword size="32px" color="#f43f5e"></icon-sword>
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
			{/* Standard icon */}
			<Home size="24px" color="white" />

			{/* Custom attributes pass through safely via props spread */}
			<Sword 
				size="2rem" 
				color="#f43f5e" 
				className="hover:scale-110 transition-transform"
				onClick={() => console.log("Equipped!")} 
			/>
		</nav>
	);
}
```

---

## 4. SSR Pure Dictionary ("ssr-dictionary.js")
**Best For:** Server-Side Frameworks (Nuxt, SvelteKit, Astro) that need to securely draw vectors on the server before shipping HTML down to the browser.

### Usage
Because this format is just a flat key-value data string database, your meta-framework templates loop or reference the keys directly to build server-painted inline SVGs.

#### Example in Vue / Nuxt:
```typescript
<script setup>
import { IconVectors } from './js_icons/ssr-dictionary.js';

const props = defineProps({
	name: String, // e.g., 'home' or 'sword'
	size: { type: String, default: '16px' },
	color: { type: String, default: 'white' }
});

const pathData = computed(() => IconVectors[props.name]);
</script>

<template>
  <svg 
	xmlns="http://www.w3.org/2000/svg" 
	viewBox="0 0 16 16" 
	shape-rendering="crispEdges" 
	fill="currentColor"
	:style="{ width: size, height: size, color: color, display: 'inline-block', verticalAlign: 'middle' }"
  >
	<path :d="pathData" />
  </svg>
</template>
```
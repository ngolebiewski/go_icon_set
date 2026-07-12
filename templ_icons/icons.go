package templ_icons

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
	cfg := currentDefault 
	if len(cfgs) > 0 {
		if cfgs[0].Size != ""  { cfg.Size = cfgs[0].Size }
		if cfgs[0].Color != "" { cfg.Color = cfgs[0].Color }
	}
	return fmt.Sprintf("style=\"width: %s; height: %s; color: %s; display: inline-block; vertical-align: middle;\"", cfg.Size, cfg.Size, cfg.Color)
}

// IntersectionFilled renders the sharp vector icon for "intersection-filled"
func IntersectionFilled(cfg ...IconConfig) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		style := buildStyleAttr(cfg)
		_, err := io.WriteString(w, fmt.Sprintf("<svg xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 16 16\" shape-rendering=\"crispEdges\" fill=\"currentColor\" %s><path d=\"M3,0h1v1h-1z M12,0h1v1h-1z M3,4h4v1h-4z M8,4h4v1h-4z M2,5h1v1h-1z M7,5h1v1h-1z M12,5h1v1h-1z M1,6h1v1h-1z M6,6h3v1h-3z M13,6h1v1h-1z M0,7h1v1h-1z M5,7h5v1h-5z M14,7h1v1h-1z M0,8h1v1h-1z M5,8h5v1h-5z M14,8h1v1h-1z M0,9h1v1h-1z M5,9h5v1h-5z M14,9h1v1h-1z M0,10h1v1h-1z M5,10h5v1h-5z M14,10h1v1h-1z M1,11h1v1h-1z M6,11h3v1h-3z M13,11h1v1h-1z M2,12h1v1h-1z M7,12h1v1h-1z M12,12h1v1h-1z M3,13h4v1h-4z M8,13h4v1h-4z\" /></svg>", style))
		return err
	})
}

// Sun2 renders the sharp vector icon for "sun-2"
func Sun2(cfg ...IconConfig) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		style := buildStyleAttr(cfg)
		_, err := io.WriteString(w, fmt.Sprintf("<svg xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 16 16\" shape-rendering=\"crispEdges\" fill=\"currentColor\" %s><path d=\"M4,1h1v1h-1z M8,1h1v1h-1z M11,1h1v1h-1z M5,2h1v1h-1z M8,2h1v1h-1z M10,2h1v1h-1z M2,3h1v1h-1z M6,3h4v1h-4z M13,3h1v1h-1z M3,4h1v1h-1z M5,4h6v1h-6z M12,4h1v1h-1z M4,5h8v1h-8z M3,6h10v1h-10z M3,7h10v1h-10z M1,8h14v1h-14z M3,9h10v1h-10z M3,10h10v1h-10z M2,11h1v1h-1z M4,11h8v1h-8z M13,11h1v1h-1z M1,12h1v1h-1z M5,12h6v1h-6z M14,12h1v1h-1z M4,13h1v1h-1z M6,13h4v1h-4z M11,13h1v1h-1z M3,14h1v1h-1z M8,14h1v1h-1z M12,14h1v1h-1z M8,15h1v1h-1z\" /></svg>", style))
		return err
	})
}

// Raster renders the sharp vector icon for "raster"
func Raster(cfg ...IconConfig) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		style := buildStyleAttr(cfg)
		_, err := io.WriteString(w, fmt.Sprintf("<svg xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 16 16\" shape-rendering=\"crispEdges\" fill=\"currentColor\" %s><path d=\"M0,4h1v1h-1z M2,4h1v1h-1z M4,4h1v1h-1z M6,4h1v1h-1z M8,4h1v1h-1z M10,4h1v1h-1z M12,4h1v1h-1z M14,4h1v1h-1z M1,5h1v1h-1z M5,5h1v1h-1z M9,5h1v1h-1z M13,5h1v1h-1z M0,6h1v1h-1z M2,6h1v1h-1z M4,6h1v1h-1z M6,6h1v1h-1z M8,6h1v1h-1z M10,6h1v1h-1z M12,6h1v1h-1z M14,6h1v1h-1z M1,7h1v1h-1z M3,7h1v1h-1z M5,7h1v1h-1z M7,7h1v1h-1z M9,7h1v1h-1z M11,7h1v1h-1z M13,7h1v1h-1z M15,7h1v1h-1z M0,8h3v1h-3z M4,8h3v1h-3z M8,8h3v1h-3z M12,8h3v1h-3z M1,9h1v1h-1z M3,9h1v1h-1z M5,9h1v1h-1z M7,9h1v1h-1z M9,9h1v1h-1z M11,9h1v1h-1z M13,9h1v1h-1z M15,9h1v1h-1z M0,10h16v1h-16z M1,11h1v1h-1z M3,11h1v1h-1z M5,11h1v1h-1z M7,11h1v1h-1z M9,11h1v1h-1z M11,11h1v1h-1z M13,11h1v1h-1z M15,11h1v1h-1z M0,12h16v1h-16z M0,13h16v1h-16z M0,14h16v1h-16z M0,15h16v1h-16z\" /></svg>", style))
		return err
	})
}

// Scribble renders the sharp vector icon for "scribble"
func Scribble(cfg ...IconConfig) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		style := buildStyleAttr(cfg)
		_, err := io.WriteString(w, fmt.Sprintf("<svg xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 16 16\" shape-rendering=\"crispEdges\" fill=\"currentColor\" %s><path d=\"M5,3h3v1h-3z M3,4h3v1h-3z M7,4h3v1h-3z M3,5h1v1h-1z M9,5h1v1h-1z M2,6h2v1h-2z M9,6h1v1h-1z M2,7h1v1h-1z M9,7h1v1h-1z M2,8h1v1h-1z M9,8h1v1h-1z M15,8h1v1h-1z M1,9h2v1h-2z M5,9h7v1h-7z M15,9h1v1h-1z M1,10h1v1h-1z M6,10h1v1h-1z M9,10h1v1h-1z M11,10h2v1h-2z M15,10h1v1h-1z M1,11h2v1h-2z M9,11h1v1h-1z M12,11h1v1h-1z M14,11h2v1h-2z M2,12h1v1h-1z M9,12h1v1h-1z M12,12h2v1h-2z M1,13h4v1h-4z M9,13h1v1h-1z M11,13h2v1h-2z M5,14h8v1h-8z M8,15h1v1h-1z\" /></svg>", style))
		return err
	})
}

// Globe renders the sharp vector icon for "globe"
func Globe(cfg ...IconConfig) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		style := buildStyleAttr(cfg)
		_, err := io.WriteString(w, fmt.Sprintf("<svg xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 16 16\" shape-rendering=\"crispEdges\" fill=\"currentColor\" %s><path d=\"M6,0h4v1h-4z M4,1h2v1h-2z M10,1h2v1h-2z M3,2h1v1h-1z M9,2h1v1h-1z M12,2h1v1h-1z M2,3h5v1h-5z M8,3h1v1h-1z M11,3h1v1h-1z M13,3h1v1h-1z M1,4h1v1h-1z M3,4h1v1h-1z M5,4h4v1h-4z M12,4h1v1h-1z M14,4h1v1h-1z M1,5h1v1h-1z M4,5h5v1h-5z M12,5h1v1h-1z M14,5h1v1h-1z M0,6h1v1h-1z M4,6h1v1h-1z M6,6h1v1h-1z M15,6h1v1h-1z M0,7h1v1h-1z M5,7h2v1h-2z M15,7h1v1h-1z M0,8h1v1h-1z M5,8h1v1h-1z M13,8h3v1h-3z M0,9h2v1h-2z M5,9h4v1h-4z M13,9h1v1h-1z M15,9h1v1h-1z M1,10h1v1h-1z M6,10h2v1h-2z M13,10h2v1h-2z M1,11h1v1h-1z M5,11h4v1h-4z M13,11h2v1h-2z M2,12h1v1h-1z M6,12h3v1h-3z M11,12h1v1h-1z M13,12h2v1h-2z M3,13h1v1h-1z M8,13h1v1h-1z M12,13h1v1h-1z M4,14h8v1h-8z M6,15h4v1h-4z\" /></svg>", style))
		return err
	})
}

// HeartEmpty renders the sharp vector icon for "heart-empty"
func HeartEmpty(cfg ...IconConfig) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		style := buildStyleAttr(cfg)
		_, err := io.WriteString(w, fmt.Sprintf("<svg xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 16 16\" shape-rendering=\"crispEdges\" fill=\"currentColor\" %s><path d=\"M4,4h2v1h-2z M10,4h2v1h-2z M3,5h1v1h-1z M6,5h1v1h-1z M9,5h1v1h-1z M12,5h1v1h-1z M3,6h1v1h-1z M7,6h2v1h-2z M12,6h1v1h-1z M4,7h1v1h-1z M11,7h1v1h-1z M4,8h1v1h-1z M10,8h2v1h-2z M5,9h1v1h-1z M10,9h1v1h-1z M5,10h2v1h-2z M9,10h2v1h-2z M7,11h2v1h-2z M7,12h2v1h-2z\" /></svg>", style))
		return err
	})
}

// Pdf renders the sharp vector icon for "pdf"
func Pdf(cfg ...IconConfig) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		style := buildStyleAttr(cfg)
		_, err := io.WriteString(w, fmt.Sprintf("<svg xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 16 16\" shape-rendering=\"crispEdges\" fill=\"currentColor\" %s><path d=\"M6,2h7v1h-7z M4,3h1v1h-1z M6,3h7v1h-7z M3,4h2v1h-2z M6,4h7v1h-7z M6,5h7v1h-7z M3,6h10v1h-10z M3,7h10v1h-10z M3,8h10v1h-10z M6,9h2v1h-2z M9,9h2v1h-2z M4,10h1v1h-1z M6,10h2v1h-2z M9,10h1v1h-1z M11,10h2v1h-2z M6,11h2v1h-2z M9,11h1v1h-1z M12,11h1v1h-1z M4,12h2v1h-2z M9,12h1v1h-1z M11,12h2v1h-2z M3,13h3v1h-3z M7,13h1v1h-1z M9,13h1v1h-1z M11,13h2v1h-2z\" /></svg>", style))
		return err
	})
}

// Facebook renders the sharp vector icon for "facebook"
func Facebook(cfg ...IconConfig) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		style := buildStyleAttr(cfg)
		_, err := io.WriteString(w, fmt.Sprintf("<svg xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 16 16\" shape-rendering=\"crispEdges\" fill=\"currentColor\" %s><path d=\"M5,2h5v1h-5z M4,3h7v1h-7z M3,4h9v1h-9z M2,5h6v1h-6z M10,5h3v1h-3z M2,6h5v1h-5z M8,6h5v1h-5z M2,7h5v1h-5z M8,7h5v1h-5z M2,8h4v1h-4z M9,8h4v1h-4z M2,9h5v1h-5z M8,9h5v1h-5z M3,10h4v1h-4z M8,10h4v1h-4z M4,11h3v1h-3z M8,11h3v1h-3z M5,12h2v1h-2z M8,12h2v1h-2z\" /></svg>", style))
		return err
	})
}

// DownTriangle renders the sharp vector icon for "down-triangle"
func DownTriangle(cfg ...IconConfig) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		style := buildStyleAttr(cfg)
		_, err := io.WriteString(w, fmt.Sprintf("<svg xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 16 16\" shape-rendering=\"crispEdges\" fill=\"currentColor\" %s><path d=\"M2,5h13v1h-13z M3,6h11v1h-11z M4,7h9v1h-9z M5,8h7v1h-7z M6,9h5v1h-5z M7,10h3v1h-3z M8,11h1v1h-1z\" /></svg>", style))
		return err
	})
}

// Mail renders the sharp vector icon for "mail"
func Mail(cfg ...IconConfig) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		style := buildStyleAttr(cfg)
		_, err := io.WriteString(w, fmt.Sprintf("<svg xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 16 16\" shape-rendering=\"crispEdges\" fill=\"currentColor\" %s><path d=\"M0,2h16v1h-16z M0,3h3v1h-3z M13,3h3v1h-3z M0,4h1v1h-1z M3,4h2v1h-2z M11,4h2v1h-2z M15,4h1v1h-1z M0,5h1v1h-1z M5,5h2v1h-2z M9,5h2v1h-2z M15,5h1v1h-1z M0,6h1v1h-1z M7,6h2v1h-2z M15,6h1v1h-1z M0,7h1v1h-1z M15,7h1v1h-1z M0,8h1v1h-1z M15,8h1v1h-1z M0,9h1v1h-1z M15,9h1v1h-1z M0,10h16v1h-16z\" /></svg>", style))
		return err
	})
}

// Coin renders the sharp vector icon for "coin"
func Coin(cfg ...IconConfig) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		style := buildStyleAttr(cfg)
		_, err := io.WriteString(w, fmt.Sprintf("<svg xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 16 16\" shape-rendering=\"crispEdges\" fill=\"currentColor\" %s><path d=\"M6,3h4v1h-4z M4,4h8v1h-8z M3,5h2v1h-2z M6,5h2v1h-2z M9,5h4v1h-4z M3,6h1v1h-1z M5,6h2v1h-2z M11,6h2v1h-2z M2,7h4v1h-4z M7,7h1v1h-1z M9,7h5v1h-5z M2,8h1v1h-1z M4,8h3v1h-3z M9,8h3v1h-3z M13,8h1v1h-1z M2,9h6v1h-6z M10,9h2v1h-2z M13,9h1v1h-1z M2,10h6v1h-6z M9,10h1v1h-1z M11,10h3v1h-3z M3,11h1v1h-1z M5,11h1v1h-1z M10,11h3v1h-3z M3,12h5v1h-5z M9,12h4v1h-4z M5,13h1v1h-1z M7,13h5v1h-5z M6,14h4v1h-4z\" /></svg>", style))
		return err
	})
}

// Home renders the sharp vector icon for "home"
func Home(cfg ...IconConfig) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		style := buildStyleAttr(cfg)
		_, err := io.WriteString(w, fmt.Sprintf("<svg xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 16 16\" shape-rendering=\"crispEdges\" fill=\"currentColor\" %s><path d=\"M7,1h1v1h-1z M6,2h3v1h-3z M5,3h5v1h-5z M4,4h7v1h-7z M3,5h9v1h-9z M2,6h11v1h-11z M1,7h3v1h-3z M11,7h3v1h-3z M3,8h1v1h-1z M11,8h1v1h-1z M3,9h1v1h-1z M11,9h1v1h-1z M3,10h1v1h-1z M6,10h3v1h-3z M11,10h1v1h-1z M3,11h1v1h-1z M6,11h3v1h-3z M11,11h1v1h-1z M3,12h1v1h-1z M6,12h3v1h-3z M11,12h1v1h-1z M3,13h1v1h-1z M6,13h3v1h-3z M11,13h1v1h-1z M3,14h9v1h-9z\" /></svg>", style))
		return err
	})
}

// Moon renders the sharp vector icon for "moon"
func Moon(cfg ...IconConfig) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		style := buildStyleAttr(cfg)
		_, err := io.WriteString(w, fmt.Sprintf("<svg xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 16 16\" shape-rendering=\"crispEdges\" fill=\"currentColor\" %s><path d=\"M8,4h3v1h-3z M10,5h2v1h-2z M11,6h2v1h-2z M12,7h2v1h-2z M11,8h3v1h-3z M11,9h3v1h-3z M10,10h4v1h-4z M5,11h2v1h-2z M9,11h4v1h-4z M6,12h6v1h-6z M7,13h4v1h-4z\" /></svg>", style))
		return err
	})
}

// QuestionMark renders the sharp vector icon for "question-mark"
func QuestionMark(cfg ...IconConfig) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		style := buildStyleAttr(cfg)
		_, err := io.WriteString(w, fmt.Sprintf("<svg xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 16 16\" shape-rendering=\"crispEdges\" fill=\"currentColor\" %s><path d=\"M5,1h6v1h-6z M5,2h6v1h-6z M5,3h6v1h-6z M2,4h3v1h-3z M11,4h3v1h-3z M2,5h3v1h-3z M11,5h3v1h-3z M2,6h3v1h-3z M11,6h3v1h-3z M8,7h3v1h-3z M8,8h3v1h-3z M8,9h3v1h-3z M8,12h3v1h-3z M8,13h3v1h-3z M8,14h3v1h-3z\" /></svg>", style))
		return err
	})
}

// In renders the sharp vector icon for "in"
func In(cfg ...IconConfig) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		style := buildStyleAttr(cfg)
		_, err := io.WriteString(w, fmt.Sprintf("<svg xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 16 16\" shape-rendering=\"crispEdges\" fill=\"currentColor\" %s><path d=\"M2,1h5v1h-5z M2,2h1v1h-1z M2,3h1v1h-1z M2,4h1v1h-1z M2,5h1v1h-1z M9,5h1v1h-1z M2,6h1v1h-1z M8,6h1v1h-1z M2,7h1v1h-1z M7,7h6v1h-6z M2,8h1v1h-1z M8,8h1v1h-1z M2,9h1v1h-1z M9,9h1v1h-1z M2,10h1v1h-1z M2,11h1v1h-1z M2,12h1v1h-1z M2,13h5v1h-5z\" /></svg>", style))
		return err
	})
}

// CheckboxUnchecked renders the sharp vector icon for "checkbox-unchecked"
func CheckboxUnchecked(cfg ...IconConfig) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		style := buildStyleAttr(cfg)
		_, err := io.WriteString(w, fmt.Sprintf("<svg xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 16 16\" shape-rendering=\"crispEdges\" fill=\"currentColor\" %s><path d=\"M4,3h9v1h-9z M3,4h1v1h-1z M12,4h1v1h-1z M3,5h1v1h-1z M12,5h1v1h-1z M3,6h1v1h-1z M12,6h1v1h-1z M3,7h1v1h-1z M12,7h1v1h-1z M3,8h1v1h-1z M12,8h1v1h-1z M3,9h1v1h-1z M12,9h1v1h-1z M3,10h1v1h-1z M12,10h1v1h-1z M3,11h1v1h-1z M12,11h1v1h-1z M3,12h9v1h-9z\" /></svg>", style))
		return err
	})
}

// CheckboxChecked renders the sharp vector icon for "checkbox-checked"
func CheckboxChecked(cfg ...IconConfig) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		style := buildStyleAttr(cfg)
		_, err := io.WriteString(w, fmt.Sprintf("<svg xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 16 16\" shape-rendering=\"crispEdges\" fill=\"currentColor\" %s><path d=\"M4,3h9v1h-9z M3,4h1v1h-1z M12,4h1v1h-1z M3,5h1v1h-1z M12,5h1v1h-1z M3,6h1v1h-1z M10,6h1v1h-1z M12,6h1v1h-1z M3,7h1v1h-1z M9,7h1v1h-1z M12,7h1v1h-1z M3,8h1v1h-1z M8,8h1v1h-1z M12,8h1v1h-1z M3,9h1v1h-1z M5,9h1v1h-1z M7,9h1v1h-1z M12,9h1v1h-1z M3,10h1v1h-1z M6,10h1v1h-1z M12,10h1v1h-1z M3,11h1v1h-1z M12,11h1v1h-1z M3,12h9v1h-9z\" /></svg>", style))
		return err
	})
}

// Building renders the sharp vector icon for "building"
func Building(cfg ...IconConfig) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		style := buildStyleAttr(cfg)
		_, err := io.WriteString(w, fmt.Sprintf("<svg xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 16 16\" shape-rendering=\"crispEdges\" fill=\"currentColor\" %s><path d=\"M4,0h9v1h-9z M4,1h1v1h-1z M6,1h1v1h-1z M8,1h1v1h-1z M10,1h1v1h-1z M12,1h1v1h-1z M4,2h9v1h-9z M4,3h1v1h-1z M6,3h1v1h-1z M8,3h1v1h-1z M10,3h1v1h-1z M12,3h1v1h-1z M4,4h9v1h-9z M4,5h1v1h-1z M6,5h1v1h-1z M8,5h1v1h-1z M10,5h1v1h-1z M12,5h1v1h-1z M4,6h9v1h-9z M4,7h1v1h-1z M6,7h1v1h-1z M8,7h1v1h-1z M10,7h1v1h-1z M12,7h1v1h-1z M4,8h9v1h-9z M4,9h1v1h-1z M6,9h1v1h-1z M8,9h1v1h-1z M10,9h1v1h-1z M12,9h1v1h-1z M4,10h9v1h-9z M4,11h1v1h-1z M6,11h1v1h-1z M8,11h1v1h-1z M10,11h1v1h-1z M12,11h1v1h-1z M4,12h9v1h-9z M4,13h1v1h-1z M6,13h1v1h-1z M9,13h1v1h-1z M12,13h1v1h-1z M4,14h1v1h-1z M6,14h1v1h-1z M9,14h1v1h-1z M12,14h1v1h-1z M4,15h1v1h-1z M6,15h1v1h-1z M9,15h1v1h-1z M12,15h1v1h-1z\" /></svg>", style))
		return err
	})
}

// Profile renders the sharp vector icon for "profile"
func Profile(cfg ...IconConfig) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		style := buildStyleAttr(cfg)
		_, err := io.WriteString(w, fmt.Sprintf("<svg xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 16 16\" shape-rendering=\"crispEdges\" fill=\"currentColor\" %s><path d=\"M6,0h4v1h-4z M4,1h2v1h-2z M10,1h2v1h-2z M3,2h1v1h-1z M12,2h1v1h-1z M2,3h1v1h-1z M7,3h2v1h-2z M13,3h1v1h-1z M1,4h1v1h-1z M6,4h4v1h-4z M14,4h1v1h-1z M1,5h1v1h-1z M5,5h6v1h-6z M14,5h1v1h-1z M0,6h1v1h-1z M5,6h6v1h-6z M15,6h1v1h-1z M0,7h1v1h-1z M6,7h4v1h-4z M15,7h1v1h-1z M0,8h1v1h-1z M7,8h2v1h-2z M15,8h1v1h-1z M0,9h1v1h-1z M15,9h1v1h-1z M1,10h1v1h-1z M5,10h6v1h-6z M14,10h1v1h-1z M1,11h1v1h-1z M3,11h10v1h-10z M14,11h1v1h-1z M2,12h12v1h-12z M3,13h10v1h-10z M4,14h8v1h-8z M6,15h4v1h-4z\" /></svg>", style))
		return err
	})
}

// Camera renders the sharp vector icon for "camera"
func Camera(cfg ...IconConfig) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		style := buildStyleAttr(cfg)
		_, err := io.WriteString(w, fmt.Sprintf("<svg xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 16 16\" shape-rendering=\"crispEdges\" fill=\"currentColor\" %s><path d=\"M1,5h14v1h-14z M1,6h11v1h-11z M14,6h1v1h-1z M1,7h11v1h-11z M14,7h1v1h-1z M1,8h5v1h-5z M9,8h6v1h-6z M1,9h4v1h-4z M10,9h5v1h-5z M1,10h3v1h-3z M11,10h4v1h-4z M1,11h3v1h-3z M11,11h4v1h-4z M1,12h4v1h-4z M10,12h5v1h-5z M1,13h5v1h-5z M9,13h6v1h-6z M1,14h14v1h-14z\" /></svg>", style))
		return err
	})
}

// PencilVertical renders the sharp vector icon for "pencil-vertical"
func PencilVertical(cfg ...IconConfig) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		style := buildStyleAttr(cfg)
		_, err := io.WriteString(w, fmt.Sprintf("<svg xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 16 16\" shape-rendering=\"crispEdges\" fill=\"currentColor\" %s><path d=\"M6,0h3v1h-3z M6,1h3v1h-3z M6,3h3v1h-3z M6,4h3v1h-3z M6,5h3v1h-3z M6,6h3v1h-3z M6,7h3v1h-3z M6,8h3v1h-3z M6,9h3v1h-3z M6,10h3v1h-3z M6,11h3v1h-3z M6,12h3v1h-3z M7,13h1v1h-1z M7,14h1v1h-1z\" /></svg>", style))
		return err
	})
}

// NirvanaFace renders the sharp vector icon for "nirvana-face"
func NirvanaFace(cfg ...IconConfig) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		style := buildStyleAttr(cfg)
		_, err := io.WriteString(w, fmt.Sprintf("<svg xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 16 16\" shape-rendering=\"crispEdges\" fill=\"currentColor\" %s><path d=\"M6,1h5v1h-5z M4,2h2v1h-2z M11,2h2v1h-2z M3,3h1v1h-1z M13,3h1v1h-1z M2,4h1v1h-1z M14,4h1v1h-1z M2,5h1v1h-1z M5,5h1v1h-1z M7,5h1v1h-1z M10,5h1v1h-1z M12,5h1v1h-1z M14,5h1v1h-1z M1,6h1v1h-1z M6,6h1v1h-1z M11,6h1v1h-1z M15,6h1v1h-1z M1,7h1v1h-1z M5,7h1v1h-1z M7,7h1v1h-1z M10,7h1v1h-1z M12,7h1v1h-1z M15,7h1v1h-1z M1,8h1v1h-1z M15,8h1v1h-1z M1,9h1v1h-1z M15,9h1v1h-1z M1,10h1v1h-1z M15,10h1v1h-1z M2,11h1v1h-1z M6,11h1v1h-1z M11,11h1v1h-1z M14,11h1v1h-1z M2,12h1v1h-1z M7,12h4v1h-4z M14,12h1v1h-1z M3,13h1v1h-1z M13,13h1v1h-1z M4,14h2v1h-2z M11,14h2v1h-2z M6,15h5v1h-5z\" /></svg>", style))
		return err
	})
}

// UpTriangle renders the sharp vector icon for "up-triangle"
func UpTriangle(cfg ...IconConfig) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		style := buildStyleAttr(cfg)
		_, err := io.WriteString(w, fmt.Sprintf("<svg xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 16 16\" shape-rendering=\"crispEdges\" fill=\"currentColor\" %s><path d=\"M8,5h1v1h-1z M7,6h3v1h-3z M6,7h5v1h-5z M5,8h7v1h-7z M4,9h9v1h-9z M3,10h11v1h-11z M2,11h13v1h-13z\" /></svg>", style))
		return err
	})
}

// Calendar renders the sharp vector icon for "calendar"
func Calendar(cfg ...IconConfig) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		style := buildStyleAttr(cfg)
		_, err := io.WriteString(w, fmt.Sprintf("<svg xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 16 16\" shape-rendering=\"crispEdges\" fill=\"currentColor\" %s><path d=\"M6,2h7v1h-7z M4,3h1v1h-1z M6,3h7v1h-7z M3,4h2v1h-2z M6,4h7v1h-7z M6,5h7v1h-7z M3,6h10v1h-10z M3,7h10v1h-10z M3,8h10v1h-10z M3,9h5v1h-5z M9,9h1v1h-1z M12,9h1v1h-1z M3,10h4v1h-4z M9,10h3v1h-3z M3,11h5v1h-5z M9,11h2v1h-2z M3,12h5v1h-5z M9,12h1v1h-1z M11,12h2v1h-2z M3,13h5v1h-5z M9,13h1v1h-1z\" /></svg>", style))
		return err
	})
}

// Out renders the sharp vector icon for "out"
func Out(cfg ...IconConfig) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		style := buildStyleAttr(cfg)
		_, err := io.WriteString(w, fmt.Sprintf("<svg xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 16 16\" shape-rendering=\"crispEdges\" fill=\"currentColor\" %s><path d=\"M2,1h5v1h-5z M2,2h1v1h-1z M2,3h1v1h-1z M2,4h1v1h-1z M2,5h1v1h-1z M8,5h1v1h-1z M2,6h1v1h-1z M9,6h1v1h-1z M2,7h1v1h-1z M5,7h6v1h-6z M2,8h1v1h-1z M9,8h1v1h-1z M2,9h1v1h-1z M8,9h1v1h-1z M2,10h1v1h-1z M2,11h1v1h-1z M2,12h1v1h-1z M2,13h5v1h-5z\" /></svg>", style))
		return err
	})
}

// HeartFilled renders the sharp vector icon for "heart-filled"
func HeartFilled(cfg ...IconConfig) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		style := buildStyleAttr(cfg)
		_, err := io.WriteString(w, fmt.Sprintf("<svg xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 16 16\" shape-rendering=\"crispEdges\" fill=\"currentColor\" %s><path d=\"M3,4h2v1h-2z M9,4h2v1h-2z M2,5h4v1h-4z M8,5h4v1h-4z M2,6h10v1h-10z M3,7h8v1h-8z M3,8h8v1h-8z M4,9h6v1h-6z M4,10h6v1h-6z M6,11h2v1h-2z M6,12h2v1h-2z\" /></svg>", style))
		return err
	})
}

// Doc renders the sharp vector icon for "doc"
func Doc(cfg ...IconConfig) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		style := buildStyleAttr(cfg)
		_, err := io.WriteString(w, fmt.Sprintf("<svg xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 16 16\" shape-rendering=\"crispEdges\" fill=\"currentColor\" %s><path d=\"M6,2h7v1h-7z M4,3h1v1h-1z M6,3h7v1h-7z M3,4h2v1h-2z M6,4h7v1h-7z M6,5h7v1h-7z M3,6h10v1h-10z M3,7h10v1h-10z M3,8h10v1h-10z M3,9h10v1h-10z M3,10h10v1h-10z M3,11h10v1h-10z M3,12h10v1h-10z M3,13h10v1h-10z\" /></svg>", style))
		return err
	})
}

// Skull renders the sharp vector icon for "skull"
func Skull(cfg ...IconConfig) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		style := buildStyleAttr(cfg)
		_, err := io.WriteString(w, fmt.Sprintf("<svg xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 16 16\" shape-rendering=\"crispEdges\" fill=\"currentColor\" %s><path d=\"M7,3h3v1h-3z M6,4h5v1h-5z M5,5h1v1h-1z M8,5h1v1h-1z M11,5h1v1h-1z M5,6h1v1h-1z M8,6h1v1h-1z M11,6h1v1h-1z M5,7h7v1h-7z M6,8h5v1h-5z M7,9h1v1h-1z M9,9h1v1h-1z M4,10h2v1h-2z M11,10h1v1h-1z M6,11h2v1h-2z M10,11h1v1h-1z M7,12h4v1h-4z M5,13h2v1h-2z M11,13h2v1h-2z M4,14h1v1h-1z\" /></svg>", style))
		return err
	})
}

// Book renders the sharp vector icon for "book"
func Book(cfg ...IconConfig) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		style := buildStyleAttr(cfg)
		_, err := io.WriteString(w, fmt.Sprintf("<svg xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 16 16\" shape-rendering=\"crispEdges\" fill=\"currentColor\" %s><path d=\"M3,2h8v1h-8z M2,3h1v1h-1z M10,3h1v1h-1z M2,4h1v1h-1z M4,4h5v1h-5z M10,4h1v1h-1z M2,5h1v1h-1z M4,5h5v1h-5z M10,5h1v1h-1z M2,6h1v1h-1z M4,6h5v1h-5z M10,6h1v1h-1z M2,7h1v1h-1z M4,7h5v1h-5z M10,7h1v1h-1z M2,8h1v1h-1z M4,8h5v1h-5z M10,8h1v1h-1z M2,9h1v1h-1z M4,9h5v1h-5z M10,9h1v1h-1z M2,10h1v1h-1z M4,10h5v1h-5z M10,10h1v1h-1z M2,11h1v1h-1z M10,11h1v1h-1z M2,12h9v1h-9z M2,13h1v1h-1z M9,13h1v1h-1z M2,14h9v1h-9z\" /></svg>", style))
		return err
	})
}

// Linkedin renders the sharp vector icon for "linkedin"
func Linkedin(cfg ...IconConfig) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		style := buildStyleAttr(cfg)
		_, err := io.WriteString(w, fmt.Sprintf("<svg xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 16 16\" shape-rendering=\"crispEdges\" fill=\"currentColor\" %s><path d=\"M4,3h9v1h-9z M3,4h10v1h-10z M3,5h10v1h-10z M3,6h1v1h-1z M5,6h8v1h-8z M3,7h10v1h-10z M3,8h1v1h-1z M5,8h1v1h-1z M9,8h4v1h-4z M3,9h1v1h-1z M5,9h1v1h-1z M7,9h2v1h-2z M10,9h3v1h-3z M3,10h1v1h-1z M5,10h1v1h-1z M7,10h2v1h-2z M10,10h3v1h-3z M3,11h1v1h-1z M5,11h1v1h-1z M7,11h2v1h-2z M10,11h3v1h-3z M3,12h9v1h-9z\" /></svg>", style))
		return err
	})
}

// Bluesky renders the sharp vector icon for "bluesky"
func Bluesky(cfg ...IconConfig) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		style := buildStyleAttr(cfg)
		_, err := io.WriteString(w, fmt.Sprintf("<svg xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 16 16\" shape-rendering=\"crispEdges\" fill=\"currentColor\" %s><path d=\"M2,2h2v1h-2z M11,2h2v1h-2z M1,3h4v1h-4z M10,3h4v1h-4z M1,4h5v1h-5z M9,4h5v1h-5z M1,5h6v1h-6z M8,5h6v1h-6z M1,6h13v1h-13z M2,7h11v1h-11z M3,8h9v1h-9z M5,9h5v1h-5z M3,10h9v1h-9z M2,11h11v1h-11z M2,12h5v1h-5z M8,12h5v1h-5z M3,13h3v1h-3z M10,13h2v1h-2z\" /></svg>", style))
		return err
	})
}

// Itchio renders the sharp vector icon for "itchio"
func Itchio(cfg ...IconConfig) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		style := buildStyleAttr(cfg)
		_, err := io.WriteString(w, fmt.Sprintf("<svg xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 16 16\" shape-rendering=\"crispEdges\" fill=\"currentColor\" %s><path d=\"M3,1h10v1h-10z M2,2h12v1h-12z M1,3h14v1h-14z M0,4h2v1h-2z M3,4h1v1h-1z M5,4h1v1h-1z M7,4h1v1h-1z M9,4h1v1h-1z M11,4h1v1h-1z M13,4h1v1h-1z M15,4h1v1h-1z M0,5h1v1h-1z M2,5h1v1h-1z M4,5h1v1h-1z M6,5h1v1h-1z M8,5h1v1h-1z M10,5h1v1h-1z M12,5h1v1h-1z M15,5h1v1h-1z M2,6h12v1h-12z M1,7h3v1h-3z M6,7h3v1h-3z M11,7h4v1h-4z M1,8h2v1h-2z M12,8h3v1h-3z M1,9h2v1h-2z M7,9h1v1h-1z M12,9h3v1h-3z M1,10h1v1h-1z M6,10h1v1h-1z M8,10h1v1h-1z M13,10h2v1h-2z M1,11h1v1h-1z M13,11h2v1h-2z M1,12h1v1h-1z M5,12h5v1h-5z M13,12h2v1h-2z M1,13h1v1h-1z M4,13h7v1h-7z M13,13h2v1h-2z M2,14h12v1h-12z\" /></svg>", style))
		return err
	})
}

// Discord renders the sharp vector icon for "discord"
func Discord(cfg ...IconConfig) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		style := buildStyleAttr(cfg)
		_, err := io.WriteString(w, fmt.Sprintf("<svg xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 16 16\" shape-rendering=\"crispEdges\" fill=\"currentColor\" %s><path d=\"M3,4h2v1h-2z M11,4h2v1h-2z M2,5h4v1h-4z M7,5h2v1h-2z M10,5h4v1h-4z M2,6h12v1h-12z M1,7h14v1h-14z M1,8h4v1h-4z M7,8h2v1h-2z M11,8h4v1h-4z M0,9h5v1h-5z M7,9h2v1h-2z M11,9h5v1h-5z M0,10h16v1h-16z M0,11h4v1h-4z M5,11h6v1h-6z M12,11h4v1h-4z M1,12h2v1h-2z M6,12h4v1h-4z M13,12h2v1h-2z M1,13h4v1h-4z M11,13h4v1h-4z\" /></svg>", style))
		return err
	})
}

// Patreon renders the sharp vector icon for "patreon"
func Patreon(cfg ...IconConfig) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		style := buildStyleAttr(cfg)
		_, err := io.WriteString(w, fmt.Sprintf("<svg xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 16 16\" shape-rendering=\"crispEdges\" fill=\"currentColor\" %s><path d=\"M8,1h4v1h-4z M6,2h7v1h-7z M5,3h9v1h-9z M4,4h10v1h-10z M3,5h11v1h-11z M3,6h11v1h-11z M3,7h11v1h-11z M3,8h11v1h-11z M3,9h10v1h-10z M3,10h7v1h-7z M3,11h6v1h-6z M3,12h5v1h-5z M3,13h5v1h-5z M4,14h3v1h-3z\" /></svg>", style))
		return err
	})
}

// Sword renders the sharp vector icon for "sword"
func Sword(cfg ...IconConfig) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		style := buildStyleAttr(cfg)
		_, err := io.WriteString(w, fmt.Sprintf("<svg xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 16 16\" shape-rendering=\"crispEdges\" fill=\"currentColor\" %s><path d=\"M13,1h2v1h-2z M12,2h3v1h-3z M11,3h3v1h-3z M10,4h3v1h-3z M9,5h3v1h-3z M8,6h3v1h-3z M7,7h3v1h-3z M6,8h3v1h-3z M2,9h1v1h-1z M5,9h3v1h-3z M2,10h5v1h-5z M4,11h2v1h-2z M3,12h1v1h-1z M5,12h1v1h-1z M2,13h1v1h-1z M5,13h2v1h-2z M1,14h1v1h-1z\" /></svg>", style))
		return err
	})
}

// Stairs renders the sharp vector icon for "stairs"
func Stairs(cfg ...IconConfig) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		style := buildStyleAttr(cfg)
		_, err := io.WriteString(w, fmt.Sprintf("<svg xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 16 16\" shape-rendering=\"crispEdges\" fill=\"currentColor\" %s><path d=\"M10,5h2v1h-2z M10,6h2v1h-2z M8,7h4v1h-4z M8,8h4v1h-4z M6,9h6v1h-6z M6,10h6v1h-6z M4,11h8v1h-8z M4,12h8v1h-8z\" /></svg>", style))
		return err
	})
}

// Search renders the sharp vector icon for "search"
func Search(cfg ...IconConfig) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		style := buildStyleAttr(cfg)
		_, err := io.WriteString(w, fmt.Sprintf("<svg xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 16 16\" shape-rendering=\"crispEdges\" fill=\"currentColor\" %s><path d=\"M9,2h3v1h-3z M8,3h1v1h-1z M12,3h1v1h-1z M7,4h1v1h-1z M13,4h1v1h-1z M7,5h1v1h-1z M13,5h1v1h-1z M7,6h1v1h-1z M13,6h1v1h-1z M8,7h1v1h-1z M12,7h1v1h-1z M7,8h1v1h-1z M9,8h3v1h-3z M6,9h1v1h-1z M5,10h1v1h-1z M4,11h1v1h-1z\" /></svg>", style))
		return err
	})
}

// Github renders the sharp vector icon for "github"
func Github(cfg ...IconConfig) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		style := buildStyleAttr(cfg)
		_, err := io.WriteString(w, fmt.Sprintf("<svg xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 16 16\" shape-rendering=\"crispEdges\" fill=\"currentColor\" %s><path d=\"M6,2h4v1h-4z M4,3h8v1h-8z M3,4h2v1h-2z M6,4h4v1h-4z M11,4h2v1h-2z M3,5h3v1h-3z M10,5h3v1h-3z M2,6h3v1h-3z M11,6h3v1h-3z M2,7h3v1h-3z M11,7h3v1h-3z M2,8h1v1h-1z M4,8h2v1h-2z M10,8h4v1h-4z M2,9h1v1h-1z M4,9h3v1h-3z M9,9h5v1h-5z M3,10h1v1h-1z M5,10h1v1h-1z M10,10h3v1h-3z M3,11h2v1h-2z M10,11h3v1h-3z M4,12h2v1h-2z M10,12h2v1h-2z M10,13h1v1h-1z\" /></svg>", style))
		return err
	})
}

// X renders the sharp vector icon for "x"
func X(cfg ...IconConfig) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		style := buildStyleAttr(cfg)
		_, err := io.WriteString(w, fmt.Sprintf("<svg xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 16 16\" shape-rendering=\"crispEdges\" fill=\"currentColor\" %s><path d=\"M2,3h3v1h-3z M14,3h1v1h-1z M2,4h1v1h-1z M5,4h1v1h-1z M13,4h1v1h-1z M3,5h1v1h-1z M6,5h1v1h-1z M12,5h1v1h-1z M4,6h1v1h-1z M7,6h1v1h-1z M11,6h1v1h-1z M5,7h1v1h-1z M8,7h1v1h-1z M10,7h1v1h-1z M6,8h1v1h-1z M9,8h1v1h-1z M7,9h1v1h-1z M10,9h1v1h-1z M6,10h1v1h-1z M8,10h1v1h-1z M11,10h1v1h-1z M5,11h1v1h-1z M9,11h1v1h-1z M12,11h1v1h-1z M4,12h1v1h-1z M10,12h1v1h-1z M13,12h1v1h-1z M3,13h1v1h-1z M11,13h1v1h-1z M14,13h1v1h-1z M2,14h1v1h-1z M12,14h3v1h-3z\" /></svg>", style))
		return err
	})
}

// Crown renders the sharp vector icon for "crown"
func Crown(cfg ...IconConfig) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		style := buildStyleAttr(cfg)
		_, err := io.WriteString(w, fmt.Sprintf("<svg xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 16 16\" shape-rendering=\"crispEdges\" fill=\"currentColor\" %s><path d=\"M2,5h1v1h-1z M6,5h1v1h-1z M10,5h1v1h-1z M14,5h1v1h-1z M2,6h1v1h-1z M6,6h1v1h-1z M10,6h1v1h-1z M14,6h1v1h-1z M2,7h1v1h-1z M4,7h1v1h-1z M6,7h1v1h-1z M8,7h1v1h-1z M10,7h1v1h-1z M12,7h1v1h-1z M14,7h1v1h-1z M3,8h1v1h-1z M5,8h1v1h-1z M7,8h1v1h-1z M9,8h1v1h-1z M11,8h1v1h-1z M13,8h1v1h-1z M3,9h1v1h-1z M5,9h1v1h-1z M7,9h1v1h-1z M9,9h1v1h-1z M11,9h1v1h-1z M13,9h1v1h-1z M3,10h1v1h-1z M13,10h1v1h-1z M3,11h1v1h-1z M13,11h1v1h-1z M4,12h1v1h-1z M12,12h1v1h-1z M4,13h9v1h-9z\" /></svg>", style))
		return err
	})
}

// Splatter renders the sharp vector icon for "splatter"
func Splatter(cfg ...IconConfig) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		style := buildStyleAttr(cfg)
		_, err := io.WriteString(w, fmt.Sprintf("<svg xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 16 16\" shape-rendering=\"crispEdges\" fill=\"currentColor\" %s><path d=\"M4,3h1v1h-1z M9,3h1v1h-1z M11,4h1v1h-1z M7,5h2v1h-2z M11,6h1v1h-1z M14,6h1v1h-1z M4,7h1v1h-1z M6,7h2v1h-2z M9,7h1v1h-1z M7,8h1v1h-1z M10,8h1v1h-1z M1,9h1v1h-1z M4,9h1v1h-1z M6,9h1v1h-1z M8,9h2v1h-2z M11,9h1v1h-1z M7,10h3v1h-3z M5,11h1v1h-1z M7,11h1v1h-1z M11,11h1v1h-1z M5,12h1v1h-1z M3,13h1v1h-1z M6,13h1v1h-1z M10,13h1v1h-1z M13,13h1v1h-1z\" /></svg>", style))
		return err
	})
}

// Bricks renders the sharp vector icon for "bricks"
func Bricks(cfg ...IconConfig) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		style := buildStyleAttr(cfg)
		_, err := io.WriteString(w, fmt.Sprintf("<svg xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 16 16\" shape-rendering=\"crispEdges\" fill=\"currentColor\" %s><path d=\"M3,1h1v1h-1z M12,1h1v1h-1z M3,2h1v1h-1z M12,2h1v1h-1z M0,3h16v1h-16z M8,4h1v1h-1z M15,4h1v1h-1z M8,5h1v1h-1z M15,5h1v1h-1z M8,6h1v1h-1z M15,6h1v1h-1z M0,7h16v1h-16z M3,8h1v1h-1z M12,8h1v1h-1z M3,9h1v1h-1z M12,9h1v1h-1z M3,10h1v1h-1z M12,10h1v1h-1z M0,11h16v1h-16z M8,12h1v1h-1z M15,12h1v1h-1z M8,13h1v1h-1z M15,13h1v1h-1z M8,14h1v1h-1z M15,14h1v1h-1z M0,15h16v1h-16z\" /></svg>", style))
		return err
	})
}

// Hamburger renders the sharp vector icon for "hamburger"
func Hamburger(cfg ...IconConfig) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		style := buildStyleAttr(cfg)
		_, err := io.WriteString(w, fmt.Sprintf("<svg xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 16 16\" shape-rendering=\"crispEdges\" fill=\"currentColor\" %s><path d=\"M1,2h14v1h-14z M1,3h14v1h-14z M1,7h14v1h-14z M1,8h14v1h-14z M1,12h14v1h-14z M1,13h14v1h-14z\" /></svg>", style))
		return err
	})
}

// Sun renders the sharp vector icon for "sun"
func Sun(cfg ...IconConfig) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		style := buildStyleAttr(cfg)
		_, err := io.WriteString(w, fmt.Sprintf("<svg xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 16 16\" shape-rendering=\"crispEdges\" fill=\"currentColor\" %s><path d=\"M7,1h1v1h-1z M5,2h1v1h-1z M7,2h1v1h-1z M11,2h1v1h-1z M7,3h1v1h-1z M10,3h1v1h-1z M2,4h1v1h-1z M6,4h4v1h-4z M3,5h8v1h-8z M4,6h8v1h-8z M4,7h10v1h-10z M1,8h11v1h-11z M4,9h8v1h-8z M3,10h8v1h-8z M13,10h1v1h-1z M2,11h2v1h-2z M5,11h7v1h-7z M5,12h1v1h-1z M8,12h1v1h-1z M11,12h1v1h-1z M8,13h1v1h-1z M5,14h1v1h-1z\" /></svg>", style))
		return err
	})
}

// Axe renders the sharp vector icon for "axe"
func Axe(cfg ...IconConfig) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		style := buildStyleAttr(cfg)
		_, err := io.WriteString(w, fmt.Sprintf("<svg xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 16 16\" shape-rendering=\"crispEdges\" fill=\"currentColor\" %s><path d=\"M10,1h1v1h-1z M9,2h2v1h-2z M8,3h3v1h-3z M7,4h4v1h-4z M12,4h1v1h-1z M6,5h6v1h-6z M10,6h6v1h-6z M9,7h1v1h-1z M11,7h4v1h-4z M8,8h1v1h-1z M11,8h3v1h-3z M7,9h1v1h-1z M11,9h2v1h-2z M6,10h1v1h-1z M11,10h1v1h-1z M5,11h1v1h-1z M4,12h1v1h-1z M3,13h1v1h-1z M2,14h1v1h-1z\" /></svg>", style))
		return err
	})
}

// IntersectionEmpty renders the sharp vector icon for "intersection-empty"
func IntersectionEmpty(cfg ...IconConfig) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		style := buildStyleAttr(cfg)
		_, err := io.WriteString(w, fmt.Sprintf("<svg xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 16 16\" shape-rendering=\"crispEdges\" fill=\"currentColor\" %s><path d=\"M3,4h4v1h-4z M8,4h4v1h-4z M2,5h1v1h-1z M7,5h1v1h-1z M12,5h1v1h-1z M1,6h1v1h-1z M6,6h1v1h-1z M8,6h1v1h-1z M13,6h1v1h-1z M0,7h1v1h-1z M5,7h1v1h-1z M9,7h1v1h-1z M14,7h1v1h-1z M0,8h1v1h-1z M5,8h1v1h-1z M9,8h1v1h-1z M14,8h1v1h-1z M0,9h1v1h-1z M5,9h1v1h-1z M9,9h1v1h-1z M14,9h1v1h-1z M0,10h1v1h-1z M5,10h1v1h-1z M9,10h1v1h-1z M14,10h1v1h-1z M1,11h1v1h-1z M6,11h1v1h-1z M8,11h1v1h-1z M13,11h1v1h-1z M2,12h1v1h-1z M7,12h1v1h-1z M12,12h1v1h-1z M3,13h4v1h-4z M8,13h4v1h-4z\" /></svg>", style))
		return err
	})
}

// QuestionBox renders the sharp vector icon for "question-box"
func QuestionBox(cfg ...IconConfig) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		style := buildStyleAttr(cfg)
		_, err := io.WriteString(w, fmt.Sprintf("<svg xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 16 16\" shape-rendering=\"crispEdges\" fill=\"currentColor\" %s><path d=\"M1,0h15v1h-15z M0,1h1v1h-1z M15,1h1v1h-1z M0,2h1v1h-1z M15,2h1v1h-1z M0,3h1v1h-1z M6,3h5v1h-5z M15,3h1v1h-1z M0,4h1v1h-1z M5,4h7v1h-7z M15,4h1v1h-1z M0,5h1v1h-1z M4,5h2v1h-2z M10,5h2v1h-2z M15,5h1v1h-1z M0,6h1v1h-1z M4,6h2v1h-2z M10,6h2v1h-2z M15,6h1v1h-1z M0,7h1v1h-1z M10,7h2v1h-2z M15,7h1v1h-1z M0,8h1v1h-1z M8,8h2v1h-2z M15,8h1v1h-1z M0,9h1v1h-1z M7,9h3v1h-3z M15,9h1v1h-1z M0,10h1v1h-1z M7,10h2v1h-2z M15,10h1v1h-1z M0,11h1v1h-1z M15,11h1v1h-1z M0,12h1v1h-1z M7,12h2v1h-2z M15,12h1v1h-1z M0,13h1v1h-1z M7,13h2v1h-2z M15,13h1v1h-1z M0,14h1v1h-1z M15,14h1v1h-1z M0,15h15v1h-15z\" /></svg>", style))
		return err
	})
}

// PencilDiaganol renders the sharp vector icon for "pencil-diaganol"
func PencilDiaganol(cfg ...IconConfig) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		style := buildStyleAttr(cfg)
		_, err := io.WriteString(w, fmt.Sprintf("<svg xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 16 16\" shape-rendering=\"crispEdges\" fill=\"currentColor\" %s><path d=\"M11,3h1v1h-1z M10,4h1v1h-1z M12,4h1v1h-1z M9,5h3v1h-3z M8,6h3v1h-3z M7,7h3v1h-3z M6,8h3v1h-3z M5,9h3v1h-3z M4,10h3v1h-3z M3,11h1v1h-1z M5,11h1v1h-1z M2,12h3v1h-3z M2,13h2v1h-2z M1,14h1v1h-1z\" /></svg>", style))
		return err
	})
}

// Instagram renders the sharp vector icon for "instagram"
func Instagram(cfg ...IconConfig) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		style := buildStyleAttr(cfg)
		_, err := io.WriteString(w, fmt.Sprintf("<svg xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 16 16\" shape-rendering=\"crispEdges\" fill=\"currentColor\" %s><path d=\"M3,3h10v1h-10z M2,4h1v1h-1z M13,4h1v1h-1z M2,5h1v1h-1z M10,5h1v1h-1z M13,5h1v1h-1z M2,6h1v1h-1z M7,6h2v1h-2z M13,6h1v1h-1z M2,7h1v1h-1z M6,7h1v1h-1z M9,7h1v1h-1z M13,7h1v1h-1z M2,8h1v1h-1z M5,8h1v1h-1z M10,8h1v1h-1z M13,8h1v1h-1z M2,9h1v1h-1z M5,9h1v1h-1z M10,9h1v1h-1z M13,9h1v1h-1z M2,10h1v1h-1z M6,10h1v1h-1z M9,10h1v1h-1z M13,10h1v1h-1z M2,11h1v1h-1z M7,11h2v1h-2z M13,11h1v1h-1z M2,12h1v1h-1z M13,12h1v1h-1z M3,13h10v1h-10z\" /></svg>", style))
		return err
	})
}

// Steam renders the sharp vector icon for "steam"
func Steam(cfg ...IconConfig) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		style := buildStyleAttr(cfg)
		_, err := io.WriteString(w, fmt.Sprintf("<svg xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 16 16\" shape-rendering=\"crispEdges\" fill=\"currentColor\" %s><path d=\"M6,0h4v1h-4z M4,1h8v1h-8z M3,2h10v1h-10z M2,3h7v1h-7z M11,3h3v1h-3z M1,4h7v1h-7z M9,4h2v1h-2z M12,4h3v1h-3z M1,5h6v1h-6z M8,5h1v1h-1z M11,5h1v1h-1z M13,5h2v1h-2z M2,6h5v1h-5z M8,6h1v1h-1z M11,6h1v1h-1z M13,6h3v1h-3z M4,7h3v1h-3z M9,7h2v1h-2z M12,7h4v1h-4z M0,8h1v1h-1z M11,8h5v1h-5z M0,9h3v1h-3z M6,9h1v1h-1z M8,9h8v1h-8z M1,10h3v1h-3z M5,10h2v1h-2z M8,10h7v1h-7z M1,11h4v1h-4z M7,11h8v1h-8z M2,12h12v1h-12z M3,13h10v1h-10z M4,14h8v1h-8z M6,15h4v1h-4z\" /></svg>", style))
		return err
	})
}

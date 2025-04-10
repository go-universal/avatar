package avatar

import "strings"

// SVGColor represents an interface for SVG color handling.
type SVGColor interface {
	// IsHex returns true if the color is in hexadecimal format.
	IsHex() bool

	// IsDefinition returns true if the color is an SVG definition.
	IsDefinition() bool

	// resolve returns the color value if it's a hex code.
	// If it's a definition, it returns the URL with defs (as "url(#id)") as the first value,
	// and the definition content with the given id replaced by {id} as the second.
	resolve(id string) (string, string)
}

// svgColor represents an implementation of the SVGColor interface.
type svgColor struct {
	isDefinition bool
	value        string
}

// NewHex creates a new SVGColor in hexadecimal format.
// The input color should be a valid hex color string.
func NewHex(color string) SVGColor {
	color = strings.TrimSpace(color)
	if color == "" || color == "#" {
		return nil
	}

	if !strings.HasPrefix(color, "#") {
		color = "#" + color
	}

	return &svgColor{
		isDefinition: false,
		value:        color,
	}
}

// NewDefinition creates a new SVGColor as an SVG definition.
// The input definition should contain the placeholder id="{id}".
func NewDefinition(definition string) SVGColor {
	if !strings.Contains(definition, `id="{id}"`) {
		return nil
	}

	return &svgColor{
		isDefinition: true,
		value:        definition,
	}
}

func (c *svgColor) IsHex() bool {
	return !c.isDefinition
}

func (c *svgColor) IsDefinition() bool {
	return c.isDefinition
}

func (c *svgColor) resolve(id string) (string, string) {
	if c.IsHex() {
		return c.value, ""
	}

	return `url(#` + id + `)`,
		strings.ReplaceAll(c.value, "{id}", id)
}

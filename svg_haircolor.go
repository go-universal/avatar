package avatar

// SVGHairColor is an interface that defines methods to resolve hair colors.
type SVGHairColor interface {
	// resolveBase resolves the base color.
	resolveBase(id string) (string, string)

	// resolveShadow resolves the shadow color.
	resolveShadow(id string) (string, string)

	// resolveHighlight resolves the highlight color.
	resolveHighlight(id string) (string, string)
}

// svgHairColor is an implementation of the SVGHairColor interface.
type svgHairColor struct {
	base      SVGColor
	shadow    SVGColor
	highlight SVGColor
}

// NewHairColor creates a new instance of SVGHairColor with the provided base, shadow, and highlight colors.
func NewHairColor(base, shadow, highlight SVGColor) SVGHairColor {
	return &svgHairColor{
		base:      base,
		shadow:    shadow,
		highlight: highlight,
	}
}

func (hc *svgHairColor) resolveBase(id string) (string, string) {
	if hc.base == nil {
		return "transparent", ""
	} else {
		return hc.base.resolve(id)
	}
}

func (hc *svgHairColor) resolveShadow(id string) (string, string) {
	if hc.shadow == nil {
		return "transparent", ""
	} else {
		return hc.shadow.resolve(id)
	}
}

func (hc *svgHairColor) resolveHighlight(id string) (string, string) {
	if hc.highlight == nil {
		return "transparent", ""
	} else {
		return hc.highlight.resolve(id)
	}
}

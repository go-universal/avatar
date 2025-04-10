package avatar

// SVGPalette defines an interface for resolving Palette colors.
type SVGPalette interface {
	// resolveShape resolves the shape color
	resolveShape(id string) (string, string)

	// resolveText resolves the text color
	resolveText(id string) (string, string)

	// resolveDress resolves the dress color
	resolveDress(id string) (string, string)

	// resolveShadow resolves the dress shadows color
	resolveShadow(id string) (string, string)

	// resolveDecorator resolves the decorator color
	resolveDecorator(id string) (string, string)
}

// svgPalette is an implementation of the SVGPalette interface.
type svgPalette struct {
	shape     SVGColor
	text      SVGColor
	dress     SVGColor
	shadow    SVGColor
	decorator SVGColor
}

// NewPalette creates a new SVGPalette with the specified colors.
func NewPalette(shape, text, dress, shadow, decorator SVGColor) SVGPalette {
	return &svgPalette{
		shape:     shape,
		text:      text,
		dress:     dress,
		shadow:    shadow,
		decorator: decorator,
	}
}

func (p *svgPalette) resolveShape(id string) (string, string) {
	if p.shape == nil {
		return "transparent", ""
	} else {
		return p.shape.resolve(id)
	}
}

func (p *svgPalette) resolveText(id string) (string, string) {
	if p.text == nil {
		return "transparent", ""
	} else {
		return p.text.resolve(id)
	}
}

func (p *svgPalette) resolveDress(id string) (string, string) {
	if p.dress == nil {
		return "transparent", ""
	} else {
		return p.dress.resolve(id)
	}
}

func (p *svgPalette) resolveShadow(id string) (string, string) {
	if p.shadow == nil {
		return "transparent", ""
	} else {
		return p.shadow.resolve(id)
	}
}

func (p *svgPalette) resolveDecorator(id string) (string, string) {
	if p.decorator == nil {
		return "transparent", ""
	} else {
		return p.decorator.resolve(id)
	}
}

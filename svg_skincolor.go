package avatar

// SVGSkinColor is an interface that defines methods to resolve skin colors.
type SVGSkinColor interface {
	// resolveSkin resolves the skin color.
	resolveSkin(id string) (string, string)

	// resolveSkin resolves the shadow color.
	resolveShadow(id string) (string, string)
}

// svgSkinColor is an implementation of the SVGSkinColor interface.
type svgSkinColor struct {
	skin   SVGColor
	shadow SVGColor
}

// NewSkinColor creates a new SVGSkinColor with the provided skin and shadow colors.
func NewSkinColor(skin, shadow SVGColor) SVGSkinColor {
	return &svgSkinColor{
		skin:   skin,
		shadow: shadow,
	}
}

func (sc *svgSkinColor) resolveSkin(id string) (string, string) {
	if sc.skin == nil {
		return "transparent", ""
	} else {
		return sc.skin.resolve(id)
	}
}

func (sc *svgSkinColor) resolveShadow(id string) (string, string) {
	if sc.shadow == nil {
		return "transparent", ""
	} else {
		return sc.shadow.resolve(id)
	}
}

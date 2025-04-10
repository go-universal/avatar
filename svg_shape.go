package avatar

// SVGShape is an interface that defines methods to resolve background shape and mask.
type SVGShape interface {
	// Shape returns the SVG shape as a string.
	Shape() string

	// Mask returns the SVG mask as a string.
	Mask() string
}

// svgShape is an implementation of the SVGShape interface.
type svgShape struct {
	shape string
	mask  string
}

// NewShape creates a new SVGShape with the given shape and mask.
// Shape support fill="{shape}" to colorize shape.
func NewShape(shape, mask string) SVGShape {
	return &svgShape{
		shape: shape,
		mask:  mask,
	}
}

func (s *svgShape) Shape() string {
	return s.shape
}

func (s *svgShape) Mask() string {
	return s.mask
}

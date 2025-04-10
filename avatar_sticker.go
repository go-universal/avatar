package avatar

import (
	"encoding/base64"
	"os"
	"strings"

	"github.com/google/uuid"
)

type StickerAvatar interface {
	// Shape returns the background shape of the avatar.
	Shape() string

	// Palette returns the color palette of the avatar.
	Palette() string

	// Sticker returns the sticker of the avatar.
	Sticker() string

	// RandomizeShape randomizes the background shape of the avatar.
	RandomizeShape(only ...Shape) StickerAvatar

	// RandomizePalette randomizes the color palette of the avatar.
	RandomizePalette(only ...Palette) StickerAvatar

	// RandomizeSticker randomizes the sticker of the avatar.
	RandomizeSticker(only ...Sticker) StickerAvatar

	// Render returns the inline SVG representation of the avatar.
	Render() string

	// SVG returns the SVG representation of the avatar.
	SVG() string

	// Base64 returns the base64 encoded SVG representation of the avatar.
	Base64() string

	// Save saves the avatar to the specified destination.
	Save(dest string) error

	// Params returns the parameters of the avatar as a map.
	Params() map[string]string
}

// newStickerAvatar creates a new instance of StickerAvatar.
func newStickerAvatar(f *factory, sticker string) StickerAvatar {
	return &stickerAV{
		f:  f,
		id: uuid.NewString(),

		shape:   "",
		palette: "",

		sticker: sticker,
	}
}

// stickerAV is a concrete implementation of the StickerAvatar interface.
type stickerAV struct {
	f  *factory
	id string

	shape   string
	palette string

	sticker string
}

func (s *stickerAV) idFor(k string) string {
	return k + "-" + s.id
}

func (s *stickerAV) Shape() string {
	return s.shape
}

func (s *stickerAV) Palette() string {
	return s.palette
}

func (s *stickerAV) Sticker() string {
	return s.sticker
}

func (s *stickerAV) RandomizeShape(only ...Shape) StickerAvatar {
	s.shape = s.f.randomShape(toStringSlice(only)...)
	return s
}

func (s *stickerAV) RandomizePalette(only ...Palette) StickerAvatar {
	s.palette = s.f.randomPalette(toStringSlice(only)...)
	return s
}

func (s *stickerAV) RandomizeSticker(only ...Sticker) StickerAvatar {
	s.sticker = s.f.randomSticker(toStringSlice(only)...)
	return s
}

func (s *stickerAV) Render() string {
	var builder strings.Builder
	builder.WriteString(`<svg viewBox="0 0 128 128" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink">`)
	builder.WriteString(`<desc>Created with https://github.com/go-universal/avatar</desc>`)

	// resolve
	shape := s.f.resolveShape(s.shape)
	palette := s.f.resolvePalette(s.palette)
	sticker := s.f.resolveSticker(s.sticker)

	shapeCol, shapeDef := palette.resolveShape(s.idFor("back"))
	textCol, textDef := palette.resolveText(s.idFor("text"))

	// create replacer
	template := strings.NewReplacer(
		"{shape}", shapeCol,
		"{text}", textCol,
	)

	// add definitions
	builder.WriteString("<defs>")
	builder.WriteString(shapeDef)
	builder.WriteString(textDef)
	builder.WriteString("</defs>")

	// add background shape
	if shape.Shape() != "" {
		builder.WriteString(template.Replace(shape.Shape()))
	}

	// Add mask and wrapper group
	if shape.Mask() != "" {
		id := s.idFor("mask")
		builder.WriteString(`<mask id="` + id + `">`)
		builder.WriteString(shape.Mask())
		builder.WriteString(`</mask>`)
		builder.WriteString(`<g mask="url(#` + id + `)">`)
	} else {
		builder.WriteString(`<g>`)
	}

	// Add text avatar
	builder.WriteString(template.Replace(sticker))

	builder.WriteString(`</g>`)
	builder.WriteString(`</svg>`)
	return builder.String()
}

func (s *stickerAV) SVG() string {
	return `<?xml version="1.0" encoding="utf-8"?><!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">` +
		s.Render()
}

func (s *stickerAV) Base64() string {
	return "data:image/svg+xml;base64," +
		string(base64.StdEncoding.EncodeToString([]byte(s.SVG())))
}

func (s *stickerAV) Save(dest string) error {
	return os.WriteFile(dest, []byte(s.SVG()), 0644)

}

func (s *stickerAV) Params() map[string]string {
	return map[string]string{
		"shape":   s.Shape(),
		"palette": s.Palette(),
		"sticker": s.Sticker(),
	}
}

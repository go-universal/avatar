package avatar

import (
	"encoding/base64"
	"os"
	"strings"
	"unicode/utf8"

	"github.com/google/uuid"
)

// TextAvatar represents an interface for generating and manipulating avatar for text.
type TextAvatar interface {
	// Shape returns the background shape of the avatar.
	Shape() string

	// Palette returns the color palette of the avatar.
	Palette() string

	// Code returns the letter code of the avatar.
	Code() rune

	// Letter returns the letter  of the avatar.
	Letter() string

	// RandomizeShape randomizes the background shape of the avatar.
	RandomizeShape(only ...Shape) TextAvatar

	// RandomizePalette randomizes the color palette of the avatar.
	RandomizePalette(only ...Palette) TextAvatar

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

// newLetterAvatar creates a new instance of TextAvatar.
func newLetterAvatar(f *factory, name string) TextAvatar {
	letter := extractLetter(name)
	if v, ok := f.transforms[letter]; ok {
		letter = v
	}

	return &textAV{
		f:  f,
		id: uuid.NewString(),

		shape:   "",
		palette: "",

		letter: letter,
	}
}

// textAV is a concrete implementation of the StickerAvatar TextAvatar.
type textAV struct {
	f  *factory
	id string

	shape   string
	palette string

	letter string
}

func (t *textAV) idFor(k string) string {
	return k + "-" + t.id
}

func (t *textAV) Shape() string {
	return t.shape
}

func (t *textAV) Palette() string {
	return t.palette
}

func (t *textAV) Code() rune {
	r, _ := utf8.DecodeRuneInString(t.letter)
	return r
}

func (t *textAV) Letter() string {
	return t.letter
}

func (t *textAV) RandomizeShape(only ...Shape) TextAvatar {
	t.shape = t.f.randomShape(toStringSlice(only)...)
	return t
}

func (t *textAV) RandomizePalette(only ...Palette) TextAvatar {
	t.palette = t.f.randomPalette(toStringSlice(only)...)
	return t
}

func (t *textAV) Render() string {
	var builder strings.Builder
	builder.WriteString(`<svg viewBox="0 0 128 128" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink">`)
	builder.WriteString(`<desc>Created with https://github.com/go-universal/avatar</desc>`)

	// resolve
	shape := t.f.resolveShape(t.shape)
	palette := t.f.resolvePalette(t.palette)
	letter := t.f.resolveLetter(t.letter)

	shapeCol, shapeDef := palette.resolveShape(t.idFor("back"))
	textCol, textDef := palette.resolveText(t.idFor("text"))

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
		id := t.idFor("mask")
		builder.WriteString(`<mask id="` + id + `">`)
		builder.WriteString(shape.Mask())
		builder.WriteString(`</mask>`)
		builder.WriteString(`<g mask="url(#` + id + `)">`)
	} else {
		builder.WriteString(`<g>`)
	}

	// Add text avatar
	builder.WriteString(template.Replace(letter))

	builder.WriteString(`</g>`)
	builder.WriteString(`</svg>`)
	return builder.String()
}

func (t *textAV) SVG() string {
	return `<?xml version="1.0" encoding="utf-8"?><!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">` +
		t.Render()
}

func (t *textAV) Base64() string {
	return "data:image/svg+xml;base64," +
		string(base64.StdEncoding.EncodeToString([]byte(t.SVG())))
}

func (t *textAV) Save(dest string) error {
	return os.WriteFile(dest, []byte(t.SVG()), 0644)
}

func (t *textAV) Params() map[string]string {
	return map[string]string{
		"shape":   t.Shape(),
		"palette": t.Palette(),
		"letter":  t.Letter(),
	}
}

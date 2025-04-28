package avatar

import (
	"encoding/base64"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/go-universal/utils"
	"github.com/google/uuid"
)

// PersonAvatar represents an interface for generating and manipulating avatar for person.
type PersonAvatar interface {
	// Shape returns the background shape of the avatar.
	Shape() string

	// Palette returns the color palette of the avatar.
	Palette() string

	// SkinColor returns the skin color of the avatar.
	SkinColor() string

	// HairColor returns the hair color of the avatar.
	HairColor() string

	// IsMale returns true if the avatar is male, false otherwise.
	IsMale() bool

	// Hair returns the hair style of the avatar.
	Hair() string

	// Beard returns the beard style of the avatar.
	Beard() string

	// Dress returns the dress style of the avatar.
	Dress() string

	// Eye returns the eye style of the avatar.
	Eye() string

	// Eyebrow returns the eyebrow style of the avatar.
	Eyebrow() string

	// Mouth returns the mouth style of the avatar.
	Mouth() string

	// Glasses returns the glasses style of the avatar.
	Glasses() string

	// Accessory returns the accessory style of the avatar.
	Accessory() string

	// RandomizeShape randomizes the shape of the avatar.
	RandomizeShape(only ...Shape) PersonAvatar

	// RandomizePalette randomizes the color palette of the avatar.
	RandomizePalette(only ...Palette) PersonAvatar

	// RandomizeSkinColor randomizes the skin color of the avatar.
	RandomizeSkinColor(only ...SkinColor) PersonAvatar

	// RandomizeHairColor randomizes the hair color of the avatar.
	RandomizeHairColor(only ...HairColor) PersonAvatar

	// RandomizeHair randomizes the hair style of the avatar.
	RandomizeHair(only ...Hair) PersonAvatar

	// RandomizeBeard randomizes the beard style of the avatar.
	RandomizeBeard(only ...Beard) PersonAvatar

	// RandomizeDress randomizes the dress style of the avatar.
	RandomizeDress(only ...Dress) PersonAvatar

	// RandomizeEye randomizes the eye style of the avatar.
	RandomizeEye(only ...Eye) PersonAvatar

	// RandomizeEyebrow randomizes the eyebrow style of the avatar.
	RandomizeEyebrow(only ...Eyebrow) PersonAvatar

	// RandomizeMouth randomizes the mouth style of the avatar.
	RandomizeMouth(only ...Mouth) PersonAvatar

	// RandomizeGlasses randomizes the glasses style of the avatar.
	RandomizeGlasses(only ...Glasses) PersonAvatar

	// RandomizeAccessory randomizes the accessory style of the avatar.
	RandomizeAccessory(only ...Accessory) PersonAvatar

	// Render returns the inline SVG representation of the avatar.
	Render() string

	// SVG returns the SVG representation of the avatar.
	SVG() string

	// Base64 returns the base64 encoded representation of the avatar.
	Base64() string

	// Params returns the parameters of the avatar as a map.
	Params() map[string]string

	// SaveAs saves the avatar to the specified destination.
	SaveAs(dest string) error

	// Save stores the avatar to storage.
	Save() error

	// Delete removes the uploaded file safely, queueing the file name on failure if queue passed to factory.
	Delete() error

	// Path returns the file path where the avatar file is stored.
	Path() string

	// URL returns the URL where the avatar file can be accessed.
	URL() string
}

// personAV is a concrete implementation of the PersonAvatar interface.
type personAV struct {
	f  *factory
	id string

	name string
	out  string

	shape   string
	palette string

	skinColor string
	hairColor string

	isMale    bool
	hair      string
	beard     string
	dress     string
	eye       string
	eyebrow   string
	mouth     string
	glasses   string
	accessory string
}

// newPersonAvatar creates a new instance of PersonAvatar.
func newPersonAvatar(f *factory, isMale bool, name string) PersonAvatar {
	name = strings.TrimSuffix(name, ".svg") + ".svg"

	return &personAV{
		f:  f,
		id: uuid.NewString(),

		name: name,
		out:  "",

		shape:   "",
		palette: "",

		skinColor: "",
		hairColor: "",

		isMale:    isMale,
		hair:      "",
		beard:     "",
		dress:     "",
		eye:       "",
		eyebrow:   "",
		mouth:     "",
		glasses:   "",
		accessory: "",
	}
}

func (p *personAV) idFor(k string) string {
	return k + "-" + p.id
}

func (p *personAV) Shape() string {
	return p.shape
}

func (p *personAV) Palette() string {
	return p.palette
}

func (p *personAV) SkinColor() string {
	return p.skinColor
}

func (p *personAV) HairColor() string {
	return p.hairColor
}

func (p *personAV) IsMale() bool {
	return p.isMale
}

func (p *personAV) Hair() string {
	return p.hair
}

func (p *personAV) Beard() string {
	return p.beard
}

func (p *personAV) Dress() string {
	return p.dress
}

func (p *personAV) Eye() string {
	return p.eye
}

func (p *personAV) Eyebrow() string {
	return p.eyebrow
}

func (p *personAV) Mouth() string {
	return p.mouth
}

func (p *personAV) Glasses() string {
	return p.glasses
}

func (p *personAV) Accessory() string {
	return p.accessory
}

func (p *personAV) RandomizeShape(only ...Shape) PersonAvatar {
	p.shape = p.f.randomShape(toStringSlice(only)...)
	return p
}

func (p *personAV) RandomizePalette(only ...Palette) PersonAvatar {
	p.palette = p.f.randomPalette(toStringSlice(only)...)
	return p
}

func (p *personAV) RandomizeSkinColor(only ...SkinColor) PersonAvatar {
	p.skinColor = p.f.randomSkinColor(toStringSlice(only)...)
	return p
}

func (p *personAV) RandomizeHairColor(only ...HairColor) PersonAvatar {
	p.hairColor = p.f.randomHairColor(toStringSlice(only)...)
	return p
}

func (p *personAV) RandomizeHair(only ...Hair) PersonAvatar {
	p.hair = p.f.randomHair(p.isMale, toStringSlice(only)...)
	return p
}

func (p *personAV) RandomizeBeard(only ...Beard) PersonAvatar {
	p.beard = p.f.randomBeard(p.isMale, toStringSlice(only)...)
	return p
}

func (p *personAV) RandomizeDress(only ...Dress) PersonAvatar {
	p.dress = p.f.randomDress(p.isMale, toStringSlice(only)...)
	return p
}

func (p *personAV) RandomizeEye(only ...Eye) PersonAvatar {
	p.eye = p.f.randomEye(p.isMale, toStringSlice(only)...)
	return p
}

func (p *personAV) RandomizeEyebrow(only ...Eyebrow) PersonAvatar {
	p.eyebrow = p.f.randomEyebrow(p.isMale, toStringSlice(only)...)
	return p
}

func (p *personAV) RandomizeMouth(only ...Mouth) PersonAvatar {
	p.mouth = p.f.randomMouth(p.isMale, toStringSlice(only)...)
	return p
}

func (p *personAV) RandomizeGlasses(only ...Glasses) PersonAvatar {
	p.glasses = p.f.randomGlasses(p.isMale, toStringSlice(only)...)
	return p
}

func (p *personAV) RandomizeAccessory(only ...Accessory) PersonAvatar {
	p.accessory = p.f.randomAccessory(p.isMale, toStringSlice(only)...)
	return p
}

func (p *personAV) Render() string {
	var builder strings.Builder
	builder.WriteString(`<svg viewBox="0 0 128 128" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink">`)
	builder.WriteString(`<desc>Created with https://github.com/go-universal/avatar</desc>`)

	// resolve
	shape := p.f.resolveShape(p.shape)
	palette := p.f.resolvePalette(p.palette)
	body := p.f.resolveBodyShape()
	skinColor := p.f.resolveSkinColor(p.skinColor)
	hairColor := p.f.resolveHairColor(p.hairColor)
	hair := p.f.resolveHair(p.isMale, p.hair)
	beard := p.f.resolveBeard(p.isMale, p.beard)
	dress := p.f.resolveDress(p.isMale, p.dress)
	eye := p.f.resolveEye(p.isMale, p.eye)
	eyebrow := p.f.resolveEyebrow(p.isMale, p.eyebrow)
	mouth := p.f.resolveMouth(p.isMale, p.mouth)
	glasses := p.f.resolveGlasses(p.isMale, p.glasses)
	accessory := p.f.resolveAccessory(p.isMale, p.accessory)

	skinCol, skinDef := skinColor.resolveSkin(p.idFor("skin"))
	skinShadowCol, skinShadowDef := skinColor.resolveShadow(p.idFor("skin-shadow"))

	hairCol, hairDef := hairColor.resolveBase(p.idFor("hair"))
	hairShadowCol, hairShadowDef := hairColor.resolveShadow(p.idFor("hair-shadow"))
	hairHighlightCol, hairHighlightDef := hairColor.resolveHighlight(p.idFor("hair-highlight"))

	shapeCol, shapeDef := palette.resolveShape(p.idFor("back"))
	dressCol, dressDef := palette.resolveDress(p.idFor("dress"))
	dressShadowCol, dressShadowDef := palette.resolveShadow(p.idFor("dress-shadow"))
	decoratorCol, decoratorDef := palette.resolveDecorator(p.idFor("decorator"))

	// create replacer
	template := strings.NewReplacer(
		"{shape}", shapeCol,
		"{skin}", skinCol,
		"{skin_shadow}", skinShadowCol,
		"{hair}", hairCol,
		"{hair_shadow}", hairShadowCol,
		"{hair_highlight}", hairHighlightCol,
		"{dress}", dressCol,
		"{dress_shadow}", dressShadowCol,
		"{decorator}", decoratorCol,
	)

	// add definitions
	builder.WriteString("<defs>")
	builder.WriteString(skinDef)
	builder.WriteString(skinShadowDef)
	builder.WriteString(hairDef)
	builder.WriteString(hairShadowDef)
	builder.WriteString(hairHighlightDef)
	builder.WriteString(shapeDef)
	builder.WriteString(dressDef)
	builder.WriteString(dressShadowDef)
	builder.WriteString(decoratorDef)
	builder.WriteString("</defs>")

	// add background shape
	if shape.Shape() != "" {
		builder.WriteString(template.Replace(shape.Shape()))
	}

	// Add mask and wrapper group
	if shape.Mask() != "" {
		id := p.idFor("mask")
		builder.WriteString(`<mask id="` + id + `">`)
		builder.WriteString(shape.Mask())
		builder.WriteString(`</mask>`)
		builder.WriteString(`<g mask="url(#` + id + `)">`)
	} else {
		builder.WriteString(`<g>`)
	}

	// Add person avatar
	builder.WriteString(template.Replace(body))
	builder.WriteString(template.Replace(dress))
	builder.WriteString(template.Replace(eye))
	builder.WriteString(template.Replace(eyebrow))
	builder.WriteString(template.Replace(mouth))
	builder.WriteString(template.Replace(glasses))
	builder.WriteString(template.Replace(accessory))
	builder.WriteString(template.Replace(hair))
	builder.WriteString(template.Replace(beard))

	builder.WriteString(`</g>`)
	builder.WriteString(`</svg>`)
	return builder.String()
}

func (p *personAV) SVG() string {
	return `<?xml version="1.0" encoding="utf-8"?><!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">` +
		p.Render()
}

func (p *personAV) Base64() string {
	return "data:image/svg+xml;base64," +
		string(base64.StdEncoding.EncodeToString([]byte(p.SVG())))
}

func (p *personAV) Params() map[string]string {
	return map[string]string{
		"shape":      p.Shape(),
		"palette":    p.Palette(),
		"skin_color": p.SkinColor(),
		"hair_color": p.HairColor(),
		"is_male":    strconv.FormatBool(p.IsMale()),
		"hair":       p.Hair(),
		"beard":      p.Beard(),
		"dress":      p.Dress(),
		"eye":        p.Eye(),
		"eyebrow":    p.Eyebrow(),
		"mouth":      p.Mouth(),
		"glasses":    p.Glasses(),
		"accessory":  p.Accessory(),
	}
}

func (p *personAV) SaveAs(dest string) error {
	return os.WriteFile(dest, []byte(p.SVG()), 0644)
}

func (p *personAV) Save() error {
	if p.out != "" || p.f.storage == "" || p.name == "" {
		return nil
	}

	var name string
	if p.f.numbered {
		if n, err := utils.NumberedFile(p.f.storage, p.name); err != nil {
			return err
		} else {
			name = n
		}
	} else {
		name = utils.TimestampedFile(p.name)
	}

	dest := utils.NormalizePath(p.f.storage, name)

	exists, err := utils.FileExists(dest)
	if err != nil {
		return err
	} else if exists {
		return fmt.Errorf("%s file exists", dest)
	}

	if err := p.SaveAs(dest); err != nil {
		return err
	}

	p.out = name
	return nil
}

func (p *personAV) Delete() error {
	path := p.Path()
	if path == "" {
		return nil
	}

	err := os.Remove(path)
	if err == nil || errors.Is(err, os.ErrNotExist) {
		return nil
	}

	if p.f.queue != nil {
		return p.f.queue.Push(path)
	}

	return err
}

func (p *personAV) Path() string {
	if p.out == "" || p.f.storage == "" {
		return ""
	}

	return utils.NormalizePath(p.f.storage, p.out)
}

func (p *personAV) URL() string {
	path := p.Path()
	if path == "" {
		return ""
	}

	return utils.AbsoluteURL(p.f.prefix, path)
}

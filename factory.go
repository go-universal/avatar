package avatar

import (
	"sync"

	"github.com/go-universal/cache"
)

// AvatarFactory defines an interface for building different types of avatars.
// It provides methods to create male, female, text and sticker avatars.
type AvatarFactory interface {
	// NewMale creates a new male avatar.
	NewMale(name string) PersonAvatar

	// NewFemale creates a new female avatar.
	NewFemale(name string) PersonAvatar

	// NewPerson creates a new person avatar based on the gender specified by the isMale parameter.
	// If isMale is true, it creates a male avatar; otherwise, a female avatar.
	NewPerson(isMale bool, name string) PersonAvatar

	// NewText creates a new text avatar with the specified name.
	NewText(title, name string) TextAvatar

	// NewSticker creates a new sticker avatar with the specified stickers.
	NewSticker(sticker Sticker, name string) StickerAvatar

	// Shapes returns a list of available avatar background shapes.
	Shapes() []string

	// Palettes returns a list of available color palettes for customization.
	Palettes() []string

	// Stickers returns a list of available stickers for the avatar.
	Stickers() []string

	// SkinColors returns a list of available skin color options.
	SkinColors() []string

	// HairColors returns a list of available hair color options.
	HairColors() []string

	// MaleHairs returns a list of available male hairstyles.
	MaleHairs() []string

	// MaleBeards returns a list of available male beard styles.
	MaleBeards() []string

	// MaleDresses returns a list of available male dresses or outfits.
	MaleDresses() []string

	// MaleEyes returns a list of available male eye styles.
	MaleEyes() []string

	// MaleEyebrows returns a list of available male eyebrow styles.
	MaleEyebrows() []string

	// MaleMouths returns a list of available male mouth styles.
	MaleMouths() []string

	// MaleGlasses returns a list of available male glasses styles.
	MaleGlasses() []string

	// MaleAccessories returns a list of available male accessories.
	MaleAccessories() []string

	// FemaleHairs returns a list of available female hairstyles.
	FemaleHairs() []string

	// FemaleDresses returns a list of available female dresses or outfits.
	FemaleDresses() []string

	// FemaleEyes returns a list of available female eye styles.
	FemaleEyes() []string

	// FemaleEyebrows returns a list of available female eyebrow styles.
	FemaleEyebrows() []string

	// FemaleMouths returns a list of available female mouth styles.
	FemaleMouths() []string

	// FemaleGlasses returns a list of available female glasses styles.
	FemaleGlasses() []string

	// FemaleAccessories returns a list of available female accessories.
	FemaleAccessories() []string
}

type factory struct {
	// FS
	numbered bool
	storage  string
	prefix   string
	queue    cache.Queue

	// Global
	mutex    sync.RWMutex
	shapes   map[string]SVGShape
	palettes map[string]SVGPalette

	// Stickers
	stickers map[string]string

	// Text
	letters    map[string]string
	transforms map[string]string

	// Person
	bodyShape  string
	skinColors map[string]SVGSkinColor
	hairColors map[string]SVGHairColor

	maleHairs       map[string]string
	maleBeards      map[string]string
	maleDresses     map[string]string
	maleEyes        map[string]string
	maleEyebrows    map[string]string
	maleMouths      map[string]string
	maleGlasses     map[string]string
	maleAccessories map[string]string

	femaleHairs       map[string]string
	femaleDresses     map[string]string
	femaleEyes        map[string]string
	femaleEyebrows    map[string]string
	femaleMouths      map[string]string
	femaleGlasses     map[string]string
	femaleAccessories map[string]string
}

package avatar

// FactoryBuilder defines the interface for building an avatar factory.
type FactoryBuilder interface {
	// AddAccessory adds a custom accessory to the avatar factory.
	// Shape can include fill="{decorator}" to match the avatar palette.
	AddAccessory(isMale bool, name, shape string) FactoryBuilder

	// DefaultAccessories add all predefined accessories to the avatar factory.
	DefaultAccessories() FactoryBuilder

	// NecklaceAccessory add necklace accessory for female to the avatar factory.
	NecklaceAccessory() FactoryBuilder

	// ChokerAccessory add choker accessory for female to the avatar factory.
	ChokerAccessory() FactoryBuilder

	// AddBeard adds a custom beard to the avatar factory.
	// Shape can include fill="{hair}", fill="{hair_shadow}" and fill="{hair_highlight}" to match the avatar palette.
	AddBeard(name, shape string) FactoryBuilder

	// DefaultBeards add all predefined beards to the avatar factory.
	DefaultBeards() FactoryBuilder

	// MustachBeard add mustache beard to the avatar factory.
	MustachBeard() FactoryBuilder

	// FancyMustachBeard add fancy mustache beard to the avatar factory.
	FancyMustachBeard() FactoryBuilder

	// NormalBeard add normal beard to the avatar factory.
	NormalBeard() FactoryBuilder

	// MediumBeard add medium-length beard to the avatar factory.
	MediumBeard() FactoryBuilder

	// LongBeard add long beard to the avatar factory.
	LongBeard() FactoryBuilder

	// SetBodyShape sets the default body shape for the avatar factory.
	// Shape can include fill="{skin}" and fill="{skin_shadow}" to match the avatar palette.
	SetBodyShape(shape string) FactoryBuilder

	// AddDress adds a custom dress to the avatar factory.
	// Shape can include fill="{dress}", fill="{dress_shadow}" and fill="{decorator}" to match the avatar palette.
	AddDress(isMale bool, name, shape string) FactoryBuilder

	// DefaultDresses add all predefined dresses to the avatar factory.
	DefaultDresses() FactoryBuilder

	// SuitDress add suit dress to the avatar factory.
	SuitDress() FactoryBuilder

	// ShirtDress add shirt dress to the avatar factory.
	ShirtDress() FactoryBuilder

	// TShirtDress add t-shirt dress to the avatar factory.
	TShirtDress() FactoryBuilder

	// AddEye adds a custom eye to the avatar factory.
	AddEye(isMale bool, name, shape string) FactoryBuilder

	// AddEyebrow adds a custom eyebrow to the avatar factory.
	AddEyebrow(isMale bool, name, shape string) FactoryBuilder

	// AddGlasses adds a custom glasses to the avatar factory.
	// Shape can include fill="{decorator}" to match the avatar palette.
	AddGlasses(isMale bool, name, shape string) FactoryBuilder

	// DefaultGlasses add all predefined glasses to the avatar factory.
	DefaultGlasses() FactoryBuilder

	// PrescriptionGlasses add prescription glasses to the avatar factory.
	PrescriptionGlasses() FactoryBuilder

	// RoundPrescriptionGlasses add round prescription glasses to the avatar factory.
	RoundPrescriptionGlasses() FactoryBuilder

	// SunglassGlasses add sunglass glasses to the avatar factory.
	SunglassGlasses() FactoryBuilder

	// RoundSunglassGlasses add round sunglass glasses to the avatar factory.
	RoundSunglassGlasses() FactoryBuilder

	// AddHair adds a custom hair to the avatar factory.
	// Shape can include fill="{hair}", fill="{hair_shadow}" and fill="{hair_highlight}" to match the avatar palette.
	AddHair(isMale bool, name, shape string) FactoryBuilder

	// DefaultHairs add all predefined hairs to the avatar factory.
	DefaultHairs() FactoryBuilder

	// ShortHair add short hair to the avatar factory.
	ShortHair() FactoryBuilder

	// MediumHair add medium hair to the avatar factory.
	MediumHair() FactoryBuilder

	// WavyHair add wavy hair to the avatar factory.
	WavyHair() FactoryBuilder

	// CurlyHair add curly hair to the avatar factory.
	CurlyHair() FactoryBuilder

	// AddHaircolor adds a custom hair color to the avatar factory.
	AddHaircolor(name string, color SVGHairColor) FactoryBuilder

	// DefaultHairColors add all predefined hair colors to the avatar factory.
	DefaultHairColors() FactoryBuilder

	// BrownHairColor add brown hair color to the avatar factory.
	BrownHairColor() FactoryBuilder

	// LightHairColor add light hair color to the avatar factory.
	LightHairColor() FactoryBuilder

	// DarkHairColor add dark hair color to the avatar factory.
	DarkHairColor() FactoryBuilder

	// AddMouth adds a custom mouth to the avatar factory.
	AddMouth(isMale bool, name, shape string) FactoryBuilder

	// AddPalette adds a custom palette to the avatar factory.
	AddPalette(name string, palette SVGPalette) FactoryBuilder

	// DefaultPalettes add all predefined palettes to the avatar factory.
	DefaultPalettes() FactoryBuilder

	// PurplePalette add purple palette to the avatar factory.
	PurplePalette() FactoryBuilder

	// GreenPalette add green palette to the avatar factory.
	GreenPalette() FactoryBuilder

	// BluePalette add blue palette to the avatar factory.
	BluePalette() FactoryBuilder

	// YellowPalette add yellow palette to the avatar factory.
	YellowPalette() FactoryBuilder

	// OrangePalette add  orange palette to  avatarthe factory.
	OrangePalette() FactoryBuilder

	// RedPalette add red palette to the avatar factory.
	RedPalette() FactoryBuilder

	// TealPalette add teal palette to the avatar factory.
	TealPalette() FactoryBuilder

	// PinkPalette add pink palette to the avatar factory.
	PinkPalette() FactoryBuilder

	// AddShape adds a custom background shape to the avatar factory.
	// Shape can include fill="{shape}" to match the avatar palette.
	AddShape(name string, shape SVGShape) FactoryBuilder

	// DefaultShapes add all predefined background shapes to the avatar factory.
	DefaultShapes() FactoryBuilder

	// FillShape add square background shape to the avatar factory.
	FillShape() FactoryBuilder

	// CircleShape add circle background shape to the avatar factory.
	CircleShape() FactoryBuilder

	// PolygonShape add polygon background shape to the avatar factory.
	PolygonShape() FactoryBuilder

	// AddSkinColor adds a custom skin color to the avatar factory.
	AddSkinColor(name string, color SVGSkinColor) FactoryBuilder

	// DefaultSkinColors add all predefined skin colors to the avatar factory.
	DefaultSkinColors() FactoryBuilder

	// WhiteSkin add white skin color to the avatar factory.
	WhiteSkin() FactoryBuilder

	// BrownSkin add brown skin color to the avatar factory.
	BrownSkin() FactoryBuilder

	// BlackSkin add black skin color to the avatar factory.
	BlackSkin() FactoryBuilder

	// AddSticker adds a custom sticker to the avatar factory.
	// Shape can include fill="{text}" to match the avatar palette.
	AddSticker(name, shape string) FactoryBuilder

	// AddLetter adds a custom letter shape to the avatar factory.
	// Shape can include fill="{text}" to match the avatar palette.
	AddLetter(letter rune, shape string) FactoryBuilder

	// AddTransformer adds a custom letter transformer to the avatar factory.
	AddTransformer(letter, replacement rune) FactoryBuilder

	// PersianLetters add all predefined persian letters to the avatar factory.
	PersianLetters() FactoryBuilder

	// PersianTransformers add all predefined persian transformers to the avatar factory.
	PersianTransformers() FactoryBuilder

	// englishLetters add all predefined english letters to the avatar factory.
	englishLetters() FactoryBuilder

	// Build creates a new AvatarFactory instance.
	Build() AvatarFactory
}

// builder struct is used to build the avatar factory.
type builder struct {
	factory *factory
}

// New creates a new FactoryBuilder instance.
func New() FactoryBuilder {
	b := &builder{
		factory: &factory{
			shapes:   make(map[string]SVGShape),
			palettes: make(map[string]SVGPalette),

			stickers: make(map[string]string),

			letters:    make(map[string]string),
			transforms: make(map[string]string),

			skinColors: make(map[string]SVGSkinColor),
			hairColors: make(map[string]SVGHairColor),

			maleHairs:       make(map[string]string),
			maleBeards:      make(map[string]string),
			maleDresses:     make(map[string]string),
			maleEyes:        make(map[string]string),
			maleEyebrows:    make(map[string]string),
			maleMouths:      make(map[string]string),
			maleGlasses:     make(map[string]string),
			maleAccessories: make(map[string]string),

			femaleHairs:       make(map[string]string),
			femaleDresses:     make(map[string]string),
			femaleEyes:        make(map[string]string),
			femaleEyebrows:    make(map[string]string),
			femaleMouths:      make(map[string]string),
			femaleGlasses:     make(map[string]string),
			femaleAccessories: make(map[string]string),
		},
	}

	b.englishLetters()
	return b
}

func (b *builder) Build() AvatarFactory {
	return b.factory
}

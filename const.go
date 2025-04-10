package avatar

const None = "::None"

// Accessory defines preset accessory types for an avatar.
type Accessory string

const (
	AccessoryNecklace Accessory = "necklace"
	AccessoryChoker   Accessory = "choker"
)

// Beard defines preset beard types for an avatar.
type Beard string

const (
	BeardMustach      Beard = "mustach"
	BeardFancyMustach Beard = "fancy-mustach"
	BeardNormal       Beard = "beard"
	BeardMedium       Beard = "medium-beard"
	BeardLong         Beard = "long-beard"
)

// Dress defines preset dress types for an avatar.
type Dress string

const (
	DressSuit   Dress = "suit"
	DressShirt  Dress = "shirt"
	DressTShirt Dress = "t-shirt"
)

// Eye defines preset eye types for an avatar.
type Eye string

// Eyebrow defines preset eyebrow types for an avatar.
type Eyebrow string

// Glasses defines preset glasses types for an avatar.
type Glasses string

const (
	GlassesPrescription      Glasses = "prescription"
	GlassesRoundPrescription Glasses = "round-prescription"
	GlassesSunglass          Glasses = "sunglass"
	GlassesRoundSunglass     Glasses = "round-sunglass"
)

// Hair defines preset hair types for an avatar.
type Hair string

const (
	HairShort  Hair = "short"
	HairMedium Hair = "medium"
	HairWavy   Hair = "wavy"
	HairCurly  Hair = "curly"
)

// HairColor defines preset hair color types for an avatar.
type HairColor string

const (
	HairBrown HairColor = "brown"
	HairLight HairColor = "light"
	HairDark  HairColor = "dark"
)

// Mouth defines preset mouth types for an avatar.
type Mouth string

// Palette defines preset palette types for an avatar.
type Palette string

const (
	Purple Palette = "purple"
	Green  Palette = "green"
	Blue   Palette = "blue"
	Yellow Palette = "yellow"
	Orange Palette = "orange"
	Red    Palette = "red"
	Teal   Palette = "teal"
	Pink   Palette = "pink"
)

// Shape defines preset background shape types for an avatar.
type Shape string

const (
	ShapeFill    Shape = "fill"
	ShapeCircle  Shape = "circle"
	ShapePolygon Shape = "polygon"
)

// SkinColor defines preset skin color types for an avatar.
type SkinColor string

const (
	SkinWhite SkinColor = "white"
	SkinBrown SkinColor = "brown"
	SkinBlack SkinColor = "black"
)

// Shape defines preset sticker types for an avatar.
type Sticker string

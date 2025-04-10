package avatar

func (b *builder) AddSkinColor(name string, color SVGSkinColor) FactoryBuilder {
	b.factory.addSkinColor(name, color)
	return b
}

func (b *builder) DefaultSkinColors() FactoryBuilder {
	b.WhiteSkin()
	b.BrownSkin()
	b.BlackSkin()
	return b
}

func (b *builder) WhiteSkin() FactoryBuilder {
	b.factory.addSkinColor(
		string(SkinWhite),
		NewSkinColor(
			NewHex("#fbdad0"),
			NewHex("#84706d"),
		),
	)
	return b
}

func (b *builder) BrownSkin() FactoryBuilder {
	b.factory.addSkinColor(
		string(SkinBrown),
		NewSkinColor(
			NewHex("#ceaa82"),
			NewHex("#896f52"),
		),
	)
	return b
}

func (b *builder) BlackSkin() FactoryBuilder {
	b.factory.addSkinColor(
		string(SkinBrown),
		NewSkinColor(
			NewHex("#b06f51"),
			NewHex("#3d261d"),
		),
	)
	return b
}

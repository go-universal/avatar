package avatar

func (b *builder) AddPalette(name string, palette SVGPalette) FactoryBuilder {
	b.factory.addPalette(name, palette)
	return b
}

func (b *builder) DefaultPalettes() FactoryBuilder {
	b.PurplePalette()
	b.GreenPalette()
	b.BluePalette()
	b.YellowPalette()
	b.OrangePalette()
	b.RedPalette()
	b.TealPalette()
	b.PinkPalette()
	return b
}

func (b *builder) PurplePalette() FactoryBuilder {
	b.factory.addPalette(
		string(Purple),
		NewPalette(
			NewHex("#c0bade"),
			NewHex("#272145"),
			NewHex("#4e428a"),
			NewHex("#312956"),
			NewHex("#272145"),
		),
	)
	return b
}

func (b *builder) GreenPalette() FactoryBuilder {
	b.factory.addPalette(
		string(Green),
		NewPalette(
			NewHex("#88d8b0"),
			NewHex("#13281d"),
			NewHex("#008d68"),
			NewHex("#044432"),
			NewHex("#02140f"),
		),
	)
	return b
}

func (b *builder) BluePalette() FactoryBuilder {
	b.factory.addPalette(
		string(Blue),
		NewPalette(
			NewHex("#a2cced"),
			NewHex("#162835"),
			NewHex("#006fc8"),
			NewHex("#045389"),
			NewHex("#121b21"),
		),
	)
	return b
}

func (b *builder) YellowPalette() FactoryBuilder {
	b.factory.addPalette(
		string(Yellow),
		NewPalette(
			NewHex("#ffeead"),
			NewHex("#725d22"),
			NewHex("#ffd869"),
			NewHex("#ecc74d"),
			NewHex("#16130a"),
		),
	)
	return b
}

func (b *builder) OrangePalette() FactoryBuilder {
	b.factory.addPalette(
		string(Orange),
		NewPalette(
			NewHex("#ffcc5c"),
			NewHex("#993700"),
			NewHex("#f47521"),
			NewHex("#a5410a"),
			NewHex("#210c00"),
		),
	)
	return b
}

func (b *builder) RedPalette() FactoryBuilder {
	b.factory.addPalette(
		string(Red),
		NewPalette(
			NewHex("#ff6f69"),
			NewHex("#5b0505"),
			NewHex("#b22e25"),
			NewHex("#892016"),
			NewHex("#280c0b"),
		),
	)
	return b
}

func (b *builder) TealPalette() FactoryBuilder {
	b.factory.addPalette(
		string(Teal),
		NewPalette(
			NewHex("#88d0ca"),
			NewHex("#002621"),
			NewHex("#0f3630"),
			NewHex("#002621"),
			NewHex("#011915"),
		),
	)
	return b
}

func (b *builder) PinkPalette() FactoryBuilder {
	b.factory.addPalette(
		string(Pink),
		NewPalette(
			NewHex("#f58f97"),
			NewHex("#6b1520"),
			NewHex("#e83b46"),
			NewHex("#b82d37"),
			NewHex("#1c0407"),
		),
	)
	return b
}

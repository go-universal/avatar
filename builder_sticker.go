package avatar

func (b *builder) AddSticker(name, shape string) FactoryBuilder {
	b.factory.addSticker(name, shape)
	return b
}

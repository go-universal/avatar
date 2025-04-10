package avatar

func (b *builder) SetBodyShape(shape string) FactoryBuilder {
	b.factory.setBodyShape(shape)
	return b
}

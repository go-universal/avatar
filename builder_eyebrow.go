package avatar

func (b *builder) AddEyebrow(isMale bool, name, shape string) FactoryBuilder {
	b.factory.addEyebrow(isMale, name, shape)
	return b
}

package avatar

func (b *builder) AddMouth(isMale bool, name, shape string) FactoryBuilder {
	b.factory.addMouth(isMale, name, shape)
	return b
}

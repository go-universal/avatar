package avatar

func (b *builder) AddEye(isMale bool, name, shape string) FactoryBuilder {
	b.factory.addEye(isMale, name, shape)
	return b
}

package avatar

func (b *builder) AddBeard(name, shape string) FactoryBuilder {
	b.factory.addBeard(name, shape)
	return b
}

func (b *builder) DefaultBeards() FactoryBuilder {
	b.MustachBeard()
	b.FancyMustachBeard()
	b.NormalBeard()
	b.MediumBeard()
	b.LongBeard()
	return b
}

func (b *builder) MustachBeard() FactoryBuilder {
	b.factory.addBeard(
		string(BeardMustach),
		`<path fill="{hair_shadow}" d="M72.45,71.28c-.2-.29-.5-.48-.85-.54-1.75-.3-5.2-1.99-5.67-1.95-.29,.03-.87,.22-1.32,.38-.4,.14-.83,.14-1.23,0-.45-.16-1.02-.35-1.32-.38-.47-.04-3.92,1.65-5.67,1.95-.34,.06-.65,.25-.85,.54-.45,.64-1.09,1.85-.84,3.22,.38,2.03,.99-2.27,9.29-2.27s8.91,4.3,9.29,2.27c.25-1.36-.39-2.57-.84-3.22Z" />`,
	)
	return b
}

func (b *builder) FancyMustachBeard() FactoryBuilder {
	b.factory.addBeard(
		string(BeardFancyMustach),
		`<path fill="{hair_shadow}" d="M67.41,69.01c-1.34-.57-2.84,.17-3.41,1.38-.57-1.21-2.07-1.95-3.41-1.38-1.63,.69-4.38,5.23-8.68,6.4,4.17,.98,7.07,.72,8.68,0,1.78-.8,2.94-1.64,3.41-2.76,.47,1.12,1.63,1.96,3.41,2.76,1.61,.72,4.51,.98,8.68,0-4.31-1.18-7.06-5.71-8.68-6.4Z" />`,
	)
	return b
}

func (b *builder) NormalBeard() FactoryBuilder {
	b.factory.addBeard(
		string(BeardNormal),
		`<path fill="{hair_shadow}" d="M78.96,73.2c-.72,.93-2.43,2.55-4.22,4.11-.37,.32-1.08,.12-1.07-.3,.12-2.76-.41-6.04-1.8-6.11-1.9-.09-5.51-1.01-6.04-.79-.16,.07-.39,.14-.64,.21-.77,.22-1.62,.22-2.39,0-.25-.07-.48-.15-.64-.21-.53-.22-4.15,.7-6.04,.79-1.38,.06-1.91,3.35-1.8,6.11,.02,.42-.69,.63-1.07,.3-1.8-1.56-3.5-3.17-4.22-4.11-1.61-2.08-3.72-8.04-3.72-8.04,.07,1.43,1.73,9.44,3.42,11.62,1.67,2.15,8.41,5.82,9.86,6.29,1.44,.47,2.69,.8,5.4,.8s3.96-.33,5.4-.8c1.45-.47,8.19-4.14,9.86-6.29,1.69-2.18,3.52-10.15,3.59-11.58,0,0-2.27,5.92-3.89,8Zm-6.67,5.01c-.06,.79-.52,1.55-1.29,2.11-.77,.56-1.38,.95-1.72,1.06-1.29,.42-2.43-.92-4.63-1.17-.12-.13-.21-.27-.23-.44-.12-1.08,.12-2.01,.79-2.1,.67-.09,2.33-1.81-1.22-1.81s-1.9,1.72-1.22,1.81c.67,.09,.91,1.02,.79,2.1-.02,.17-.11,.31-.23,.44-2.2,.24-3.34,1.59-4.63,1.17-.33-.11-.95-.5-1.72-1.06-.77-.56-1.22-1.31-1.29-2.11-.19-2.34-.13-5.31,1.37-5.52,2.2-.31,4.22-.48,6.17-.09,.5,.1,1.02,.1,1.52,0,1.95-.39,3.97-.22,6.17,.09,1.49,.21,1.56,3.18,1.37,5.52Z" />`,
	)
	return b
}

func (b *builder) MediumBeard() FactoryBuilder {
	b.factory.addBeard(
		string(BeardMedium),
		`<path fill="{hair_shadow}" d="M78.96,70.2c-.72,.93-2.43,2.55-4.22,4.11-.37,.32-1.08,.12-1.07-.3,.12-2.76-.41-3.04-1.8-3.11-1.9-.09-5.51-1.01-6.04-.79-.16,.07-.39,.14-.64,.21-.77,.22-1.62,.22-2.39,0-.25-.07-.48-.15-.64-.21-.53-.22-4.15,.7-6.04,.79-1.38,.06-1.91,.35-1.8,3.11,.02,.42-.69,.63-1.07,.3-1.8-1.56-3.5-3.17-4.22-4.11-1.61-2.08-3.72-5.04-3.72-5.04,.07,1.43,1.73,9.44,3.42,11.62,1.67,2.15,8.41,5.82,9.86,6.29,1.44,.47,2.69,.8,5.4,.8s3.96-.33,5.4-.8c1.45-.47,8.19-4.14,9.86-6.29,1.69-2.18,3.52-10.15,3.59-11.58,0,0-2.27,2.92-3.89,5Zm-6.67,6.01c-.06,.79-.52,1.55-1.29,2.11-.77,.56-1.38,.95-1.72,1.06-1.29,.42-2.43,.08-4.63-.17-.12-.13-.21-.27-.23-.44-.12-1.08,.12-1.01,.79-1.1,.67-.09,2.33-1.81-1.22-1.81s-1.9,1.72-1.22,1.81c.67,.09,.91,.02,.79,1.1-.02,.17-.11,.31-.23,.44-2.2,.24-3.34,.59-4.63,.17-.33-.11-.95-.5-1.72-1.06-.77-.56-1.22-1.31-1.29-2.11-.19-2.34-.13-3.31,1.37-3.52,2.2-.31,4.22-.48,6.17-.09,.5,.1,1.02,.1,1.52,0,1.95-.39,3.97-.22,6.17,.09,1.49,.21,1.56,1.18,1.37,3.52Z" />`,
	)
	return b
}

func (b *builder) LongBeard() FactoryBuilder {
	b.factory.addBeard(
		string(BeardLong),
		`<path fill="{hair_shadow}" d="M78.96,72.2c-.72,.93-5.53,2.64-5.29,1.8-.12-2.27-.44-2.99-1.8-3.11-.97,.01-2.36-.16-3.53-.51-1.23-.08-2.26-.24-2.51-.28-.18,.22-.41,.34-.64,.21-.78,.42-1.62,.43-2.39,0-.24,.14-.47,.02-.64-.21-.25,.05-1.28,.22-2.51,.28-1.17,.37-2.56,.55-3.53,.51-1.35,.15-1.65,.93-1.8,3.11,.27,.92-4.56-.87-5.29-1.8-.81-1.04-1.73-2.8-2.46-4.3s-1.25-2.74-1.25-2.74c0,2.52,.49,5.2,1.38,7.8s3.35,7.63,4.86,7.33c.08,2.05,2.1,3.93,4.91,5.44,3.84,3.02,5.77,2.98,7.52,2.13,1.71,.55,3.54,.94,7.55-1.8,4.17-3.39,3.61-3.94,3.97-4.9,1.23,.41,5.01-4.51,5.92-7.21s1.41-5.67,1.41-8.76c0,0-2.27,4.92-3.89,7Zm-6.67,4.01c-.06,.79-.52,1.55-1.29,2.11-.77,.56-1.38,.95-1.72,1.06-1.29,.42-2.43,.08-4.63-.17-.12-.13-.21-.27-.23-.44-.12-1.08,.12-1.01,.79-1.1,.67-.09,2.33-1.81-1.22-1.81s-1.9,1.72-1.22,1.81c.67,.09,.91,.02,.79,1.1-.02,.17-.11,.31-.23,.44-2.2,.24-3.34,.59-4.63,.17-.33-.11-.95-.5-1.72-1.06-.77-.56-1.22-1.31-1.29-2.11-.19-2.34-.13-3.31,1.37-3.52,2.2-.31,4.22-.48,6.17-.09,.5,.1,1.02,.1,1.52,0,1.95-.39,3.97-.22,6.17,.09,1.49,.21,1.56,1.18,1.37,3.52Z" />`,
	)
	return b
}

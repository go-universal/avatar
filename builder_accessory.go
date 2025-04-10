package avatar

func (b *builder) AddAccessory(isMale bool, name, shape string) FactoryBuilder {
	b.factory.addAccessory(isMale, name, shape)
	return b
}

func (b *builder) DefaultAccessories() FactoryBuilder {
	b.NecklaceAccessory()
	b.ChokerAccessory()
	return b
}

func (b *builder) NecklaceAccessory() FactoryBuilder {
	b.factory.addAccessory(
		false,
		string(AccessoryNecklace),
		`<g>
			<rect fill="{decorator}" x="62.88" y="96.36" width="1.58" height="2.69" rx=".79" ry=".79" />
			<path d="M73.76,89.3c-1.83,4.11-5.68,6.95-10.11,6.95-4.21,0-7.88-2.55-9.81-6.32l-.66,.9c2.15,3.85,6.04,6.43,10.47,6.43s8.66-2.81,10.74-6.95l-.63-1Z" />
		</g>`,
	)
	return b
}

func (b *builder) ChokerAccessory() FactoryBuilder {
	b.factory.addAccessory(
		false,
		string(AccessoryChoker),
		`<path d="M54.98,86.01c.02,.54,.02,.89,.02,.89,0,.1-.04,.2-.05,.3h18.12c-.01-.3-.02-.72-.02-1.19h-18.07Z" />`,
	)
	return b
}

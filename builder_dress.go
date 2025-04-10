package avatar

func (b *builder) AddDress(isMale bool, name, shape string) FactoryBuilder {
	b.factory.addDress(isMale, name, shape)
	return b
}

func (b *builder) DefaultDresses() FactoryBuilder {
	b.SuitDress()
	b.ShirtDress()
	b.TShirtDress()
	return b
}

func (b *builder) SuitDress() FactoryBuilder {
	b.factory.addDress(
		true,
		string(DressSuit),
		`<g>
			<path fill="{dress}" d="M99.9,115.8c-1.1-8.2-7.5-13-14.3-17.1-3.1-1.8-5.8-3.8-8.1-5.6-.2,.3-.4,.6-.7,.9l-12.7,3-13.7-3c-.1-.2-.3-.3-.4-.5-2.1,1.7-4.7,3.5-7.6,5.3-6.6,4.1-13.1,9-14.3,17.1-1.1,8.2-1.4,12.1-1.4,12.1H101.3c0-.1-.2-4-1.4-12.2Z" />
			<polygon fill="{decorator}" points="66.3 102.2 66.3 102.2 67.6 100 66 95.5 64.1 96.1 62.2 95.5 60.6 100 61.9 102.2 56.8 128 71.4 128 66.3 102.2" />
			<polygon fill="{dress_shadow}" points="52.5 90.2 49.4 94.1 59 101.9 64.1 96.1 52.5 90.2" />
			<polygon fill="{dress_shadow}" points="75.7 90.2 78.8 94.1 69.2 101.9 64.1 96.1 75.7 90.2" />
		</g>`,
	)
	b.factory.addDress(
		false,
		string(DressSuit),
		`<g>
			<path fill="white" d="M62.3,103.5l-.7-4.5-11.8-7.2c-2.3,3.6-5.6,8.1-10.1,11.7h22.6Z" />
			<path fill="white" d="M88.1,103.5c-4.8-3.7-8.2-8.6-10.5-12.5l-11.7,8-.6,4.5h22.8Z" />
			<path fill="{dress}" d="M91.7,103.3c-2.5-2-6.6-4.6-10.9-7.4-3.3,2.2-9,3.7-15.5,3.9h-3c-6.4-.2-12-1.7-15.3-3.8-5,2.8-9.6,5.6-12.3,8.4-5.1,4.3-7.8,11.9-7.7,23.6H101c-.4-15.9-4.5-21.4-9.3-24.7Z" />
		</g>`,
	)
	return b
}

func (b *builder) ShirtDress() FactoryBuilder {
	b.factory.addDress(
		true,
		string(DressShirt),
		`<g>
			<path fill="{dress}" d="M101.05,128c-1.2-14.4-.3-20-15.7-29.1-5.7-3.4-8.9-6.9-11.3-10.8l-8.8,6-1.6,10.8-1.6-10.8-8.8-5.4c-3.7,5.6-6.6,7.8-12.5,11.3-12.6,7.4-12.8,14-13.8,28H101.05Z" />
			<path d="M67.65,98.4c0,.5-.4,.9-.9,.9s-.9-.4-.9-.9,.4-.9,.9-.9,.9,.4,.9,.9Z" />
			<path d="M67.05,103c0,.5-.4,.9-.9,.9s-.9-.4-.9-.9,.4-.9,.9-.9,.9,.4,.9,.9Z" />
		</g>`,
	)
	b.factory.addDress(
		false,
		string(DressSuit),
		`<g>
			<path fill="{dress}" d="M99.9,115.8c-1.1-8.2-7.5-13-14.3-17.1-3.1-1.8-5.8-3.8-8.1-5.6-.2,.3-.4,.6-.7,.9l-12.7,3-13.7-3c-.1-.2-.3-.3-.4-.5-2.1,1.7-4.7,3.5-7.6,5.3-6.6,4.1-13.1,9-14.3,17.1-1.1,8.2-1.4,12.1-1.4,12.1H101.3c0-.1-.2-4-1.4-12.2Z" />
			<polygon fill="{decorator}" points="66.3 102.2 66.3 102.2 67.6 100 66 95.5 64.1 96.1 62.2 95.5 60.6 100 61.9 102.2 56.8 128 71.4 128 66.3 102.2" />
			<polygon fill="{dress_shadow}" points="52.5 90.2 49.4 94.1 59 101.9 64.1 96.1 52.5 90.2" />
			<polygon fill="{dress_shadow}" points="75.7 90.2 78.8 94.1 69.2 101.9 64.1 96.1 75.7 90.2" />
		</g>`,
	)
	return b
}

func (b *builder) TShirtDress() FactoryBuilder {
	b.factory.addDress(
		true,
		string(DressTShirt),
		`<path fill="{dress}" d="M99.9,115.8c-1.1-8.2-7.5-13-14.3-17.1-3.1-1.8-5.8-3.8-8.1-5.6l-13.9,7.2-13.6-6.8c-2.1,1.7-4.7,3.5-7.6,5.3-6.6,4.1-13.1,9-14.3,17.1-1.1,8.2-1.4,12.1-1.4,12.1H101.3c0-.1-.2-4-1.4-12.2Z" />`,
	)
	b.factory.addDress(
		false,
		string(DressSuit),
		`<path fill="{dress}" d="M90.26,101.74c-17.48,6.01-34.9,6.05-52.26,0-4.86,3.55-8.97,7.89-9.89,14.16-1.1,8.2-1.4,12.1-1.4,12.1H101.3c0-.1-.2-4-1.4-12.2-.84-6.27-4.79-10.56-9.64-14.06Z" />`,
	)
	return b
}

package avatar

func (b *builder) AddShape(name string, shape SVGShape) FactoryBuilder {
	b.factory.addShape(name, shape)
	return b
}

func (b *builder) DefaultShapes() FactoryBuilder {
	b.FillShape()
	b.CircleShape()
	b.PolygonShape()
	return b
}

func (b *builder) FillShape() FactoryBuilder {
	b.factory.addShape(
		string(ShapeFill),
		NewShape(
			`<rect fill="{shape}" x="4" y="4" width="128" height="128"/>`,
			``,
		),
	)
	return b
}

func (b *builder) CircleShape() FactoryBuilder {
	b.factory.addShape(
		string(ShapeCircle),
		NewShape(
			`<circle fill="{shape}" cx="64" cy="64" r="60"/>`,
			`<path fill="white" d="M0,0V64H4c0,33.14,26.86,60,60,60s60-26.86,60-60h4V0H0Z"/>`,
		),
	)
	return b
}

func (b *builder) PolygonShape() FactoryBuilder {
	b.factory.addShape(
		string(ShapePolygon),
		NewShape(
			`<path fill="{shape}" d="M117.5,87.16V40.75c0-4.76-2.61-9.24-6.63-11.58L70.72,6.02c-4.2-2.33-9.24-2.33-13.35,0L17.22,29.27c-4.2,2.33-6.72,6.72-6.72,11.58v46.31c0,4.76,2.61,9.24,6.63,11.58l40.15,23.25c4.2,2.33,9.24,2.33,13.35,0l40.05-23.34c4.29-2.33,6.82-6.72,6.82-11.48Z"/>`,
			`<path fill="white" d="M0,0V64H10.5v23.16c0,4.76,2.61,9.24,6.63,11.58l40.15,23.25c4.2,2.33,9.24,2.33,13.35,0l40.05-23.34c4.29-2.33,6.82-6.72,6.82-11.48v-23.16h10.5V0H0Z"/>`,
		),
	)
	return b
}

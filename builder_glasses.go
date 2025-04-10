package avatar

func (b *builder) AddGlasses(isMale bool, name, shape string) FactoryBuilder {
	b.factory.addGlasses(isMale, name, shape)
	return b
}

func (b *builder) DefaultGlasses() FactoryBuilder {
	b.PrescriptionGlasses()
	b.RoundPrescriptionGlasses()
	b.SunglassGlasses()
	b.RoundSunglassGlasses()
	return b
}

func (b *builder) PrescriptionGlasses() FactoryBuilder {
	b.factory.addGlasses(
		true,
		string(GlassesPrescription),
		`<path fill="{decorator}" d="M81.2,55.86c-1.09-1.65-3.26-2.57-6.45-2.73h0s-.03,0-.03,0c-.33-.01-.66-.03-1.01-.03s-.68,.01-1,.03h-17.43c-.33-.01-.67-.03-1.02-.03-3.71,0-6.2,.92-7.42,2.74-1.06,1.58-1.08,3.74-.07,6.25,1.77,4.39,3.46,5.66,7.49,5.66s5.63-1.21,7.48-5.65c.67-1.61,.91-3.07,.72-4.34,.99-.41,2.09-.41,3.08,0-.2,1.26,.02,2.72,.67,4.33,1.77,4.39,3.45,5.66,7.49,5.66s5.63-1.21,7.48-5.65c1.05-2.51,1.05-4.66,0-6.24Zm-20.48,5.61c-1.69,4.03-2.9,4.96-6.46,4.96s-4.87-1.02-6.46-4.96c-.61-1.5-1.11-3.62-.04-5.21,.9-1.34,2.75-2.08,5.52-2.22h2.01c2.79,.15,4.66,.89,5.54,2.23,1.04,1.59,.52,3.7-.11,5.2Zm5.15-4.86c-1.21-.41-2.52-.41-3.73,0-.11-.26-.24-.51-.4-.75-.43-.66-1.04-1.19-1.8-1.62h8.12c-.76,.42-1.36,.95-1.79,1.6-.16,.24-.3,.5-.41,.76Zm14.29,4.86c-1.69,4.03-2.9,4.96-6.46,4.96s-4.87-1.02-6.46-4.96c-.61-1.5-1.11-3.62-.04-5.21,.89-1.34,2.74-2.08,5.49-2.22h2.03c2.79,.15,4.66,.89,5.54,2.23,1.05,1.59,.52,3.7-.11,5.2Z" />`,
	)
	b.factory.addGlasses(
		false,
		string(GlassesPrescription),
		`<path d="M84.38,54.28c-1.3-1.96-3.86-3.05-7.64-3.24h0s-.04,0-.04,0c-.39-.02-.79-.03-1.2-.03s-.8,.01-1.19,.03h-20.65c-.39-.02-.79-.03-1.2-.03-4.39,0-7.35,1.09-8.8,3.25-1.25,1.87-1.28,4.43-.08,7.41,2.1,5.2,4.09,6.7,8.87,6.7s6.67-1.44,8.87-6.69c.8-1.91,1.08-3.64,.86-5.14,1.17-.48,2.48-.48,3.65,0-.24,1.49,.03,3.23,.8,5.13,2.1,5.2,4.09,6.7,8.87,6.7s6.67-1.44,8.87-6.69c1.24-2.97,1.25-5.53,.01-7.4Zm-24.26,6.65c-2,4.78-3.43,5.88-7.66,5.88s-5.77-1.21-7.65-5.87c-.72-1.78-1.31-4.29-.05-6.17,1.06-1.59,3.26-2.47,6.54-2.64h2.38c3.31,.17,5.52,1.06,6.57,2.65,1.24,1.88,.62,4.38-.13,6.16Zm6.11-5.76c-1.43-.49-2.99-.49-4.42,0-.13-.31-.28-.6-.47-.88-.51-.78-1.23-1.41-2.14-1.91h9.63c-.9,.5-1.61,1.13-2.12,1.9-.19,.28-.35,.59-.48,.9Zm16.93,5.76c-2,4.78-3.43,5.88-7.65,5.88s-5.76-1.21-7.65-5.87c-.72-1.78-1.31-4.29-.05-6.17,1.06-1.58,3.25-2.46,6.51-2.64h2.41c3.31,.17,5.52,1.06,6.57,2.65,1.24,1.88,.62,4.38-.13,6.16Z" />`,
	)
	return b
}

func (b *builder) RoundPrescriptionGlasses() FactoryBuilder {
	b.factory.addGlasses(
		true,
		string(GlassesRoundPrescription),
		`<g>
			<path fill="white" opacity="0.25" d="M53.86,51.77l7.61,10.33s-2.29,4.39-5.95,3.9l-7.68-10.07s2.08-4.43,6.02-4.15Z" />
			<path fill="white" opacity="0.25" d="M72.57,51.77l7.61,10.33s-2.29,4.39-5.95,3.9l-7.68-10.07s2.08-4.43,6.02-4.15Z" />
			<path fill="{decorator}" d="M73.36,51.18c-4.12,0-7.5,3.24-7.74,7.3-1.26-.4-2.6-.15-3.26,.03-.22-4.08-3.6-7.33-7.74-7.33s-7.76,3.48-7.76,7.76,3.48,7.76,7.76,7.76,7.62-3.35,7.75-7.52c.51-.15,1.96-.49,3.24,0,.12,4.17,3.55,7.53,7.75,7.53s7.76-3.48,7.76-7.76-3.48-7.76-7.76-7.76Zm-18.73,14.88c-3.92,0-7.12-3.19-7.12-7.12s3.19-7.12,7.12-7.12,7.12,3.19,7.12,7.12-3.19,7.12-7.12,7.12Zm18.73,0c-3.92,0-7.12-3.19-7.12-7.12s3.19-7.12,7.12-7.12,7.12,3.19,7.12,7.12-3.19,7.12-7.12,7.12Z" />
		</g>`,
	)
	b.factory.addGlasses(
		false,
		string(GlassesRoundPrescription),
		`<g>
			<path fill="white" opacity="0.25" d="M52.15,49.82l8.89,12.07s-2.68,5.13-6.95,4.56l-8.98-11.78s2.43-5.18,7.04-4.85Z" />
			<path fill="white" opacity="0.25" d="M74.02,49.82l8.89,12.07s-2.68,5.13-6.95,4.56l-8.98-11.78s2.43-5.18,7.04-4.85Z" />
			<path fill="{decorator}" d="M74.95,49.14c-4.82,0-8.77,3.79-9.04,8.54-1.48-.47-3.04-.18-3.81,.03-.26-4.77-4.21-8.57-9.05-8.57s-9.07,4.07-9.07,9.07,4.07,9.07,9.07,9.07,8.91-3.92,9.06-8.79c.59-.18,2.3-.58,3.78,0,.15,4.88,4.15,8.8,9.06,8.8s9.07-4.07,9.07-9.07-4.07-9.07-9.07-9.07Zm-21.9,17.39c-4.59,0-8.32-3.73-8.32-8.32s3.73-8.32,8.32-8.32,8.32,3.73,8.32,8.32-3.73,8.32-8.32,8.32Zm21.9,0c-4.59,0-8.32-3.73-8.32-8.32s3.73-8.32,8.32-8.32,8.32,3.73,8.32,8.32-3.73,8.32-8.32,8.32Z" />
		</g>`,
	)
	return b
}

func (b *builder) SunglassGlasses() FactoryBuilder {
	b.factory.addGlasses(
		true,
		string(GlassesSunglass),
		`<path opacity="0.85" d="M80.85,54.97c-1.07-1.61-3.2-2.51-6.32-2.66h-.03c-.32-.02-.65-.03-1-.03s-.66,.01-.98,.03h-17.08c-.32-.02-.65-.03-.99-.03-3.63,0-6.08,.9-7.27,2.68-1.04,1.54-1.06,3.65-.06,6.09,1.74,4.28,3.39,5.51,7.34,5.51s5.52-1.18,7.34-5.51c.66-1.57,.89-3,.71-4.23,.96-.4,2.05-.4,3.02,0-.2,1.23,.02,2.66,.66,4.22,1.74,4.28,3.38,5.51,7.34,5.51s5.52-1.18,7.34-5.51c1.03-2.44,1.03-4.55,0-6.09Zm-15.02,.73c-1.18-.4-2.47-.4-3.65,0-.11-.25-.24-.5-.39-.73-.43-.64-1.02-1.16-1.77-1.58h7.96c-.74,.41-1.33,.93-1.75,1.56-.16,.23-.29,.48-.4,.74Z" />`,
	)
	b.factory.addGlasses(
		false,
		string(GlassesSunglass),
		`<path opacity="0.85" d="M83.88,53.44c-1.27-1.9-3.77-2.96-7.45-3.14h-.04c-.38-.02-.77-.03-1.17-.03s-.78,.01-1.16,.03h-20.15c-.38-.02-.77-.03-1.17-.03-4.29,0-7.17,1.06-8.58,3.16-1.22,1.82-1.25,4.31-.08,7.19,2.05,5.05,4,6.51,8.66,6.51s6.51-1.4,8.66-6.5c.78-1.85,1.05-3.54,.83-4.99,1.14-.47,2.42-.47,3.56,0-.23,1.45,.02,3.13,.78,4.98,2.05,5.05,3.99,6.51,8.66,6.51s6.51-1.4,8.66-6.5c1.21-2.88,1.21-5.36,0-7.18Zm-17.72,.86c-1.4-.47-2.92-.47-4.31,0-.13-.3-.28-.59-.46-.86-.5-.75-1.2-1.37-2.09-1.86h9.39c-.87,.48-1.57,1.1-2.07,1.84-.19,.28-.34,.57-.47,.88Z" />`,
	)
	return b
}

func (b *builder) RoundSunglassGlasses() FactoryBuilder {
	b.factory.addGlasses(
		true,
		string(GlassesRoundSunglass),
		`<path opacity="0.85" d="M73.57,51.92c-3.1,0-5.7,2.1-6.48,4.95-.72-.87-1.82-1.43-3.04-1.43s-2.38,.59-3.1,1.51c-.76-2.89-3.38-5.03-6.5-5.03-3.72,0-6.73,3.01-6.73,6.73s3.01,6.73,6.73,6.73c3.46,0,6.31-2.61,6.68-5.98h.03c0-1.59,1.3-2.89,2.89-2.89,1.36,0,2.49,.94,2.8,2.21,.04,3.68,3.03,6.66,6.73,6.66s6.73-3.01,6.73-6.73-3.01-6.73-6.73-6.73Z" />`,
	)
	b.factory.addGlasses(
		false,
		string(GlassesRoundSunglass),
		`<path opacity="0.85" d="M75.52,49.82c-3.73,0-6.87,2.53-7.81,5.96-.87-1.05-2.19-1.72-3.66-1.72s-2.86,.71-3.74,1.81c-.91-3.48-4.06-6.05-7.83-6.05-4.48,0-8.1,3.63-8.1,8.1s3.63,8.11,8.1,8.11c4.17,0,7.59-3.15,8.05-7.2h.04c0-1.92,1.56-3.48,3.48-3.48,1.63,0,3,1.14,3.37,2.66,.05,4.43,3.65,8.01,8.1,8.01s8.11-3.63,8.11-8.11-3.63-8.1-8.11-8.1Z" />`,
	)
	return b
}

package avatar

func (f *factory) Shapes() []string {
	f.mutex.RLock()
	defer f.mutex.RUnlock()
	return mapKeys(f.shapes)
}

func (f *factory) Palettes() []string {
	f.mutex.RLock()
	defer f.mutex.RUnlock()
	return mapKeys(f.palettes)
}

func (f *factory) Stickers() []string {
	f.mutex.RLock()
	defer f.mutex.RUnlock()
	return mapKeys(f.stickers)
}

func (f *factory) SkinColors() []string {
	f.mutex.RLock()
	defer f.mutex.RUnlock()
	return mapKeys(f.skinColors)
}

func (f *factory) HairColors() []string {
	f.mutex.RLock()
	defer f.mutex.RUnlock()
	return mapKeys(f.hairColors)
}

func (f *factory) MaleHairs() []string {
	f.mutex.RLock()
	defer f.mutex.RUnlock()
	return mapKeys(f.maleHairs)
}

func (f *factory) MaleBeards() []string {
	f.mutex.RLock()
	defer f.mutex.RUnlock()
	return mapKeys(f.maleBeards)
}

func (f *factory) MaleDresses() []string {
	f.mutex.RLock()
	defer f.mutex.RUnlock()
	return mapKeys(f.maleDresses)
}

func (f *factory) MaleEyes() []string {
	f.mutex.RLock()
	defer f.mutex.RUnlock()
	return mapKeys(f.maleEyes)
}

func (f *factory) MaleEyebrows() []string {
	f.mutex.RLock()
	defer f.mutex.RUnlock()
	return mapKeys(f.maleEyebrows)
}

func (f *factory) MaleMouths() []string {
	f.mutex.RLock()
	defer f.mutex.RUnlock()
	return mapKeys(f.maleMouths)
}

func (f *factory) MaleGlasses() []string {
	f.mutex.RLock()
	defer f.mutex.RUnlock()
	return mapKeys(f.maleGlasses)
}

func (f *factory) MaleAccessories() []string {
	f.mutex.RLock()
	defer f.mutex.RUnlock()
	return mapKeys(f.maleAccessories)
}

func (f *factory) FemaleHairs() []string {
	f.mutex.RLock()
	defer f.mutex.RUnlock()
	return mapKeys(f.femaleHairs)
}

func (f *factory) FemaleDresses() []string {
	f.mutex.RLock()
	defer f.mutex.RUnlock()
	return mapKeys(f.femaleDresses)
}

func (f *factory) FemaleEyes() []string {
	f.mutex.RLock()
	defer f.mutex.RUnlock()
	return mapKeys(f.femaleEyes)
}

func (f *factory) FemaleEyebrows() []string {
	f.mutex.RLock()
	defer f.mutex.RUnlock()
	return mapKeys(f.femaleEyebrows)
}

func (f *factory) FemaleMouths() []string {
	f.mutex.RLock()
	defer f.mutex.RUnlock()
	return mapKeys(f.femaleMouths)
}

func (f *factory) FemaleGlasses() []string {
	f.mutex.RLock()
	defer f.mutex.RUnlock()
	return mapKeys(f.femaleGlasses)
}

func (f *factory) FemaleAccessories() []string {
	f.mutex.RLock()
	defer f.mutex.RUnlock()
	return mapKeys(f.femaleAccessories)
}

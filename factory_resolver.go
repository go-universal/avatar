package avatar

func (f *factory) resolveShape(k string) SVGShape {
	f.mutex.RLock()
	defer f.mutex.RUnlock()
	if v, ok := f.shapes[k]; ok && !isNilOrZero(v) {
		return v
	}
	return fallbackShape()
}

func (f *factory) resolvePalette(k string) SVGPalette {
	f.mutex.RLock()
	defer f.mutex.RUnlock()
	if v, ok := f.palettes[k]; ok && !isNilOrZero(v) {
		return v
	}
	return fallbackPalette()
}

func (f *factory) resolveSticker(k string) string {
	f.mutex.RLock()
	defer f.mutex.RUnlock()
	if v, ok := f.stickers[k]; ok && !isNilOrZero(v) {
		return v
	}
	return fallbackSticker()
}

func (f *factory) resolveLetter(k string) string {
	f.mutex.RLock()
	defer f.mutex.RUnlock()
	if v, ok := f.letters[k]; ok && !isNilOrZero(v) {
		return v
	}
	return fallbackLetter()
}

func (f *factory) resolveBodyShape() string {
	f.mutex.RLock()
	defer f.mutex.RUnlock()
	if !isNilOrZero(f.bodyShape) {
		return f.bodyShape
	}
	return fallbackBodyShape()
}

func (f *factory) resolveSkinColor(k string) SVGSkinColor {
	f.mutex.RLock()
	defer f.mutex.RUnlock()
	if v, ok := f.skinColors[k]; ok && !isNilOrZero(v) {
		return v
	}
	return fallbackSkinColor()
}

func (f *factory) resolveHairColor(k string) SVGHairColor {
	f.mutex.RLock()
	defer f.mutex.RUnlock()
	if v, ok := f.hairColors[k]; ok && !isNilOrZero(v) {
		return v
	}
	return fallbackHairColor()
}

func (f *factory) resolveHair(isMale bool, k string) string {
	f.mutex.RLock()
	defer f.mutex.RUnlock()
	if isMale {
		if v, ok := f.maleHairs[k]; ok && !isNilOrZero(v) {
			return v
		}
	} else {
		if v, ok := f.femaleHairs[k]; ok && !isNilOrZero(v) {
			return v
		}
	}
	return fallbackHair(isMale)
}

func (f *factory) resolveBeard(isMale bool, k string) string {
	if !isMale {
		return ""
	}

	f.mutex.RLock()
	defer f.mutex.RUnlock()
	if v, ok := f.maleBeards[k]; ok && !isNilOrZero(v) {
		return v
	}
	return fallbackBeard()
}

func (f *factory) resolveDress(isMale bool, k string) string {
	f.mutex.RLock()
	defer f.mutex.RUnlock()
	if isMale {
		if v, ok := f.maleDresses[k]; ok && !isNilOrZero(v) {
			return v
		}
	} else {
		if v, ok := f.femaleDresses[k]; ok && !isNilOrZero(v) {
			return v
		}
	}
	return fallbackDress(isMale)
}

func (f *factory) resolveEye(isMale bool, k string) string {
	f.mutex.RLock()
	defer f.mutex.RUnlock()
	if isMale {
		if v, ok := f.maleEyes[k]; ok && !isNilOrZero(v) {
			return v
		}
	} else {
		if v, ok := f.femaleEyes[k]; ok && !isNilOrZero(v) {
			return v
		}
	}
	return fallbackEye(isMale)
}

func (f *factory) resolveEyebrow(isMale bool, k string) string {
	f.mutex.RLock()
	defer f.mutex.RUnlock()
	if isMale {
		if v, ok := f.maleEyebrows[k]; ok && !isNilOrZero(v) {
			return v
		}
	} else {
		if v, ok := f.femaleEyebrows[k]; ok && !isNilOrZero(v) {
			return v
		}
	}
	return fallbackEyebrow(isMale)
}

func (f *factory) resolveMouth(isMale bool, k string) string {
	f.mutex.RLock()
	defer f.mutex.RUnlock()
	if isMale {
		if v, ok := f.maleMouths[k]; ok && !isNilOrZero(v) {
			return v
		}
	} else {
		if v, ok := f.femaleMouths[k]; ok && !isNilOrZero(v) {
			return v
		}
	}
	return fallbackMouth(isMale)
}

func (f *factory) resolveGlasses(isMale bool, k string) string {
	f.mutex.RLock()
	defer f.mutex.RUnlock()
	if isMale {
		if v, ok := f.maleGlasses[k]; ok && !isNilOrZero(v) {
			return v
		}
	} else {
		if v, ok := f.femaleGlasses[k]; ok && !isNilOrZero(v) {
			return v
		}
	}
	return fallbackGlasses(isMale)
}

func (f *factory) resolveAccessory(isMale bool, k string) string {
	f.mutex.RLock()
	defer f.mutex.RUnlock()
	if isMale {
		if v, ok := f.maleAccessories[k]; ok && !isNilOrZero(v) {
			return v
		}
	} else {
		if v, ok := f.femaleAccessories[k]; ok && !isNilOrZero(v) {
			return v
		}
	}
	return fallbackAccessory(isMale)
}

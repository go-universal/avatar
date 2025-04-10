package avatar

import "strings"

func (f *factory) addShape(name string, shape SVGShape) {
	name = strings.TrimSpace(name)
	if name == "" || shape == nil {
		return
	}

	f.mutex.Lock()
	defer f.mutex.Unlock()
	f.shapes[name] = shape
}

func (f *factory) addPalette(name string, palette SVGPalette) {
	name = strings.TrimSpace(name)
	if name == "" || palette == nil {
		return
	}

	f.mutex.Lock()
	defer f.mutex.Unlock()
	f.palettes[name] = palette
}

func (f *factory) addSticker(name, shape string) {
	name = strings.TrimSpace(name)
	shape = strings.TrimSpace(shape)
	if name == "" || shape == "" {
		return
	}

	f.mutex.Lock()
	defer f.mutex.Unlock()
	f.stickers[name] = shape
}

func (f *factory) addLetter(letter, shape string) {
	letter = strings.TrimSpace(letter)
	shape = strings.TrimSpace(shape)
	if letter == "" || shape == "" {
		return
	}

	f.mutex.Lock()
	defer f.mutex.Unlock()
	f.letters[letter] = shape
}

func (f *factory) addTransformer(letter, replacement string) {
	letter = strings.TrimSpace(letter)
	replacement = strings.TrimSpace(replacement)
	if letter == "" || replacement == "" {
		return
	}

	f.mutex.Lock()
	defer f.mutex.Unlock()
	f.transforms[letter] = strings.ToLower(replacement)
}

func (f *factory) setBodyShape(shape string) {
	shape = strings.TrimSpace(shape)
	if shape == "" {
		return
	}

	f.mutex.Lock()
	defer f.mutex.Unlock()
	f.bodyShape = shape
}

func (f *factory) addSkinColor(name string, color SVGSkinColor) {
	name = strings.TrimSpace(name)
	if name == "" || color == nil {
		return
	}

	f.mutex.Lock()
	defer f.mutex.Unlock()
	f.skinColors[name] = color
}

func (f *factory) addHairColor(name string, color SVGHairColor) {
	name = strings.TrimSpace(name)
	if name == "" || color == nil {
		return
	}

	f.mutex.Lock()
	defer f.mutex.Unlock()
	f.hairColors[name] = color
}

func (f *factory) addHair(isMale bool, name, shape string) {
	name = strings.TrimSpace(name)
	shape = strings.TrimSpace(shape)
	if name == "" || shape == "" {
		return
	}

	f.mutex.Lock()
	defer f.mutex.Unlock()
	if isMale {
		f.maleHairs[name] = shape
	} else {
		f.femaleHairs[name] = shape
	}
}

func (f *factory) addBeard(name, shape string) {
	name = strings.TrimSpace(name)
	shape = strings.TrimSpace(shape)
	if name == "" || shape == "" {
		return
	}

	f.mutex.Lock()
	defer f.mutex.Unlock()
	f.maleBeards[name] = shape
}

func (f *factory) addDress(isMale bool, name, shape string) {
	name = strings.TrimSpace(name)
	shape = strings.TrimSpace(shape)
	if name == "" || shape == "" {
		return
	}

	f.mutex.Lock()
	defer f.mutex.Unlock()
	if isMale {
		f.maleDresses[name] = shape
	} else {
		f.femaleDresses[name] = shape
	}
}

func (f *factory) addEye(isMale bool, name, shape string) {
	name = strings.TrimSpace(name)
	shape = strings.TrimSpace(shape)
	if name == "" || shape == "" {
		return
	}

	f.mutex.Lock()
	defer f.mutex.Unlock()
	if isMale {
		f.maleEyes[name] = shape
	} else {
		f.femaleEyes[name] = shape
	}
}

func (f *factory) addEyebrow(isMale bool, name, shape string) {
	name = strings.TrimSpace(name)
	shape = strings.TrimSpace(shape)
	if name == "" || shape == "" {
		return
	}

	f.mutex.Lock()
	defer f.mutex.Unlock()
	if isMale {
		f.maleEyebrows[name] = shape
	} else {
		f.femaleEyebrows[name] = shape
	}
}

func (f *factory) addMouth(isMale bool, name, shape string) {
	name = strings.TrimSpace(name)
	shape = strings.TrimSpace(shape)
	if name == "" || shape == "" {
		return
	}

	f.mutex.Lock()
	defer f.mutex.Unlock()
	if isMale {
		f.maleMouths[name] = shape
	} else {
		f.femaleMouths[name] = shape
	}
}

func (f *factory) addGlasses(isMale bool, name, shape string) {
	name = strings.TrimSpace(name)
	shape = strings.TrimSpace(shape)
	if name == "" || shape == "" {
		return
	}

	f.mutex.Lock()
	defer f.mutex.Unlock()
	if isMale {
		f.maleGlasses[name] = shape
	} else {
		f.femaleGlasses[name] = shape
	}
}

func (f *factory) addAccessory(isMale bool, name, shape string) {
	name = strings.TrimSpace(name)
	shape = strings.TrimSpace(shape)
	if name == "" || shape == "" {
		return
	}

	// Safe add to factory
	f.mutex.Lock()
	defer f.mutex.Unlock()
	if isMale {
		f.maleAccessories[name] = shape
	} else {
		f.femaleAccessories[name] = shape
	}
}

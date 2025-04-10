package avatar

import "slices"

func (f *factory) randomItem(items []string, only ...string) string {
	if len(only) > 0 {
		items = filter(items, func(s string) bool { return slices.Contains(only, s) })
	}

	if len(items) == 1 {
		return items[0]
	} else if len(items) > 1 {
		return items[randomize(len(items))]
	}

	return ""
}

func (f *factory) randomShape(only ...string) string {
	return f.randomItem(f.Shapes(), only...)
}

func (f *factory) randomPalette(only ...string) string {
	return f.randomItem(f.Palettes(), only...)
}

func (f *factory) randomSticker(only ...string) string {
	return f.randomItem(f.Stickers(), only...)
}

func (f *factory) randomSkinColor(only ...string) string {
	return f.randomItem(f.SkinColors(), only...)
}

func (f *factory) randomHairColor(only ...string) string {
	return f.randomItem(f.HairColors(), only...)
}

func (f *factory) randomHair(isMale bool, only ...string) string {
	if isMale {
		return f.randomItem(f.MaleHairs(), only...)
	} else {
		return f.randomItem(f.FemaleHairs(), only...)
	}
}

func (f *factory) randomBeard(isMale bool, only ...string) string {
	if isMale {
		return f.randomItem(f.MaleBeards(), only...)
	} else {
		return ""
	}
}

func (f *factory) randomDress(isMale bool, only ...string) string {
	if isMale {
		return f.randomItem(f.MaleDresses(), only...)
	} else {
		return f.randomItem(f.FemaleDresses(), only...)
	}
}

func (f *factory) randomEye(isMale bool, only ...string) string {
	if isMale {
		return f.randomItem(f.MaleEyes(), only...)
	} else {
		return f.randomItem(f.FemaleEyes(), only...)
	}
}

func (f *factory) randomEyebrow(isMale bool, only ...string) string {
	if isMale {
		return f.randomItem(f.MaleEyebrows(), only...)
	} else {
		return f.randomItem(f.FemaleEyebrows(), only...)
	}
}

func (f *factory) randomMouth(isMale bool, only ...string) string {
	if isMale {
		return f.randomItem(f.MaleMouths(), only...)
	} else {
		return f.randomItem(f.FemaleMouths(), only...)
	}
}

func (f *factory) randomGlasses(isMale bool, only ...string) string {
	if isMale {
		return f.randomItem(f.MaleGlasses(), only...)
	} else {
		return f.randomItem(f.FemaleGlasses(), only...)
	}
}

func (f *factory) randomAccessory(isMale bool, only ...string) string {
	if isMale {
		return f.randomItem(f.MaleAccessories(), only...)
	} else {
		return f.randomItem(f.FemaleAccessories(), only...)
	}
}

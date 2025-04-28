package avatar

func (f *factory) NewMale(name string) PersonAvatar {
	return newPersonAvatar(f, true, name)
}

func (f *factory) NewFemale(name string) PersonAvatar {
	return newPersonAvatar(f, false, name)
}

func (f *factory) NewPerson(isMale bool, name string) PersonAvatar {
	return newPersonAvatar(f, isMale, name)
}

func (f *factory) NewText(title, name string) TextAvatar {
	return newLetterAvatar(f, title, name)
}

func (f *factory) NewSticker(sticker Sticker, name string) StickerAvatar {
	return newStickerAvatar(f, string(sticker), name)
}

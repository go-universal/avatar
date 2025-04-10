package avatar

func (f *factory) NewMale() PersonAvatar {
	return newPersonAvatar(f, true)
}

func (f *factory) NewFemale() PersonAvatar {
	return newPersonAvatar(f, false)
}

func (f *factory) NewPerson(isMale bool) PersonAvatar {
	return newPersonAvatar(f, isMale)
}

func (f *factory) NewText(name string) TextAvatar {
	return newLetterAvatar(f, name)
}

func (f *factory) NewSticker(sticker Sticker) StickerAvatar {
	return newStickerAvatar(f, string(sticker))
}

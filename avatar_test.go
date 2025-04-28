package avatar_test

import (
	"testing"

	"github.com/go-universal/avatar"
	"github.com/stretchr/testify/require"
)

func TestTextAvatar(t *testing.T) {
	factory := avatar.New().
		DefaultPalettes().
		CircleShape().
		Build()

	johnDoe := factory.NewText("John Doe", "")

	johnDoe.
		RandomizeShape().
		RandomizePalette()

	err := johnDoe.SaveAs("demo/john_doe.svg")
	require.NoError(t, err, "failed to save text avatar")
}

func TestStickerAvatar(t *testing.T) {
	factory := avatar.New().
		DefaultPalettes().
		CircleShape().
		Build()

	star := factory.NewSticker("", "")

	star.
		RandomizeShape().
		RandomizePalette()

	err := star.SaveAs("demo/star.svg")
	require.NoError(t, err, "failed to save sticker avatar")
}

func TestMaleAvatar(t *testing.T) {
	factory := avatar.New().
		DefaultAccessories().
		DefaultBeards().
		SuitDress().
		PrescriptionGlasses().
		DarkHairColor().
		DefaultHairs().
		DefaultPalettes().
		CircleShape().
		WhiteSkin().
		Build()

	male := factory.NewMale("")

	male.
		RandomizeShape().
		RandomizePalette().
		RandomizeSkinColor().
		RandomizeHairColor().
		RandomizeHair().
		RandomizeBeard().
		RandomizeDress().
		RandomizeEye().
		RandomizeEyebrow().
		RandomizeMouth().
		RandomizeGlasses().
		RandomizeAccessory()

	err := male.SaveAs("demo/male.svg")
	require.NoError(t, err, "failed to save male avatar")
}

func TestFemaleAvatar(t *testing.T) {
	factory := avatar.New().
		DefaultAccessories().
		DefaultBeards().
		DefaultDresses().
		RoundPrescriptionGlasses().
		// DefaultHairColors(). // uncomment if needed
		DefaultHairs().
		DefaultPalettes().
		CircleShape().
		WhiteSkin().
		Build()

	female := factory.NewFemale("")

	female.
		RandomizeShape().
		RandomizePalette().
		RandomizeSkinColor().
		RandomizeHairColor().
		RandomizeHair().
		RandomizeBeard().
		RandomizeDress().
		RandomizeEye().
		RandomizeEyebrow().
		RandomizeMouth().
		RandomizeGlasses().
		RandomizeAccessory()

	err := female.SaveAs("demo/female.svg")
	require.NoError(t, err, "failed to save female avatar")
}

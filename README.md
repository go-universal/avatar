# Avatar

![GitHub Tag](https://img.shields.io/github/v/tag/go-universal/avatar?sort=semver&label=version)
[![Go Reference](https://pkg.go.dev/badge/github.com/go-universal/avatar.svg)](https://pkg.go.dev/github.com/go-universal/avatar)
[![License](https://img.shields.io/badge/license-ISC-blue.svg)](https://github.com/go-universal/avatar/blob/main/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-universal/avatar)](https://goreportcard.com/report/github.com/go-universal/avatar)
![Contributors](https://img.shields.io/github/contributors/go-universal/avatar)
![Issues](https://img.shields.io/github/issues/go-universal/avatar)

`avatar` is a Go library for generating customizable avatars. It supports creating avatars for text, stickers, and persons with various customization options.

| Letter                                                                                   | Sticker                                                                               | Male                                                                               | Female                                                                                 |
| ---------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| ![Letter](https://github.com/go-universal/avatar/blob/master/demo/john_doe.svg?raw=true) | ![Sticker](https://github.com/go-universal/avatar/blob/master/demo/star.svg?raw=true) | ![Male](https://github.com/go-universal/avatar/blob/master/demo/male.svg?raw=true) | ![Female](https://github.com/go-universal/avatar/blob/master/demo/female.svg?raw=true) |

## Installation

To install Avatar, use `go get`:

```sh
go get github.com/go-universal/avatar
```

## Usage

### Creating a Text Avatar

```go
package main

import (
    "github.com/go-universal/avatar"
    "log"
)

func main() {
    factory := avatar.
        New().
        DefaultPalettes().
        CircleShape().
        Build()

    johnDoe := factory.NewText("John Doe", "")

    johnDoe.
        RandomizeShape().
        RandomizePalette()

    err := johnDoe.SaveAs("john_doe.svg")
    if err != nil {
        log.Fatal(err)
    }
}
```

### Creating a Sticker Avatar

```go
package main

import (
    "github.com/go-universal/avatar"
    "log"
)

func main() {
    factory := avatar.
        New().
        DefaultPalettes().
        CircleShape().
        Build()

    star := factory.NewSticker("", "")

    star.
        RandomizeShape().
        RandomizePalette()

    err := star.SaveAs("star.svg")
    if err != nil {
        log.Fatal(err)
    }
}
```

### Creating a Person Avatar

```go
package main

import (
    "github.com/go-universal/avatar"
    "log"
)

func main() {
    factory := avatar.
        New().
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
    err := male.SaveAs("male.svg")
    if err != nil {
        log.Fatal(err)
    }

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
    err := female.SaveAs("female.svg")
    if err != nil {
        log.Fatal(err)
    }
}
```

## Extending

Avatar generation in this package is based on modular SVG elements. You can easily customize the existing avatar styles or add new ones using the provided helper methods.

To create or modify styles:

- Start by editing the [template.ai](https://github.com/go-universal/avatar/blob/master/template.ai) file.
- Save your design as an SVG.
- Extract the relevant parts (e.g., shapes, hair, accessories) and integrate them into the avatar package.

Each SVG shape must include placeholders to enable dynamic replacement of colors in the final rendering.

### Template Variables for Custom Palettes

Use the following template variables in your SVG elements to enable dynamic styling:

- `{shape}` — Background shape color
- `{skin}` — Base skin tone
- `{skin_shadow}` — Shadowed areas of the skin
- `{hair}` — Hair color
- `{hair_shadow}` — Shadowed areas of the hair
- `{hair_highlight}` — Highlighted areas of the hair
- `{dress}` — Clothing color
- `{dress_shadow}` — Shadowed areas of the clothing
- `{decorator}` — Accessories (e.g., glasses, necklace)
- `{text}` — Text and sticker color

These variables allow your avatar elements to adapt to different themes and palettes dynamically.

## License

This project is licensed under the ISC License. See the [LICENSE](LICENSE) file for details.

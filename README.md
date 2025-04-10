# Avatar

![GitHub Tag](https://img.shields.io/github/v/tag/go-universal/avatar?sort=semver&label=version)
[![Go Reference](https://pkg.go.dev/badge/github.com/go-universal/avatar.svg)](https://pkg.go.dev/github.com/go-universal/avatar)
[![License](https://img.shields.io/badge/license-ISC-blue.svg)](https://github.com/go-universal/avatar/blob/main/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-universal/avatar)](https://goreportcard.com/report/github.com/go-universal/avatar)
![Contributors](https://img.shields.io/github/contributors/go-universal/avatar)
![Issues](https://img.shields.io/github/issues/go-universal/avatar)

`avatar` is a Go library for generating customizable avatars. It supports creating avatars for text, stickers, and persons with various customization options.

| Letter                                                                                | Sticker                                                                            | Male                                                                            | Female                                                                              |
| ------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- |
| ![Letter](https://github.com/go-universal/avatar/blob/main/demo/john_doe.svg?raw=true) | ![Sticker](https://github.com/go-universal/avatar/blob/main/demo/star.svg?raw=true) | ![Male](https://github.com/go-universal/avatar/blob/main/demo/male.svg?raw=true) | ![Female](https://github.com/go-universal/avatar/blob/main/demo/female.svg?raw=true) |

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

    johnDoe := factory.NewText("John Doe")

    johnDoe.
        RandomizeShape().
        RandomizePalette()

    err := johnDoe.Save("john_doe.svg")
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

    star := factory.NewSticker("")

    star.
        RandomizeShape().
        RandomizePalette()

    err := star.Save("star.svg")
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

    male := factory.NewMale()
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
    err := male.Save("male.svg")
    if err != nil {
        log.Fatal(err)
    }

    female := factory.NewFemale()
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
    err := female.Save("female.svg")
    if err != nil {
        log.Fatal(err)
    }
}
```

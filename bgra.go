// Copyright (c) 2012, James Helferty. All rights reserved.
// Use of this source code is governed by a Clear BSD License
// that can be found in the LICENSE file.

package glimage

import "image"
import "image/color"
import glcolor "github.com/spate/glimage/color"

// BGRA format
//
// Bits:
// BBBBBBBB GGGGGGGG RRRRRRRR AAAAAAAA
// 0        8        16       24
type BGRA struct {
	Pix    []uint8
	Stride int
	Rect   image.Rectangle
}

func NewBGRA(r image.Rectangle) *BGRA {
	pix := make([]uint8, 4*r.Dx()*r.Dy())
	return &BGRA{pix, 4 * r.Dx(), r}
}

func (p *BGRA) At(x, y int) color.Color {
	if !(image.Point{x, y}.In(p.Rect)) {
		return glcolor.BGRA{}
	}
	i := p.PixOffset(x, y)
	return glcolor.BGRA{p.Pix[i+0], p.Pix[i+1], p.Pix[i+2], p.Pix[i+3]}
}

func (p *BGRA) Bounds() image.Rectangle {
	return p.Rect
}

func (p *BGRA) ColorModel() color.Model {
	return glcolor.BGRAModel
}

func (p *BGRA) PixOffset(x, y int) int {
	return (y-p.Rect.Min.Y)*p.Stride + (x-p.Rect.Min.X)*4
}

func (p *BGRA) Set(x, y int, c color.Color) {
	if !(image.Point{x, y}.In(p.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	c1 := glcolor.BGRAModel.Convert(c).(color.RGBA)
	p.Pix[i+0] = c1.B
	p.Pix[i+1] = c1.G
	p.Pix[i+2] = c1.R
	p.Pix[i+3] = c1.A
}

// BGR565 format, aka R5G6B5
//
// Bits:
// BBBBBGGG GGGRRRRR
// 0        8
type BGR565 struct {
	Pix    []uint16
	Stride int
	Rect   image.Rectangle
}

func NewBGR565(r image.Rectangle) *BGR565 {
	pix := make([]uint16, r.Dx()*r.Dy())
	return &BGR565{pix, r.Dx(), r}
}

func (p *BGR565) At(x, y int) color.Color {
	if !(image.Point{x, y}.In(p.Rect)) {
		return glcolor.BGR565{}
	}
	i := p.PixOffset(x, y)
	c := glcolor.BGR565{p.Pix[i]}
	return c
}

func (p *BGR565) Bounds() image.Rectangle {
	return p.Rect
}

func (p *BGR565) ColorModel() color.Model {
	return glcolor.BGR565Model
}

func (p *BGR565) PixOffset(x, y int) int {
	return (y-p.Rect.Min.Y)*p.Stride + (x - p.Rect.Min.X)
}

func (p *BGR565) Set(x, y int, c color.Color) {
	if !(image.Point{x, y}.In(p.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	cn := glcolor.BGR565Model.Convert(c).(glcolor.BGR565)
	p.Pix[i] = cn.BGR
}

// BGRA5551 format, aka A1R5G5B5
//
// Bits:
// BBBBBGGG GGRRRRRA
// 0        8
type BGRA5551 struct {
	Pix    []uint16
	Stride int
	Rect   image.Rectangle
}

func NewBGRA5551(r image.Rectangle) *BGRA5551 {
	pix := make([]uint16, r.Dx()*r.Dy())
	return &BGRA5551{pix, r.Dx(), r}
}

func (p *BGRA5551) At(x, y int) color.Color {
	if !(image.Point{x, y}.In(p.Rect)) {
		return glcolor.BGRA5551{}
	}
	i := p.PixOffset(x, y)
	c := glcolor.BGRA5551{p.Pix[i]}
	return c
}

func (p *BGRA5551) Bounds() image.Rectangle {
	return p.Rect
}

func (p *BGRA5551) ColorModel() color.Model {
	return glcolor.BGRA5551Model
}

func (p *BGRA5551) PixOffset(x, y int) int {
	return (y-p.Rect.Min.Y)*p.Stride + (x - p.Rect.Min.X)
}

func (p *BGRA5551) Set(x, y int, c color.Color) {
	if !(image.Point{x, y}.In(p.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	cn := glcolor.BGRA5551Model.Convert(c).(glcolor.BGRA5551)
	p.Pix[i] = cn.BGRA
}

// BGRA4444 format, aka A4R4G4B4
//
// Bits:
// BBBBGGGG RRRRAAAA
// 0        8
type BGRA4444 struct {
	Pix    []uint16
	Stride int
	Rect   image.Rectangle
}

func NewBGRA4444(r image.Rectangle) *BGRA4444 {
	pix := make([]uint16, r.Dx()*r.Dy())
	return &BGRA4444{pix, r.Dx(), r}
}

func (p *BGRA4444) At(x, y int) color.Color {
	if !(image.Point{x, y}.In(p.Rect)) {
		return glcolor.BGRA4444{}
	}
	i := p.PixOffset(x, y)
	c := glcolor.BGRA4444{p.Pix[i]}
	return c
}

func (p *BGRA4444) Bounds() image.Rectangle {
	return p.Rect
}

func (p *BGRA4444) ColorModel() color.Model {
	return glcolor.BGRA4444Model
}

func (p *BGRA4444) PixOffset(x, y int) int {
	return (y-p.Rect.Min.Y)*p.Stride + (x - p.Rect.Min.X)
}

func (p *BGRA4444) Set(x, y int, c color.Color) {
	if !(image.Point{x, y}.In(p.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	cn := glcolor.BGRA4444Model.Convert(c).(glcolor.BGRA4444)
	p.Pix[i] = cn.BGRA
}

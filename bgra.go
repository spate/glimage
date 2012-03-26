// Copyright (c) 2012, James Helferty. All rights reserved.
// Use of this source code is governed by a Clear BSD License
// that can be found in the LICENSE file.

package glimage

import "image"
import "image/color"
import glcolor "github.com/spate/glimage/color"

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

func (p *BGRA) Opaque() bool {
	// FIXME
	return false
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

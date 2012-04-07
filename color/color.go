// Copyright (c) 2012, James Helferty. All rights reserved.
// Use of this source code is governed by a Clear BSD License
// that can be found in the LICENSE file.

// Package color implements additional color formats common in GL/D3D.
package color

import "image/color"

type BGRA struct {
	B, G, R, A uint8
}

func (c BGRA) RGBA() (r, g, b, a uint32) {
	r = uint32(c.R)
	r |= r << 8
	g = uint32(c.G)
	g |= g << 8
	b = uint32(c.B)
	b |= b << 8
	a = uint32(c.A)
	a |= a << 8
	return
}

type BGR565 struct {
	BGR uint16
}

func (c BGR565) RGBA() (r, g, b, a uint32) {
	r = uint32(c.BGR)<<0 & 0xf800
	r |= r>>5 | r>>10 | r>>15
	g = uint32(c.BGR)<<5 & 0xfc00
	g |= g>>6 | g>>12
	b = uint32(c.BGR)<<11 & 0xf800
	b |= b>>5 | b>>10 | b>>15
	a = 0xffff
	return r, g, b, a
}

type BGRA5551 struct {
	BGRA uint16
}

func (c BGRA5551) RGBA() (r, g, b, a uint32) {
	r = uint32(c.BGRA)<<1 & 0xf800
	r |= r>>5 | r>>10 | r>>15
	g = uint32(c.BGRA)<<6 & 0xf800
	g |= g>>5 | g>>10 | r>>15
	b = uint32(c.BGRA)<<11 & 0xf800
	b |= b>>5 | b>>10 | b>>15
	if (c.BGRA & 0x8000) == 0x8000 {
		a = 0xffff
	} else {
		a = 0x0000
	}
	return r, g, b, a
}

type BGRA4444 struct {
	BGRA uint16
}

func (c BGRA4444) RGBA() (r, g, b, a uint32) {
	r = uint32(c.BGRA)<<4 & 0xf000
	r |= r>>4
	r |= r>>8
	g = uint32(c.BGRA)<<8 & 0xf000
	g |= g>>4
	g |= g>>8
	b = uint32(c.BGRA)<<12 & 0xf000
	b |= b>>4
	b |= b>>8
	a = uint32(c.BGRA)<<0 & 0xf000
	a |= a>>4
	a |= a>>8
	return r, g, b, a
}

// Model for RGB565 and BGRA used by Dxt and GL
var (
	BGRAModel   color.Model = color.ModelFunc(bgraModel)
	BGR565Model color.Model = color.ModelFunc(bgr565Model)
	BGRA5551Model color.Model = color.ModelFunc(bgra5551Model)
	BGRA4444Model color.Model = color.ModelFunc(bgra4444Model)
)

func bgraModel(c color.Color) color.Color {
	if _, ok := c.(BGRA); ok {
		return c
	}
	r, g, b, a := c.RGBA()
	return BGRA{uint8(b >> 8), uint8(g >> 8), uint8(r >> 8), uint8(a >> 8)}
}

func bgr565Model(c color.Color) color.Color {
	if _, ok := c.(BGR565); ok {
		return c
	}
	r, g, b, _ := c.RGBA()
	bgr := uint16(r>>0) & 0xf800
	bgr |= uint16(g>>5) & 0x07e0
	bgr |= uint16(b>>11) & 0x001f
	return BGR565{bgr}
}

func bgra5551Model(c color.Color) color.Color {
	if _, ok := c.(BGRA5551); ok {
		return c
	}
	r, g, b, a := c.RGBA()
	bgra := uint16(r>>1) & 0x7c00
	bgra |= uint16(g>>6) & 0x03e0
	bgra |= uint16(b>>11) & 0x001f
	bgra |= uint16(a>>0) & 0x8000
	return BGRA5551{bgra}
}

func bgra4444Model(c color.Color) color.Color {
	if _, ok := c.(BGRA4444); ok {
		return c
	}
	r, g, b, a := c.RGBA()
	bgra := uint16(r>>4) & 0x0f00
	bgra |= uint16(g>>8) & 0x00f0
	bgra |= uint16(b>>12) & 0x000f
	bgra |= uint16(a>>0) & 0xf000
	return BGRA4444{bgra}
}

// Copyright (c) 2012, James Helferty. All rights reserved.
// Use of this source code is governed by a Clear BSD License
// that can be found in the LICENSE file.

package glimage

import "image/color"

type RGB565 struct {
	RGB uint16
}

func (c RGB565) RGBA() (r, g, b, a uint32) {
	r = uint32(c.RGB<<16) & 0xf800
	r |= r>>5 | r>>10 | r>>15
	g = uint32(c.RGB<<21) & 0xfc00
	g |= g>>6 | g>>12
	b = uint32(c.RGB<<26) & 0xf800
	b |= b>>5 | b>>10 | b>>15
	a = 0xffff
	return r, g, b, a
}

// Model for RGB565 used by Dxt and GL
var (
	RGB565Model color.Model = color.ModelFunc(rgb565Model)
)

func rgb565Model(c color.Color) color.Color {
	if _, ok := c.(RGB565); ok {
		return c
	}
	r, g, b, _ := c.RGBA()
	rgb := uint16(r>>16) & 0xf800
	rgb |= uint16(g>>21) & 0x07e0
	rgb |= uint16(b>>26) & 0x001f
	return RGB565{rgb}
}

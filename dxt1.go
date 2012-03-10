// Copyright (c) 2012, James Helferty. All rights reserved.
// Use of this source code is governed by a Clear BSD License
// that can be found in the LICENSE file.

package glimage

import "image"
import "image/color"

// Dxt1 is an in-memory image whose At method returns color.RGBA values.
type Dxt1 struct {
	// Pix holds the image's pixels in block format. For details, see
	// http://www.opengl.org/registry/specs/EXT/texture_compression_s3tc.txt
	// Note that this is the RGB encoding where A=1 (always opaque)
	Pix []uint8
	// Stride is the Pix stride (in bytes) between vertically adjacent blocks
	Stride int
	// Rect is the image's bounds.
	Rect image.Rectangle
}

// NewDxt1 returns a new Dxt1 with the given bounds
func NewDxt1(r image.Rectangle) *Dxt1 {
	w, h := r.Dx(), r.Dy()
	pix := make([]uint8, ((w+3)/4)*((h+3)/4)*8)
	return &Dxt1{pix, w + 3/4*8, r}
}

func (p *Dxt1) ColorModel() color.Model {
	return color.RGBAModel
}

func (p *Dxt1) Bounds() image.Rectangle {
	return p.Rect
}

func (p *Dxt1) At(x, y int) color.Color {
	if !(image.Point{x, y}.In(p.Rect)) {
		return color.RGBA{}
	}
	i := p.BlockOffset(x, y)
	r, g, b, _ := ConvertDxt1BlockAt(p.Pix[i:i+8], x%4, y%4)
	return color.RGBA{uint8(r >> 8), uint8(g >> 8), uint8(b >> 8), 0xFF}
}

// Opaque returns whether entire image is opaque. Dxt1 is always opaque.
func (p *Dxt1) Opaque() bool {
	return true
}

func (p *Dxt1) BlockOffset(x, y int) int {
	return p.Stride*(y/4) + (x / 4)
}

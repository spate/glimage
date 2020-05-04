// Copyright (c) 2012, James Helferty. All rights reserved.
// Use of this source code is governed by a Clear BSD License
// that can be found in the LICENSE file.

package glimage

import "image"
import "image/color"

// Dxt3 is an in-memory image whose At method returns color.RGBA values.
type Dxt3 struct {
	// Pix holds the image's pixels in block format. For details, see
	// http://www.opengl.org/registry/specs/EXT/texture_compression_s3tc.txt
	Pix []uint8
	// Stride is the Pix stride (in bytes) between vertically adjacent blocks
	Stride int
	// Rect is the image's bounds.
	Rect image.Rectangle
}

// NewDxt3 returns a new Dxt3 with the given bounds
func NewDxt3(r image.Rectangle) *Dxt3 {
	w, h := r.Dx(), r.Dy()
	pix := make([]uint8, ((w+3)/4)*((h+3)/4)*16)
	return &Dxt3{pix, (w + 3) / 4 * 16, r}
}

func (p *Dxt3) ColorModel() color.Model {
	return color.RGBAModel
}

func (p *Dxt3) Bounds() image.Rectangle {
	return p.Rect
}

func (p *Dxt3) At(x, y int) color.Color {
	if !(image.Point{x, y}.In(p.Rect)) {
		return color.RGBA{}
	}
	i := p.BlockOffset(x, y)
	r, g, b, _ := ConvertDxt3BlockAt(p.Pix[i:i+16], x%4, y%4)
	return color.RGBA{uint8(r >> 8), uint8(g >> 8), uint8(b >> 8), 0xFF}
}

func (p *Dxt3) BlockOffset(x, y int) int {
	return p.Stride*(y/4) + ((x / 4) * 16)
}

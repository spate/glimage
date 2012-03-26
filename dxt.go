// Copyright (c) 2012, James Helferty. All rights reserved.
// Use of this source code is governed by a Clear BSD License
// that can be found in the LICENSE file.

package glimage

import "github.com/spate/glimage/color"

func ConvertDxt1BlockAt(pix []uint8, x, y int) (r, g, b, a uint32) {
	color0 := color.RGB565{uint16(pix[0] | pix[1]<<8)}
	color1 := color.RGB565{uint16(pix[2] | pix[3]<<8)}
	bits := uint32(pix[4]) | uint32(pix[5])<<8 | uint32(pix[6])<<16 | uint32(pix[7])<<24

	code := bits >> (2 * (uint8(y)*4 + uint8(x))) & 0x3
	switch code {
	case 0:
		return color0.RGBA()
	case 1:
		return color1.RGBA()
	case 2:
		if color0.RGB > color1.RGB {
			r0, g0, b0, _ := color0.RGBA()
			r1, g1, b1, _ := color1.RGBA()
			return (2*r0 + r1) / 3, (2*g0 + g1) / 3, (2*b0 + b1) / 3, 0xFFFF
		} else {
			r0, g0, b0, _ := color0.RGBA()
			r1, g1, b1, _ := color1.RGBA()
			return (r0 + r1) / 2, (g0 + g1) / 2, (b0 + b1) / 2, 0xFFFF
		}
	case 3:
		if color0.RGB > color1.RGB {
			r0, g0, b0, _ := color0.RGBA()
			r1, g1, b1, _ := color1.RGBA()
			return (r0 + 2*r1) / 3, (g0 + 2*g1) / 3, (b0 + 2*b1) / 3, 0xFFFF
		} else {
			return 0, 0, 0, 0xFFFF
		}
	}
	// should never get here
	return 0, 0, 0, 0
}

func ConvertDxt3BlockAt(pix []uint8, x, y int) (r, g, b, a uint32) {
	// RGB determined same as DXT1
	r, g, b, _ = ConvertDxt1BlockAt(pix[8:], x, y)

	// Alpha is quantized to 4 bits
	alpha := uint64(pix[0]) | uint64(pix[1])<<8 | uint64(pix[2])<<16 | uint64(pix[3])<<24
	alpha |= uint64(pix[4])<<32 | uint64(pix[5])<<40 | uint64(pix[6])<<48 | uint64(pix[7])<<56
	a = uint32(alpha >> (4*uint8(y)*4 + uint8(x)) & 0xF)
	a |= a<<4 | a<<8 | a<<12
	return
}

func ConvertDxt5BlockAt(pix []uint8, x, y int) (r, g, b, a uint32) {
	// RGB determined same as DXT1
	r, g, b, _ = ConvertDxt1BlockAt(pix[8:], x, y)

	// Unpack alpha
	alpha0 := uint32(pix[0])
	alpha0 |= alpha0 << 8
	alpha1 := uint32(pix[1])
	alpha1 |= alpha1 << 8
	bits := uint64(pix[2]) | uint64(pix[3])<<8 | uint64(pix[4])<<16
	bits |= uint64(pix[5])<<24 | uint64(pix[6])<<32 | uint64(pix[7])<<40

	code := bits >> (3 * ((uint8(y) * 4) + uint8(x)))
	switch code {
	case 0:
		a = alpha0
	case 1:
		a = alpha1
	default:
		if alpha0 > alpha1 {
			switch code {
			case 2:
				a = (6*alpha0 + 1*alpha1) / 7
			case 3:
				a = (5*alpha0 + 2*alpha1) / 7
			case 4:
				a = (4*alpha0 + 3*alpha1) / 7
			case 5:
				a = (3*alpha0 + 4*alpha1) / 7
			case 6:
				a = (2*alpha0 + 5*alpha1) / 7
			case 7:
				a = (1*alpha0 + 6*alpha1) / 7
			}
		} else {
			switch code {
			case 2:
				a = (4*alpha0 + 1*alpha1) / 5
			case 3:
				a = (3*alpha0 + 2*alpha1) / 5
			case 4:
				a = (2*alpha0 + 3*alpha1) / 5
			case 5:
				a = (1*alpha0 + 4*alpha1) / 5
			case 6:
				a = 0x0000
			case 7:
				a = 0xFFFF
			}
		}
	}

	return
}

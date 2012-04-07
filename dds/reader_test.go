// Copyright (c) 2012, James Helferty. All rights reserved.
// Use of this source code is governed by a Clear BSD License
// that can be found in the LICENSE file.

package dds

import "testing"
import "os"
import "fmt"
import "image"
import "image/color"

func testColor(t *testing.T, fmt string, target color.RGBA, img image.Image, x, y int) {
	sample := color.RGBAModel.Convert(img.At(x, y)).(color.RGBA)
	if target != sample {
		t.Errorf("%s, loc (%v,%v): sample %v != target %v", fmt, x, y, sample, target)
	}
}

func testDDS(t *testing.T, format string, test_transparent bool) {
	filename := fmt.Sprintf("testdata/test%v.dds", format)
	f, err := os.Open(filename)
	if err != nil {
		t.Errorf("can't open file %v", filename)
		return
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	//fmt.Printf("%v: %v\n\n", format, img)

	// opaque
	testColor(t, format, color.RGBA{0xff, 0x00, 0x00, 0xff}, img, 0, 0)
	testColor(t, format, color.RGBA{0x00, 0x00, 0xff, 0xff}, img, 2, 0)
	testColor(t, format, color.RGBA{0xff, 0xff, 0xff, 0xff}, img, 0, 4)
	testColor(t, format, color.RGBA{0x00, 0xff, 0x00, 0xff}, img, 2, 4)

	// transparent
	if test_transparent {
		testColor(t, format, color.RGBA{0xff, 0x00, 0x00, 0x00}, img, 4, 0)
		testColor(t, format, color.RGBA{0x00, 0x00, 0xff, 0x00}, img, 6, 0)
		testColor(t, format, color.RGBA{0xff, 0xff, 0xff, 0x00}, img, 4, 4)
		testColor(t, format, color.RGBA{0x00, 0xff, 0x00, 0x00}, img, 6, 4)
	}
}

func TestDDSFiles(t *testing.T) {
	testDDS(t, "A8R8G8B8", true)
	testDDS(t, "A4R4G4B4", true)
	testDDS(t, "A1R5G5B5", true)
	testDDS(t, "R5G6B5", false)
	testDDS(t, "DXT1", false)
	testDDS(t, "DXT3", false)
	testDDS(t, "DXT5", false)
}

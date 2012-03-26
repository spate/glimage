// Copyright (c) 2012, James Helferty
// All rights reserved.

// Package dds implements a DDS image decoder.
package dds

import . "github.com/spate/glimage/dds/types"
import "github.com/spate/glimage"
import "image"
import "image/color"
import "encoding/binary"
import "bufio"
import "io"
import "fmt"

// main decoder struct
type decoder struct {
	r   io.Reader
	h   DDS_HEADER
	tmp [128]byte
	img []image.Image
}

type reader interface {
	io.Reader
	ReadByte() (c byte, err error)
}

func (d *decoder) decode(r io.Reader, full bool) error {
	if rr, ok := r.(reader); ok {
		d.r = rr
	} else {
		d.r = bufio.NewReader(r)
	}

	// Check for DDS magic number
	_, err := io.ReadFull(d.r, d.tmp[:4])
	if err != nil {
		return err
	}
	ident := string(d.tmp[0:4])
	if ident != "DDS " {
		return fmt.Errorf("dds: wrong magic number")
	}

	// Decode the DDS header
	err = d.decodeHeader()
	if err != nil {
		return err
	}

	// Check if it's a supported format
	// For now, we'll only support DXT1,DXT3,DXT5
	neededFlags := uint32(DDSD_HEIGHT | DDSD_WIDTH | DDSD_PIXELFORMAT)
	if d.h.Flags&neededFlags != neededFlags {
		return fmt.Errorf("dds: file header is missing necessary dds flags")
	}

	// Sanitize mipmap count
	if d.h.Flags&DDSD_MIPMAPCOUNT == 0 {
		d.h.MipMapCount = 1
	}

	if !full {
		return nil
	}

	switch {
	case d.h.Ddspf.Flags&DDPF_FOURCC != 0:
		switch d.h.Ddspf.FourCC {
		case FOURCC_DXT1:
			d.img = make([]image.Image, d.h.MipMapCount)
			w, h := int(d.h.Width), int(d.h.Height)
			for i := 0; i < int(d.h.MipMapCount); i++ {
				//fmt.Printf("mipmap %v is %vx%v\n", i, w, h)
				img := glimage.NewDxt1(image.Rect(0, 0, w, h))
				_, err = io.ReadFull(d.r, img.Pix)
				if err != nil {
					return err
				}
				d.img[i] = image.Image(img)
				w >>= 1
				h >>= 1
			}
		case FOURCC_DXT3:
			d.img = make([]image.Image, d.h.MipMapCount)
			w, h := int(d.h.Width), int(d.h.Height)
			for i := 0; i < int(d.h.MipMapCount); i++ {
				//fmt.Printf("mipmap %v is %vx%v\n", i, w, h)
				img := glimage.NewDxt3(image.Rect(0, 0, w, h))
				_, err = io.ReadFull(d.r, img.Pix)
				if err != nil {
					return err
				}
				d.img[i] = image.Image(img)
				w >>= 1
				h >>= 1
			}
		case FOURCC_DXT5:
			d.img = make([]image.Image, d.h.MipMapCount)
			w, h := int(d.h.Width), int(d.h.Height)
			for i := 0; i < int(d.h.MipMapCount); i++ {
				//fmt.Printf("mipmap %v is %vx%v\n", i, w, h)
				img := glimage.NewDxt5(image.Rect(0, 0, w, h))
				_, err = io.ReadFull(d.r, img.Pix)
				if err != nil {
					return err
				}
				d.img[i] = image.Image(img)
				w >>= 1
				h >>= 1
			}
		default:
			return fmt.Errorf("dds: unrecognized format %v", d.h.Ddspf)
		}
	case d.h.Ddspf.Flags&DDPF_RGB != 0:
		if d.h.Ddspf.Flags&DDPF_ALPHAPIXELS != 0 {
			switch {
			case d.h.Ddspf.RBitMask == 0x00FF0000 && d.h.Ddspf.GBitMask == 0x0000FF00 &&
				d.h.Ddspf.BBitMask == 0x000000FF && d.h.Ddspf.ABitMask == 0xFF000000:
				d.img = make([]image.Image, d.h.MipMapCount)
				w, h := int(d.h.Width), int(d.h.Height)
				for i := 0; i < int(d.h.MipMapCount); i++ {
					img := glimage.NewBGRA(image.Rect(0, 0, w, h))
					_, err = io.ReadFull(d.r, img.Pix)
					if err != nil {
						return err
					}
					d.img[i] = image.Image(img)
					w >>= 1
					h >>= 1
				}
			default:
				return fmt.Errorf("dds: unrecognized format %v", d.h.Ddspf)
			}
		} else {
			return fmt.Errorf("dds: unrecognized format %v", d.h.Ddspf)
		}
	default:
		return fmt.Errorf("dds: unrecognized format %v", d.h.Ddspf)
	}

	return nil
}

func (d *decoder) decodeHeader() error {
	// read in header
	err := binary.Read(d.r, binary.LittleEndian, &d.h)
	if err != nil {
		return err
	}

	if d.h.Size != 124 {
		return fmt.Errorf("dds: invalid DDS header")
	}

	if d.h.Ddspf.FourCC == FOURCC_DX10 {
		return fmt.Errorf("dds: unsupported DX10 header")
	}

	//fmt.Printf("header:\n%v\n",d.h)

	return nil
}

// FIXME: Add a DecodeAll that returns a DDS structure capable of
// holding mipmaps, volume and cubemap textures.
// func DecodeAll(r io.Reader) (dds.DDS, error)

// Decode reads a DDS image from r and returns it as an image.Image.
// The type of Image returned depends on the DDS contents.
func Decode(r io.Reader) (image.Image, error) {
	var d decoder
	err := d.decode(r, true)
	if err != nil {
		return nil, err
	}
	return d.img[0], nil
}

// DecodeConfig gets configuration information about the DDS file
func DecodeConfig(r io.Reader) (image.Config, error) {
	var d decoder
	err := d.decode(r, false)
	if err != nil {
		return image.Config{}, err
	}
	return image.Config{
		ColorModel: color.RGBAModel,
		Width:      int(d.h.Width),
		Height:     int(d.h.Height),
	}, nil
}

func init() {
	image.RegisterFormat("dds", "DDS ", Decode, DecodeConfig)
}

// Copyright (c) 2012, James Helferty
// All rights reserved.

package dds

import "github.com/spate/glimage"
import "image"
import "image/color"
import "encoding/binary"
import "bufio"
import "io"
import "fmt"
import "strings"

const (
	cDDSD_CAPS        = 0x1
	cDDSD_HEIGHT      = 0x2
	cDDSD_WIDTH       = 0x4
	cDDSD_PITCH       = 0x8
	cDDSD_PIXELFORMAT = 0x1000
	cDDSD_MIPMAPCOUNT = 0x20000
	cDDSD_LINEARSIZE  = 0x80000
	cDDSD_DEPTH       = 0x800000
)

const (
	cDDPF_ALPHAPIXELS = 0x1
	cDDPF_ALPHA       = 0x2
	cDDPF_FOURCC      = 0x4
	cDDPF_RGB         = 0x40
	cDDPF_YUV         = 0x200
	cDDPF_LUMINANCE   = 0x20000
)

const (
	cDDSCAPS_COMPLEX = 0x8
	cDDSCAPS_MIPMAP  = 0x400000
	cDDSCAPS_TEXTURE = 0x1000
)

const (
	cDDSCAPS2_CUBEMAP           = 0x200
	cDDSCAPS2_CUBEMAP_POSITIVEX = 0x400
	cDDSCAPS2_CUBEMAP_NEGATIVEX = 0x800
	cDDSCAPS2_CUBEMAP_POSITIVEY = 0x1000
	cDDSCAPS2_CUBEMAP_NEGATIVEY = 0x2000
	cDDSCAPS2_CUBEMAP_POSITIVEZ = 0x4000
	cDDSCAPS2_CUBEMAP_NEGATIVEZ = 0x8000
	cDDSCAPS2_VOLUME            = 0x200000
)

const (
	fourccDXT1 = 0x31545844
	fourccDXT3 = 0x33545844
	fourccDXT5 = 0x35545844
	fourccDX10 = 0x30315844
)

type sDDS_HEADER struct {
	Size              uint32
	Flags             uint32
	Height            uint32
	Width             uint32
	PitchOrLinearSize uint32
	Depth             uint32
	MipMapCount       uint32
	Reserved1	[11]uint32
	Ddspf       sDDS_PIXELFORMAT
	Caps        uint32
	Caps2       uint32
	Caps3		uint32
	Caps4		uint32
	Reserved2	uint32
}

func (d sDDS_HEADER) String() string {
	// flags
	var f []string
	if d.Flags&cDDSD_CAPS != 0 {
		f = append(f, "DDSD_CAPS")
	}
	if d.Flags&cDDSD_HEIGHT != 0 {
		f = append(f, "DDSD_HEIGHT")
	}
	if d.Flags&cDDSD_WIDTH != 0 {
		f = append(f, "DDSD_WIDTH")
	}
	if d.Flags&cDDSD_PITCH != 0 {
		f = append(f, "DDSD_PITCH")
	}
	if d.Flags&cDDSD_PIXELFORMAT != 0 {
		f = append(f, "DDSD_PIXELFORMAT")
	}
	if d.Flags&cDDSD_MIPMAPCOUNT != 0 {
		f = append(f, "DDSD_MIPMAPCOUNT")
	}
	if d.Flags&cDDSD_LINEARSIZE != 0 {
		f = append(f, "DDSD_LINEARSIZE")
	}
	if d.Flags&cDDSD_DEPTH != 0 {
		f = append(f, "DDSD_DEPTH")
	}
	// caps
	var c []string
	if d.Caps&cDDSCAPS_COMPLEX != 0 {
		c = append(c, "DDSCAPS_COMPLEX")
	}
	if d.Caps&cDDSCAPS_MIPMAP != 0 {
		c = append(c, "DDSCAPS_MIPMAP")
	}
	if d.Caps&cDDSCAPS_TEXTURE != 0 {
		c = append(c, "DDSCAPS_TEXTURE")
	}
	// caps2
	var c2 []string
	if d.Caps2&cDDSCAPS2_CUBEMAP != 0 {
		c2 = append(c2, "DDSCAPS2_CUBEMAP")
	}
	if d.Caps2&cDDSCAPS2_CUBEMAP_POSITIVEX != 0 {
		c2 = append(c2, "DDSCAPS2_CUBEMAP_POSITIVEX")
	}
	if d.Caps2&cDDSCAPS2_CUBEMAP_NEGATIVEX != 0 {
		c2 = append(c2, "DDSCAPS2_CUBEMAP_NEGATIVEX")
	}
	if d.Caps2&cDDSCAPS2_CUBEMAP_POSITIVEY != 0 {
		c2 = append(c2, "DDSCAPS2_CUBEMAP_POSITIVEY")
	}
	if d.Caps2&cDDSCAPS2_CUBEMAP_NEGATIVEY != 0 {
		c2 = append(c2, "DDSCAPS2_CUBEMAP_NEGATIVEY")
	}
	if d.Caps2&cDDSCAPS2_CUBEMAP_POSITIVEZ != 0 {
		c2 = append(c2, "DDSCAPS2_CUBEMAP_POSITIVEZ")
	}
	if d.Caps2&cDDSCAPS2_CUBEMAP_NEGATIVEZ != 0 {
		c2 = append(c2, "DDSCAPS2_CUBEMAP_NEGATIVEZ")
	}
	if d.Caps2&cDDSCAPS2_VOLUME != 0 {
		c2 = append(c2, "DDSCAPS2_VOLUME")
	}
	// build string, return
	return fmt.Sprintf("<flags=%08x(%s) height=%d width=%d pitch=%d depth=%d mipmaps=%d Ddspf=%v caps=%08x(%s) caps2=%08x(%s)>",
		d.Flags, strings.Join(f, "|"), d.Height, d.Width,
		d.PitchOrLinearSize, d.Depth, d.MipMapCount,
		d.Ddspf,
		d.Caps, strings.Join(c, "|"),
		d.Caps2, strings.Join(c2, "|"))
}

type sDDS_PIXELFORMAT struct {
	Size                                   uint32
	Flags                                  uint32
	FourCC                                 uint32
	RGBBitCount                            uint32
	RBitMask, GBitMask, BBitMask, ABitMask uint32
}

func (d sDDS_PIXELFORMAT)  String() string {
	var s []string
	if d.Flags&cDDPF_ALPHAPIXELS != 0 {
		s = append(s, "DDPF_ALPHAPIXELS")
	}
	if d.Flags&cDDPF_ALPHA != 0 {
		s = append(s, "DDPF_ALPHA")
	}
	if d.Flags&cDDPF_FOURCC != 0 {
		s = append(s, "DDPF_FOURCC")
	}
	if d.Flags&cDDPF_RGB != 0 {
		s = append(s, "DDPF_RGB")
	}
	if d.Flags&cDDPF_YUV != 0 {
		s = append(s, "DDPF_YUV")
	}
	if d.Flags&cDDPF_LUMINANCE != 0 {
		s = append(s, "DDPF_LUMINANCE")
	}
	return fmt.Sprintf("<flags=%08x(%s) fourcc=%08x rgbcount=%d rmask=%08x gmask=%08x bmask=%08x amask=%08x>",
		d.Flags, strings.Join(s, "|"), d.FourCC, d.RGBBitCount,
		d.RBitMask, d.GBitMask, d.BBitMask, d.ABitMask)
}

// main decoder struct
type decoder struct {
	r   io.Reader
	h   sDDS_HEADER
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
	neededFlags := uint32(cDDSD_HEIGHT | cDDSD_WIDTH | cDDSD_PIXELFORMAT)
	if d.h.Flags&neededFlags != neededFlags {
		return fmt.Errorf("dds: file header is missing necessary dds flags")
	}

	// Sanitize mipmap count
	if d.h.Flags&cDDSD_MIPMAPCOUNT == 0 {
		d.h.MipMapCount = 1
	}

	if !full {
		return nil
	}

	switch {
	case d.h.Ddspf.Flags&cDDPF_FOURCC != 0:
		switch d.h.Ddspf.FourCC {
		case fourccDXT1:
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
		case fourccDXT3:
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
		case fourccDXT5:
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

	if d.h.Ddspf.FourCC == fourccDX10 {
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

// Get configuration information about the DDS file
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

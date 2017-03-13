package imaging

import (
	"image"
	"testing"
)

func TestResize(t *testing.T) {
	td := []struct {
		desc string
		src  image.Image
		w, h int
		f    ResampleFilter
		want *image.NRGBA
	}{
		{
			"Resize 2x2 1x1 box",
			&image.NRGBA{
				Rect:   image.Rect(-1, -1, 1, 1),
				Stride: 2 * 4,
				Pix: []uint8{
					0x00, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0xff,
					0x00, 0xff, 0x00, 0xff, 0x00, 0x00, 0xff, 0xff,
				},
			},
			1, 1,
			Box,
			&image.NRGBA{
				Rect:   image.Rect(0, 0, 1, 1),
				Stride: 1 * 4,
				Pix:    []uint8{0x55, 0x55, 0x55, 0xc0},
			},
		},
		{
			"Resize 2x2 2x2 box",
			&image.NRGBA{
				Rect:   image.Rect(-1, -1, 1, 1),
				Stride: 2 * 4,
				Pix: []uint8{
					0x00, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0xff,
					0x00, 0xff, 0x00, 0xff, 0x00, 0x00, 0xff, 0xff,
				},
			},
			2, 2,
			Box,
			&image.NRGBA{
				Rect:   image.Rect(0, 0, 2, 2),
				Stride: 2 * 4,
				Pix: []uint8{
					0x00, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0xff,
					0x00, 0xff, 0x00, 0xff, 0x00, 0x00, 0xff, 0xff,
				},
			},
		},
		{
			"Resize 3x1 1x1 nearest",
			&image.NRGBA{
				Rect:   image.Rect(-1, -1, 2, 0),
				Stride: 3 * 4,
				Pix: []uint8{
					0xff, 0x00, 0x00, 0xff, 0x00, 0xff, 0x00, 0xff, 0x00, 0x00, 0xff, 0xff,
				},
			},
			1, 1,
			NearestNeighbor,
			&image.NRGBA{
				Rect:   image.Rect(0, 0, 1, 1),
				Stride: 1 * 4,
				Pix:    []uint8{0x00, 0xff, 0x00, 0xff},
			},
		},
		{
			"Resize 2x2 0x4 box",
			&image.NRGBA{
				Rect:   image.Rect(-1, -1, 1, 1),
				Stride: 2 * 4,
				Pix: []uint8{
					0x00, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0xff,
					0x00, 0xff, 0x00, 0xff, 0x00, 0x00, 0xff, 0xff,
				},
			},
			0, 4,
			Box,
			&image.NRGBA{
				Rect:   image.Rect(0, 0, 4, 4),
				Stride: 4 * 4,
				Pix: []uint8{
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0xff, 0xff, 0x00, 0x00, 0xff,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0xff, 0xff, 0x00, 0x00, 0xff,
					0x00, 0xff, 0x00, 0xff, 0x00, 0xff, 0x00, 0xff, 0x00, 0x00, 0xff, 0xff, 0x00, 0x00, 0xff, 0xff,
					0x00, 0xff, 0x00, 0xff, 0x00, 0xff, 0x00, 0xff, 0x00, 0x00, 0xff, 0xff, 0x00, 0x00, 0xff, 0xff,
				},
			},
		},
		{
			"Resize 2x2 4x0 linear",
			&image.NRGBA{
				Rect:   image.Rect(-1, -1, 1, 1),
				Stride: 2 * 4,
				Pix: []uint8{
					0x00, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0xff,
					0x00, 0xff, 0x00, 0xff, 0x00, 0x00, 0xff, 0xff,
				},
			},
			4, 0,
			Linear,
			&image.NRGBA{
				Rect:   image.Rect(0, 0, 4, 4),
				Stride: 4 * 4,
				Pix: []uint8{
					0x00, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0x3f, 0xff, 0x00, 0x00, 0xc0, 0xff, 0x00, 0x00, 0xff,
					0x00, 0xff, 0x00, 0x3f, 0x6d, 0x6e, 0x24, 0x6f, 0xb1, 0x13, 0x3a, 0xd0, 0xc0, 0x00, 0x3f, 0xff,
					0x00, 0xff, 0x00, 0xc0, 0x13, 0xb2, 0x3a, 0xcf, 0x33, 0x32, 0x9a, 0xef, 0x3f, 0x00, 0xc0, 0xff,
					0x00, 0xff, 0x00, 0xff, 0x00, 0xc0, 0x3f, 0xff, 0x00, 0x3f, 0xc0, 0xff, 0x00, 0x00, 0xff, 0xff,
				},
			},
		},
		{
			"Resize 0x0 1x1 box",
			&image.NRGBA{
				Rect:   image.Rect(-1, -1, -1, -1),
				Stride: 0,
				Pix:    []uint8{},
			},
			1, 1,
			Box,
			&image.NRGBA{},
		},
		{
			"Resize 2x2 0x0 box",
			&image.NRGBA{
				Rect:   image.Rect(-1, -1, 1, 1),
				Stride: 2 * 4,
				Pix: []uint8{
					0x00, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0xff,
					0x00, 0xff, 0x00, 0xff, 0x00, 0x00, 0xff, 0xff,
				},
			},
			0, 0,
			Box,
			&image.NRGBA{},
		},
		{
			"Resize 2x2 -1x0 box",
			&image.NRGBA{
				Rect:   image.Rect(-1, -1, 1, 1),
				Stride: 2 * 4,
				Pix: []uint8{
					0x00, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0xff,
					0x00, 0xff, 0x00, 0xff, 0x00, 0x00, 0xff, 0xff,
				},
			},
			-1, 0,
			Box,
			&image.NRGBA{},
		},
	}
	for _, d := range td {
		got := Resize(d.src, d.w, d.h, d.f)
		want := d.want
		if !compareNRGBA(got, want, 1) {
			t.Errorf("test [%s] failed: %#v", d.desc, got)
		}
	}

	for i, filter := range []ResampleFilter{
		NearestNeighbor,
		Box,
		Linear,
		Hermite,
		MitchellNetravali,
		CatmullRom,
		BSpline,
		Gaussian,
		Lanczos,
		Hann,
		Hamming,
		Blackman,
		Bartlett,
		Welch,
		Cosine,
	} {
		src := image.NewNRGBA(image.Rect(-1, -1, 2, 3))
		got := Resize(src, 5, 6, filter)
		want := image.NewNRGBA(image.Rect(0, 0, 5, 6))
		if !compareNRGBA(got, want, 0) {
			t.Errorf("test [Resize all filters #%d] failed: %#v", i, got)
		}

		if filter.Kernel != nil {
			x := filter.Kernel(filter.Support + 0.0001)
			if x != 0 {
				t.Errorf("test [ResampleFilter edge cases #%d] failed: %f", i, x)
			}
		}
	}

	bcs2 := bcspline(2, 1, 0)
	if bcs2 != 0 {
		t.Errorf("test [bcspline 2] failed: %f", bcs2)
	}
}

func TestFit(t *testing.T) {
	td := []struct {
		desc string
		src  image.Image
		w, h int
		f    ResampleFilter
		want *image.NRGBA
	}{
		{
			"Fit 2x2 1x10 box",
			&image.NRGBA{
				Rect:   image.Rect(-1, -1, 1, 1),
				Stride: 2 * 4,
				Pix: []uint8{
					0x00, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0xff,
					0x00, 0xff, 0x00, 0xff, 0x00, 0x00, 0xff, 0xff,
				},
			},
			1, 10,
			Box,
			&image.NRGBA{
				Rect:   image.Rect(0, 0, 1, 1),
				Stride: 1 * 4,
				Pix:    []uint8{0x55, 0x55, 0x55, 0xc0},
			},
		},
		{
			"Fit 2x2 10x1 box",
			&image.NRGBA{
				Rect:   image.Rect(-1, -1, 1, 1),
				Stride: 2 * 4,
				Pix: []uint8{
					0x00, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0xff,
					0x00, 0xff, 0x00, 0xff, 0x00, 0x00, 0xff, 0xff,
				},
			},
			10, 1,
			Box,
			&image.NRGBA{
				Rect:   image.Rect(0, 0, 1, 1),
				Stride: 1 * 4,
				Pix:    []uint8{0x55, 0x55, 0x55, 0xc0},
			},
		},
		{
			"Fit 2x2 10x10 box",
			&image.NRGBA{
				Rect:   image.Rect(-1, -1, 1, 1),
				Stride: 2 * 4,
				Pix: []uint8{
					0x00, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0xff,
					0x00, 0xff, 0x00, 0xff, 0x00, 0x00, 0xff, 0xff,
				},
			},
			10, 10,
			Box,
			&image.NRGBA{
				Rect:   image.Rect(0, 0, 2, 2),
				Stride: 2 * 4,
				Pix: []uint8{
					0x00, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0xff,
					0x00, 0xff, 0x00, 0xff, 0x00, 0x00, 0xff, 0xff,
				},
			},
		},
		{
			"Fit 0x0 1x1 box",
			&image.NRGBA{
				Rect:   image.Rect(-1, -1, -1, -1),
				Stride: 0,
				Pix:    []uint8{},
			},
			1, 1,
			Box,
			&image.NRGBA{},
		},
		{
			"Fit 2x2 0x0 box",
			&image.NRGBA{
				Rect:   image.Rect(-1, -1, 1, 1),
				Stride: 2 * 4,
				Pix: []uint8{
					0x00, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0xff,
					0x00, 0xff, 0x00, 0xff, 0x00, 0x00, 0xff, 0xff,
				},
			},
			0, 0,
			Box,
			&image.NRGBA{},
		},
		{
			"Fit 2x2 -1x0 box",
			&image.NRGBA{
				Rect:   image.Rect(-1, -1, 1, 1),
				Stride: 2 * 4,
				Pix: []uint8{
					0x00, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0xff,
					0x00, 0xff, 0x00, 0xff, 0x00, 0x00, 0xff, 0xff,
				},
			},
			-1, 0,
			Box,
			&image.NRGBA{},
		},
	}
	for _, d := range td {
		got := Fit(d.src, d.w, d.h, d.f)
		want := d.want
		if !compareNRGBA(got, want, 0) {
			t.Errorf("test [%s] failed: %#v", d.desc, got)
		}
	}
}

func TestFill(t *testing.T) {
	td := []struct {
		desc string
		src  image.Image
		w, h int
		a    Anchor
		f    ResampleFilter
		want *image.NRGBA
	}{
		{
			"Fill 4x4 2x2 Center Nearest",
			&image.NRGBA{
				Rect:   image.Rect(-1, -1, 3, 3),
				Stride: 4 * 4,
				Pix: []uint8{
					0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f,
					0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x1a, 0x1b, 0x1c, 0x1d, 0x1e, 0x1f,
					0x20, 0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29, 0x2a, 0x2b, 0x2c, 0x2d, 0x2e, 0x2f,
					0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39, 0x3a, 0x3b, 0x3c, 0x3d, 0x3e, 0x3f,
				},
			},
			2, 2,
			Center,
			NearestNeighbor,
			&image.NRGBA{
				Rect:   image.Rect(0, 0, 2, 2),
				Stride: 2 * 4,
				Pix: []uint8{
					0x14, 0x15, 0x16, 0x17, 0x1c, 0x1d, 0x1e, 0x1f,
					0x34, 0x35, 0x36, 0x37, 0x3c, 0x3d, 0x3e, 0x3f,
				},
			},
		},
		{
			"Fill 4x4 1x4 TopLeft Nearest",
			&image.NRGBA{
				Rect:   image.Rect(-1, -1, 3, 3),
				Stride: 4 * 4,
				Pix: []uint8{
					0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f,
					0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x1a, 0x1b, 0x1c, 0x1d, 0x1e, 0x1f,
					0x20, 0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29, 0x2a, 0x2b, 0x2c, 0x2d, 0x2e, 0x2f,
					0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39, 0x3a, 0x3b, 0x3c, 0x3d, 0x3e, 0x3f,
				},
			},
			1, 4,
			TopLeft,
			NearestNeighbor,
			&image.NRGBA{
				Rect:   image.Rect(0, 0, 1, 4),
				Stride: 1 * 4,
				Pix: []uint8{
					0x00, 0x01, 0x02, 0x03,
					0x10, 0x11, 0x12, 0x13,
					0x20, 0x21, 0x22, 0x23,
					0x30, 0x31, 0x32, 0x33,
				},
			},
		},
		{
			"Fill 4x4 8x2 Bottom Nearest",
			&image.NRGBA{
				Rect:   image.Rect(-1, -1, 3, 3),
				Stride: 4 * 4,
				Pix: []uint8{
					0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f,
					0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x1a, 0x1b, 0x1c, 0x1d, 0x1e, 0x1f,
					0x20, 0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29, 0x2a, 0x2b, 0x2c, 0x2d, 0x2e, 0x2f,
					0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39, 0x3a, 0x3b, 0x3c, 0x3d, 0x3e, 0x3f,
				},
			},
			8, 2,
			Bottom,
			NearestNeighbor,
			&image.NRGBA{
				Rect:   image.Rect(0, 0, 8, 2),
				Stride: 8 * 4,
				Pix: []uint8{
					0x30, 0x31, 0x32, 0x33, 0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39, 0x3a, 0x3b, 0x38, 0x39, 0x3a, 0x3b, 0x3c, 0x3d, 0x3e, 0x3f, 0x3c, 0x3d, 0x3e, 0x3f,
					0x30, 0x31, 0x32, 0x33, 0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39, 0x3a, 0x3b, 0x38, 0x39, 0x3a, 0x3b, 0x3c, 0x3d, 0x3e, 0x3f, 0x3c, 0x3d, 0x3e, 0x3f,
				},
			},
		},
		{
			"Fill 4x4 2x8 Top Nearest",
			&image.NRGBA{
				Rect:   image.Rect(-1, -1, 3, 3),
				Stride: 4 * 4,
				Pix: []uint8{
					0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f,
					0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x1a, 0x1b, 0x1c, 0x1d, 0x1e, 0x1f,
					0x20, 0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29, 0x2a, 0x2b, 0x2c, 0x2d, 0x2e, 0x2f,
					0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39, 0x3a, 0x3b, 0x3c, 0x3d, 0x3e, 0x3f,
				},
			},
			2, 8,
			Top,
			NearestNeighbor,
			&image.NRGBA{
				Rect:   image.Rect(0, 0, 2, 8),
				Stride: 2 * 4,
				Pix: []uint8{
					0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b,
					0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b,
					0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x1a, 0x1b,
					0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x1a, 0x1b,
					0x24, 0x25, 0x26, 0x27, 0x28, 0x29, 0x2a, 0x2b,
					0x24, 0x25, 0x26, 0x27, 0x28, 0x29, 0x2a, 0x2b,
					0x34, 0x35, 0x36, 0x37, 0x38, 0x39, 0x3a, 0x3b,
					0x34, 0x35, 0x36, 0x37, 0x38, 0x39, 0x3a, 0x3b,
				},
			},
		},
		{
			"Fill 4x4 4x4 TopRight Box",
			&image.NRGBA{
				Rect:   image.Rect(-1, -1, 3, 3),
				Stride: 4 * 4,
				Pix: []uint8{
					0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f,
					0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x1a, 0x1b, 0x1c, 0x1d, 0x1e, 0x1f,
					0x20, 0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29, 0x2a, 0x2b, 0x2c, 0x2d, 0x2e, 0x2f,
					0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39, 0x3a, 0x3b, 0x3c, 0x3d, 0x3e, 0x3f,
				},
			},
			4, 4,
			TopRight,
			Box,
			&image.NRGBA{
				Rect:   image.Rect(0, 0, 4, 4),
				Stride: 4 * 4,
				Pix: []uint8{
					0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f,
					0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x1a, 0x1b, 0x1c, 0x1d, 0x1e, 0x1f,
					0x20, 0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29, 0x2a, 0x2b, 0x2c, 0x2d, 0x2e, 0x2f,
					0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39, 0x3a, 0x3b, 0x3c, 0x3d, 0x3e, 0x3f,
				},
			},
		},
		{
			"Fill 4x4 0x4 Left Box",
			&image.NRGBA{
				Rect:   image.Rect(-1, -1, 3, 3),
				Stride: 4 * 4,
				Pix: []uint8{
					0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f,
					0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x1a, 0x1b, 0x1c, 0x1d, 0x1e, 0x1f,
					0x20, 0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29, 0x2a, 0x2b, 0x2c, 0x2d, 0x2e, 0x2f,
					0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39, 0x3a, 0x3b, 0x3c, 0x3d, 0x3e, 0x3f,
				},
			},
			0, 4,
			Left,
			Box,
			&image.NRGBA{},
		},
		{
			"Fill 0x0 4x4 Right Box",
			&image.NRGBA{},
			4, 4,
			Right,
			Box,
			&image.NRGBA{},
		},
	}
	for _, d := range td {
		got := Fill(d.src, d.w, d.h, d.a, d.f)
		want := d.want
		if !compareNRGBA(got, want, 0) {
			t.Errorf("test [%s] failed: %#v", d.desc, got)
		}
	}
}

func TestThumbnail(t *testing.T) {
	td := []struct {
		desc string
		src  image.Image
		w, h int
		f    ResampleFilter
		want *image.NRGBA
	}{
		{
			"Thumbnail 6x2 1x1 box",
			&image.NRGBA{
				Rect:   image.Rect(-1, -1, 5, 1),
				Stride: 6 * 4,
				Pix: []uint8{
					0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x00, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
					0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x00, 0xff, 0x00, 0xff, 0x00, 0x00, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
				},
			},
			1, 1,
			Box,
			&image.NRGBA{
				Rect:   image.Rect(0, 0, 1, 1),
				Stride: 1 * 4,
				Pix:    []uint8{0x55, 0x55, 0x55, 0xc0},
			},
		},
		{
			"Thumbnail 2x6 1x1 box",
			&image.NRGBA{
				Rect:   image.Rect(-1, -1, 1, 5),
				Stride: 2 * 4,
				Pix: []uint8{
					0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
					0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
					0x00, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0xff,
					0x00, 0xff, 0x00, 0xff, 0x00, 0x00, 0xff, 0xff,
					0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
					0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
				},
			},
			1, 1,
			Box,
			&image.NRGBA{
				Rect:   image.Rect(0, 0, 1, 1),
				Stride: 1 * 4,
				Pix:    []uint8{0x55, 0x55, 0x55, 0xc0},
			},
		},
		{
			"Thumbnail 1x3 2x2 box",
			&image.NRGBA{
				Rect:   image.Rect(-1, -1, 0, 2),
				Stride: 1 * 4,
				Pix: []uint8{
					0x00, 0x00, 0x00, 0x00,
					0xff, 0x00, 0x00, 0xff,
					0xff, 0xff, 0xff, 0xff,
				},
			},
			2, 2,
			Box,
			&image.NRGBA{
				Rect:   image.Rect(0, 0, 2, 2),
				Stride: 2 * 4,
				Pix: []uint8{
					0xff, 0x00, 0x00, 0xff, 0xff, 0x00, 0x00, 0xff,
					0xff, 0x00, 0x00, 0xff, 0xff, 0x00, 0x00, 0xff,
				},
			},
		},
	}
	for _, d := range td {
		got := Thumbnail(d.src, d.w, d.h, d.f)
		want := d.want
		if !compareNRGBA(got, want, 0) {
			t.Errorf("test [%s] failed: %#v", d.desc, got)
		}
	}
}

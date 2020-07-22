package dynaimg

import (
	"github.com/reiver/go-rgba32"

	"image"
	"image/color"
)

// RGBA represents a image frame of dimensions ‘Width’ × ‘Height’.
//
// The pel / pixel in this image frame is 4 bytes.
//
// That measn that len(‘Pix’) must be exactly 4 × ‘Width’ × ‘Height’.
//
// So with ‘Width’ 10 & ‘Height’ is 5, then len(‘Pix’) =
// 4 × ‘Width’ × ‘Height’ = 4 × 10 × 5 = 4 × 50 = 200.
type RGBA struct{
	Pix []uint8
	Width  int
	Height int
}

func (receiver RGBA) At(x, y int) color.Color {
	if nil == receiver.Pix {
		return color.NRGBA{}
	}

	if !(image.Point{X:x,Y:x}).In(receiver.Bounds()) {
		return color.NRGBA{}
	}

	pix := receiver.Pix

	if receiver.expectedByteSize() != len(pix) {
		return color.NRGBA{}
	}

	var p []uint8
	{
		depth := receiver.depth()

		low  := receiver.PixOffset(x,y)
		high := low + depth

		p = pix[low:high]
	}

	rgba := rgba32.Slice(p)

	return rgba
}

func (receiver RGBA) Bounds() image.Rectangle {
	const x = 0
	const y = 0

	return image.Rectangle{
		Min:image.Point{
			X: x,
			Y: y,
		},
		Max:image.Point{
			X: x+receiver.Width,
			Y: y+receiver.Height,
		},
	}
}

func (receiver RGBA) ColorModel() color.Model {
	return color.NRGBAModel
}

func (receiver RGBA) depth() int {
	return rgba32.ByteSize
}

func (receiver RGBA) expectedByteSize() int {
	return receiver.Width * receiver.Height * receiver.depth()
}

func (receiver RGBA) PixOffset(x, y int) int {
	b := receiver.Bounds()
	width := receiver.Width
	depth := receiver.depth()

	return ((y-b.Min.Y)*width*depth) + ((x-b.Min.X)*depth)
}

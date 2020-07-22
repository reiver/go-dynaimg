package dynaimg_test

import (
	"github.com/reiver/go-dynaimg"

	"image"
	"math/rand"
	"time"

	"testing"
)

func TestRGBA_Bounds(t *testing.T) {

	randomness := rand.New( rand.NewSource( time.Now().UTC().UnixNano() ) )

	tests := []struct{
		Width int
		Height int
		Expected image.Rectangle
	}{
		{
			Width:             1,
			Height:            1,
			Expected: image.Rectangle{
				Min:image.Point{
					X: 0,
					Y: 0,
				},
				Max:image.Point{
					X: 1,
					Y: 1,
				},
			},
		},
		{
			Width:             1,
			Height:            2,
			Expected: image.Rectangle{
				Min:image.Point{
					X: 0,
					Y: 0,
				},
				Max:image.Point{
					X: 1,
					Y: 2,
				},
			},
		},
		{
			Width:             2,
			Height:            1,
			Expected: image.Rectangle{
				Min:image.Point{
					X: 0,
					Y: 0,
				},
				Max:image.Point{
					X: 2,
					Y: 1,
				},
			},
		},
		{
			Width:             2,
			Height:            2,
			Expected: image.Rectangle{
				Min:image.Point{
					X: 0,
					Y: 0,
				},
				Max:image.Point{
					X: 2,
					Y: 2,
				},
			},
		},
		{
			Width:             5,
			Height:            9,
			Expected: image.Rectangle{
				Min:image.Point{
					X: 0,
					Y: 0,
				},
				Max:image.Point{
					X: 5,
					Y: 9,
				},
			},
		},
	}

	for i:=0; i<10; i++ {

		w := randomness.Intn(256)
		if w < 1 {
			w = 1
		}

		h := randomness.Intn(256)
		if h < 1 {
			h = 1
		}

		test := struct{
			Width int
			Height int
			Expected image.Rectangle
		}{
			Width:             w,
			Height:            h,
			Expected: image.Rectangle{
				Min:image.Point{
					X: 0,
					Y: 0,
				},
				Max:image.Point{
					X: w,
					Y: h,
				},
			},
		}

		tests = append(tests, test)
	}

	for testNumber, test := range tests {

		byteSize := test.Width * test.Height

		pix := make([]uint8, byteSize)

		var img dynaimg.RGBA = dynaimg.RGBA{
			Pix:pix,
			Width:test.Width,
			Height:test.Height,
		}

		{
			expected := test.Expected
			actual   := img.Bounds()

			if expected != actual {
				t.Errorf("For test #%d, the actual bounds is not what was expected.", testNumber)
				t.Logf("EXPECTED: %#v", expected)
				t.Logf("ACTUAL:   %#v", actual)
				continue
			}
		}
	}
}

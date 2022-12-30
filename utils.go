package rendercard

import (
	"image"
	"image/color"

	"github.com/Coloured-glaze/gg"
)

// Fillet 将矩形图片裁切为圆角矩形
func Fillet(dst image.Image, r int) image.Image {
	dstr := gg.ImageToNRGBA(dst)
	mx, my := dst.Bounds().Max.X, dst.Bounds().Max.Y
	var xx, yy, rr float64
	for y := 0; y < my/2; y++ {
		for x := 0; x < mx/2; x++ {
			if x <= r && y <= r {
				xx, yy, rr = float64(r-x), float64(y-r), float64(r)
				if xx*xx+yy*yy >= rr*rr {
					dstr.Set(x, y, color.NRGBA{})
					dstr.Set(mx-1-x, y, color.NRGBA{})
					dstr.Set(x, my-1-y, color.NRGBA{})
					dstr.Set(mx-1-x, my-1-y, color.NRGBA{})
				}
			}

		}
	}
	return dstr
}

// Transparency 更改透明度 magnification 倍率值
func Transparency(dst image.Image, magnification float64) image.Image {
	dstr := gg.ImageToNRGBA(dst)
	mx, my := dst.Bounds().Max.X, dst.Bounds().Max.Y
	for y := 0; y < my; y++ {
		for x := 0; x < mx; x++ {
			r, g, b, a := dstr.At(x, y).RGBA()
			dstr.Set(x, y, color.NRGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: uint8(float64(a) * magnification)})
		}
	}
	return dstr
}

package rendercard

import (
	"image"
	"image/color"
	"strings"

	"github.com/FloatTech/gg"
)

// Fillet 将矩形图片裁切为圆角矩形
func Fillet(dst image.Image, r float64) image.Image {
	canvas := gg.NewContext(dst.Bounds().Dx(), dst.Bounds().Dy())
	canvas.DrawRoundedRectangle(0, 0, float64(dst.Bounds().Dx()), float64(dst.Bounds().Dy()), r)
	canvas.Clip()
	canvas.DrawImage(dst, 0, 0)
	return canvas.Image()
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

// Truncate 截断文字
func Truncate(fontdata []byte, texts []string, maxW, fontsize float64) (newtexts []string, err error) {
	one := gg.NewContext(1, 1)
	err = one.ParseFontFace(fontdata, fontsize)
	if err != nil {
		return
	}
	newtexts = make([]string, 0, len(texts)*2)
	for i := 0; i < len(texts); i++ {
		for len(texts[i]) > 0 {
			var tmp strings.Builder
			tmp.Grow(len(texts[i]))
			res := make([]rune, 0, len(texts[i]))
			for _, t := range texts[i] {
				tmp.WriteRune(t)
				width, _ := one.MeasureString(tmp.String())
				if width > maxW {
					break
				}
				res = append(res, t)

			}
			newlinetext := string(res)
			newtexts = append(newtexts, newlinetext)
			if len(newlinetext) >= len(texts[i]) {
				break
			}
			texts[i] = texts[i][len(newlinetext):]
		}
	}
	return
}

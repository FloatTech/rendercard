package rendercard

import (
	"image"
	"image/color"
	"strings"

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

// Truncate 截断文字
func Truncate(fontfile string, texts []string, maxW, fontsize float64) ([]string, error) {
	one := gg.NewContext(1, 1)
	err := one.LoadFontFace(fontfile, fontsize)
	if err != nil {
		return nil, err
	}
	newtexts := make([]string, 0, len(texts)*2)
	for i := 0; i < len(texts); i++ {
		newlinetext, textw, tmpw := "", 0.0, 0.0
		text := texts[i]
		for len(texts[i]) > 0 {
			var tmp strings.Builder
			tmp.Grow(len(text))
			res := make([]rune, 0, len(text))
			for _, r := range text {
				tmp.WriteRune(r)
				width, _ := one.MeasureString(tmp.String()) // 获取文字宽度
				if width > maxW {                           // 如果宽度大于文字边距
					break // 跳出
				} else {
					res = append(res, r) // 写入
				}
				newlinetext = string(res)
			}
			newtexts = append(newtexts, newlinetext)
			if tmpw > textw {
				textw = tmpw
			}
			if len(newlinetext) >= len(texts[i]) {
				break
			}
			texts[i] = texts[i][len(newlinetext):]
		}
	}
	return newtexts, nil
}

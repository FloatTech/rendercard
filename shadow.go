package rendercard

import (
	"image/color"

	"github.com/Coloured-glaze/gg"
	"github.com/disintegration/imaging"
)

// DrawRoundShadowedRectangle 绘制带阴影的圆角矩形
func DrawRoundShadowedRectangle(canvas *gg.Context, x, y, w, h, r, sr, sigma float64, ox, oy int, rectanglecolor, shadowcolor color.Color) {
	DrawShadowedRectangle(canvas, x, y, w, h, sr, sigma, ox, oy, shadowcolor)
	canvas.DrawRoundedRectangle(x, y, w, h, r)
	canvas.SetColor(rectanglecolor)
	canvas.Fill()
}

// DrawShadowedRectangle 绘制阴影
func DrawShadowedRectangle(canvas *gg.Context, x, y, w, h, r, sigma float64, ox, oy int, shadowcolor color.Color) {
	one := gg.NewContext(canvas.W(), canvas.H())
	one.DrawRoundedRectangle(x, y, w, h, r)
	one.SetColor(shadowcolor)
	one.Fill()
	canvas.DrawImage(imaging.Blur(one.Image(), sigma), int(x)+ox, int(y)+oy)
}

// DrawShadowedString 绘制带阴影的文字
func DrawShadowedString(canvas *gg.Context, s, fontfile string, x, y, point, sigma float64, ox, oy int, stringcolor, shadowcolor color.Color) (err error) {
	one := gg.NewContext(canvas.W(), canvas.H())
	err = one.LoadFontFace(fontfile, point)
	if err != nil {
		return
	}
	one.SetColor(shadowcolor)
	one.DrawString(s, x, y)
	canvas.DrawImage(imaging.Blur(one.Image(), sigma), int(x)+ox, int(y)+oy)
	err = canvas.LoadFontFace(fontfile, point)
	if err != nil {
		return
	}
	canvas.SetColor(stringcolor)
	canvas.DrawString(s, x, y)
	return
}

// DrawShadowedStringAnchored 在锚点上绘制带阴影的文字
func DrawShadowedStringAnchored(canvas *gg.Context, s, fontfile string, x, y, point, sigma float64, ox, oy int, stringcolor, shadowcolor color.Color, ax, ay float64) (err error) {
	one := gg.NewContext(canvas.W(), canvas.H())
	err = one.LoadFontFace(fontfile, point)
	if err != nil {
		return
	}
	one.SetColor(shadowcolor)
	one.DrawStringAnchored(s, x, y, ax, ay)
	canvas.DrawImage(imaging.Blur(one.Image(), sigma), int(x)+ox, int(y)+oy)
	err = canvas.LoadFontFace(fontfile, point)
	if err != nil {
		return
	}
	canvas.SetColor(stringcolor)
	canvas.DrawStringAnchored(s, x, y, ax, ay)
	return
}

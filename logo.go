package rendercard

import (
	"image"

	"github.com/FloatTech/gg"
)

// RenderServerListLogo ...
func RenderServerListLogo(fontdata []byte) (img image.Image, err error) {
	const w, h = 400, 200
	canvas := gg.NewContext(w, h)
	canvas.SetRGBA255(187, 122, 132, 255)
	err = canvas.ParseFontFace(fontdata, 72)
	if err != nil {
		return
	}
	canvas.DrawStringAnchored("Server", 22, 45, 0, 1)
	canvas.SetRGBA255(246, 166, 171, 255)

	canvas.DrawStringAnchored("List", 219, 112, 0, 1)

	canvas.SetRGBA255(135, 145, 173, 255)
	err = canvas.ParseFontFace(fontdata, 48)
	if err != nil {
		return
	}
	canvas.DrawStringAnchored("服務列表", 23, 120, 0, 1)

	canvas.SetRGBA255(103, 178, 240, 255)
	err = canvas.ParseFontFace(fontdata, 36)
	if err != nil {
		return
	}

	canvas.DrawStringAnchored("サーバー", 82, 25, 0, 1)
	canvas.DrawStringAnchored("リスト", 280, 83, 0, 1)

	mask := canvas.AsMask()

	stroked := gg.NewContext(w, h)
	err = stroked.SetMask(mask)
	if err != nil {
		return
	}
	stroked.SetRGBA255(255, 255, 255, 255)
	stroked.DrawRectangle(0, 0, float64(stroked.W()), float64(stroked.H()))
	stroked.Fill()

	strokedimg := stroked.Image()
	coloredimg := canvas.Image()

	canvas = gg.NewContext(w, h)

	canvas.DrawImage(strokedimg, -3, 0)
	canvas.DrawImage(strokedimg, 3, 0)
	canvas.DrawImage(strokedimg, 0, -3)
	canvas.DrawImage(strokedimg, 0, 3)
	canvas.DrawImage(strokedimg, -3, -3)
	canvas.DrawImage(strokedimg, 3, 3)
	canvas.DrawImage(strokedimg, 3, -3)
	canvas.DrawImage(strokedimg, -3, 3)
	canvas.DrawImage(coloredimg, 0, 0)

	img = canvas.Image()
	return
}

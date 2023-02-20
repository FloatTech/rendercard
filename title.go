package rendercard

import (
	"image"

	"github.com/FloatTech/gg"
)

// DrawTitle 绘制标题
func (t *Title) DrawTitle() (imgs image.Image, err error) {
	// 创建图像
	canvas := gg.NewContext(int(DefaultWidth), 30+30+300+(t.Line*(256+30)))
	canvas.SetRGBA255(245, 245, 245, 255)
	canvas.Clear()

	// 标题背景1
	canvas.DrawRectangle(0, 30, DefaultWidth, 300)
	canvas.SetRGBA255(0, 0, 0, 153)
	canvas.Fill()

	// 标题背景2
	canvas.DrawRectangle(0, 30+40, DefaultWidth, 220)
	canvas.SetRGBA255(0, 0, 0, 153)
	canvas.Fill()

	fontsize1, fontsize2, fontsize3 := 108.0+t.TitleFontOffsetPoint, 72+t.TextFontOffsetPoint, 54.0+t.TextFontOffsetPoint
	// 加载size为108的字体
	err = canvas.ParseFontFace(t.TitleFontData, fontsize1)
	if err != nil {
		return
	}

	// 绘制标题
	canvas.SetRGBA255(250, 250, 250, 255)
	canvas.DrawStringAnchored(t.LeftTitle, (220-fontsize1-fontsize3)/3+t.OffsetX, 30+40+(220-fontsize1-fontsize3)/3+fontsize1/2+t.OffsetY, 0, 0.5)

	// 加载size为54的字体
	err = canvas.ParseFontFace(t.TextFontData, fontsize3)
	if err != nil {
		return
	}

	canvas.SetRGBA255(250, 250, 250, 255)
	// 绘制副标题
	canvas.DrawStringAnchored(t.LeftSubtitle, 3+(220-fontsize1-fontsize3)/3+t.OffsetX, 30+40+(220-fontsize1-fontsize3)/3*2+fontsize1+fontsize3/2+t.OffsetY, 0, 0.5)

	// 加载icon并绘制 高限制220
	var icon image.Image
	icon, err = gg.LoadImage(t.ImagePath)
	if err != nil {
		return
	}
	sc := 220 / float64(icon.Bounds().Dy())
	canvas.ScaleAbout(sc, sc, DefaultWidth-float64(icon.Bounds().Dx())*sc/2, 40+30+220/2)
	canvas.DrawImageAnchored(icon, int(DefaultWidth)-int(float64(icon.Bounds().Dx())*sc/2), 40+30+220/2, 0.5, 0.5)
	canvas.Identity()
	// 加载size为72的字体
	err = canvas.ParseFontFace(t.TextFontData, fontsize2)
	if err != nil {
		return
	}
	canvas.DrawStringAnchored(t.RightTitle, DefaultWidth-float64(icon.Bounds().Dx())*sc-25+t.OffsetX, 30+40+(220-fontsize2*2)/3+fontsize2/2+t.OffsetY, 1, 0.5)
	canvas.DrawStringAnchored(t.RightSubtitle, DefaultWidth-float64(icon.Bounds().Dx())*sc-25+t.OffsetX, 30+40+(220-fontsize2*2)/3*2+fontsize2+fontsize2/2+t.OffsetY, 1, 0.5)

	imgs = canvas.Image()
	return
}

// DrawTitleWithText 绘制标题正文
func (t *Title) DrawTitleWithText(info []string) (imgs image.Image, err error) {
	line := len(info)
	if line < 8 {
		line = 8
	}
	imgh := line*(28+20) + 220 + 50

	// 创建图像
	canvas := gg.NewContext(int(DefaultWidth), imgh)
	canvas.SetRGBA255(250, 250, 250, 255)
	canvas.Clear()

	// 加载icon
	var icon image.Image
	icon, err = gg.LoadImage(t.ImagePath)
	if err != nil {
		return
	}
	sc := 0.0
	if float64(icon.Bounds().Dy())/float64(icon.Bounds().Dx()) > 1 {
		sc = 512 / float64(icon.Bounds().Dy())
	} else {
		sc = DefaultWidth / 2 / float64(icon.Bounds().Dx())
	}
	canvas.ScaleAbout(sc, sc, float64(canvas.W()), float64(canvas.H()))
	canvas.DrawImageAnchored(icon, canvas.W(), canvas.H(), 1, 1)
	canvas.Identity()

	// 加载size为108的字体
	fontsize1, fontsize2 := 108.0+t.TitleFontOffsetPoint, 54.0+t.TextFontOffsetPoint
	fsp1, fsp2 := fontsize1*72/96, fontsize2*72/96
	err = canvas.ParseFontFace(t.TitleFontData, fontsize1)
	if err != nil {
		return
	}

	canvas.SetRGBA255(15, 15, 15, 255)

	// 绘制标题
	canvas.DrawStringAnchored(t.LeftTitle, 25+t.OffsetX, 25+fsp1/2+t.OffsetY, 0, 0.5)

	// 加载size为54的字体
	err = canvas.ParseFontFace(t.TextFontData, fontsize2)
	if err != nil {
		return
	}

	// 绘制一系列标题
	canvas.DrawStringAnchored(t.LeftSubtitle, 25+3+t.OffsetX, 25+fsp1+25+fsp2/2+t.OffsetY, 0, 0.5)
	stringwight, _ := canvas.MeasureString(t.LeftSubtitle)
	canvas.DrawRectangle(25+3+t.OffsetX, 25+fsp1+25+fsp2+7+t.OffsetY, stringwight, 6)
	// 绘制插件开启状态
	if t.IsEnabled {
		canvas.SetRGBA255(35, 235, 35, 255)
	} else {
		canvas.SetRGBA255(235, 35, 35, 255)
	}
	canvas.Fill()
	canvas.SetRGBA255(15, 15, 15, 255)

	// 加载size为54的字体
	err = canvas.ParseFontFace(t.TitleFontData, fontsize2)
	if err != nil {
		return
	}

	canvas.DrawStringAnchored(t.RightTitle, DefaultWidth-40+t.OffsetX, 40+fsp2/2+t.OffsetY, 1, 0.5)
	canvas.DrawStringAnchored(t.RightSubtitle, DefaultWidth-40+t.OffsetX, 40+fsp2+25+fsp2/2+t.OffsetY, 1, 0.5)

	// 加载size为38的字体
	err = canvas.ParseFontFace(t.TextFontData, 38+t.TextFontOffsetPoint)
	if err != nil {
		return
	}

	y := 25 + fsp1 + 25 + fsp2
	for _, text := range info {
		canvas.DrawString(text, 25.0, y+canvas.FontHeight()*2)
		y += 20 + canvas.FontHeight()
	}
	imgs = canvas.Image()
	return
}

// DrawCard 绘制卡片
func (t *Title) DrawCard() (imgs image.Image, err error) {
	recw, rech := 384.0, 256.0
	canvas := gg.NewContext(int(recw), int(rech))
	// 绘制图片
	var banner image.Image
	banner, err = gg.LoadImage(t.ImagePath)
	if err == nil {
		sc := 0.0
		if float64(banner.Bounds().Dy())/float64(banner.Bounds().Dx()) < rech/recw {
			sc = rech / float64(banner.Bounds().Dy())
		} else {
			sc = recw / float64(banner.Bounds().Dx())
		}
		canvas.ScaleAbout(sc, sc, recw/2, rech/2)
		canvas.DrawImageAnchored(banner, int(recw)/2, int(rech)/2, 0.5, 0.5)
		canvas.Identity()
	} else {
		canvas.SetRGB255(RandJPColor())
		canvas.Clear()
	}
	if t.IsEnabled {
		canvas.DrawRectangle(0, rech*0.54, recw, rech-rech*0.54)
	} else {
		canvas.DrawRectangle(0, 0, recw, rech)
	}
	canvas.SetRGBA255(0, 0, 0, 183)
	canvas.Fill()

	// 绘制插件信息
	canvas.SetRGBA255(240, 240, 240, 255)
	fontsize1, fontsize2 := 64.0+t.TitleFontOffsetPoint, 32.0+t.TextFontOffsetPoint
	fsp1, fsp2 := fontsize1*72/96, fontsize2*72/96
	err = canvas.ParseFontFace(t.TitleFontData, fontsize1)
	if err != nil {
		return
	}
	canvas.DrawStringAnchored(t.LeftTitle, recw/2+t.OffsetX, rech*0.54+(rech-rech*0.54-(fsp1+fsp2))/3+fsp1/2+t.OffsetY, 0.5, 0.5)

	err = canvas.ParseFontFace(t.TextFontData, fontsize2)
	if err != nil {
		return
	}
	canvas.DrawStringAnchored(t.LeftSubtitle, recw/2+t.OffsetX, rech*0.54+(rech-rech*0.54-(fsp1+fsp2))/3*2+fsp1+fsp2/2+t.OffsetY, 0.5, 0.5)

	imgs = canvas.Image()
	return
}

package rendercard

import (
	"image"

	"github.com/Coloured-glaze/gg"
	"github.com/FloatTech/zbputils/img"
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

	fontsize1, fontsize2 := 108.0, 54.0
	// 加载size为108的字体
	err = canvas.LoadFontFace(t.TitleFont, fontsize1)
	if err != nil {
		return
	}

	// 绘制标题
	canvas.SetRGBA255(250, 250, 250, 255)
	stringwight, _ := canvas.MeasureString(t.LeftTitle)
	canvas.DrawStringAnchored(t.LeftTitle, (220-(fontsize1+fontsize2)*72/96)*0.33+stringwight/2+t.OffsetX, 30+40+(220-(fontsize1+fontsize2)*72/96)*0.33+fontsize1*72/96*0.5+t.OffsetY, 0.5, 0.5)

	// 加载size为54的字体
	err = canvas.LoadFontFace(t.TextFont, fontsize2)
	if err != nil {
		return
	}

	canvas.SetRGBA255(250, 250, 250, 255)
	// 绘制副标题
	stringwight, _ = canvas.MeasureString(t.LeftSubtitle)
	canvas.DrawStringAnchored(t.LeftSubtitle, 3+(220-(fontsize1+fontsize2)*72/96)*0.33+stringwight/2+t.OffsetX, 30+40+(220-(fontsize1+fontsize2)*72/96)*0.66+fontsize1*72/96+fontsize2*72/96*0.5+t.OffsetY, 0.5, 0.5)

	// 加载icon并绘制
	var icon *img.Factory
	icon, err = img.LoadFirstFrame(t.ImagePath, 220, 220)
	if err != nil {
		return
	}
	canvas.DrawImage(icon.Im, int(DefaultWidth)-icon.W, 40+30)
	// 加载size为72的字体
	fontsize1 = 72
	err = canvas.LoadFontFace(t.TextFont, fontsize1)
	if err != nil {
		return
	}
	stringwight, _ = canvas.MeasureString(t.RightTitle)
	canvas.DrawStringAnchored(t.RightTitle, DefaultWidth-25-float64(icon.W)-stringwight/2+t.OffsetX, 30+40+(220-fontsize1*72/96*2)*0.33+fontsize1*72/96*0.5+t.OffsetY, 0.5, 0.5)
	stringwight, _ = canvas.MeasureString(t.RightSubtitle)
	canvas.DrawStringAnchored(t.RightSubtitle, DefaultWidth-25-float64(icon.W)-stringwight/2+t.OffsetX, 30+40+(220-fontsize1*72/96*2)*0.66+fontsize1*72/96*1.5+t.OffsetY, 0.5, 0.5)

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
	var icon *img.Factory
	icon, err = img.LoadFirstFrame(t.ImagePath, 512, 512)
	if err != nil {
		return
	}
	canvas.DrawImage(icon.Im, DefaultWidth-icon.W, imgh-icon.H)

	// 加载size为108的字体
	fontsize1, fontsize2 := 108.0, 54.0
	err = canvas.LoadFontFace(t.TitleFont, fontsize1)
	if err != nil {
		return
	}

	canvas.SetRGBA255(15, 15, 15, 255)

	// 绘制标题
	stringwight, _ := canvas.MeasureString(t.LeftTitle)
	canvas.DrawStringAnchored(t.LeftTitle, 25+stringwight/2+t.OffsetX, 25+fontsize1*72/96*0.5+t.OffsetY, 0.5, 0.5)

	// 加载size为54的字体
	err = canvas.LoadFontFace(t.TextFont, fontsize2)
	if err != nil {
		return
	}

	// 绘制一系列标题
	stringwight, _ = canvas.MeasureString(t.LeftSubtitle)
	canvas.DrawStringAnchored(t.LeftSubtitle, 25+3+stringwight/2+t.OffsetX, 25+fontsize1*72/96+25+fontsize2*72/96*0.5+t.OffsetY, 0.5, 0.5)

	canvas.DrawRectangle(25+3+t.OffsetX, 25+fontsize1*72/96+25+fontsize2*72/96+7+t.OffsetY, stringwight, 6)
	// 绘制插件开启状态
	if t.IsEnabled {
		canvas.SetRGBA255(35, 235, 35, 255)
	} else {
		canvas.SetRGBA255(235, 35, 35, 255)
	}
	canvas.Fill()
	canvas.SetRGBA255(15, 15, 15, 255)

	// 加载size为54的字体
	err = canvas.LoadFontFace(t.TitleFont, fontsize2)
	if err != nil {
		return
	}

	stringwight, _ = canvas.MeasureString(t.RightTitle)
	canvas.DrawStringAnchored(t.RightTitle, DefaultWidth-40-stringwight/2+t.OffsetX, 40+fontsize2*72/96*0.5+t.OffsetY, 0.5, 0.5)
	stringwight, _ = canvas.MeasureString(t.RightSubtitle)
	canvas.DrawStringAnchored(t.RightSubtitle, DefaultWidth-40-stringwight/2+t.OffsetX, 40+25+fontsize2*72/96*1.5+t.OffsetY, 0.5, 0.5)

	// 加载size为38的字体
	err = canvas.LoadFontFace(t.TextFont, 38)
	if err != nil {
		return
	}

	y := 25 + fontsize1*72/96 + 25 + fontsize2*72/96
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
	var banner *img.Factory
	banner, err = img.LoadFirstFrame(t.ImagePath, int(recw)*2, int(rech)*2)
	if err == nil {
		canvas.DrawImage(img.Size(banner.Im, int(recw), int(rech)).Im, 0, 0)
	} else {
		canvas.DrawRectangle(0, 0, recw, rech)
		canvas.SetRGB255(RandJPColor())
		canvas.Fill()
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
	fontsize1, fontsize2 := 64.0, 32.0
	err = canvas.LoadFontFace(t.TitleFont, fontsize1)
	if err != nil {
		return
	}
	stringwight, _ := canvas.MeasureString(t.LeftTitle)
	canvas.DrawStringAnchored(t.LeftTitle, stringwight/2+(rech-rech*0.54-(fontsize1+fontsize2)*72/96)*0.33+t.OffsetX, rech*0.54+(rech-rech*0.54-(fontsize1+fontsize2)*72/96)*0.33+fontsize1*72/96*0.5+t.OffsetY, 0.5, 0.5)

	err = canvas.LoadFontFace(t.TextFont, fontsize2)
	if err != nil {
		return
	}
	stringwight, _ = canvas.MeasureString(t.LeftSubtitle)
	canvas.DrawStringAnchored(t.LeftSubtitle, 3+stringwight/2+(rech-rech*0.54-(fontsize1+fontsize2)*72/96)*0.33+t.OffsetX, rech*0.54+(rech-rech*0.54-(fontsize1+fontsize2)*72/96)*0.66+fontsize1*72/96+fontsize2*72/96*0.5+t.OffsetY, 0.5, 0.5)

	imgs = Fillet(canvas.Image(), 16)
	return
}

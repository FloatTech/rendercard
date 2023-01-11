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

	// 加载size为108的字体
	err = canvas.LoadFontFace(t.TitleFont, 108)
	if err != nil {
		return
	}

	// 绘制标题
	canvas.SetRGBA255(250, 250, 250, 255)
	canvas.DrawString(t.LeftTitle, 25, 30+40+55+canvas.FontHeight()-canvas.FontHeight()/3)

	// 加载size为54的字体
	err = canvas.LoadFontFace(t.TitleFont, 54)
	if err != nil {
		return
	}

	canvas.SetRGBA255(250, 250, 250, 255)
	// 绘制副标题
	canvas.DrawString(t.LeftSubtitle, 25+3, 30+40+165+canvas.FontHeight()/3)

	// 加载icon并绘制
	var icon *img.Factory
	icon, err = img.LoadFirstFrame(t.ImagePath, 220, 220)
	if err != nil {
		return
	}
	canvas.DrawImage(icon.Im, int(DefaultWidth)-icon.W, 40+30)
	// 加载size为54的字体
	err = canvas.LoadFontFace(t.TitleFont, 72)
	if err != nil {
		return
	}
	fw, _ := canvas.MeasureString(t.RightTitle)
	canvas.DrawString(t.RightTitle, DefaultWidth-25-fw-float64(icon.W), 30+40+15+canvas.FontHeight()*1.25)
	fw1, _ := canvas.MeasureString(t.RightSubtitle)
	canvas.DrawString(t.RightSubtitle, DefaultWidth-25-fw1-float64(icon.W), 30+40+15+canvas.FontHeight()*2.75)

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
	icon, err = img.LoadFirstFrame(t.ImagePath, 768, 768)
	if err != nil {
		return
	}
	canvas.DrawImage(icon.Im, DefaultWidth-icon.W, imgh-icon.H)

	// 绘制标题与内容的分割线
	/*canvas.DrawRectangle(0, 220, Imgwight, 10)
	canvas.SetRGBA255(240, 240, 240, 255)
	canvas.Fill()*/

	// 加载size为108的字体
	err = canvas.LoadFontFace(t.TitleFont, 108)
	if err != nil {
		return
	}

	canvas.SetRGBA255(15, 15, 15, 255)

	// 绘制标题
	titley := 35 + canvas.FontHeight()*0.66
	canvas.DrawString(t.LeftTitle, 25, titley)
	// 加载size为54的字体
	err = canvas.LoadFontFace(t.TitleFont, 54)
	if err != nil {
		return
	}

	// 绘制一系列标题
	canvas.DrawString(t.LeftSubtitle, 25+3, titley+canvas.FontHeight()*1.6)

	lefttitlewight, _ := canvas.MeasureString(t.LeftSubtitle)
	canvas.DrawRectangle(25, titley+canvas.FontHeight()*1.85, lefttitlewight, 6)
	// 绘制插件开启状态
	if t.IsEnabled {
		canvas.SetRGBA255(35, 235, 35, 255)
	} else {
		canvas.SetRGBA255(235, 35, 35, 255)
	}
	canvas.Fill()
	canvas.SetRGBA255(15, 15, 15, 255)

	fw, _ := canvas.MeasureString(t.RightTitle)
	canvas.DrawString(t.RightTitle, DefaultWidth-40-fw, 30+canvas.FontHeight()*1.25)
	fw1, _ := canvas.MeasureString(t.RightSubtitle)
	canvas.DrawString(t.RightSubtitle, DefaultWidth-40-fw1, 30+canvas.FontHeight()*2.5)

	// 加载size为38的字体
	err = canvas.LoadFontFace(t.TextFont, 38)
	if err != nil {
		return
	}

	y := titley
	for _, text := range info {
		canvas.DrawString(text, 25.0, 1.5*titley+y+canvas.FontHeight())
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

	// 绘制遮罩
	/*canvas.DrawRectangle(0, rech/3*2, recw, rech/3)
	canvas.SetRGBA255(0, 0, 0, 153)
	canvas.Fill()*/

	// 绘制排名
	/*canvas.DrawRectangle(recw/10, 0, recw/10, (rech/4)-10)
	canvas.DrawRoundedRectangle(recw/10, 0, recw/10, (rech / 4), 8)*/
	if t.IsEnabled {
		canvas.DrawRectangle(0, rech*0.54, recw, rech-rech*0.54)
		// canvas.SetRGBA255(15, 175, 15, 255)
	} else {
		canvas.DrawRectangle(0, 0, recw, rech)
		// canvas.SetRGBA255(200, 15, 15, 255)
	}
	canvas.SetRGBA255(0, 0, 0, 183)
	canvas.Fill()

	// 绘制插件排名
	/*canvas.SetRGBA255(240, 240, 240, 255)
	var fw2 float64
	i, _ := strconv.Atoi(t.Rightsubtitle)
	if i > 99 {
		err = canvas.LoadFontFace(t.Fontpath, 24)
	} else {
		err = canvas.LoadFontFace(t.Fontpath, 28)
	}
	if err != nil {
		return
	}
	fw2, _ = canvas.MeasureString(t.Rightsubtitle)
	canvas.DrawString(t.Rightsubtitle, recw/10+((recw/10-fw2)/2), canvas.FontHeight()*3/8+(rech/8))*/

	// 绘制插件信息
	canvas.SetRGBA255(240, 240, 240, 255)
	err = canvas.LoadFontFace(t.TitleFont, 64)
	if err != nil {
		return
	}
	y := (rech * 0.56) + canvas.FontHeight()*0.95
	canvas.DrawString(t.LeftTitle, recw*0.04, y)

	err = canvas.LoadFontFace(t.TitleFont, 32)
	if err != nil {
		return
	}
	canvas.DrawString(t.LeftSubtitle, recw*0.04, y+canvas.FontHeight()*1.85)

	imgs = Fillet(canvas.Image(), 16)
	return
}

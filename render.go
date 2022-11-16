// Package rendercard ...
package rendercard

import (
	"image"
	"math/rand"
	"strconv"

	"github.com/Coloured-glaze/gg"
	"github.com/FloatTech/floatbox/img/writer"
	"github.com/FloatTech/zbputils/img"
)

const (
	Imgwight = 1272.0
)

// Titleinfo ...
type Titleinfo struct {
	Line int

	Lefttitle    string
	Leftsubtitle string

	Righttitle    string
	Rightsubtitle string

	Imgpath string

	Textpath  string
	Textpath2 string

	Status bool
}

// Drawtitle ...
func (t Titleinfo) Drawtitle() (imgs image.Image, err error) {
	// 创建图像
	canvas := gg.NewContext(int(Imgwight), 30+30+300+(t.Line*(256+30)))
	canvas.SetRGBA255(240, 240, 240, 255)
	canvas.Clear()

	// 标题背景1
	canvas.DrawRectangle(0, 30, Imgwight, 300)
	canvas.SetRGBA255(0, 0, 0, 153)
	canvas.Fill()

	// 标题背景2
	canvas.DrawRectangle(0, 30+40, Imgwight, 220)
	canvas.SetRGBA255(0, 0, 0, 153)
	canvas.Fill()

	// 加载size为108的字体
	err = canvas.LoadFontFace(t.Textpath, 108)
	if err != nil {
		return
	}

	// 绘制标题
	canvas.SetRGBA255(240, 240, 240, 255)
	canvas.DrawString(t.Lefttitle, 25, 30+40+55+canvas.FontHeight()-canvas.FontHeight()/3)

	// 加载size为54的字体
	err = canvas.LoadFontFace(t.Textpath, 54)
	if err != nil {
		return
	}

	// 绘制一系列标题
	canvas.DrawString(t.Leftsubtitle, 25+3, 30+40+165+canvas.FontHeight()/3)

	fw, _ := canvas.MeasureString(t.Righttitle)
	canvas.DrawString(t.Righttitle, Imgwight-25-fw-170-25, 30+40+25+15+canvas.FontHeight()+canvas.FontHeight()/4)
	fw1, _ := canvas.MeasureString(t.Rightsubtitle)
	canvas.DrawString(t.Rightsubtitle, Imgwight-25-fw1-170-25, 30+40+25+15+canvas.FontHeight()*2+canvas.FontHeight()/2)
	canvas.SetRGBA255(240, 240, 240, 255)

	// 加载icon并绘制
	var icon *img.Factory
	icon, err = img.LoadFirstFrame(t.Imgpath, 170, 170)
	if err != nil {
		return
	}
	canvas.DrawImage(icon.Im, int(Imgwight)-25-170, 30+40+25)
	imgs = canvas.Image()
	return
}

// Drawtitledtext ...
func (t Titleinfo) Drawtitledtext(info []string) (imgs []byte, err error) {
	line := len(info) - 1
	if line < 5 {
		line = 5
	}
	imgh := line*(38+20) + 220 + 10 + 30 + 10 + 50

	// 创建图像
	canvas := gg.NewContext(int(Imgwight), imgh)
	canvas.SetRGBA255(15, 15, 15, 204)
	canvas.Clear()

	// 加载icon
	var icon *img.Factory
	icon, err = img.LoadFirstFrame(t.Imgpath, 170, 170)
	if err != nil {
		return
	}
	canvas.DrawImage(icon.Im, int(Imgwight)-25-170, 25)

	// 绘制标题与内容的分割线
	canvas.DrawRectangle(0, 220, Imgwight, 10)
	canvas.SetRGBA255(240, 240, 240, 255)
	canvas.Fill()

	// 加载size为108的字体
	err = canvas.LoadFontFace(t.Textpath, 108)
	if err != nil {
		return
	}

	// 绘制标题
	canvas.SetRGBA255(240, 240, 240, 255)
	canvas.DrawString(t.Lefttitle, 25+40+25, 55+canvas.FontHeight()-canvas.FontHeight()/3)

	// 加载size为54的字体
	err = canvas.LoadFontFace(t.Textpath, 54)
	if err != nil {
		return
	}

	// 绘制插件开启状态
	canvas.DrawRectangle(25, 25, 40, 170)
	if t.Status {
		canvas.SetRGBA255(15, 175, 15, 255)
	} else {
		canvas.SetRGBA255(200, 15, 15, 255)
	}
	canvas.Fill()
	canvas.SetRGBA255(240, 240, 240, 255)

	// 绘制一系列标题
	canvas.DrawString(t.Leftsubtitle, 25+3+40+25, 165+canvas.FontHeight()/3)
	fw, _ := canvas.MeasureString(t.Righttitle)
	canvas.DrawString(t.Righttitle, Imgwight-25-fw-170-25, 25+15+canvas.FontHeight()+canvas.FontHeight()/4)
	fw1, _ := canvas.MeasureString(t.Rightsubtitle)
	canvas.DrawString(t.Rightsubtitle, Imgwight-25-fw1-170-25, 25+15+canvas.FontHeight()*2+canvas.FontHeight()/2)

	// 加载size为38的字体
	err = canvas.LoadFontFace(t.Textpath2, 38)
	if err != nil {
		return
	}

	x, y := 25.0, 25.0
	for i := 0; i < len(info); i++ {
		canvas.DrawString(info[i], x, y+220+10+canvas.FontHeight())
		y += 20 + canvas.FontHeight()
	}
	imgs, cl := writer.ToBytes(canvas.Image())
	cl()
	return
}

// Drawcard ...
func (t Titleinfo) Drawcard() (imgs image.Image, err error) {
	recw, rech := 384.0, 256.0
	canvas := gg.NewContext(int(recw), int(rech))
	// 绘制图片
	var banner *img.Factory
	banner, err = img.LoadFirstFrame(t.Imgpath, int(recw)*2, int(rech)*2)
	if err == nil {
		canvas.DrawImage(img.Size(banner.Im, int(recw), int(rech)).Im, 0, 0)
	} else {
		canvas.DrawRectangle(0, 0, recw, rech)
		canvas.SetRGBA255(rand.Intn(45)+165, rand.Intn(45)+165, rand.Intn(45)+165, 255)
		canvas.Fill()
	}

	// 绘制遮罩
	canvas.DrawRectangle(0, rech/3*2, recw, rech/3)
	canvas.SetRGBA255(0, 0, 0, 153)
	canvas.Fill()

	// 绘制排名
	canvas.DrawRectangle(recw/10, 0, recw/10, (rech/4)-10)
	canvas.DrawRoundedRectangle(recw/10, 0, recw/10, (rech / 4), 8)
	if t.Status {
		canvas.SetRGBA255(15, 175, 15, 255)
	} else {
		canvas.SetRGBA255(200, 15, 15, 255)
	}
	canvas.Fill()

	// 绘制插件排名
	canvas.SetRGBA255(240, 240, 240, 255)
	var fw2 float64
	i, _ := strconv.Atoi(t.Rightsubtitle)
	if i > 99 {
		err = canvas.LoadFontFace(t.Textpath, 24)
	} else {
		err = canvas.LoadFontFace(t.Textpath, 28)
	}
	if err != nil {
		return
	}
	fw2, _ = canvas.MeasureString(t.Rightsubtitle)
	canvas.DrawString(t.Rightsubtitle, recw/10+((recw/10-fw2)/2), canvas.FontHeight()*3/8+(rech/8))

	// 绘制插件信息
	canvas.SetRGBA255(240, 240, 240, 255)
	err = canvas.LoadFontFace(t.Textpath, 48)
	if err != nil {
		return
	}
	canvas.DrawString(t.Lefttitle, recw/32, (recw*0.475)+canvas.FontHeight()-canvas.FontHeight()/4)

	err = canvas.LoadFontFace(t.Textpath, 24)
	if err != nil {
		return
	}
	canvas.DrawString(t.Leftsubtitle, recw/32, (recw*0.475)+recw/6-canvas.FontHeight()/4)

	imgs = canvas.Image()
	return
}

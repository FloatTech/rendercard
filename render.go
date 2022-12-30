// Package rendercard ...
package rendercard

import (
	"errors"
	"image"
	"math/rand"
	"strings"

	"github.com/Coloured-glaze/gg"
	jpcolor "github.com/FloatTech/rendercard/color"
	"github.com/FloatTech/zbputils/img"
)

const (
	// Imgwight 图像宽
	Imgwight = 1272.0
)

// Titleinfo ...
type Titleinfo struct {
	// 行数
	Line int
	// 左侧两行文本
	Lefttitle    string
	Leftsubtitle string
	// 右侧两行文本
	Righttitle    string
	Rightsubtitle string
	// 图片路径
	Imgpath string
	// 字体路径
	Fontpath  string
	Fontpath2 string
	// 状态
	Status bool
}

// TextCardInfo ...
type TextCardInfo struct {
	// 卡片规格:宽度,默认600
	Width int
	// 卡片规格:高度,默认由Title+Text内容决定
	High int
	// 卡片规格:背景图
	Imgpath string
	// 标题字体
	FontOfTitle string
	// 正文的字体
	FontOfText string
	// 标题规格:标题内容
	Title string
	// 是否显示标题
	DisplayTitle bool
	// 标题规格:标题布局[Left|Center|Right],默认Left
	TitleSetting string
	// 正文规格:正文内容
	Text []string
	// 正文规格:正文要求
	//
	// true为每个元素按行显示,false按空格分割显示;
	TextSetting bool
}

// Drawtitle ...
func (t Titleinfo) Drawtitle() (imgs image.Image, err error) {
	// 创建图像
	canvas := gg.NewContext(int(Imgwight), 30+30+300+(t.Line*(256+30)))
	canvas.SetRGBA255(245, 245, 245, 255)
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
	err = canvas.LoadFontFace(t.Fontpath, 108)
	if err != nil {
		return
	}

	// 绘制标题
	canvas.SetRGBA255(250, 250, 250, 255)
	canvas.DrawString(t.Lefttitle, 25, 30+40+55+canvas.FontHeight()-canvas.FontHeight()/3)

	// 加载size为54的字体
	err = canvas.LoadFontFace(t.Fontpath, 54)
	if err != nil {
		return
	}

	canvas.SetRGBA255(250, 250, 250, 255)
	// 绘制副标题
	canvas.DrawString(t.Leftsubtitle, 25+3, 30+40+165+canvas.FontHeight()/3)

	// 加载icon并绘制
	var icon *img.Factory
	icon, err = img.LoadFirstFrame(t.Imgpath, 220, 220)
	if err != nil {
		return
	}
	canvas.DrawImage(icon.Im, int(Imgwight)-icon.W, 40+30)
	// 加载size为54的字体
	err = canvas.LoadFontFace(t.Fontpath, 72)
	if err != nil {
		return
	}
	fw, _ := canvas.MeasureString(t.Righttitle)
	canvas.DrawString(t.Righttitle, Imgwight-25-fw-float64(icon.W), 30+40+15+canvas.FontHeight()*1.25)
	fw1, _ := canvas.MeasureString(t.Rightsubtitle)
	canvas.DrawString(t.Rightsubtitle, Imgwight-25-fw1-float64(icon.W), 30+40+15+canvas.FontHeight()*2.75)

	imgs = canvas.Image()
	return
}

// Drawtitledtext ...
func (t Titleinfo) Drawtitledtext(info []string) (imgs image.Image, err error) {
	line := len(info)
	if line < 8 {
		line = 8
	}
	imgh := line*(28+20) + 220 + 50

	// 创建图像
	canvas := gg.NewContext(int(Imgwight), imgh)
	canvas.SetRGBA255(250, 250, 250, 255)
	canvas.Clear()

	// 加载icon
	var icon *img.Factory
	icon, err = img.LoadFirstFrame(t.Imgpath, 768, 768)
	if err != nil {
		return
	}
	canvas.DrawImage(icon.Im, Imgwight-icon.W, imgh-icon.H)

	// 绘制标题与内容的分割线
	/*canvas.DrawRectangle(0, 220, Imgwight, 10)
	canvas.SetRGBA255(240, 240, 240, 255)
	canvas.Fill()*/

	// 加载size为108的字体
	err = canvas.LoadFontFace(t.Fontpath, 108)
	if err != nil {
		return
	}

	canvas.SetRGBA255(15, 15, 15, 255)

	// 绘制标题
	titley := 35 + canvas.FontHeight()*0.66
	canvas.DrawString(t.Lefttitle, 25, titley)
	// 加载size为54的字体
	err = canvas.LoadFontFace(t.Fontpath, 54)
	if err != nil {
		return
	}

	// 绘制一系列标题
	canvas.DrawString(t.Leftsubtitle, 25+3, titley+canvas.FontHeight()*1.6)

	lefttitlewight, _ := canvas.MeasureString(t.Leftsubtitle)
	canvas.DrawRectangle(25, titley+canvas.FontHeight()*1.85, lefttitlewight, 6)
	// 绘制插件开启状态
	if t.Status {
		canvas.SetRGBA255(35, 235, 35, 255)
	} else {
		canvas.SetRGBA255(235, 35, 35, 255)
	}
	canvas.Fill()
	canvas.SetRGBA255(15, 15, 15, 255)

	fw, _ := canvas.MeasureString(t.Righttitle)
	canvas.DrawString(t.Righttitle, Imgwight-40-fw, 30+canvas.FontHeight()*1.25)
	fw1, _ := canvas.MeasureString(t.Rightsubtitle)
	canvas.DrawString(t.Rightsubtitle, Imgwight-40-fw1, 30+canvas.FontHeight()*2.5)

	// 加载size为38的字体
	err = canvas.LoadFontFace(t.Fontpath2, 38)
	if err != nil {
		return
	}

	x, y := 25.0, titley
	for _, text := range info {
		canvas.DrawString(text, x, 1.5*titley+y+canvas.FontHeight())
		y += 20 + canvas.FontHeight()
	}
	imgs = canvas.Image()
	return
}

// Drawcard ...
func (t Titleinfo) Drawcard() (imgs image.Image, err error) {
	recw, rech := 384.0, 256.0
	canvas := gg.NewContext(int(recw), int(rech))
	// 绘制图片
	var banner *img.Factory
	banner, err = img.LoadFirstFrame(t.Imgpath, int(recw)*2, int(rech)*2)
	r, g, b := jpcolor.Randcolor()

	if err == nil {
		canvas.DrawImage(img.Size(banner.Im, int(recw), int(rech)).Im, 0, 0)
	} else {
		canvas.DrawRectangle(0, 0, recw, rech)
		canvas.SetRGBA255(r, g, b, 255)
		canvas.Fill()
	}

	// 绘制遮罩
	/*canvas.DrawRectangle(0, rech/3*2, recw, rech/3)
	canvas.SetRGBA255(0, 0, 0, 153)
	canvas.Fill()*/

	// 绘制排名
	/*canvas.DrawRectangle(recw/10, 0, recw/10, (rech/4)-10)
	canvas.DrawRoundedRectangle(recw/10, 0, recw/10, (rech / 4), 8)*/
	if t.Status {
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
	err = canvas.LoadFontFace(t.Fontpath, 64)
	if err != nil {
		return
	}
	y := (rech * 0.56) + canvas.FontHeight()*0.9
	canvas.DrawString(t.Lefttitle, recw*0.04, y)

	err = canvas.LoadFontFace(t.Fontpath, 32)
	if err != nil {
		return
	}
	canvas.DrawString(t.Leftsubtitle, recw*0.04, y+canvas.FontHeight()*1.75)

	imgs = Fillet(canvas.Image(), 14)
	return
}

// DrawTextCard 绘制文字卡片
func (g TextCardInfo) DrawTextCard() (imgForCard image.Image, err error) {
	width := g.Width
	if width == 0 {
		width = 600
	}
	// 根据宽度获取高度
	fontOfText := g.FontOfText
	if fontOfText == "" {
		return nil, errors.New("请输入FontOfText参数")
	}
	// 正文数据
	textString := ""
	if g.TextSetting {
		textString = strings.Join(g.Text, "\n")
	} else {
		textString = strings.Join(g.Text, " ")
	}
	textImg, err := RenderText(textString, fontOfText, width-80, 38)
	if err != nil {
		return
	}
	textHigh := textImg.Bounds().Dy()
	// 计算图片高度
	imgHigh := g.High
	if imgHigh == 0 {
		if g.DisplayTitle {
			imgHigh = 30 + 100 + textHigh + 20
		} else {
			imgHigh = 20 + textHigh + 20
		}
	}
	// 创建画布
	canvas := gg.NewContext(width, imgHigh)
	// 随机背景色
	if g.Imgpath == "" {
		canvas.DrawRectangle(0, 0, float64(width), float64(imgHigh))
		canvas.SetRGBA255(rand.Intn(45)+165, rand.Intn(45)+165, rand.Intn(45)+165, 255)
		canvas.Fill()
	} else {
		banner, err := img.LoadFirstFrame(g.Imgpath, width, imgHigh)
		if err == nil {
			canvas.DrawImage(img.Size(banner.Im, width, imgHigh).Im, 0, 0)
		}
	}
	// 标题
	if g.DisplayTitle {
		fontOfTitle := g.FontOfTitle
		if fontOfTitle == "" {
			return nil, errors.New("请输入FontOfTitle参数")
		}
		err = canvas.LoadFontFace(fontOfTitle, 103)
		if err != nil {
			return
		}
		canvas.SetRGB(0, 0, 0)
		titleDx := 10.0
		widthOfTilte, titleDy := canvas.MeasureString(g.Title)
		switch g.TitleSetting {
		case "Left":
		case "Center":
			titleDx = (float64(width) - widthOfTilte) / 2
		case "Right":
			titleDx = float64(width) - widthOfTilte
		default:
			return nil, errors.New("TitleSetting 参数错误")
		}
		canvas.DrawString(g.Title, titleDx, titleDy+10)
		// 画横线
		canvas.DrawRoundedRectangle(10, 115, 580, 10, 2.5)
		canvas.SetRGB(0, 0, 0)
		canvas.Fill()
		// 内容
		canvas.DrawImage(textImg, 10, 130)
	} else {
		// 内容
		canvas.DrawImage(textImg, 10, 20)
	}
	// 制图
	imgForCard = canvas.Image()
	return
}

// RenderText 文字转图片 width 是图片宽度
func RenderText(text, font string, width, fontSize int) (txtPic image.Image, err error) {
	canvas := gg.NewContext(width, fontSize) // fake
	if err = canvas.LoadFontFace(font, float64(fontSize)); err != nil {
		return
	}
	buff := make([]string, 0)
	for _, s := range strings.Split(text, "\n") {
		line := ""
		for _, v := range s {
			length, _ := canvas.MeasureString(line)
			if int(length) <= width {
				line += string(v)
			} else {
				buff = append(buff, line)
				line = string(v)
			}
		}
		buff = append(buff, line)
	}
	_, h := canvas.MeasureString("好")
	canvas = gg.NewContext(width+int(h*2+0.5), int(float64(len(buff)*3+1)/2*h+0.5))
	canvas.SetRGB(0, 0, 0)
	if err = canvas.LoadFontFace(font, float64(fontSize)); err != nil {
		return
	}
	for i, v := range buff {
		if v != "" {
			canvas.DrawString(v, float64(width)*0.01, float64((i+1)*3)/2*h)
		}
	}
	return canvas.Image(), nil
}

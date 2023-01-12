// Package rendercard 渲染卡片
package rendercard

import (
	"errors"
	"image"
	"math/rand"
	"strings"

	"github.com/Coloured-glaze/gg"
	"github.com/FloatTech/zbputils/img"
	"github.com/FloatTech/zbputils/img/text"
)

var (
	// ErrNilTextFont ...
	ErrNilTextFont = errors.New("nil TextFont")
	// ErrNilTitleFont ...
	ErrNilTitleFont = errors.New("nil TitleFont")
)

// DrawTextCard 绘制文字卡片
func (c *Card) DrawTextCard() (imgForCard image.Image, err error) {
	width := c.Width
	if width == 0 {
		width = 600
	}
	// 根据宽度获取高度
	fontOfText := c.TextFont
	if fontOfText == "" {
		return nil, ErrNilTextFont
	}
	// 正文数据
	textString := ""
	if c.IsTextSplitPerElement {
		textString = strings.Join(c.Text, "\n")
	} else {
		textString = strings.Join(c.Text, " ")
	}
	textImg, err := text.Render(textString, fontOfText, width-80, 38)
	if err != nil {
		return
	}
	textHigh := textImg.Bounds().Dy()
	// 计算图片高度
	imgHigh := c.Height
	if imgHigh == 0 {
		if c.CanTitleShown {
			imgHigh = 30 + 100 + textHigh + 20
		} else {
			imgHigh = 20 + textHigh + 20
		}
	}
	// 创建画布
	canvas := gg.NewContext(width, imgHigh)
	// 随机背景色
	if c.BackgroundImage == "" {
		canvas.DrawRectangle(0, 0, float64(width), float64(imgHigh))
		canvas.SetRGBA255(rand.Intn(45)+165, rand.Intn(45)+165, rand.Intn(45)+165, 255)
		canvas.Fill()
	} else {
		banner, err := img.LoadFirstFrame(c.BackgroundImage, width, imgHigh)
		if err == nil {
			canvas.DrawImage(img.Size(banner.Im, width, imgHigh).Im, 0, 0)
		}
	}
	// 标题
	if c.CanTitleShown {
		fontOfTitle := c.TitleFont
		if fontOfTitle == "" {
			return nil, ErrNilTitleFont
		}
		err = canvas.LoadFontFace(fontOfTitle, 103)
		if err != nil {
			return
		}
		canvas.SetRGB(0, 0, 0)
		titleDx := 10.0
		widthOfTilte, titleDy := canvas.MeasureString(c.Title)
		switch c.TitleAlign {
		case NilAlign:
		case AlignLeft:
		case AlignCenter:
			titleDx = (float64(width) - widthOfTilte) / 2
		case AlignRight:
			titleDx = float64(width) - widthOfTilte
		}
		canvas.DrawString(c.Title, titleDx, titleDy+10)
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

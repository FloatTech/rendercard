// Package rendercard 渲染卡片
package rendercard

import (
	"errors"
	"image"
	"math/rand"

	"github.com/FloatTech/gg"
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
	if c.TextFontData == nil {
		return nil, ErrNilTextFont
	}
	// 正文数据
	texts, err := Truncate(c.TextFontData, c.Text, float64(width)-80, 38)
	if err != nil {
		return
	}
	textHigh := float64(len(texts)+1) * 38 * 72 / 96 * 1.5
	// 计算图片高度
	imgHigh := c.Height
	if imgHigh == 0 {
		if c.CanTitleShown {
			imgHigh = 30 + 100 + int(textHigh) + 20
		} else {
			imgHigh = 20 + int(textHigh) + 20
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
		banner, err := gg.LoadImage(c.BackgroundImage)
		if err == nil {
			if float64(banner.Bounds().Dy())/float64(banner.Bounds().Dx()) < float64(canvas.H())/float64(canvas.W()) {
				sc := float64(canvas.H()) / float64(banner.Bounds().Dy())
				canvas.ScaleAbout(sc, sc, float64(canvas.W())/2, float64(canvas.H())/2)
				canvas.DrawImageAnchored(banner, canvas.W()/2, canvas.H()/2, 0.5, 0.5)
			} else {
				sc := float64(canvas.W()) / float64(banner.Bounds().Dx())
				canvas.Scale(sc, sc)
				canvas.DrawImage(banner, 0, 0)
			}
			canvas.Identity()
		}
	}
	y := 0.0
	// 标题
	if c.CanTitleShown {
		if c.TitleFontData == nil {
			return nil, ErrNilTitleFont
		}
		err = canvas.ParseFontFace(c.TitleFontData, 103)
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
		y = 130 + 38*72/96*1.5
	} else {
		// 内容
		y = 20 + 38*72/96*1.5
	}
	err = canvas.ParseFontFace(c.TextFontData, 38)
	if err != nil {
		return
	}
	for _, s := range texts {
		canvas.DrawStringAnchored(s, 10, y, 0, 0.5)
		y += canvas.FontHeight() * 1.5
	}
	// 制图
	imgForCard = canvas.Image()
	return
}

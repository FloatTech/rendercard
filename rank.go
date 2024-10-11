package rendercard

import (
	"image"
	"image/color"
	"strconv"
	"sync"

	"github.com/FloatTech/gg"
)

// DrawRankingCard ...
func DrawRankingCard(fontdata []byte, title string, toplefttext, bottomlefttext, righttext []string, avatars []image.Image) (img image.Image, err error) {
	line := len(avatars)
	const w = 672
	h := 64 + (80+14)*line + 20 - 14
	canvas := gg.NewContext(w, h)
	canvas.SetRGBA255(255, 255, 255, 255)
	canvas.Clear()

	ac := gg.NewLinearGradient(16, float64(h)/2, w-16, float64(h)/2)
	ac.AddColorStop(0, color.NRGBA{255, 255, 255, 95})
	ac.AddColorStop(0.167, color.NRGBA{255, 255, 255, 127})
	ac.AddColorStop(0.334, color.NRGBA{255, 255, 255, 159})
	ac.AddColorStop(0.5, color.NRGBA{255, 255, 255, 255})
	ac.AddColorStop(1, color.NRGBA{255, 255, 255, 255})

	cardh, cardw := 80.0, 640.0
	cardspac := 14.0
	hspac := 64.0

	wg := &sync.WaitGroup{}
	wg.Add(line)
	cardimgs := make([]image.Image, line)
	for i := 0; i < line; i++ {
		go func(i int) {
			defer wg.Done()
			card := gg.NewContext(w, h)

			card.NewSubPath()

			card.MoveTo(16+cardh/2, hspac+(cardspac+cardh)*float64(i))

			card.LineTo(16+cardw-16, hspac+(cardspac+cardh)*float64(i))
			card.DrawArc(16+cardw-16, hspac+(cardspac+cardh)*float64(i)+16, 16, gg.Radians(-90), gg.Radians(0))
			card.LineTo(16+cardw, hspac+(cardspac+cardh)*float64(i)+cardh-16)
			card.DrawArc(16+cardw-16, hspac+(cardspac+cardh)*float64(i)+cardh-16, 16, gg.Radians(0), gg.Radians(90))
			card.LineTo(16+cardh/2, hspac+(cardspac+cardh)*float64(i)+cardh)

			card.DrawArc(16+cardh/2, hspac+(cardspac+cardh)*float64(i)+cardh-cardh/2, cardh/2, gg.Radians(90), gg.Radians(270))

			card.ClosePath()

			card.ClipPreserve()

			avatar := avatars[i]

			card.ScaleAbout(cardw/2/float64(avatar.Bounds().Dx()), cardw/2/float64(avatar.Bounds().Dy()), 16, hspac+(cardspac+cardh)*float64(i)+cardh/2)

			card.DrawImageAnchored(avatar, 16, int(hspac+(cardspac+cardh)*float64(i)+cardh/2), 0, 0.5)
			card.Identity()
			card.ResetClip()

			card.SetFillStyle(ac)
			card.FillPreserve()

			card.SetRGBA255(0, 0, 0, 255)
			card.Stroke()

			card.DrawCircle(16+10+30, hspac+(cardspac+cardh)*float64(i)+cardh/2, 30)
			card.ClipPreserve()
			card.ScaleAbout(60.0/float64(avatar.Bounds().Dx()), 60.0/float64(avatar.Bounds().Dy()), 16+10+30, hspac+(cardspac+cardh)*float64(i)+cardh/2)
			card.DrawImageAnchored(avatar, 16+10+30, int(hspac+(cardspac+cardh)*float64(i)+cardh/2), 0.5, 0.5)
			card.Identity()

			card.ResetClip()
			card.Stroke()

			card.SetRGB255(RandJPColor())
			card.DrawCircle(w-16-8-25, hspac+(cardspac+cardh)*float64(i)+cardh/2, 25)
			card.Fill()
			cardimgs[i] = card.Image()
		}(i)
	}

	canvas.SetRGBA255(0, 0, 0, 255)
	canvas.ParseFontFace(fontdata, 32)
	canvas.DrawStringAnchored(title, w/2, 64/2, 0.5, 0.5)

	err = canvas.ParseFontFace(fontdata, 20)
	if err != nil {
		return
	}
	wg.Wait()
	for i := 0; i < 10; i++ {
		canvas.DrawImage(cardimgs[i], 0, 0)
		canvas.DrawStringAnchored(toplefttext[i], 16+10+60+10, hspac+(cardspac+cardh)*float64(i)+cardh*3/8, 0, 0.5)
	}

	canvas.SetRGBA255(63, 63, 63, 255)
	err = canvas.ParseFontFace(fontdata, 12)
	if err != nil {
		return
	}
	for i := 0; i < 10; i++ {
		canvas.DrawStringAnchored(bottomlefttext[i], 16+10+60+10, hspac+(cardspac+cardh)*float64(i)+cardh*5/8, 0, 0.5)
	}
	canvas.SetRGBA255(0, 0, 0, 255)
	err = canvas.ParseFontFace(fontdata, 24)
	if err != nil {
		return
	}
	for i := 0; i < 10; i++ {
		canvas.DrawStringAnchored(righttext[i], w-16-8-50-8, hspac+(cardspac+cardh)*float64(i)+cardh/2, 1, 0.5)
	}

	canvas.SetRGBA255(255, 255, 255, 255)
	err = canvas.ParseFontFace(fontdata, 28)
	if err != nil {
		return
	}
	for i := 0; i < 10; i++ {
		canvas.DrawStringAnchored(strconv.Itoa(i+1), w-16-8-25, hspac+(cardspac+cardh)*float64(i)+cardh/2, 0.5, 0.5)
	}

	img = canvas.Image()
	return
}

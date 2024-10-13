package rendercard

import (
	"image"
	"image/color"
	"strconv"
	"sync"

	"github.com/FloatTech/gg"
)

// RankInfo ...
type RankInfo struct {
	Avatar         image.Image
	TopLeftText    string
	BottomLeftText string
	RightText      string
}

// DrawRankingCard ...
func DrawRankingCard(fontdata []byte, title string, rankinfo []*RankInfo) (img image.Image, err error) {
	line := len(rankinfo)
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

	cardh, cardw := 80, 640
	cardspac := 14
	hspac, wspac := 64.0, 16.0
	r := 16.0

	wg := &sync.WaitGroup{}
	wg.Add(line)
	cardimgs := make([]image.Image, line)
	for i := 0; i < line; i++ {
		go func(i int) {
			defer wg.Done()
			card := gg.NewContext(w, cardh)

			card.NewSubPath()
			card.MoveTo(wspac+float64(cardh)/2, 0)

			card.LineTo(wspac+float64(cardw)-r, 0)
			card.DrawArc(wspac+float64(cardw)-r, r, r, gg.Radians(-90), gg.Radians(0))
			card.LineTo(wspac+float64(cardw), float64(cardh)-r)
			card.DrawArc(wspac+float64(cardw)-r, float64(cardh)-r, r, gg.Radians(0), gg.Radians(90))
			card.LineTo(wspac+float64(cardh)/2, float64(cardh))

			card.DrawArc(wspac+float64(cardh)/2, float64(cardh)/2, float64(cardh)/2, gg.Radians(90), gg.Radians(270))
			card.ClosePath()

			card.ClipPreserve()

			avatar := rankinfo[i].Avatar

			card.ScaleAbout(float64(cardw)/2/float64(avatar.Bounds().Dx()), float64(cardw)/2/float64(avatar.Bounds().Dy()), wspac, float64(cardh)/2)

			card.DrawImageAnchored(avatar, int(wspac), cardh/2, 0, 0.5)
			card.Identity()
			card.ResetClip()

			card.SetFillStyle(ac)
			card.FillPreserve()

			card.SetRGBA255(0, 0, 0, 255)
			card.Stroke()
			card.DrawLine(wspac+float64(cardh)/2, float64(cardh), wspac+float64(cardw)-r, float64(cardh))
			card.Stroke()

			card.DrawCircle(wspac+10+30, float64(cardh)/2, 30)
			card.ClipPreserve()
			card.ScaleAbout(60.0/float64(avatar.Bounds().Dx()), 60.0/float64(avatar.Bounds().Dy()), wspac+10+30, float64(cardh)/2)
			card.DrawImageAnchored(avatar, int(wspac)+10+30, cardh/2, 0.5, 0.5)
			card.Identity()

			card.ResetClip()
			card.SetRGBA255(0, 0, 0, 127)
			card.Stroke()

			card.SetRGB255(RandJPColor())
			card.DrawCircle(wspac+float64(cardw-8-25), float64(cardh)/2, 25)
			card.Fill()
			cardimgs[i] = card.Image()
		}(i)
	}

	canvas.SetRGBA255(0, 0, 0, 255)
	err = canvas.ParseFontFace(fontdata, 32)
	if err != nil {
		return
	}
	canvas.DrawStringAnchored(title, w/2, 64/2, 0.5, 0.5)

	err = canvas.ParseFontFace(fontdata, 22)
	if err != nil {
		return
	}
	wg.Wait()
	for i := 0; i < line; i++ {
		canvas.DrawImageAnchored(cardimgs[i], w/2, int(hspac)+((cardh+cardspac)*i), 0.5, 0)
		canvas.DrawStringAnchored(rankinfo[i].TopLeftText, wspac+10+60+10, hspac+float64((cardspac+cardh)*i+cardh*3/8), 0, 0.5)
	}

	canvas.SetRGBA255(63, 63, 63, 255)
	err = canvas.ParseFontFace(fontdata, 12)
	if err != nil {
		return
	}
	for i := 0; i < line; i++ {
		canvas.DrawStringAnchored(rankinfo[i].BottomLeftText, wspac+10+60+10, hspac+float64((cardspac+cardh)*i+cardh*5/8), 0, 0.5)
	}
	canvas.SetRGBA255(0, 0, 0, 255)
	err = canvas.ParseFontFace(fontdata, 20)
	if err != nil {
		return
	}
	for i := 0; i < line; i++ {
		canvas.DrawStringAnchored(rankinfo[i].RightText, w-wspac-8-50-8, hspac+float64((cardspac+cardh)*i+cardh/2), 1, 0.5)
	}

	canvas.SetRGBA255(255, 255, 255, 255)
	err = canvas.ParseFontFace(fontdata, 28)
	if err != nil {
		return
	}
	for i := 0; i < line; i++ {
		canvas.DrawStringAnchored(strconv.Itoa(i+1), w-wspac-8-25, hspac+float64((cardspac+cardh)*i+cardh/2), 0.5, 0.5)
	}

	img = canvas.Image()
	return
}

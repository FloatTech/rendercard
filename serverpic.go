package rendercard

import (
	"image"
	"sync"

	"github.com/FloatTech/floatbox/math"

	"github.com/FloatTech/gg"
)

// PluginInfo ...
type PluginInfo struct {
	Name   string
	Brief  string
	Status bool
}

// RenderServerPic ...
func RenderServerPic(pluginlist []*PluginInfo, torussd, glowsd []byte, zbplogopath string, serverlistlogo image.Image) (img image.Image, err error) {
	logo, err := gg.LoadImage(zbplogopath)
	if err != nil {
		return
	}
	listnum := len(pluginlist)
	ln := math.Ceil(listnum, 3)
	w := (290+24)*3 + 24
	h := serverlistlogo.Bounds().Dy() + ln*(80+16) + serverlistlogo.Bounds().Dy()/3
	canvas := gg.NewContext(w, h)

	canvas.SetRGBA255(235, 235, 235, 255)
	canvas.Clear()

	canvas.SetRGBA255(135, 144, 173, 255)
	canvas.NewSubPath()
	canvas.MoveTo(0, 0)
	canvas.LineTo(float64(canvas.W()), 140)
	canvas.LineTo(float64(canvas.W()), 0)
	canvas.ClosePath()
	canvas.Fill()

	canvas.NewSubPath()
	canvas.MoveTo(float64(canvas.W()), float64(canvas.H()))
	canvas.LineTo(0, float64(canvas.H()))
	canvas.LineTo(0, float64(canvas.H())-140)
	canvas.ClosePath()
	canvas.Fill()

	canvas.SetRGBA255(247, 177, 170, 255)
	canvas.NewSubPath()
	canvas.MoveTo(0, 0)
	canvas.LineTo(float64(canvas.W()), 0)
	canvas.LineTo(float64(canvas.W()), 70)
	canvas.LineTo(0, 270)
	canvas.ClosePath()
	canvas.Fill()

	canvas.NewSubPath()
	canvas.MoveTo(float64(canvas.W()), float64(canvas.H()))
	canvas.LineTo(0, float64(canvas.H()))
	canvas.LineTo(0, float64(canvas.H())-70)
	canvas.LineTo(float64(canvas.W()), float64(canvas.H())-270)
	canvas.ClosePath()
	canvas.Fill()

	canvas.SetRGBA255(186, 113, 132, 255)
	canvas.NewSubPath()
	canvas.MoveTo(0, 0)
	canvas.LineTo(float64(canvas.W()), 0)
	canvas.LineTo(float64(canvas.W()), 35)
	canvas.LineTo(0, 160)
	canvas.ClosePath()
	canvas.Fill()

	canvas.NewSubPath()
	canvas.MoveTo(float64(canvas.W()), float64(canvas.H()))
	canvas.LineTo(0, float64(canvas.H()))
	canvas.LineTo(0, float64(canvas.H())-35)
	canvas.LineTo(float64(canvas.W()), float64(canvas.H())-160)
	canvas.ClosePath()
	canvas.Fill()

	canvas.ScaleAbout(0.5, 0.5, float64(canvas.W()/4), 0)
	canvas.DrawImageAnchored(logo, canvas.W()/4, 0, 0.5, 0)
	canvas.Identity()

	canvas.DrawImageAnchored(serverlistlogo, canvas.W()/4*3, 0, 0.5, 0)

	cardimgs := make([]image.Image, 4)

	wg := sync.WaitGroup{}
	cardsnum := ln / 4 * 3
	wg.Add(4)
	for i := 0; i < 3; i++ {
		a := i * cardsnum
		b := (i + 1) * cardsnum
		go func(i int, list []*PluginInfo) {
			defer wg.Done()
			if b == 0 {
				return
			}
			cardimgs[i], err = renderinfocards(torussd, glowsd, list)
			if err != nil {
				return
			}
		}(i, pluginlist[a:b])
	}
	go func() {
		defer wg.Done()
		cardimgs[3], err = renderinfocards(torussd, glowsd, pluginlist[cardsnum*3:])
		if err != nil {
			return
		}
	}()
	wg.Wait()
	spacing := 0
	for i := 0; i < len(cardimgs); i++ {
		if cardimgs != nil {
			canvas.DrawImage(cardimgs[i], 0, serverlistlogo.Bounds().Dy()+spacing)
			spacing += cardimgs[i].Bounds().Dy()
		}
	}

	img = canvas.Image()
	return
}

func renderinfocards(torussd, glowsd []byte, plugininfos []*PluginInfo) (img image.Image, err error) {
	w := (290+24)*3 + 24
	cardnum := len(plugininfos)
	h := math.Ceil(cardnum, 3) * (80 + 16)
	cardw, cardh := 290.0, 80.0
	spacingw, spacingh := 24.0, 16.0
	canvas := gg.NewContext(w, h)
	beginw, beginh := 24.0, 0.0
	for i := 0; i < cardnum; i++ {
		canvas.SetRGBA255(204, 51, 51, 255)
		if plugininfos[i].Status {
			canvas.SetRGBA255(136, 178, 0, 255)
		}
		canvas.DrawRoundedRectangle(beginw, beginh, cardw/2, cardh, 16)
		canvas.Fill()

		canvas.SetRGBA255(34, 26, 33, 255)
		canvas.DrawRoundedRectangle(beginw+10, beginh, cardw-10, cardh, 16)
		canvas.Fill()
		beginw += cardw + spacingw
		if (i+1)%3 == 0 {
			beginw = spacingw
			beginh += cardh + spacingh
		}
	}

	err = canvas.ParseFontFace(torussd, 36)
	if err != nil {
		return
	}
	canvas.SetRGBA255(235, 235, 235, 255)
	beginw, beginh = 24.0, 0.0
	for i := 0; i < cardnum; i++ {
		canvas.DrawStringAnchored(plugininfos[i].Name, beginw+14, beginh+canvas.FontHeight()/2+4, 0, 0.5)
		beginw += cardw + spacingw
		if (i+1)%3 == 0 {
			beginw = spacingw
			beginh += cardh + spacingh
		}
	}
	err = canvas.ParseFontFace(glowsd, 16)
	if err != nil {
		return
	}
	beginw, beginh = 24.0, 0.0
	for i := 0; i < cardnum; i++ {
		canvas.DrawStringAnchored(plugininfos[i].Brief, beginw+14, beginh+cardh-canvas.FontHeight()-4, 0, 0.5)
		beginw += cardw + spacingw
		if (i+1)%3 == 0 {
			beginw = spacingw
			beginh += cardh + spacingh
		}
	}
	img = canvas.Image()
	return
}

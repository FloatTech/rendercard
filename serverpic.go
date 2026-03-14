package rendercard

import (
	"image"
	"image/color"
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

	canvas.SetRGBA255(235, 235, 235, 127)
	canvas.Clear()

	halfalphamask := canvas.AsMask()

	canvas.SetRGBA255(235, 235, 235, 255)
	canvas.Clear()

	err = canvas.SetMask(halfalphamask)
	if err != nil {
		return
	}

	canvas.DrawImageAnchored(logo, canvas.W()/2, canvas.H()/2, 0.5, 0.5)

	canvas.ResetClip()

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
		fade := gg.NewLinearGradient(beginw+10, 0, beginw+cardw-10, 0)
		fade.AddColorStop(1, color.NRGBA{204, 51, 51, 255})
		fade.AddColorStop(0.8, color.NRGBA{204, 51, 51, 191})
		fade.AddColorStop(0.7, color.NRGBA{204, 51, 51, 63})
		fade.AddColorStop(0.6, color.NRGBA{204, 51, 51, 0})
		fade.AddColorStop(0, color.NRGBA{204, 51, 51, 0})
		statusimage := drawEnableOrDisableImage(plugininfos[i].Status)

		if plugininfos[i].Status {
			canvas.SetRGBA255(136, 178, 0, 255)
			fade = gg.NewLinearGradient(beginw+10, 0, beginw+cardw-10, 0)
			fade.AddColorStop(1, color.NRGBA{136, 178, 0, 255})
			fade.AddColorStop(0.8, color.NRGBA{136, 178, 0, 191})
			fade.AddColorStop(0.7, color.NRGBA{136, 178, 0, 63})
			fade.AddColorStop(0.6, color.NRGBA{136, 178, 0, 0})
			fade.AddColorStop(0, color.NRGBA{136, 178, 0, 0})
			statusimage = drawEnableOrDisableImage(plugininfos[i].Status)
		}

		canvas.DrawRoundedRectangle(beginw+10, beginh, cardw-10, cardh, 16)
		canvas.Clip()
		canvas.InvertMask()

		canvas.DrawRoundedRectangle(beginw, beginh, cardw/2, cardh, 16)
		canvas.Fill()

		canvas.ResetClip()

		canvas.SetRGBA255(255, 255, 255, 217)
		canvas.DrawRoundedRectangle(beginw+10, beginh, cardw-10, cardh, 16)
		canvas.FillPreserve()
		canvas.SetFillStyle(fade)
		canvas.Fill()
		canvas.SetFillStyle(nil)

		alphacanvas := gg.NewContext(w, h)
		alphacanvas.SetRGBA255(255, 255, 255, 93)
		alphacanvas.Clear()

		err = canvas.SetMask(alphacanvas.AsMask())
		if err != nil {
			return
		}
		canvas.DrawImageAnchored(statusimage, int(beginw+cardw)-10-5, int(beginh+cardh/2), 1, 0.5)
		canvas.ResetClip()

		err = canvas.SetMask(alphacanvas.AsMask())
		if err != nil {
			return
		}
		canvas.DrawRoundedRectangle(beginw+10, beginh, cardw-10, cardh, 16)
		canvas.Clip()
		canvas.DrawCircle(beginw+cardw-10-5-50/2, beginh+cardh/2, float64(cardh-10)/2)
		canvas.SetLineWidth(50 * 0.185)
		canvas.Stroke()
		canvas.ResetClip()

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
	canvas.SetRGBA255(20, 20, 20, 255)
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

func drawEnableOrDisableImage(status bool) image.Image {
	dc := gg.NewContext(50, 50)
	dcwf, dchf := float64(dc.W()), float64(dc.H())
	fac := 0.185
	if status {
		dc.RotateAbout(gg.Radians(-45), dcwf/2, dchf/2)
		dc.DrawRectangle(dcwf*fac+dcwf*fac/2, (dchf-dchf/2)/2-1, dcwf, dcwf/2-dcwf*fac)
		dc.Clip()
		dc.InvertMask()
		dc.DrawRectangle(dcwf*fac/2, (dchf-dchf/2)/2, dcwf-dcwf*fac, dchf/2)
		dc.SetRGBA255(255, 255, 255, 255)
		dc.Fill()
		dc.ResetClip()
		return dc.Image()
	}
	dc.RotateAbout(gg.Radians(-45), dcwf/2, dchf/2)
	dc.DrawRectangle((dcwf-dcwf*fac)/2, 0, dcwf*fac, dchf)
	dc.SetRGBA255(255, 255, 255, 255)
	dc.Fill()
	dc.DrawRectangle(0, (dcwf-dcwf*fac)/2, dchf, dcwf*fac)
	dc.Fill()
	return dc.Image()
}

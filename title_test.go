package rendercard

import (
	"crypto/md5"
	"encoding/hex"
	"image/png"
	"os"
	"testing"

	"github.com/FloatTech/gg"
	"github.com/stretchr/testify/assert"
)

func TestDrawTitle(t *testing.T) {
	glows, err := os.ReadFile("GlowSansSC-Normal-ExtraBold.ttf")
	if err != nil {
		t.Fatal(err)
	}
	impact, err := os.ReadFile("Impact.ttf")
	if err != nil {
		t.Fatal(err)
	}
	img, err := (&Title{
		Line:          0,
		LeftTitle:     "服务列表",
		LeftSubtitle:  "service_list",
		RightTitle:    "FloatTech",
		RightSubtitle: "ZeroBot-Plugin",
		TitleFontData: glows,
		TextFontData:  impact,
		ImagePath:     ".github/warma.png",
	}).DrawTitle()
	if err != nil {
		t.Fatal(err)
	}
	err = gg.SavePNG(".github/DrawTitle.png", img)
	if err != nil {
		t.Fatal(err)
	}
	h := md5.New()
	err = png.Encode(h, img)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "5450734a22186aa176387ff2a91d9a7e", hex.EncodeToString(h.Sum(nil)))
}

func TestDrawTitleWithText(t *testing.T) {
	glows, err := os.ReadFile("GlowSansSC-Normal-ExtraBold.ttf")
	if err != nil {
		t.Fatal(err)
	}
	impact, err := os.ReadFile("Impact.ttf")
	if err != nil {
		t.Fatal(err)
	}
	img, err := (&Title{
		Line:          0,
		IsEnabled:     true,
		LeftTitle:     "ServiceName",
		LeftSubtitle:  "简介",
		RightTitle:    "FloatTech",
		RightSubtitle: "ZeroBot-Plugin",
		TitleFontData: impact,
		TextFontData:  glows,
		ImagePath:     ".github/warma.png",
	}).DrawTitleWithText([]string{"one", "two", "san", "si"})
	if err != nil {
		t.Fatal(err)
	}
	err = gg.SavePNG(".github/DrawTitleWithText.png", img)
	if err != nil {
		t.Fatal(err)
	}
	h := md5.New()
	err = png.Encode(h, img)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "49c0bb06f1b8e16833e73d2bdcbe01ce", hex.EncodeToString(h.Sum(nil)))
}

func TestDrawCard(t *testing.T) {
	glows, err := os.ReadFile("GlowSansSC-Normal-ExtraBold.ttf")
	if err != nil {
		t.Fatal(err)
	}
	impact, err := os.ReadFile("Impact.ttf")
	if err != nil {
		t.Fatal(err)
	}
	img, err := (&Title{
		Line:                 0,
		IsEnabled:            true,
		LeftTitle:            "ServiceName",
		LeftSubtitle:         "简介",
		TitleFontData:        impact,
		TextFontData:         glows,
		ImagePath:            ".github/warma.png",
		TitleFontOffsetPoint: -6,
		TextFontOffsetPoint:  -6,
	}).DrawCard()
	if err != nil {
		t.Fatal(err)
	}
	err = gg.SavePNG(".github/DrawCard.png", img)
	if err != nil {
		t.Fatal(err)
	}
	h := md5.New()
	err = png.Encode(h, img)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "deba289674306c6d674c0c07f6ae458c", hex.EncodeToString(h.Sum(nil)))
}

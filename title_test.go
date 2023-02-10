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
	assert.Equal(t, "28f9f68f3ef7987a9d189b52b9a8b399", hex.EncodeToString(h.Sum(nil)))
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
	assert.Equal(t, "bf5ce310c0a565477fa83904a506c08d", hex.EncodeToString(h.Sum(nil)))
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
	assert.Equal(t, "4e1b4c439a78d833b0931655d736b02f", hex.EncodeToString(h.Sum(nil)))
}

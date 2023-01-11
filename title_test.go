package rendercard

import (
	"crypto/md5"
	"encoding/hex"
	"image/png"
	"testing"

	"github.com/Coloured-glaze/gg"
	"github.com/stretchr/testify/assert"
)

func TestDrawTitle(t *testing.T) {
	img, err := (&Title{
		Line:          0,
		LeftTitle:     "服务列表",
		LeftSubtitle:  "service_list",
		RightTitle:    "FloatTech",
		RightSubtitle: "ZeroBot-Plugin",
		TitleFont:     "GlowSansSC-Normal-ExtraBold.ttf",
		TextFont:      "Impact.ttf",
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
	img, err := (&Title{
		Line:          0,
		IsEnabled:     true,
		LeftTitle:     "服务名",
		LeftSubtitle:  "instruction",
		RightTitle:    "FloatTech",
		RightSubtitle: "ZeroBot-Plugin",
		TitleFont:     "GlowSansSC-Normal-ExtraBold.ttf",
		TextFont:      "Impact.ttf",
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
	assert.Equal(t, "3f18e919e69e38b2302d8b2fc31c4cb2", hex.EncodeToString(h.Sum(nil)))
}

func TestDrawCard(t *testing.T) {
	img, err := (&Title{
		Line:         0,
		IsEnabled:    true,
		LeftTitle:    "服务名",
		LeftSubtitle: "instruction",
		TitleFont:    "GlowSansSC-Normal-ExtraBold.ttf",
		TextFont:     "Impact.ttf",
		ImagePath:    ".github/warma.png",
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
	assert.Equal(t, "596186f97c3a4126a9d3a7b30b959810", hex.EncodeToString(h.Sum(nil)))
}

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
		LeftTitle:     "ServiceName",
		LeftSubtitle:  "简介",
		RightTitle:    "FloatTech",
		RightSubtitle: "ZeroBot-Plugin",
		TitleFont:     "Impact.ttf",
		TextFont:      "GlowSansSC-Normal-ExtraBold.ttf",
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
	img, err := (&Title{
		Line:         0,
		IsEnabled:    true,
		LeftTitle:    "ServiceName",
		LeftSubtitle: "简介",
		TitleFont:    "Impact.ttf",
		TextFont:     "GlowSansSC-Normal-ExtraBold.ttf",
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
	assert.Equal(t, "c7669d5cc76599abbd8d22b507d2102a", hex.EncodeToString(h.Sum(nil)))
}

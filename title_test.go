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
		LeftTitle:     "FUWULIEBIAO",
		LeftSubtitle:  "service_list",
		RightTitle:    "FloatTech",
		RightSubtitle: "ZeroBot-Plugin",
		TitleFont:     "Impact.ttf",
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
	assert.Equal(t, "7a7ccec658b1acaf1a018d47357b1f62", hex.EncodeToString(h.Sum(nil)))
}

func TestDrawTitleWithText(t *testing.T) {
	img, err := (&Title{
		Line:          0,
		IsEnabled:     true,
		LeftTitle:     "NAME",
		LeftSubtitle:  "instruction",
		RightTitle:    "FloatTech",
		RightSubtitle: "ZeroBot-Plugin",
		TitleFont:     "Impact.ttf",
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
	assert.Equal(t, "fdff61acd070ab6806cb587f5fad527c", hex.EncodeToString(h.Sum(nil)))
}

func TestDrawCard(t *testing.T) {
	img, err := (&Title{
		Line:         0,
		IsEnabled:    true,
		LeftTitle:    "NAME",
		LeftSubtitle: "instruction",
		TitleFont:    "Impact.ttf",
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
	assert.Equal(t, "84d9e5f39b921c3866785c9ae39f74d7", hex.EncodeToString(h.Sum(nil)))
}

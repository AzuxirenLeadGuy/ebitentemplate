package coregame

import (
	"fmt"
	"image/color"
	rd "math/rand"

	v2 "github.com/hajimehoshi/ebiten/v2"
	util "github.com/hajimehoshi/ebiten/v2/ebitenutil"
	ip "github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	r, g, b       byte
	Width, Height int
}

func reflect_0_255(val int32) byte {
	if val < 0 {
		val = -val
	} else if val > 255 {
		val = 510 - val
	}
	return byte(val)
}

func get_part(partid, bit_size, val int32) int32 {
	portion := (val >> (bit_size * partid))
	portion = portion & ((1 << bit_size) - 1)
	portion = portion - (1 << (bit_size - 1))
	return int32(portion)
}

const rbits = 4

func (g *Game) Update() error {
	if ip.IsMouseButtonJustReleased(v2.MouseButtonLeft) {
		x := rd.Int31n(1 << (rbits * 3))

		rx := get_part(0, rbits, x)
		gx := get_part(1, rbits, x)
		bx := get_part(2, rbits, x)

		g.r = reflect_0_255(int32(g.r) + rx)
		g.g = reflect_0_255(int32(g.g) + gx)
		g.b = reflect_0_255(int32(g.b) + bx)
	} else if ip.IsMouseButtonJustReleased(v2.MouseButtonRight) {
		g.r = 128
		g.g = 128
		g.b = 128
	}
	return nil
}

func (g *Game) Draw(screen *v2.Image) {
	screen.Fill(color.RGBA{R: g.r, G: g.g, B: g.b, A: 255})
	util.DebugPrint(screen, fmt.Sprintf("Background color (%d %d %d)", g.r, g.g, g.b))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.Width, g.Height
}

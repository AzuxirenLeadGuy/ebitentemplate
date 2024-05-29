package coregame

import (
	"coregame"
	"github.com/hajimehoshi/ebiten/v2/mobile"
)

func init()  { mobile.SetGame(&coregame.Game{}) }
func Dummy() {}

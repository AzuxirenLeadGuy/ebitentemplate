package main

import (
	"coregame"
	v2 "github.com/hajimehoshi/ebiten/v2"
	"log"
)

func main() {
	if err := v2.RunGame(&coregame.Game{Width: 640, Height: 480}); err != nil {
		log.Fatal(err)
	}
}

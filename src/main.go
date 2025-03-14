package main

import (
	"fmt"
	"log"

	"image/color"

	"github.com/Giordano26/chip8/core"
	"github.com/Giordano26/chip8/core/graphics"
	"github.com/Giordano26/chip8/core/keyboard"
	"github.com/hajimehoshi/ebiten/v2"
)

type Chip8 = core.Chip8

type Game struct {
	chip8 *Chip8
}

func (g *Game) Update() error {
	keyboard.CheckKeys(&g.chip8.Chip8Keyboard)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for x := 0; x < graphics.ScreenWidth; x++ {
		for y := 0; y < graphics.ScreenHeight; y++ {
			if graphics.IsScreenSet(&g.chip8.Chip8Screen, x, y) {
				screen.Set(x, y, color.White)
			}
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return graphics.ScreenWidth, graphics.ScreenHeight
}

func main() {
	fmt.Println("CHIP-8 Emulator starting...")

	ebiten.SetWindowSize(graphics.ScreenWidth*graphics.WindowScale, graphics.ScreenHeight*graphics.WindowScale)
	ebiten.SetWindowTitle(graphics.WindowTitle)

	chip8 := &Chip8{}

	graphics.ScreenSet(&chip8.Chip8Screen, 10, 1)

	if err := ebiten.RunGame(&Game{chip8}); err != nil {
		log.Fatal(err)
	}
}

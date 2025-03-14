package main

import (
	"fmt"
	"log"

	"github.com/Giordano26/chip8/core"
	"github.com/Giordano26/chip8/core/graphics"
	"github.com/Giordano26/chip8/core/keyboard"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
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
	ebitenutil.DebugPrint(screen, "Hello, World!")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return graphics.ScreenWidth * graphics.WindowScale, graphics.ScreenHeight * graphics.WindowScale
}

func main() {
	fmt.Println("CHIP-8 Emulator starting...")

	ebiten.SetWindowSize(graphics.ScreenWidth*graphics.WindowScale, graphics.ScreenHeight*graphics.WindowScale)
	ebiten.SetWindowTitle(graphics.WindowTitle)

	chip8 := &Chip8{}

	core.StackPush(chip8, 0x200)
	fmt.Println(core.StackPop(chip8))

	if err := ebiten.RunGame(&Game{chip8}); err != nil {
		log.Fatal(err)
	}
}

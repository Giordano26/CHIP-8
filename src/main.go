package main

import (
	"fmt"
	"log"
	"os"

	"image/color"

	"github.com/Giordano26/chip8/core"
	"github.com/Giordano26/chip8/core/graphics"
	"github.com/Giordano26/chip8/core/keyboard"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/sqweek/dialog"
)

type Chip8 = core.Chip8

type Game struct {
	chip8 *Chip8
}

func (g *Game) Update() error {
	keyboard.CheckKeys(&g.chip8.Chip8Keyboard)

	core.CheckDelayTimer(g.chip8)
	core.CheckSoundTimer(g.chip8)

	core.CheckNextInstruction(g.chip8)

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

func openFile() []byte {
	filename, err := dialog.File().
		Title("Open Chip-8 ROM").
		Filter("Chip-8 ROM", "ch8").
		Filter("All Files", "*").
		Load()

	if err != nil {
		panic("[ERROR] Failed to open file")
	}

	content, err := os.ReadFile(filename)

	if err != nil {
		panic("[ERROR] Failed to read file")
	}

	return content

}

func main() {
	fmt.Println("CHIP-8 Emulator starting...")

	ebiten.SetWindowSize(graphics.ScreenWidth*graphics.WindowScale, graphics.ScreenHeight*graphics.WindowScale)
	ebiten.SetWindowTitle(graphics.WindowTitle)
	ebiten.SetTPS(240)

	chip8 := &Chip8{}
	core.Chip8Init(chip8)

	rom := openFile()
	core.LoadRom(chip8, rom)

	if err := ebiten.RunGame(&Game{chip8}); err != nil {
		log.Fatal(err)
	}
}

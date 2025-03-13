package main

import (
	"fmt"

	"github.com/Giordano26/chip8/core"
	"github.com/Giordano26/chip8/core/graphics"
	"github.com/Giordano26/chip8/core/memory"
)

type Chip8 = core.Chip8

func main() {
	chip8 := Chip8{}
	fmt.Println("CHIP-8 Emulator starting...")
	graphics.Run()
	memory.Chip8MemorySet(&chip8.Chip8Memory, 400, 'z')
	fmt.Println(memory.Chip8MemoryGet(&chip8.Chip8Memory, 400))
}

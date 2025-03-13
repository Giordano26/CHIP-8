package core

import (
	"github.com/Giordano26/chip8/core/memory"
	"github.com/Giordano26/chip8/core/registers"
)

type Chip8 struct {
	Chip8Memory    memory.Memory
	Chip8Registers registers.Registers
}

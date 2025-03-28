package core

import (
	"time"

	"github.com/Giordano26/chip8/core/graphics"
	"github.com/Giordano26/chip8/core/keyboard"
	"github.com/Giordano26/chip8/core/memory"
	"github.com/Giordano26/chip8/core/registers"
	"github.com/Giordano26/chip8/core/stack"
)

const (
	FontSetLoad = 0x00
)

type Chip8 struct {
	Chip8Memory    memory.Memory
	Chip8Registers registers.Registers
	Chip8Stack     stack.Stack
	Chip8Keyboard  keyboard.Keyboard
	Chip8Screen    graphics.Screen
}

func stackInBounds(chip8 *Chip8) {
	if chip8.Chip8Registers.SP >= stack.StackSize {
		panic("[WARNING] Stack pointer out of bounds")
	}
}

func StackPush(chip8 *Chip8, value uint16) {
	stackInBounds(chip8)

	chip8.Chip8Stack.Stack[chip8.Chip8Registers.SP] = value
	chip8.Chip8Registers.SP += 1
}

func StackPop(chip8 *Chip8) uint16 {

	chip8.Chip8Registers.SP -= 1
	return chip8.Chip8Stack.Stack[chip8.Chip8Registers.SP]
}

func CheckDelayTimer(chip8 *Chip8) {
	if chip8.Chip8Registers.DelayTimer > 0 {
		chip8.Chip8Registers.DelayTimer--
		time.Sleep(100 * time.Millisecond)
	}
}

func Chip8Init(chip *Chip8) {
	copy(chip.Chip8Memory.Memory[FontSetLoad:], graphics.FontSet[:])
	chip.Chip8Registers.PC = 0x200
	chip.Chip8Registers.SP = 0
	chip.Chip8Registers.I = 0
}

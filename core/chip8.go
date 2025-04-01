package core

import (
	"fmt"
	"time"

	"github.com/Giordano26/chip8/core/audio"
	"github.com/Giordano26/chip8/core/graphics"
	"github.com/Giordano26/chip8/core/keyboard"
	"github.com/Giordano26/chip8/core/memory"
	"github.com/Giordano26/chip8/core/registers"
	"github.com/Giordano26/chip8/core/stack"
)

const (
	FontSetLoad     = 0x00
	ProgramLoadAddr = 0x200
)

type Chip8 struct {
	Chip8Memory    memory.Memory
	Chip8Registers registers.Registers
	Chip8Stack     stack.Stack
	Chip8Keyboard  keyboard.Keyboard
	Chip8Screen    graphics.Screen
	Chip8Audio     audio.Audio
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

func CheckSoundTimer(chip8 *Chip8) {
	if chip8.Chip8Registers.SoundTimer > 0 {
		chip8.Chip8Registers.SoundTimer--
		if !chip8.Chip8Audio.IsPlaying() {
			chip8.Chip8Audio.PlayBeep(440, 100*time.Millisecond)
		}
	}
}

func Chip8Init(chip8 *Chip8) {
	copy(chip8.Chip8Memory.Memory[FontSetLoad:], graphics.FontSet[:])

	chip8.Chip8Registers.SP = 0
	chip8.Chip8Registers.I = 0

	chip8.Chip8Audio = *audio.NewSoundPlayer()
}

func LoadRom(chip8 *Chip8, rom []byte) {

	if len(rom)+ProgramLoadAddr > memory.MemorySize {
		panic("[WARNING] ROM size exceeds memory size")
	}
	copy(chip8.Chip8Memory.Memory[ProgramLoadAddr:], rom)
	chip8.Chip8Registers.PC = 0x200
}

func CheckNextInstruction(chip8 *Chip8) {
	opcode := memory.GetOpCode(&chip8.Chip8Memory, int(chip8.Chip8Registers.PC))

	Chip8Exec(chip8, opcode)

	chip8.Chip8Registers.PC += 2

	fmt.Println(opcode)
}

func Chip8Exec(chip8 *Chip8, opcode uint16) {

}

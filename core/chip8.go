package core

import (
	"crypto/rand"
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
		time.Sleep(1 * time.Millisecond)
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
	chip8.Chip8Registers.PC += 2
	Chip8Exec(chip8, opcode)

}

func Chip8Exec(chip8 *Chip8, opcode uint16) {

	switch opcode {
	//cls
	case 0x00E0:
		graphics.ScreenClear(&chip8.Chip8Screen)
	//ret
	case 0x00EE:
		chip8.Chip8Registers.PC = StackPop(chip8)

	default:
		executeBitwiseInstruction(chip8, opcode)
	}
}

func executeBitwiseInstruction(chip8 *Chip8, opcode uint16) {

	nnn := opcode & 0x0FFF //ignoring first 4 bits
	nibble := opcode & 0x000F
	x := (opcode & 0x0F00) >> 8
	y := (opcode & 0x00F0) >> 4
	kk := opcode & 0x00FF

	switch opcode & 0xF000 { //get first 4 bits to check type of instruction

	// JMP TO LOCATION nnn
	case 0x1000:
		chip8.Chip8Registers.PC = nnn
	//CALL SUBROUTINE AT nnn
	case 0x2000:
		{
			StackPush(chip8, chip8.Chip8Registers.PC)
			chip8.Chip8Registers.PC = nnn
		}
	//SE Vx, kk
	case 0x3000:
		{
			if chip8.Chip8Registers.V[x] == uint8(kk) {
				chip8.Chip8Registers.PC += 2
			}
		}
	//SNE Vx, kk
	case 0x4000:
		{
			if chip8.Chip8Registers.V[x] != uint8(kk) {
				chip8.Chip8Registers.PC += 2
			}
		}
	//SE Vx, Vy
	case 0x5000:
		{
			if chip8.Chip8Registers.V[x] == chip8.Chip8Registers.V[y] {
				chip8.Chip8Registers.PC += 2
			}
		}
	//LD Vx, kk
	case 0x6000:
		chip8.Chip8Registers.V[x] = uint8(kk)

	//ADD Vx, kk
	case 0x7000:
		chip8.Chip8Registers.V[x] += uint8(kk)

	case 0x8000:
		execExtendedEight(chip8, opcode, x, y)

	case 0x9000:
		execExtendedF(chip8, opcode, x)
	//LD I, nnn
	case 0xA000:
		chip8.Chip8Registers.I = nnn
	//JP V0, nnn
	case 0xB000:
		chip8.Chip8Registers.PC = nnn + uint16(chip8.Chip8Registers.V[0])
	//RND Vx, kk
	case 0xC000:
		rndByte := make([]uint8, 1)

		_, err := rand.Read(rndByte)
		if err != nil {
			panic("[WARNING] Error generating random byte")
		}

		chip8.Chip8Registers.V[x] = rndByte[0] & uint8(kk)

	//DRW Vx, Vy, nibble
	case 0xD000:
		chip8.Chip8Registers.V[0xF] = graphics.DrawSprite(&chip8.Chip8Screen, int(chip8.Chip8Registers.V[x]), int(chip8.Chip8Registers.V[y]), chip8.Chip8Memory.Memory[chip8.Chip8Registers.I:chip8.Chip8Registers.I+uint16(nibble)])

	case 0xE000:
		{
			switch opcode & 0x00FF {
			//SKP Vx
			case 0x009E:
				{
					if keyboard.IsKeyDown(&chip8.Chip8Keyboard, int(chip8.Chip8Registers.V[x])) {
						chip8.Chip8Registers.PC += 2
					}
				}
			//SKNP Vx
			case 0x00A1:
				{
					if !keyboard.IsKeyDown(&chip8.Chip8Keyboard, int(chip8.Chip8Registers.V[x])) {
						chip8.Chip8Registers.PC += 2
					}
				}
			}
		}
	//LD Vx, DT
	case 0xF000:
		execExtendedF(chip8, opcode, x)
	}

}

func execExtendedEight(chip8 *Chip8, opcode, x, y uint16) {

	switch opcode & 0x000F {
	//LD Vx, Vy
	case 0x00:
		chip8.Chip8Registers.V[x] = chip8.Chip8Registers.V[y]
	//OR Vx, Vy
	case 0x01:
		chip8.Chip8Registers.V[x] |= chip8.Chip8Registers.V[y]

	//AND Vx, Vy
	case 0x02:
		chip8.Chip8Registers.V[x] &= chip8.Chip8Registers.V[y]

	//XOR Vx, Vy
	case 0x03:
		chip8.Chip8Registers.V[x] ^= chip8.Chip8Registers.V[y]

	//ADD Vx, Vy
	case 0x04:
		{
			sum := uint16(chip8.Chip8Registers.V[x]) + uint16(chip8.Chip8Registers.V[y])
			chip8.Chip8Registers.V[x] = uint8(sum)

			chip8.Chip8Registers.V[0xF] = 0

			if sum > 0xFF {
				chip8.Chip8Registers.V[0xF] = 1
			}

		}

	//SUB Vx, Vy
	case 0x05:
		{
			chip8.Chip8Registers.V[0xF] = 0

			if chip8.Chip8Registers.V[x] > chip8.Chip8Registers.V[y] {
				chip8.Chip8Registers.V[0xF] = 1
			}

			chip8.Chip8Registers.V[x] -= chip8.Chip8Registers.V[y]
		}

	//SHR Vx {, Vy}
	case 0x06:
		{
			chip8.Chip8Registers.V[0xF] = chip8.Chip8Registers.V[x] & 0x01

			chip8.Chip8Registers.V[x] >>= 1 //divide by 2
		}

	//SUBN Vx, Vy
	case 0x07:
		{
			chip8.Chip8Registers.V[0xF] = 0

			if chip8.Chip8Registers.V[y] > chip8.Chip8Registers.V[x] {
				chip8.Chip8Registers.V[0xF] = 1
			}

			chip8.Chip8Registers.V[x] = chip8.Chip8Registers.V[y] - chip8.Chip8Registers.V[x]
		}

	//SHL Vx {, Vy}
	case 0x0E:
		{
			chip8.Chip8Registers.V[0xF] = chip8.Chip8Registers.V[x] & 0x80 >> 7

			chip8.Chip8Registers.V[x] <<= 1 //multiplied by 2
		}
	}
}

func execExtendedF(chip8 *Chip8, opcode, x uint16) {

	switch opcode & 0x00ff {
	//LD Vx, DT
	case 0x07:
		chip8.Chip8Registers.V[x] = chip8.Chip8Registers.DelayTimer

	//LD Vx, K
	case 0x0A:
		{
			pressedKey := keyboard.WaitForKeyPress()
			chip8.Chip8Registers.V[x] = pressedKey
		}

	//LD DT, Vx
	case 0x15:
		chip8.Chip8Registers.DelayTimer = chip8.Chip8Registers.V[x]

	//LD ST, Vx
	case 0x18:
		chip8.Chip8Registers.SoundTimer = chip8.Chip8Registers.V[x]

	//ADD I, Vx
	case 0x1E:
		chip8.Chip8Registers.I += uint16(chip8.Chip8Registers.V[x])

	//LD F, Vx
	case 0x29:
		chip8.Chip8Registers.I = uint16(chip8.Chip8Registers.V[x]) * 5

	//LD B, Vx
	case 0x33:
		{
			hundreds := chip8.Chip8Registers.V[x] / 100
			tens := (chip8.Chip8Registers.V[x] / 10) % 10
			units := chip8.Chip8Registers.V[x] % 10

			memory.Chip8MemorySet(&chip8.Chip8Memory, int(chip8.Chip8Registers.I), hundreds)
			memory.Chip8MemorySet(&chip8.Chip8Memory, int(chip8.Chip8Registers.I)+1, tens)
			memory.Chip8MemorySet(&chip8.Chip8Memory, int(chip8.Chip8Registers.I)+2, units)
		}

	//LD[I], Vx
	case 0x55:
		{
			for i := uint16(0); i <= x; i++ {
				memory.Chip8MemorySet(&chip8.Chip8Memory, int(chip8.Chip8Registers.I+i), chip8.Chip8Registers.V[i])
			}
		}

	//LD Vx, [I]
	case 0x65:
		{
			for i := uint16(0); i <= x; i++ {
				chip8.Chip8Registers.V[i] = memory.Chip8MemoryGet(&chip8.Chip8Memory, int(chip8.Chip8Registers.I+i))
			}
		}
	}
}

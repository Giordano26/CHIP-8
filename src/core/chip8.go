package core

import (
	"os"

	"github.com/hajimehoshi/ebiten"
)

const (
	MEMORY_SIZE = 4096
)

type Chip8 struct {
	display [64][32]byte // display size

	memory [4096]byte // memory size 4k
	vx     [16]byte   // cpu registers V0-VF
	keys   [16]bool   // input key
	stack  [16]uint16 // program counter stack

	oc uint16 // current opcode
	pc uint16 // program counter
	sp uint16 // stack pointer
	ir uint16 // index register

	delayTimer byte
	soundTImer byte

	drawFlag bool
}

type Graphics struct {
	chip8 *Chip8
}

func (c *Chip8) HandleInput() {
	keyMap := map[ebiten.Key]int{
		ebiten.Key1: 0x1, ebiten.Key2: 0x2, ebiten.Key3: 0x3, ebiten.Key4: 0xC,
		ebiten.KeyQ: 0x4, ebiten.KeyW: 0x5, ebiten.KeyE: 0x6, ebiten.KeyR: 0xD,
		ebiten.KeyA: 0x7, ebiten.KeyS: 0x8, ebiten.KeyD: 0x9, ebiten.KeyF: 0xE,
		ebiten.KeyZ: 0xA, ebiten.KeyX: 0x0, ebiten.KeyC: 0xB, ebiten.KeyV: 0xF,
	}

	for key, value := range keyMap {
		if ebiten.IsKeyPressed(key) {
			c.keys[value] = true
		} else {
			c.keys[value] = false
		}
	}
}

func (c *Chip8) Initialize() {
	c.pc = 0x200 //start the program counter at the right address

	//reset current opcode and registers
	c.oc = 0
	c.ir = 0
	c.sp = 0

	c.delayTimer = 0
	c.soundTImer = 0

	c.drawFlag = true

	for i := range 80 {
		c.memory[i] = fontSet[i]
	}

	for i := range 16 {
		c.vx[i] = 0
	}
}

func (c *Chip8) LoadROM(path string) error {
	rom, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	// Load ROM into memory starting at 0x200
	for i := range rom {
		c.memory[0x200+i] = rom[i]
	}

	return nil
}

func (c *Chip8) Fetch() uint16 {
	//The high byte is cast to a uint16 (16-bit unsigned integer) and shifted left by 8 bits.
	//Same for the low part of instruction located on the next address and ORs it with the high part

	opcode := uint16(c.memory[c.pc])<<8 | uint16(c.memory[c.pc+1])
	c.pc += 2
	return opcode
}

func (c *Chip8) EmulateCycle() {
	opcode := c.Fetch()

	//0xF000 is a bitmask used to extract the first nibble (4 bits) of the opcode.
	//CHIP-8 opcodes are designed so that the first nibble often determines the type of instruction.
	//0x1YYY jmp to YYY
	//0x6YYY set VY to YY
	//0xAYYY set I to YYY

	switch opcode & 0xF000 {

	}
}

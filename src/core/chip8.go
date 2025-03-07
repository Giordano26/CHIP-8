package core

import (
	"os"
)

var fontSet = []uint8{
	0xF0, 0x90, 0x90, 0x90, 0xF0, //0
	0x20, 0x60, 0x20, 0x20, 0x70, //1
	0xF0, 0x10, 0xF0, 0x80, 0xF0, //2
	0xF0, 0x10, 0xF0, 0x10, 0xF0, //3
	0x90, 0x90, 0xF0, 0x10, 0x10, //4
	0xF0, 0x80, 0xF0, 0x10, 0xF0, //5
	0xF0, 0x80, 0xF0, 0x90, 0xF0, //6
	0xF0, 0x10, 0x20, 0x40, 0x40, //7
	0xF0, 0x90, 0xF0, 0x90, 0xF0, //8
	0xF0, 0x90, 0xF0, 0x10, 0xF0, //9
	0xF0, 0x90, 0xF0, 0x90, 0x90, //A
	0xE0, 0x90, 0xE0, 0x90, 0xE0, //B
	0xF0, 0x80, 0x80, 0x80, 0xF0, //C
	0xE0, 0x90, 0x90, 0x90, 0xE0, //D
	0xF0, 0x80, 0xF0, 0x80, 0xF0, //E
	0xF0, 0x80, 0xF0, 0x80, 0x80, //F
}

type Chip8 struct {
	display [32][64]byte // display size

	memory [4096]byte // memory size 4k
	vx     [16]byte   // cpu registers V0-VF
	key    [16]byte   // input key
	stack  [16]uint16 // program counter stack

	oc uint16 // current opcode
	pc uint16 // program counter
	sp uint16 // stack pointer
	ir uint16 // index register

	delayTimer byte
	soundTImer byte

	drawFlag bool
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
	for i := 0; i < len(rom); i++ {
		c.memory[0x200+i] = rom[i]
	}

	return nil
}

package memory

const (
	MemorySize = 4096
)

type Memory struct {
	Memory [MemorySize]uint8 // memory size 4k
}

func isMemoryIndexValid(index int) {
	if index < 0 || index >= MemorySize {
		panic("[WARNING] Memory index out of bounds")
	}
}

func Chip8MemorySet(m *Memory, index int, value uint8) {
	isMemoryIndexValid(index)
	m.Memory[index] = value
}

func Chip8MemoryGet(m *Memory, index int) uint8 {
	isMemoryIndexValid(index)
	return m.Memory[index]
}

func GetOpCode(m *Memory, index int) uint16 {
	isMemoryIndexValid(index)

	byte1 := Chip8MemoryGet(m, index)
	byte2 := Chip8MemoryGet(m, index+1)

	return uint16(byte1)<<8 | uint16(byte2)
}

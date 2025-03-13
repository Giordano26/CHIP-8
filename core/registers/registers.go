package registers

const (
	TotalDataRegisters = 16
)

type Registers struct {
	V [TotalDataRegisters]uint8
	I uint16

	delayTimer uint8
	soundTimer uint8

	PC uint16
	SP uint8
}

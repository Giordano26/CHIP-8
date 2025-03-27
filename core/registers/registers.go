package registers

const (
	TotalDataRegisters = 16
)

type Registers struct {
	V [TotalDataRegisters]uint8
	I uint16

	DelayTimer uint8
	SoundTimer uint8

	PC uint16
	SP uint8
}

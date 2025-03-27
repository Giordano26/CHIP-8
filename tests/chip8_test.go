package tests

import (
	"testing"

	"github.com/Giordano26/chip8/core"
)

func TestStackPush(t *testing.T) {
	type testCase struct {
		chip8 *core.Chip8
		value uint16

		expectedValue uint16
		expectedSP    uint8
		shouldPanic   bool
	}

	tests := []testCase{
		{chip8: &core.Chip8{}, value: 0x1234, expectedValue: 0x1234, expectedSP: 1, shouldPanic: false},
		{chip8: &core.Chip8{}, value: 0x5678, expectedValue: 0x5678, expectedSP: 1, shouldPanic: false},
	}

	for _, test := range tests {
		if test.shouldPanic {
			defer func() {
				if r := recover(); r == nil {
					t.Errorf("Expected panic for value %d, but did not panic", test.value)
				}
			}()
		}

		core.StackPush(test.chip8, test.value)

		if !test.shouldPanic {
			actualValue := test.chip8.Chip8Stack.Stack[test.chip8.Chip8Registers.SP-1]
			if actualValue != test.expectedValue {
				t.Errorf("For value %d, expected %d, but got %d", test.value, test.expectedValue, actualValue)
			}

			actualSP := test.chip8.Chip8Registers.SP
			if actualSP != test.expectedSP {
				t.Errorf("For value %d, expected SP %d, but got %d", test.value, test.expectedSP, actualSP)
			}
		}
	}
}
func TestStackPop(t *testing.T) {
	type testCase struct {
		chip8 *core.Chip8
		value uint16

		expectedValue uint16
		expectedSP    uint8
		shouldPanic   bool
	}

	tests := []testCase{
		{chip8: &core.Chip8{}, value: 0x1234, expectedValue: 0x1234, expectedSP: 0, shouldPanic: false},
		{chip8: &core.Chip8{}, value: 0x5678, expectedValue: 0x5678, expectedSP: 0, shouldPanic: false},
	}

	for _, test := range tests {
		if test.shouldPanic {
			defer func() {
				if r := recover(); r == nil {
					t.Errorf("Expected panic for value %d, but did not panic", test.value)
				}
			}()
		}

		core.StackPush(test.chip8, test.value)
		actualValue := core.StackPop(test.chip8)

		if !test.shouldPanic {
			if actualValue != test.expectedValue {
				t.Errorf("For value %d, expected %d, but got %d", test.value, test.expectedValue, actualValue)
			}

			actualSP := test.chip8.Chip8Registers.SP
			if actualSP != test.expectedSP {
				t.Errorf("For value %d, expected SP %d, but got %d", test.value, test.expectedSP, actualSP)
			}
		}
	}
}
func TestChip8Init(t *testing.T) {
	type testCase struct {
		chip8 *core.Chip8

		expectedPC uint16
		expectedI  uint16
		expectedSP uint8
	}

	tests := []testCase{
		{chip8: &core.Chip8{}, expectedPC: 0x200, expectedI: 0, expectedSP: 0},
	}

	for _, test := range tests {
		core.Chip8Init(test.chip8)

		if test.chip8.Chip8Registers.PC != test.expectedPC {
			t.Errorf("Expected PC %d, but got %d", test.expectedPC, test.chip8.Chip8Registers.PC)
		}

		if test.chip8.Chip8Registers.I != test.expectedI {
			t.Errorf("Expected I %d, but got %d", test.expectedI, test.chip8.Chip8Registers.I)
		}

		if test.chip8.Chip8Registers.SP != test.expectedSP {
			t.Errorf("Expected SP %d, but got %d", test.expectedSP, test.chip8.Chip8Registers.SP)
		}
	}
}

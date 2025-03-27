package tests

import (
	"testing"

	"github.com/Giordano26/chip8/core/memory"
)

func TestChip8MemorySet(t *testing.T) {
	type testCase struct {
		m           *memory.Memory
		index       int
		value       uint8
		expected    uint8
		shouldPanic bool
	}

	tests := []testCase{
		{m: &memory.Memory{}, index: 0, value: 1, expected: 1, shouldPanic: false},
		{m: &memory.Memory{}, index: 4095, value: 255, expected: 255, shouldPanic: false},
		{m: &memory.Memory{}, index: -1, value: 10, expected: 0, shouldPanic: true},
		{m: &memory.Memory{}, index: 4096, value: 10, expected: 0, shouldPanic: true},
	}

	for _, test := range tests {
		if test.shouldPanic {
			defer func() {
				if r := recover(); r == nil {
					t.Errorf("Expected panic for index %d, but did not panic", test.index)
				}
			}()
		}

		memory.Chip8MemorySet(test.m, test.index, test.value)

		if !test.shouldPanic {
			actual := test.m.Memory[test.index]
			if actual != test.expected {
				t.Errorf("For index %d, expected %d, but got %d", test.index, test.expected, actual)
			}
		}
	}
}

func TestChip8MemoryGet(t *testing.T) {
	type testCase struct {
		m           *memory.Memory
		index       int
		expected    uint8
		shouldPanic bool
	}

	tests := []testCase{
		{m: &memory.Memory{}, index: 0, expected: 255, shouldPanic: false},
		{m: &memory.Memory{}, index: 4095, expected: 123, shouldPanic: false},
		{m: &memory.Memory{}, index: -1, expected: 0, shouldPanic: true},
		{m: &memory.Memory{}, index: 4096, expected: 0, shouldPanic: true},
	}

	for _, test := range tests {
		if test.shouldPanic {
			defer func() {
				if r := recover(); r == nil {
					t.Errorf("Expected panic for index %d, but did not panic", test.index)
				}
			}()
		}

		test.m.Memory[test.index] = test.expected
		actual := memory.Chip8MemoryGet(test.m, test.index)

		if !test.shouldPanic {
			if actual != test.expected {
				t.Errorf("For index %d, expected %d, but got %d", test.index, test.expected, actual)
			}
		}
	}
}

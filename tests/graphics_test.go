package tests

import (
	"testing"

	"github.com/Giordano26/chip8/core/graphics"
)

func TestScreenSet(t *testing.T) {
	type testCase struct {
		screen      *graphics.Screen
		x           int
		y           int
		expected    bool
		shouldPanic bool
	}

	tests := []testCase{
		{screen: &graphics.Screen{}, x: 0, y: 0, expected: true, shouldPanic: false},
		{screen: &graphics.Screen{}, x: 63, y: 31, expected: true, shouldPanic: false},
		{screen: &graphics.Screen{}, x: -1, y: 0, expected: false, shouldPanic: true},
		{screen: &graphics.Screen{}, x: 0, y: -1, expected: false, shouldPanic: true},
		{screen: &graphics.Screen{}, x: 64, y: 0, expected: false, shouldPanic: true},
		{screen: &graphics.Screen{}, x: 0, y: 32, expected: false, shouldPanic: true},
	}

	for _, test := range tests {
		if test.shouldPanic {
			defer func() {
				if r := recover(); r == nil {
					t.Errorf("Expected panic for x %d y %d, but did not panic", test.x, test.y)
				}
			}()
		}

		graphics.ScreenSet(test.screen, test.x, test.y)

		if !test.shouldPanic {
			actual := test.screen.Pixels[test.y][test.x]
			if actual != test.expected {
				t.Errorf("For x %d and %d, expected %t, but got %t", test.x, test.y, test.expected, actual)
			}
		}
	}
}

func TestIsScreenSet(t *testing.T) {
	type testCase struct {
		screen      *graphics.Screen
		x           int
		y           int
		expected    bool
		shouldPanic bool
	}

	tests := []testCase{
		{screen: &graphics.Screen{}, x: 0, y: 0, expected: true, shouldPanic: false},
		{screen: &graphics.Screen{}, x: 0, y: 0, expected: false, shouldPanic: false},
		{screen: &graphics.Screen{}, x: 63, y: 31, expected: true, shouldPanic: false},
		{screen: &graphics.Screen{}, x: -1, y: 0, expected: false, shouldPanic: true},
		{screen: &graphics.Screen{}, x: 0, y: -1, expected: false, shouldPanic: true},
		{screen: &graphics.Screen{}, x: 64, y: 0, expected: false, shouldPanic: true},
		{screen: &graphics.Screen{}, x: 0, y: 32, expected: false, shouldPanic: true},
	}

	for _, test := range tests {
		if test.shouldPanic {
			defer func() {
				if r := recover(); r == nil {
					t.Errorf("Expected panic for x %d y %d, but did not panic", test.x, test.y)
				}
			}()
		}
		for _, test := range tests {
			if test.shouldPanic {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("Expected panic for x %d y %d, but did not panic", test.x, test.y)
					}
				}()
			}

			test.screen.Pixels[test.y][test.x] = test.expected

			actual := graphics.IsScreenSet(test.screen, test.x, test.y)
			if !test.shouldPanic {
				if actual != test.expected {
					t.Errorf("For x %d and y %d, expected %t, but got %t", test.x, test.y, test.expected, actual)
				}
			}
		}
	}
}

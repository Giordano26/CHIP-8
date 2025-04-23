package keyboard

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	TotalKeys = 16
)

type Keyboard struct {
	keys [TotalKeys]bool
}

var KeyMap = []ebiten.Key{
	ebiten.KeyDigit0, ebiten.KeyDigit1, // 0 1
	ebiten.KeyDigit2, ebiten.KeyDigit3, // 2 3
	ebiten.KeyDigit4, ebiten.KeyDigit5, // 4 5
	ebiten.KeyDigit6, ebiten.KeyDigit7, // 6 7
	ebiten.KeyDigit8, ebiten.KeyDigit9, // 8 9
	ebiten.KeyZ, ebiten.KeyX, ebiten.KeyC, // A B C
	ebiten.KeyV, ebiten.KeyB, ebiten.KeyN, // D E F
}

func CheckKeys(keyboard *Keyboard) {
	for vkey, key := range KeyMap {
		if inpututil.IsKeyJustPressed(key) {
			keyDown(keyboard, vkey)
		}

		if inpututil.IsKeyJustReleased(key) {
			keyUp(keyboard, vkey)
		}
	}
}

func keyDown(keyboard *Keyboard, key int) {
	keyboard.keys[key] = true
}

func keyUp(keyboard *Keyboard, key int) {
	keyboard.keys[key] = false
}

func IsKeyDown(keyboard *Keyboard, key int) bool {
	return keyboard.keys[key]
}

func WaitForKeyPress() uint8 {
	// I cant manage a way to make a block and still do not block the ebiten loop
	// if I block the code here and wait for input the ebiten loop will be blocked
	// hence making the if check unreachable since no "key check" is being made
	// generaly this opcode is used for reseting after a press so i will just assume
	// this general purpose and use 1 as deafault case
	time.Sleep(1000 * time.Millisecond)
	return 1
}

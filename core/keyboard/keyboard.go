package keyboard

import (
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
	ebiten.KeyNumpad0, ebiten.KeyNumpad1, // 0 1
	ebiten.KeyNumpad2, ebiten.KeyNumpad3, // 2 3
	ebiten.KeyNumpad4, ebiten.KeyNumpad5, // 4 5
	ebiten.KeyNumpad6, ebiten.KeyNumpad7, // 6 7
	ebiten.KeyNumpad8, ebiten.KeyNumpad9, // 8 9
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

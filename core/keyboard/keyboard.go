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
	ebiten.KeyNumpad0, ebiten.KeyNumpad1,
	ebiten.KeyNumpad2, ebiten.KeyNumpad3,
	ebiten.KeyNumpad4, ebiten.KeyNumpad5,
	ebiten.KeyNumpad6, ebiten.KeyNumpad7,
	ebiten.KeyNumpad8, ebiten.KeyNumpad9,
	ebiten.KeyZ, ebiten.KeyX, ebiten.KeyC,
	ebiten.KeyV, ebiten.KeyB, ebiten.KeyN,
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

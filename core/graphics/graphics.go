package graphics

const (
	ScreenWidth  = 64
	ScreenHeight = 32
	WindowTitle  = "LD-8 (Chip-8 Emulator)"
	WindowScale  = 10

	BackgroundColor = 0
	PixelColor      = 0xFFFFFFFF
)

type Screen struct {
	Pixels [ScreenHeight][ScreenWidth]bool
}

var FontSet = []uint8{
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

func isPixelIndexValid(x, y int) {
	if x < 0 || x >= ScreenWidth || y < 0 || y >= ScreenHeight {
		panic("[WARNING] Pixel index out of bounds")
	}
}

func ScreenSet(screen *Screen, x, y int) {
	isPixelIndexValid(x, y)
	screen.Pixels[y][x] = true
}

func IsScreenSet(screen *Screen, x, y int) bool {
	isPixelIndexValid(x, y)
	return screen.Pixels[y][x]
}

func DrawSprite(screen *Screen, x, y int, sprite []uint8) uint8 {
	var collision uint8 = 0

	for ly, spriteByte := range sprite {
		for lx := range 8 {
			if (spriteByte & (0x80 >> lx)) == 0 {
				continue
			}
			py := (ly + y) % ScreenHeight
			px := (lx + x) % ScreenWidth

			if screen.Pixels[py][px] {
				collision = 1
			}

			screen.Pixels[py][px] = !screen.Pixels[py][px]
		}
	}
	return collision
}

func ScreenClear(screen *Screen) {
	pixels := &screen.Pixels

	for y := range pixels {
		for x := range pixels[y] {
			pixels[y][x] = false
		}
	}

}

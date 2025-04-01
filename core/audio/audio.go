package audio

import (
	"bytes"
	"errors"
	"math"
	"time"

	"github.com/hajimehoshi/ebiten/v2/audio"
)

const (
	sampleRate     = 44100
	bytesPerSample = 4 // 2 channels * 2 bytes (int16)
)

type Audio struct {
	audioContext *audio.Context
	player       *audio.Player
}

func NewSoundPlayer() *Audio {
	return &Audio{
		audioContext: audio.NewContext(sampleRate),
	}
}

func (s *Audio) GenerateBeep(frequency float64, duration time.Duration) *bytes.Reader {
	bufferSize := int(float64(sampleRate) * duration.Seconds())
	data := make([]byte, bufferSize*bytesPerSample)

	for i := range bufferSize {
		t := float64(i) / sampleRate
		val := math.Sin(2 * math.Pi * frequency * t)
		v := int16(val * math.MaxInt16)

		data[i*4] = byte(v)
		data[i*4+1] = byte(v >> 8)
		data[i*4+2] = byte(v)
		data[i*4+3] = byte(v >> 8)
	}

	return bytes.NewReader(data)
}

func (s *Audio) PlaySound(data *bytes.Reader) error {
	if s.audioContext == nil {
		return errors.New("audio context not initialized")
	}

	if s.player != nil && s.player.IsPlaying() {
		s.player.Pause()
	}

	p, err := s.audioContext.NewPlayer(data)
	if err != nil {
		return err
	}

	s.player = p
	s.player.Play()
	return nil
}

func (s *Audio) PlayBeep(frequency float64, duration time.Duration) error {
	data := s.GenerateBeep(frequency, duration)
	return s.PlaySound(data)
}

func (s *Audio) Stop() {
	if s.player != nil {
		s.player.Pause()
	}
}

func (s *Audio) IsPlaying() bool {
	return s.player != nil && s.player.IsPlaying()
}

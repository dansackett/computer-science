package adapter

import (
	"errors"
	"fmt"
	"strings"
)

type MediaPlayer interface {
	Play(audioType, filename string) (string, error)
}

type AdvancedMediaPlayer interface {
	PlayVLC(filename string) string
	PlayMP3(filename string) string
}

type VLCPlayer struct{}

func (vlc *VLCPlayer) PlayVLC(filename string) string {
	return fmt.Sprintf("Playing VLC file '%s'", filename)
}

func (vlc *VLCPlayer) PlayMP3(filename string) string {
	return ""
}

type MP3Player struct{}

func (mp3 *MP3Player) PlayVLC(filename string) string {
	return ""
}

func (mp3 *MP3Player) PlayMP3(filename string) string {
	return fmt.Sprintf("Playing MP3 file '%s'", filename)
}

type MediaAdapter struct {
	player AdvancedMediaPlayer
}

func NewMediaAdapter(audioType string) (*MediaAdapter, error) {
	var player AdvancedMediaPlayer
	if strings.ToLower(audioType) == "vlc" {
		player = &VLCPlayer{}
	} else if strings.ToLower(audioType) == "mp3" {
		player = &MP3Player{}
	} else {
		return nil, errors.New("Unsupported media type")
	}
	return &MediaAdapter{player: player}, nil
}

func (a *MediaAdapter) Play(audioType, filename string) (string, error) {
	if strings.ToLower(audioType) == "vlc" {
		return a.player.PlayVLC(filename), nil
	} else if strings.ToLower(audioType) == "mp3" {
		return a.player.PlayMP3(filename), nil
	} else {
		return "", errors.New("Unsupported media type")
	}
}

type AudioPlayer struct{}

func NewAudioPlayer() *AudioPlayer {
	return &AudioPlayer{}
}

func (a *AudioPlayer) Play(audioType, filename string) (string, error) {
	adapter, err := NewMediaAdapter(audioType)
	if err != nil {
		return "", err
	}
	return adapter.Play(audioType, filename)
}

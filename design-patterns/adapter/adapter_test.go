package adapter

import "testing"

func TestAdapter(t *testing.T) {
	player := NewAudioPlayer()

	_, mp3err := player.Play("mp3", "some_file.mp3")
	_, vlcerr := player.Play("vlc", "some_file.vlc")
	_, mp4err := player.Play("mp4", "some_file.mp4")

	if mp3err != nil {
		t.Errorf("MP3 tracks should be playable")
	}

	if vlcerr != nil {
		t.Errorf("VLC tracks should be playable")
	}

	if mp4err == nil {
		t.Errorf("MP4 tracks should NOT be playable")
	}
}

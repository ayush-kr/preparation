package main

import "log"

//MediaPlayer interface for playing a file
type MediaPlayer interface {
	play(string, string)
}

//AdvancedMediaPlayer interface for playing advance files
type AdvancedMediaPlayer interface {
	playVlc(string)
	playMp4(string)
}

type vlcPlayer struct{}

func (v *vlcPlayer) playVlc(filename string) {
	log.Println("Playing vlc file", filename)
}

func (v *vlcPlayer) playMp4(filename string) {}

type mp4Player struct{}

func (v *mp4Player) playVlc(filename string) {}

func (v *mp4Player) playMp4(filename string) {
	log.Println("Playing mp4 file", filename)
}

//MediaAdapter struct playing different kinds of media
type MediaAdapter struct {
	player AdvancedMediaPlayer
}

func (ma *MediaAdapter) play(audiotype, filename string) {
	if audiotype == "vlc" {
		ma.player = &vlcPlayer{}
		ma.player.playVlc(filename)
	} else if audiotype == "mp4" {
		ma.player = &mp4Player{}
		ma.player.playMp4(filename)
	}
}

type audioPlayer struct {
	adapter MediaAdapter
}

func (ap *audioPlayer) play(audiotype, filename string) {
	if audiotype == "mp3" {
		log.Println("Playing mp3", filename)
		return
	}
	if audiotype == "mp4" || audiotype == "vlc" {
		ap.adapter = MediaAdapter{}
		ap.adapter.play(audiotype, filename)
		return
	}
	log.Println(audiotype, "is unsupported")
}

func main() {
	ap := &audioPlayer{}
	ap.play("mp3", "ayush.mp3")
	ap.play("mp4", "mms.mp4")
	ap.play("vlc", "xxx.vlc")
	ap.play("avi", "daiyya.avi")
}

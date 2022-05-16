package main

import (
	"Popcorn/download"
	"Popcorn/webCtrl"
	vlc "github.com/adrg/libvlc-go/v3"
	"log"
	"os"
	"time"
)

func main() {
	// Initialize libVLC. Additional command line arguments can be passed in
	// to libVLC by specifying them in the Init function.

	messageBuffer := webCtrl.MessageBuffer

	go webCtrl.ServeHttp()

	if err := vlc.Init("--fullscreen","--no-autoscale", "-L"); err != nil {
		log.Fatal(err)
	}
	defer vlc.Release()

	// Create a new list player.
	player, err := vlc.NewListPlayer()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		player.Stop()
		player.Release()
	}()

	// Create a new media list.
	list, err := vlc.NewMediaList()
	if err != nil {
		log.Fatal(err)
	}
	defer list.Release()

	initLoop := true

	for initLoop {
	if _, err := os.Stat("video/video.mp4"); err == nil {
		initLoop = false
	} else {
		time.Sleep(time.Millisecond * 100)
	}
}
	err = list.AddMediaFromPath("video/video.mp4")
	if err != nil {
		log.Fatal(err)
	}


	// Set player media list.
	if err = player.SetMediaList(list); err != nil {
		log.Fatal(err)
	}

	player.Play()

for {
	select {
	case message := <-messageBuffer:
		switch message.Message {
		case "index":
			println(message.Data)
		case "add":

			if err != nil {
				log.Fatal(err)
			}
			list.AddMediaFromPath("placeholder.png")
			player.PlayNext()
			download.GetVideo(message.Data)
			time.Sleep(time.Millisecond * 100)
			list.AddMediaFromPath("video/video.mp4")
			list.RemoveMediaAtIndex(0)
			player.PlayNext()


		}
	}
}
}
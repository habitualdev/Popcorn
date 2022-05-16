package download

import (
	"github.com/kkdai/youtube/v2"
	"io"
	"net/url"
	"os"
)


func GetVideo(id string) {
	videoID, err := url.QueryUnescape(id)
	client := youtube.Client{}
	println("Downloading video...")
	video, err := client.GetVideo(videoID)
	if err != nil {
		panic(err)
	}

	formats := video.Formats.WithAudioChannels() // only get videos with audio
	stream, _, err := client.GetStream(video, &formats[0])
	if err != nil {
		panic(err)
	}

	file, err := os.Create("video/video.mp4")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = io.Copy(file, stream)
	if err != nil {
		panic(err)
	}
	println("Download complete")
}

package ffmpeg

import (
	"log"

	ffmpeg "github.com/u2takey/ffmpeg-go"
)

func ConvertToMp3(inputFile string, outputFile string, bitrate string) {
	err := ffmpeg.Input(inputFile).
		Output(outputFile, ffmpeg.KwArgs{"acodec": "libmp3lame", "b:a": bitrate}).
		OverWriteOutput().ErrorToStdOut().Run()

	if err != nil {
		log.Fatalf("error: %s", err)
	}

	// Conversion completed successfully
}

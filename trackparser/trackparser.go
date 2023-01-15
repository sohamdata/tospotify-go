package trackparser

import (
	"errors"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
)

type Track struct {
	Artist string
	Name   string
}

var allowedExtensions = map[string]bool{
	// ".ext" : true,
	".mp3": true,
	".flac": true,
	".wav": true,
}

func GetTracks(path string, tracks *[]Track) {
	files, err := ioutil.ReadDir(path)

	if err != nil {
		panic(err)
	}

	for _, file := range files {
		if file.IsDir() {
			GetTracks(path+"/"+file.Name(), tracks)
			continue
		}

		fileName := file.Name()

		if isMusicTrack(fileName) {
			track, err := parseMusicTrack(fileName)

			if err != nil {
				log.Print(err, fileName)
				continue
			}

			*tracks = append(*tracks, track)
		}

	}
}

func isMusicTrack(fileName string) bool {
	ext := strings.ToLower(filepath.Ext(fileName))

	// return strings.HasSuffix(ext, ".mp3") ||
	// 	strings.HasSuffix(ext, ".flac") ||
	// 	strings.HasSuffix(ext, ".wav")
	_, ok := allowedExtensions[ext]
	return ok

}

// Song must be "[artist] - [track]" e.g. "Mac Miller - The Spins"
func parseMusicTrack(fileName string) (Track, error) {
	out := strings.SplitN(fileName, " - ", 2)
	if len(out) != 2 {
		return Track{}, errors.New("cannot parse track, please use [artist] - [track] format")
	}

	return Track{
		Artist: out[0],
		Name:   strings.TrimSuffix(out[1], filepath.Ext(out[1])),
	}, nil
}
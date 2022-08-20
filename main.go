package main

import (
	"github.com/zmb3/spotify"
)

type App struct {
	Auth           spotify.Authenticator
	Client         spotify.Client
	MusicLibPath   string
	Playlist       *spotify.FullPlaylist
	RemoteTrackIds []spotify.ID
}

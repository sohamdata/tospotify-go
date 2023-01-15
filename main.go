package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
	"tospotify-go/trackparser"

	_ "github.com/joho/godotenv/autoload"
	"github.com/pkg/browser"
	"github.com/zmb3/spotify"
)

type App struct {
	Auth           spotify.Authenticator
	Client         spotify.Client
	MusicLibPath   string
	Playlist       *spotify.FullPlaylist
	RemoteTrackIds []spotify.ID
	Tracks         []trackparser.Track
}

var redirectUrl = "http://localhost:3000/callback"
var state = ""

func main() {
	path, _ := filepath.Abs(os.Args[1])
	app := App{
		Auth:         spotify.NewAuthenticator(redirectUrl, spotify.ScopeUserLibraryModify, spotify.ScopeUserLibraryRead, spotify.ScopePlaylistModifyPrivate),
		MusicLibPath: path,
	}
	url := app.Auth.AuthURL(state)
	browser.OpenURL(url)

	mux := http.NewServeMux()
	mux.HandleFunc("/callback", app.CallbackHandler)
	// err := http.ListenAndServe("localhost:3000", mux)
	// if err != nil {
	// 	panic(err)
	// }
	if err := http.ListenAndServe("localhost:3000", mux); err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
}

func (app *App) CallbackHandler(w http.ResponseWriter, r *http.Request) {
	token, err := app.Auth.Token(state, r)
	if err != nil {
		log.Print(err)
		http.Error(w, "Couldn't get token", http.StatusNotFound)
		return
	}

	app.Client = app.Auth.NewClient(token)
	trackparser.GetTracks(app.MusicLibPath, &app.Tracks)

	app.CreateSpotifyPlaylist()
	app.FindSpotifyTracks()
	app.AddSpotifyTracks()

	log.Print("Total mp3 tracks: " + strconv.Itoa(len(app.Tracks)))
	log.Print("Added Spotify tracks: " + strconv.Itoa(len(app.RemoteTrackIds)))
}

func (app *App) FindSpotifyTracks() {
	for _, track := range app.Tracks {
		// query like "artist:" don't work unless the artist is the _exact_ name.
		query := track.Name + " " + track.Artist
		searchResult, err := app.Client.Search(query, spotify.SearchTypeTrack)
		if err != nil {
			log.Print(err)
		}
		newtrack := track.Artist + " - " + track.Name
		if len(searchResult.Tracks.Tracks) == 0 {
			log.Println(" :( FAILED " + newtrack)
		} else {
			log.Println(" :) PASSED " + newtrack)
			app.RemoteTrackIds = append(app.RemoteTrackIds, searchResult.Tracks.Tracks[0].ID)
		}
	}
}

func (app *App) CreateSpotifyPlaylist() {
	user, _ := app.Client.CurrentUser()
	playlistName := "L2S " + time.Now().String()
	app.Playlist, _ = app.Client.CreatePlaylistForUser(user.ID, playlistName, "this works", false)
}

func (app *App) AddSpotifyTracks() {
	totalCount := len(app.RemoteTrackIds)
	index := 0

	for index < totalCount {
		var i int
		if totalCount-index >= 100 {
			i = 100
		} else {
			i = totalCount % 100
		}
		ids := app.RemoteTrackIds[index : index+i]
		_, err := app.Client.AddTracksToPlaylist(app.Playlist.ID, ids...)
		if err != nil {
			log.Panic("error adding track id", err)
		}
		index += i
	}
}

# toSpotify
A command-line tool written in Go that allows users to quickly create a Spotify playlist from songs stored locally(.mp3, .wav, .flac, etc.)

### The purpose of this project is to learn and practice concepts related to:
- Go Language
- The Spotify Web API

Used: [A Go wrapper for the Spotify Web API](https://github.com/zmb3/spotify) by zmb3

## Getting Started

1) Create a app on [Spotify's developer dashboard](https://developer.spotify.com/dashboard/applications).

2) Add `http://localhost:3000/callback` in Redirect URIs in the edit settings of the app.

3) Create `.env` and fill in your app id and secret (`example.env`  for reference).

4) Run  `go run main.go path/to/music`

> Song must be in the format **[Artist] - [Track]** e.g. "Post Malone - Money Made Me Do It"

> To add more file extensions, add them to the `allowedExtensions` in trackparser.go

demo:

![demo](https://user-images.githubusercontent.com/78294692/212536072-81cc88b9-d262-446a-8628-8bb066e3eef9.gif)

# toSpotify
A command-line tool written in Go that allows users to quickly create a Spotify playlist from songs stored locally (.mp3, .wav, .flac, etc.)

I made this because I had way too many songs on my old computer, adding them to a playlist manually would take forever. Also, I wanted to learn Go and with this I got to play around with it and the Spotify Web API.

Thanks to the [Go wrapper](https://github.com/zmb3/spotify) for the Spotify Web API by zmb3.

## Getting Started

1) Create a app on [Spotify's developer dashboard](https://developer.spotify.com/dashboard/applications).

2) Add `http://localhost:3000/callback` in Redirect URIs in the edit settings of the app.

3) Create `.env` and fill in your app id and secret (`.env.example` for reference).

4) Run  `go run main.go path/to/music`

> Song must be in the format **[Artist] - [Track]** e.g. "Post Malone - Money Made Me Do It"

> To add more file extensions, add them to the `allowedExtensions` in trackparser.go

demo:

![demo](https://user-images.githubusercontent.com/78294692/212536072-81cc88b9-d262-446a-8628-8bb066e3eef9.gif)

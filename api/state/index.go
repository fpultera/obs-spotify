package handler

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"

	"github.com/zmb3/spotify/v2"
	"golang.org/x/oauth2"
	"obs-spotify/pkg/spotifyclient"
)

// simpleError es una función helper para enviar errores como JSON.
func simpleError(w http.ResponseWriter, message string, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}

func Handler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("spotify_token")
	if err != nil {
		simpleError(w, "Not authenticated. Please go to /api/login", http.StatusUnauthorized)
		return
	}

	// Decodificamos el valor de la cookie de Base64 a JSON.
	decodedJSON, err := base64.StdEncoding.DecodeString(cookie.Value)
	if err != nil {
		simpleError(w, "Could not decode cookie", http.StatusBadRequest)
		return
	}

	log.Printf("DEBUG: Decoded cookie value: %s", string(decodedJSON))

	var token oauth2.Token
	err = json.Unmarshal(decodedJSON, &token)
	if err != nil {
		log.Printf("ERROR: Unmarshalling cookie failed: %v", err)
		simpleError(w, "Invalid token format in cookie", http.StatusBadRequest)
		return
	}

	auth := spotifyclient.GetAuthenticator()
	client := spotify.New(auth.Client(r.Context(), &token))

	current, err := client.PlayerCurrentlyPlaying(r.Context())
	if err != nil {
		simpleError(w, "Could not get current track from Spotify. Token might be expired.", http.StatusUnauthorized)
		return
	}

	resp := map[string]interface{}{
		"is_playing": false,
	}

	if current != nil && current.Item != nil {
		resp["is_playing"] = current.Playing
		resp["title"] = current.Item.Name
		var artists string
		for i, a := range current.Item.Artists {
			if i > 0 {
				artists += ", "
			}
			artists += a.Name
		}
		resp["artists"] = artists
		resp["album"] = current.Item.Album.Name
		if len(current.Item.Album.Images) > 0 {
			resp["cover_url"] = current.Item.Album.Images[0].URL
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}


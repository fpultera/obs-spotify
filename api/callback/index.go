package handler

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"time"

	"obs-spotify/pkg/spotifyclient"
)

// Handler handles the Spotify callback after login.
func Handler(w http.ResponseWriter, r *http.Request) {
	auth := spotifyclient.GetAuthenticator()

	code := r.URL.Query().Get("code")
	if code == "" {
		http.Error(w, "Error: 'code' parameter not found in URL.", http.StatusBadRequest)
		return
	}

	token, err := auth.Exchange(r.Context(), code)
	if err != nil {
		http.Error(w, "Error exchanging code for token: "+err.Error(), http.StatusInternalServerError)
		return
	}

	tokenData := map[string]interface{}{
		"access_token":  token.AccessToken,
		"token_type":    token.TokenType,
		"refresh_token": token.RefreshToken,
		"expiry":        token.Expiry,
	}

	tokenJSON, err := json.Marshal(tokenData)
	if err != nil {
		http.Error(w, "Error processing token.", http.StatusInternalServerError)
		return
	}

	// Codificamos el JSON a Base64 para que sea seguro para la cookie.
	encodedToken := base64.StdEncoding.EncodeToString(tokenJSON)

	cookie := http.Cookie{
		Name:     "spotify_token",
		Value:    encodedToken, // Usamos el valor codificado
		Path:     "/",
		Expires:  time.Now().Add(time.Hour * 1),
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}

	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/api/widget", http.StatusFound)
}


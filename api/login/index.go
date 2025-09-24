package handler

import (
	"net/http"
	"obs-spotify/pkg/spotifyclient" // Importamos nuestro paquete compartido
)

// Handler redirige al usuario al login de Spotify
func Handler(w http.ResponseWriter, r *http.Request) {
	// Forzando una nueva compilación v2
	auth := spotifyclient.GetAuthenticator()
	url := auth.AuthURL("state-token") // El state debería ser aleatorio y verificado
	http.Redirect(w, r, url, http.StatusFound)
}



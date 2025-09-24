package spotifyclient

import (
	"log"
	"os"
	"sync" // Importamos el paquete de sincronización

	spotifyauth "github.com/zmb3/spotify/v2/auth"
)

// Usaremos estas variables para asegurar que la inicialización ocurra solo una vez.
var (
	authenticator *spotifyauth.Authenticator
	once          sync.Once
)

// GetAuthenticator ahora se encarga de la inicialización de forma segura.
func GetAuthenticator() *spotifyauth.Authenticator {
	// sync.Once.Do se asegura de que el código dentro de la función
	// se ejecute exactamente una vez, sin importar cuántas veces
	// se llame a GetAuthenticator() o desde cuántas peticiones simultáneas.
	once.Do(func() {
		// La lógica de leer las variables de entorno se mueve aquí.
		// Ahora se ejecuta la primera vez que un handler llama a esta función.
		redirectURI := os.Getenv("REDIRECT_URI")
		clientID := os.Getenv("CLIENT_ID")
		clientSecret := os.Getenv("CLIENT_SECRET")

		// La comprobación de errores se mantiene igual.
		if redirectURI == "" || clientID == "" || clientSecret == "" {
			log.Fatalf("FATAL: Required environment variables are missing (REDIRECT_URI, CLIENT_ID, CLIENT_SECRET)")
		}

		// Creamos la instancia del autenticador.
		authenticator = spotifyauth.New(
			spotifyauth.WithRedirectURL(redirectURI),
			spotifyauth.WithClientID(clientID),
			spotifyauth.WithClientSecret(clientSecret),
			spotifyauth.WithScopes(
				spotifyauth.ScopeUserReadCurrentlyPlaying,
				spotifyauth.ScopeUserReadPlaybackState,
			),
		)
	})

	return authenticator
}


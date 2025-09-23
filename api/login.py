import os
from spotipy.oauth2 import SpotifyOAuth
from flask import redirect

# Variables de entorno
CLIENT_ID = os.getenv("CLIENT_ID")
CLIENT_SECRET = os.getenv("CLIENT_SECRET")
REDIRECT_URI = os.getenv("REDIRECT_URI", "https://obs-spotify.vercel.app/api/callback")
SCOPE = "user-read-currently-playing user-read-playback-state"

def handler(request):
    # Crear el objeto OAuth sin levantar servidor local ni navegador
    sp_oauth = SpotifyOAuth(
        client_id=CLIENT_ID,
        client_secret=CLIENT_SECRET,
        redirect_uri=REDIRECT_URI,
        scope=SCOPE,
        cache_path="/tmp/.cache-spotify-widget",
        show_dialog=True  # siempre muestra login si no hay token cached
    )

    # Generar la URL de autorización
    auth_url = sp_oauth.get_authorize_url()
    # Redirigir al usuario a Spotify para autorizar
    return redirect(auth_url)

import os
from spotipy.oauth2 import SpotifyOAuth
from flask import redirect

# Variables de entorno
CLIENT_ID = os.getenv("CLIENT_ID")
CLIENT_SECRET = os.getenv("CLIENT_SECRET")
REDIRECT_URI = os.getenv("REDIRECT_URI", "https://obs-spotify.vercel.app/api/callback")
SCOPE = "user-read-currently-playing user-read-playback-state"
CACHE_PATH = "/tmp/.cache-spotify-widget"

def handler(request):
    # Limpiar cache previo para evitar que Spotipy intente abrir servidor local
    if os.path.exists(CACHE_PATH):
        os.remove(CACHE_PATH)

    # Crear objeto OAuth seguro para serverless
    sp_oauth = SpotifyOAuth(
        client_id=CLIENT_ID,
        client_secret=CLIENT_SECRET,
        redirect_uri=REDIRECT_URI,
        scope=SCOPE,
        cache_path=CACHE_PATH,
        show_dialog=True  # fuerza el login si no hay token cached
    )

    # Obtener URL de autorización de Spotify
    auth_url = sp_oauth.get_authorize_url()
    # Redirigir al usuario a Spotify
    return redirect(auth_url)

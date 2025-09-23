import os
from spotipy.oauth2 import SpotifyOAuth
from flask import redirect

CLIENT_ID = os.getenv("CLIENT_ID")
CLIENT_SECRET = os.getenv("CLIENT_SECRET")
REDIRECT_URI = os.getenv("REDIRECT_URI", "https://tuapp.vercel.app/api/callback")
SCOPE = "user-read-currently-playing user-read-playback-state"

def handler(request):
    # Crear el objeto OAuth sin intentar abrir navegador ni server local
    sp_oauth = SpotifyOAuth(
        client_id=CLIENT_ID,
        client_secret=CLIENT_SECRET,
        redirect_uri=REDIRECT_URI,
        scope=SCOPE,
        cache_path="/tmp/.cache-spotify-widget",
        show_dialog=True  # fuerza mostrar login si no hay token cached
    )

    auth_url = sp_oauth.get_authorize_url()
    return redirect(auth_url)

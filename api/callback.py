import os
from spotipy.oauth2 import SpotifyOAuth
from flask import Response

# Variables de entorno
CLIENT_ID = os.getenv("CLIENT_ID")
CLIENT_SECRET = os.getenv("CLIENT_SECRET")
REDIRECT_URI = os.getenv("REDIRECT_URI", "https://obs-spotify.vercel.app/api/callback")
SCOPE = "user-read-currently-playing user-read-playback-state"
CACHE_PATH = "/tmp/.cache-spotify-widget"

def handler(request):
    # Crear objeto OAuth seguro para serverless
    sp_oauth = SpotifyOAuth(
        client_id=CLIENT_ID,
        client_secret=CLIENT_SECRET,
        redirect_uri=REDIRECT_URI,
        scope=SCOPE,
        cache_path=CACHE_PATH
    )

    # Obtener el 'code' que envía Spotify a la callback
    code = request.args.get("code")
    if code:
        # Intercambiar el code por un access token
        token_info = sp_oauth.get_access_token(code, as_dict=True)
        if token_info and "access_token" in token_info:
            return Response("✅ Autenticado con Spotify. Ahora podés ir a /api/widget")

    return Response("❌ Error: falta el code o no se pudo obtener el token")

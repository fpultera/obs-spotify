import os
from spotipy.oauth2 import SpotifyOAuth
from flask import Response

CLIENT_ID = os.getenv("CLIENT_ID")
CLIENT_SECRET = os.getenv("CLIENT_SECRET")
REDIRECT_URI = os.getenv("REDIRECT_URI", "https://tuapp.vercel.app/api/callback")
SCOPE = "user-read-currently-playing user-read-playback-state"

def handler(request):
    sp_oauth = SpotifyOAuth(client_id=CLIENT_ID,
                            client_secret=CLIENT_SECRET,
                            redirect_uri=REDIRECT_URI,
                            scope=SCOPE,
                            cache_path="/tmp/.cache-spotify-widget")
    code = request.args.get("code")
    if code:
        sp_oauth.get_access_token(code)
        return Response("✅ Autenticado con Spotify. Ahora podés ir a /api/widget")
    return Response("❌ Error: falta el code")

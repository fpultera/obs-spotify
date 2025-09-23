import os
import spotipy
from spotipy.oauth2 import SpotifyOAuth
from flask import jsonify

CLIENT_ID = os.getenv("CLIENT_ID")
CLIENT_SECRET = os.getenv("CLIENT_SECRET")
REDIRECT_URI = os.getenv("REDIRECT_URI", "https://tuapp.vercel.app/api/callback")
SCOPE = "user-read-currently-playing user-read-playback-state"

def handler(request):
    sp_oauth = SpotifyOAuth(
        client_id=CLIENT_ID,
        client_secret=CLIENT_SECRET,
        redirect_uri=REDIRECT_URI,
        scope=SCOPE,
        cache_path="/tmp/.cache-spotify-widget"
    )

    token_info = sp_oauth.get_cached_token()
    if not token_info:
        return jsonify({"error": "Not authenticated. Go to /api/login"})

    spotify = spotipy.Spotify(auth_manager=sp_oauth)
    now = spotify.current_user_playing_track()
    if now and now.get("item"):
        item = now["item"]
        return jsonify({
            "is_playing": now.get("is_playing", False),
            "title": item["name"],
            "artists": ", ".join(a["name"] for a in item["artists"]),
            "album": item["album"]["name"],
            "cover_url": item["album"]["images"][0]["url"] if item["album"]["images"] else None
        })
    return jsonify({"is_playing": False})

package handler

import (
	"net/http"
)

var htmlTemplate = `
<!doctype html>
<html lang="es">
<head>
<meta charset="utf-8" />
<title>Spotify Now Playing</title>
<style>
body {margin:0;font-family:Arial,sans-serif;background:transparent;display:flex;justify-content:center;align-items:center;height:100vh;}
.card {display:flex;gap:16px;align-items:center;padding:12px;border-radius:12px;background:rgba(0,0,0,0.6);color:#fff;min-width:300px;box-shadow:0 4px 12px rgba(0,0,0,0.3);}
.cover {width:96px;height:96px;overflow:hidden;border-radius:6px;flex-shrink:0;}
.cover img {width:100%;height:100%;object-fit:cover;}
.meta {display:flex;flex-direction:column;gap:6px;}
.title {font-weight:700;font-size:1.1rem;}
.artists, .album {opacity:0.9;font-size:0.9rem;}
.status {opacity:0.7;font-size:0.75rem;margin-top:6px;}
.status a {color: #1DB954; text-decoration: none;}
</style>
</head>
<body>
<div class="card">
  <div class="cover"><img id="cover" src="https://placehold.co/96x96/000000/FFFFFF?text=Spotify" alt="Album cover"/></div>
  <div class="meta">
    <div id="title" class="title"></div>
    <div id="artists" class="artists"></div>
    <div id="album" class="album"></div>
    <div id="status" class="status">Cargando...</div>
  </div>
</div>
<script>
async function fetchState(){
  try {
    const res = await fetch('/api/state');

    if (res.status === 401) {
      document.getElementById('status').innerHTML = 'No autenticado. <a href="/api/login" target="_blank">Iniciar sesión</a>.';
      document.getElementById('title').textContent = '';
      document.getElementById('artists').textContent = '';
      document.getElementById('album').textContent = '';
      return;
    }

    if (!res.ok) {
      throw new Error('La respuesta del servidor no fue OK');
    }

    const j = await res.json();
    
    document.getElementById('title').textContent = j.title || 'Canción Desconocida';
    document.getElementById('artists').textContent = j.artists || 'Artista Desconocido';
    document.getElementById('album').textContent = j.album || '';
    document.getElementById('status').textContent = j.is_playing ? '▶ Reproduciendo' : '❚❚ Pausado';
    if(j.cover_url) {
      document.getElementById('cover').src = j.cover_url;
    }

  } catch (error) {
    console.error("Error al obtener el estado:", error);
    document.getElementById('status').textContent = 'Error al cargar los datos.';
  }
}
setInterval(fetchState, 3000); // Call every 3 seconds
fetchState();
</script>
</body>
</html>
`

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(htmlTemplate))
}


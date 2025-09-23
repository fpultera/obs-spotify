from flask import Response

HTML_TEMPLATE = """
<!doctype html>
<html lang="es">
<head>
<meta charset="utf-8" />
<style>
body {margin:0;font-family:Arial,sans-serif;background:transparent;display:flex;justify-content:center;align-items:center;height:100vh;}
.card {display:flex;gap:16px;align-items:center;padding:12px;border-radius:12px;background: rgba(0,0,0,0.45); color:#fff;}
.cover {width:128px;height:128px;overflow:hidden;border-radius:6px;}
.cover img {width:100%;height:100%;object-fit:cover;}
.meta {display:flex;flex-direction:column;gap:6px;}
.title {font-weight:700;}
.artists, .album {opacity:0.9;}
.small {opacity:0.7;font-size:0.75rem;margin-top:6px;}
</style>
</head>
<body>
<div class="card">
  <div class="cover"><img id="cover" src=""/></div>
  <div class="meta">
    <div class="title" id="title"></div>
    <div class="artists" id="artists"></div>
    <div class="album" id="album"></div>
    <div class="small" id="status"></div>
  </div>
</div>
<script>
async function fetchState(){
  const res = await fetch('/api/state');
  const j = await res.json();
  document.getElementById('title').textContent = j.title||'';
  document.getElementById('artists').textContent = j.artists||'';
  document.getElementById('album').textContent = j.album||'';
  document.getElementById('status').textContent = j.is_playing?'Reproduciendo':'No reproduciendo';
  if(j.cover_url) document.getElementById('cover').src = j.cover_url;
}
setInterval(fetchState,2000);
fetchState();
</script>
</body>
</html>
"""

def handler(request):
    return Response(HTML_TEMPLATE, mimetype="text/html")

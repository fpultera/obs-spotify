package handler

import (
	"encoding/json"
	"net/http"
	"os"
)

// Handler es un endpoint de depuración para verificar las variables de entorno.
func Handler(w http.ResponseWriter, r *http.Request) {
	// Creamos un mapa para guardar el estado de las variables.
	vars := map[string]string{
		"CLIENT_ID":     os.Getenv("CLIENT_ID"),
		"REDIRECT_URI":  os.Getenv("REDIRECT_URI"),
		"CLIENT_SECRET": os.Getenv("CLIENT_SECRET"),
	}

	// Por seguridad, no mostramos el secret completo, solo si existe.
	if vars["CLIENT_SECRET"] != "" {
		vars["CLIENT_SECRET"] = "[CARGADO]"
	} else {
		vars["CLIENT_SECRET"] = "[NO ENCONTRADO]"
	}

	// Hacemos lo mismo para las otras variables para que la salida sea clara.
	if vars["CLIENT_ID"] == "" {
		vars["CLIENT_ID"] = "[NO ENCONTRADO]"
	}
	if vars["REDIRECT_URI"] == "" {
		vars["REDIRECT_URI"] = "[NO ENCONTRADO]"
	}

	// Devolvemos el resultado como JSON.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(vars)
}

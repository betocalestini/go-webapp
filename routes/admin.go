package routes

import "net/http"

func adminGetHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Você está logado!"))
}

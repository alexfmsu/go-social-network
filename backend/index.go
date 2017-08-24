package backend

import (
	// "github.com/gorilla/websocket"
	"html/template"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "text/plain")
	t, _ := template.ParseFiles("./templates/index.gtpl")
	t.Execute(w, nil)
}

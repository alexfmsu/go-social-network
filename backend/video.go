package backend

import (
	"../settings"
	// "fmt"
	"html/template"
	"net/http"
)

func Video(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles(
		settings.Template["video"],
		settings.Template["header"],
		settings.Template["nav"],
		settings.Template["footer"],
	)

	t.Execute(w, nil)
}

package backend

import (
	"../settings"
	"fmt"
	"html/template"
	"net/http"
)

func Photos(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(
		settings.Template["photos"],
		settings.Template["header"],
		settings.Template["nav"],
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	t.Execute(w, nil)
}

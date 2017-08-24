package backend

import (
	"../settings"
	"fmt"
	"html/template"
	"net/http"
)

func Messages(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(
		settings.Template["messages"],
		settings.Template["header"],
		settings.Template["nav"],
		settings.Template["footer"],
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	t.Execute(w, nil)
}

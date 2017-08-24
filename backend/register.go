package backend

import (
	"../settings"
	"fmt"
	"html/template"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles(
		settings.Template["title"],
		settings.Template["nav"],
		settings.Template["friends"],
	)

	if r.Method == "GET" {
		t.Execute(w, "")
	} else if r.Method == "POST" {

		//запрос в БД на проверку наличия юзера

		var result_obj *settings.User_obj

		result_obj = settings.Check_login(r.FormValue("username"), GetMD5Hash(r.FormValue("password")))
		//fmt.Print(result_obj.id)

		if result_obj == nil { // eсли юзер существует
			http.Redirect(w, r, "/login", 301)
			settings.CreateUser(r.FormValue("username"), r.FormValue("password"), r.FormValue("email"), r.FormValue("first_name"), r.FormValue("last_name"))

		} else {
			// t.Execute(w, map[string]string{"Title": "My title", "Body": "Hi this is my body"})
			// t.ExecuteTemplate(w, map[string]string{"Title": "My title", "Body": "Hi this is my body"})
			// t.Execute(w, map[string]string{"Title": "My title", "Body": "Hi this is my body"})
			// t.Execute(w, "html")
			t.Execute(w, "name")
			// t.ExecuteTemplate(w, "html", "User already exist!")
			// t.ExecuteTemplate(w, "bar", "User already exist!")

		}
	}

	fmt.Print(r.Method)
}

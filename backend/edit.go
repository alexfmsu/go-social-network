package backend

import (
	"net/http"
	"../settings"
	"html/template"
	"fmt"
)

func Edit(w http.ResponseWriter, r *http.Request) {

	//вернет обьект юзера или редиректнет
	user := settings.CheckAuth(w,r)

	t, _ := template.ParseFiles("./templates/edit.gtpl")




	if(r.Method=="GET") {
		fmt.Print(user.Id, "\n")
		t.Execute(w, *user,)
	} else if (r.Method == "POST"){

		if r.FormValue("username") != "" && r.FormValue("email") != "" && r.FormValue("first_name") != "" && r.FormValue("last_name") != ""{


			settings.UpdateUserData(user.Id,r.FormValue("username"),r.FormValue("email"),r.FormValue("first_name"),r.FormValue("last_name"))

			http.Redirect(w, r, "/profile", 301)

		}

	}

}
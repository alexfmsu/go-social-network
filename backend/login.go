package backend

import (
	"../settings"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"html/template"
	"net/http"
)

func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))

	return hex.EncodeToString(hasher.Sum(nil))
}

func Login(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(
		settings.Template["login"],

		settings.Template["header"],
		settings.Template["footer"],
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	if r.Method == "GET" {
		t.Execute(w, "")
	} else if r.Method == "POST" {

		//запрос в БД на проверку наличия юзера

		var result_obj *settings.User_obj

		result_obj = settings.Check_login(r.FormValue("username"), GetMD5Hash(r.FormValue("password")))
		t, _ := template.ParseFiles("./templates/login.gtpl")

		//fmt.Print(result_obj.id)

		if result_obj == nil { // если пароль неверен то ошибка
			t.Execute(w, "Your login: "+r.FormValue("username")+
				" and password:"+r.FormValue("password")+" not found!")
		} else {
			//если все верно то редирект
			http.SetCookie(w, settings.GetCookieFromID(result_obj.Id))

			http.Redirect(w, r, "/profile", 301)

			//создаем токен и добавляем хедер с куками

			//hex(aes(user_id))

			t.Execute(w, "")

		}

		//запрос в БД

	}
}

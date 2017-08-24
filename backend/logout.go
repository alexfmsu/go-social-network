package backend

import (
	"net/http"
)

func Logout(w http.ResponseWriter, r *http.Request) {

	//вернет обьект юзера или редиректнет
	cookie := http.Cookie{Name: "token", Value: ""}

	http.SetCookie(w, &cookie)

	http.Redirect(w, r, "/login", 301)

}

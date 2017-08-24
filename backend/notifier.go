package backend

import (
	// "../settings"
	"fmt"
	"net/http"
)

func Notifier2(w http.ResponseWriter, r *http.Request) {
	// if r.Method == "GET" {
	// fmt.Println(r.FormValue("user_2"))

	w.Write([]byte("sdd"))
	fmt.Println("333")
	// return
	// }
	// db := settings.GetDB()

	// u := settings.CheckAuth(w, r)

	// // var new_friends int
	// // var new_messages int

	// var err error
	// err = db.QueryRow("SELECT new_friends, new_messages FROM users where id = ?", u.Id).Scan(&u.NewFriends, &u.NewMessages)

	// fmt.Println("ABCD")
	// fmt.Println("n-f:", u.NewFriends)
	// fmt.Println("n-m:", u.NewMessages)

	// if err != nil {
	// 	fmt.Println("err")
	// }

	// if u.NewFriends > 0
	// return
	// var w2 http.ResponseWriter
	// w2.Write([]byte("ABCD"))

	// w = w2
	// w.Write([]byte(fmt.Sprint(u.NewFriends)))
}

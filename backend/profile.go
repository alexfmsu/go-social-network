package backend

import (
	"../settings"
	"io"
	"os"
	// "../user"
	"fmt"
	"html/template"
	"net/http"
	// "time"
)

func Profile(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles(
		settings.Template["profile"],
		settings.Template["header"],
		settings.Template["nav"],
		settings.Template["footer"],
	)

	//вернет обьект юзера или редиректнет
	user := settings.CheckAuth(w, r)

	if r.Method == "POST" {
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("uploadfile")
		// fmt.Println(file)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		// fmt.Fprintf(w, "%v", handler.Header)
		f, err := os.OpenFile("./templates/images/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)

		// return
	}
	// if err == sql.ErrNoRows {
	// 	return nil //если не найдет юзера
	// } else if err != nil {
	// 	return nil //если другая ошибка
	// }

	// go func(u *settings.User_obj) {
	// 	db := settings.GetDB()
	// 	var err error

	// 	for {
	// 		timeout := time.Tick(1 * time.Second)

	// 		select {

	// 		case <-timeout:
	// 			err = db.QueryRow("SELECT new_friends, new_messages FROM users where id = ?", u.Id).Scan(&u.NewFriends, &u.NewMessages)

	// 			// fmt.Println("n-m:", u.NewMessages)

	// 			if err != nil {
	// 				fmt.Println("err")
	// 			}

	// 			if u.NewFriends > 0 {
	// 				fmt.Println("n-f:", u.NewFriends)
	// 				// fmt.Print(user.Id, "\n")
	// 				// 	t, _ := template.ParseFiles("./templates/friends.gtpl")

	// 				// 	t.Execute(w, "sd")
	// 				return
	// 			}
	// 		}
	// 	}
	// }(user)

	fmt.Print(user.Id, "\n")
	t.Execute(w, *user)
}

// func notifier(u *settings.User_obj) {
// 	db := settings.GetDB()

// 	// var new_friends int
// 	// var new_messages int

// 	var err error

// 	for {
// 		timeout := time.Tick(1 * time.Second)

// 		select {

// 		case <-timeout:
// 			err = db.QueryRow("SELECT new_friends, new_messages FROM users where id = ?", u.Id).Scan(&u.NewFriends, &u.NewMessages)

// 			fmt.Println("n-f:", u.NewFriends)
// 			fmt.Println("n-m:", u.NewMessages)

// 			if err != nil {
// 				fmt.Println("err")
// 			}
// 		}
// 	}
// }

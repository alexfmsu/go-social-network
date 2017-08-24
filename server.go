// <
package main

import (
	"./backend" // rename
	"./db"      // rename
	"fmt"
	// "github.com/gorilla/websocket"
	"net/http"
)

// var upgrader = websocket.Upgrader{
// 	ReadBufferSize:  1024,
// 	WriteBufferSize: 1024,
// }

// >

func main() {
	db.ConnectDB()
	db.InitStmt()

	http.HandleFunc("/", backend.Index)
	fmt.Println("123")
	http.HandleFunc("/register", backend.Register)
	http.HandleFunc("/recover", backend.Recover)
	http.HandleFunc("/login", backend.Login)
	http.HandleFunc("/logout", backend.Logout)

	http.HandleFunc("/profile", backend.Profile)
	http.HandleFunc("/edit", backend.Edit)
	http.HandleFunc("/friends", backend.Friends)
	http.HandleFunc("/messages", backend.Messages)
	http.HandleFunc("/photos", backend.Photos)
	http.HandleFunc("/audio", backend.Audio)
	http.HandleFunc("/video", backend.Video)
	http.HandleFunc("/add_friend", backend.AddFriend)
	http.HandleFunc("/confirm_friend", backend.ConfirmFriend)

	http.HandleFunc("/nofifier2", backend.Notifier2)

	// http.HandleFunc("/websocket", func(w http.ResponseWriter, r *http.Request) {
	// 	conn, err := upgrader.Upgrade(w, r, nil)

	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	}
	// 	// myJson := map[string]string{"av": "s"}
	// 	msgType, msg, _ := conn.ReadMessage()

	// 	conn.WriteMessage(websocket.TextMessage, []byte("sd"))

	// 	fmt.Println(msgType, msg)
	// 	fmt.Println("Client subscribed")

	// })

	http.ListenAndServe(":8999", nil)
}

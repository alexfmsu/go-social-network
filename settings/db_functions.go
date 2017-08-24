package settings

import (
	"../db"
	"database/sql"
	"fmt"
	"log"
)

type User_obj struct {
	Id         int
	Login      string
	Pass       string
	Email      string
	First_name string
	Last_name  string
	Gender     string

	NewFriends  int
	NewMessages int
}

func Get_user_by_id(id int) *User_obj {
	db := db.GetDB()

	row := db.QueryRow(
		"SELECT id, login, pass, email, firstname, lastname, gender FROM users WHERE id=?",
		id,
	)

	var user User_obj

	err := row.Scan(&user.Id, &user.Login, &user.Pass, &user.Email, &user.First_name, &user.Last_name, &user.Gender)

	if err == sql.ErrNoRows {
		return nil //если не найдет юзера
	} else if err != nil {
		return nil //если другая ошибка
	}

	return &user
}

func (u *User_obj) GetUpdates() map[string]int {
	db := db.GetDB()

	err := db.QueryRow("SELECT new_friends, new_messages FROM users where id = ?", u.Id).Scan(&u.NewFriends, &u.NewMessages)
	if err != nil {
		fmt.Println("err")
	}

	return map[string]int{
		"new_friends":  u.NewFriends,
		"new_messages": u.NewMessages,
	}
}

//для авторизации
func Check_login(login string, pass string) *User_obj {
	db := db.GetDB()

	row := db.QueryRow(
		"SELECT id, login, pass, email, firstname, lastname, gender FROM users WHERE login = ? AND pass = ?",
		login, pass,
	)

	var user User_obj

	err := row.Scan(&user.Id, &user.Login, &user.Pass, &user.Email, &user.First_name, &user.Last_name, &user.Gender)
	if err == sql.ErrNoRows {
		return nil //если не найдет юзера
	} else if err != nil {
		log.Fatal(err)
		return nil //возвращает обьект пользователя
	}

	return &user
}

func Get_user_friends(id int) {

	//возвращает массив id друзей

}

func CreateUser(login string, password string, email string, first_name string, last_name string) {
	db := db.GetDB()

	res, err := db.Exec(
		"INSERT INTO users (login, pass, email, firstname, lastname) VALUES (?, md5(?), ?, ?, ?)",
		login, password, email, first_name, last_name,
	)

	fmt.Print(res, err)
}

func UpdateUserData(id int, login string, email string, first_name string, last_name string) {
	db := db.GetDB()

	res, err := db.Exec(
		"UPDATE users  SET login = ?, email = ?, firstname = ?, lastname = ? WHERE id= ?",
		login, email, first_name, last_name, id,
	)

	fmt.Print(res, err)
}

//и подобные функции

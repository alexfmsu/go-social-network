package backend

import (
	"../db"
	"../settings"
	"database/sql"
	"html/template"
	"log"
	"net/http"
)

type dict map[string]string

func get_info(rows *sql.Rows) (map[string]dict, error) {
	res := make(map[string]dict, 0)

	for rows.Next() {
		var login string
		var email string
		var firstname string
		var lastname string
		var gender string

		err := rows.Scan(&login, &email, &firstname, &lastname, &gender)

		if err != nil {
			return res, err
		}

		res[login] = dict{
			"email":     email,
			"firstname": firstname,
			"lastname":  lastname,
			"gender":    gender,
		}
	}

	return res, nil
}

func Friends(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles(
		settings.Template["friends"],
		settings.Template["header"],
		settings.Template["nav"],
		settings.Template["footer"],
	)

	s := r.FormValue("search")

	u := settings.CheckAuth(w, r)
	user_1 := settings.CheckAuth(w, r).Login

	msg := make(map[string]map[string]dict, 0)
	msg["friends"] = make(map[string]dict, 0)

	if s == "" {
		updates := u.GetUpdates()

		log.Println(updates)

		if updates["new_friends"] != 0 {
			res, err := db.InviterInfo.Query(u.Login, "0")
			db.CheckQueryOnErr(err)

			msg["invite"] = make(map[string]dict, 0)

			msg["invite"], _ = get_info(res)
		}

		rows, err := db.GetFriends.Query(u.Login, u.Login)
		db.CheckQueryOnErr(err)

		for rows.Next() {
			var login string

			err := rows.Scan(&login)
			db.CheckQueryOnErr(err)

			msg["friends"][login] = dict{
				"email": login,
			}
		}
	} else {
		res, err := db.SearchUser.Query(s, s, s)
		db.CheckQueryOnErr(err)

		msg["found_users"] = make(map[string]dict, 0)
		for res.Next() {
			var login string
			var email string
			var firstname string
			var lastname string
			// var gender string

			err = res.Scan(&login, &email, &firstname, &lastname)

			// user := User{login, email, firstname, lastname, ""}

			if user_1 != login {
				msg["found_users"][login] = dict{
					"email":     email,
					"firstname": firstname,
					"lastname":  lastname,
					// "gender":    gender,
				}
			}
		}

		if len(msg["found_users"]) == 0 {
			// log.Println("len:", len(mapfound_users))

			// msg["err"] = make(map[string]dict, 0)
			// msg["err"] = map[string]string{
			// 	"err_msg": "Nothing's found",
			// }

			// t.Execute(w, msg)
		} else {
		}
	}

	t.Execute(w, msg)
}

func AddFriend(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		user_1 := settings.CheckAuth(w, r).Login
		user_2 := r.FormValue("user_2")

		// log.Println("user_1:", user_1)
		// log.Println("user_2:", user_2)

		res, err := db.AddInvite.Exec(user_1, user_2)
		db.CheckQueryOnErr(err)

		affect, err := res.RowsAffected()

		ans := "0"

		if affect == 1 {
			ans = "1"
		}

		res, err = db.IncInviters.Exec(user_2)
		db.CheckQueryOnErr(err)

		affect, err = res.RowsAffected()

		db.CheckQueryOnErr(err)

		log.Println("affect:", affect)

		// var new_friends int
		// var new_messages int

		// err := row.Scan(&new_messages, &new_friends)

		// if err != nil {
		// 	log.Println("ERR")
		// 	// 	return
		// }
		// log.Println("new_friends:", new_friends)
		// log.Println("new_messages:", new_messages)

		w.Write([]byte(ans))
	}

}

func ConfirmFriend(w http.ResponseWriter, r *http.Request) {
	log.Println("confirm")

	if r.Method == "GET" {
		user_1 := r.FormValue("user_1")
		user_2 := settings.CheckAuth(w, r).Login

		// log.Println("user_1:", user_1)
		// log.Println("user_2:", user_2)

		res, err := db.SetFriendshipStatus.Exec("1", user_1, user_2)
		db.CheckQueryOnErr(err)

		affect, err := res.RowsAffected()

		ans := "0"

		if affect == 1 {
			ans = "1"
		}

		res, err = db.DecInviters.Exec(user_2)
		db.CheckQueryOnErr(err)

		affect4, err := res.RowsAffected()
		db.CheckQueryOnErr(err)

		log.Println("affect4:", affect4)
		// var new_friends int
		// var new_messages int

		// err := row.Scan(&new_messages, &new_friends)

		w.Write([]byte(ans))
	}
}

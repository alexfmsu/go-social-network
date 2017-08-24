package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func ConnectDB() {
	host := "31.31.196.159"
	db_name := "u0386911_social"
	pass := "^SocialNetwork$"
	login := "u0386911_user"
	port := "3306"

	db, _ = sql.Open("mysql", login+":"+pass+"@tcp("+host+":"+port+")/"+db_name)

	db.SetMaxOpenConns(10)
}

func GetDB() *sql.DB {
	return db
}

func prepareStmt(db *sql.DB, stmt string) *sql.Stmt {
	res, err := db.Prepare(stmt)

	if err != nil {
		log.Fatal("Could not prepare `" + stmt + "`: " + err.Error())
	}

	return res
}

func CheckQueryOnErr(err error) {
	if err != nil {
		log.Fatal("Cannot execute query:", err.Error())
	}
}

var (
	db *sql.DB

	InviterInfo         *sql.Stmt
	GetFriends          *sql.Stmt
	SearchUser          *sql.Stmt
	IncInviters         *sql.Stmt
	DecInviters         *sql.Stmt
	SetFriendshipStatus *sql.Stmt
	AddInvite           *sql.Stmt
)

func InitStmt() {
	InviterInfo = prepareStmt(db, "SELECT login, email, firstname, lastname, gender FROM users WHERE login = (SELECT user_1 FROM friends WHERE user_2 = ? AND status = ?)")
	GetFriends = prepareStmt(db, "SELECT user_1 FROM (SELECT user_1, user_2 FROM `friends` where status='1') as users WHERE user_2 = ? UNION SELECT user_2 FROM (SELECT user_1, user_2 FROM `friends` where status='1') as users WHERE user_1 = ?")
	SearchUser = prepareStmt(db, "SELECT login, email, firstname, lastname FROM users WHERE login = ? OR firstname = ? OR lastname = ?")
	IncInviters = prepareStmt(db, "UPDATE users SET new_friends = new_friends+1 WHERE login=?")
	DecInviters = prepareStmt(db, "UPDATE users SET new_friends = new_friends-1 WHERE login=?")
	SetFriendshipStatus = prepareStmt(db, "update friends set status = ? where user_1 = ? AND user_2 = ?")
	AddInvite = prepareStmt(db, "insert into friends(user_1, user_2) values(?, ?)")
}

package main

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"io/ioutil"
	"time"
)

// La fonction insert prend une table et une interface de données pour insérer ces données dans la BDD
func insert(table Table, value interface{}) error {
	// call example: insert(USERS, Users{0, "Alecs", "alecs@ynov.com", "Admin", "I am a Pokemon fanboy", "shorturl.at/qtxIN", ""})
	var stmt *sql.Stmt
	var res sql.Result
	var err error
	// insert
	switch table {
	case USERS:
		if fmt.Sprintf("%T", value) != "main.Users" {
			return errors.New("wrong table type")
		}
		stmt, err = db.Prepare(fmt.Sprintf("%s %s %s", "INSERT INTO ", table, "(nickname, email, role, biography, profileImage, status) values(?,?,?,?,?,?)"))
		checkErr(err)

		res, err = stmt.Exec(value.(Users).Nickname, value.(Users).Email, value.(Users).Role, value.(Users).Biography, value.(Users).ProfileImage, value.(Users).Status)
		checkErr(err)
	case CATEGORY:
		if fmt.Sprintf("%T", value) != "main.Category" {
			return errors.New("wrong table type")
		}
		stmt, err = db.Prepare(fmt.Sprintf("%s %s %s", "INSERT INTO ", table, "(name, description) values(?,?)"))
		checkErr(err)

		res, err = stmt.Exec(value.(Category).Name, value.(Category).Description)
		checkErr(err)
	case USERSCAT:
		if fmt.Sprintf("%T", value) != "main.UsersCat" {
			return errors.New("wrong table type")
		}
		stmt, err = db.Prepare(fmt.Sprintf("%s %s %s", "INSERT INTO ", table, "(id_users, id_category) values(?,?)"))
		checkErr(err)

		res, err = stmt.Exec(value.(UsersCat).id_users, value.(UsersCat).id_category)
		checkErr(err)
	case POSTS:
		if fmt.Sprintf("%T", value) != "main.Posts" {
			return errors.New("wrong table type")
		}
		stmt, err = db.Prepare(fmt.Sprintf("%s %s %s", "INSERT INTO ", table, "(title, creationDate, modificationDate, deleteDate, likes, dislikes, nickname_users) values(?,?,?,?,?,?,?)"))
		checkErr(err)

		res, err = stmt.Exec(value.(Posts).Title, value.(Posts).CreationDate, value.(Posts).ModificationDate, value.(Posts).DeleteDate, value.(Posts).Likes, value.(Posts).Dislikes, value.(Posts).Nickname_users)
		checkErr(err)
	case COMMENTS:
		if fmt.Sprintf("%T", value) != "main.Comments" {
			return errors.New("wrong table type")
		}
		stmt, err = db.Prepare(fmt.Sprintf("%s %s %s", "INSERT INTO ", table, "(content, creationDate, modificationDate, deleteDate, likes, dislikes, id_users, id_posts) values(?,?,?,?,?,?,?,?)"))
		checkErr(err)

		res, err = stmt.Exec(value.(Comments).Content, value.(Comments).CreationDate, value.(Comments).ModificationDate, value.(Comments).DeleteDate, value.(Comments).Likes, value.(Comments).Dislikes, value.(Comments).NicknameUsers, value.(Comments).IdPosts)
		checkErr(err)
	case BADGE:
		if fmt.Sprintf("%T", value) != "main.Badge" {
			return errors.New("wrong table type")
		}
		stmt, err = db.Prepare(fmt.Sprintf("%s %s %s", "INSERT INTO ", table, "(name, image, description) values(?,?,?)"))
		checkErr(err)

		res, err = stmt.Exec(value.(Badge).name, value.(Badge).image, value.(Badge).description)
		checkErr(err)
	case USERSBADGE:
		if fmt.Sprintf("%T", value) != "main.UsersBadge" {
			return errors.New("wrong table type")
		}
		stmt, err = db.Prepare(fmt.Sprintf("%s %s %s", "INSERT INTO ", table, "(id_users, id_badge) values(?,?)"))
		checkErr(err)

		res, err = stmt.Exec(value.(UsersBadge).id_users, value.(UsersBadge).id_badge)
		checkErr(err)
	}

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println("Value inserted into table=", table, " at id=", id)
	return nil
}

// Encore à modifier, mais pour l'instant permet de changer un username en un autre, en prenant une table spécifique
func update(table Table, value interface{}, id int) {
	var stmt *sql.Stmt
	var res sql.Result
	var err error

	// update
	stmt, err = db.Prepare(fmt.Sprintf("%s %s %s", "UPDATE ", table, " SET nickname=? WHERE id=?"))
	checkErr(err)

	res, err = stmt.Exec("astaxieupdate", id)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)
}

// La fonction delete prend une table et un id en paramètre pour retirer les informations correspondantes.
func delete(table Table, id int) {
	var stmt *sql.Stmt
	var res sql.Result
	var affect int64
	var err error

	stmt, err = db.Prepare(fmt.Sprintf("%s %s %s", "DELETE FROM ", table, " WHERE id=?"))
	checkErr(err)

	res, err = stmt.Exec(id)
	checkErr(err)

	affect, err = res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	db.Close()
}

// La fonction CreateTables permet de réinitialiser les tables en les supprimant et en les recréant, seules les tables précisées dans le code sont touchées
func CreateTables() string {

	dropTable, err := ioutil.ReadFile("data/sql/drop_tables.sql")
	checkErr(err)

	createTable, err := ioutil.ReadFile("data/sql/create_tables.sql")
	checkErr(err)

	basicdata, err := ioutil.ReadFile("data/sql/insert_data.sql")
	checkErr(err)

	tableList := string(dropTable) + string(createTable) + string(basicdata)
	return tableList
}

func registerBDD(nickname string, email string, hashedpwd string) {
	stmt, err := db.Prepare("INSERT INTO Users (nickname, email, hashedpwd, role) VALUES (?, ?, ?, ?)")
	checkErr(err)
	_, err = stmt.Exec(nickname, email, hashedpwd, "User")
	checkErr(err)
}

func CreateSession(id_user int, uuid string) {
	stmt, err := db.Prepare("INSERT INTO SessionControl (uuid, id_user) VALUES (?, ?)")
	checkErr(err)
	_, err = stmt.Exec(uuid, id_user)
	checkErr(err)
}

func CreatePost(nickname_users string, title string, content string, category string) {
	creationDate := time.Now()
	stmt, err := db.Prepare("INSERT INTO Posts (title, content, creationDate, nickname_users, category, status) VALUES (?, ?, ?, ?, ?, ?)")
	checkErr(err)
	_, err = stmt.Exec(title, content, creationDate, nickname_users, category, "Actif")
	checkErr(err)
}

func ModifiedPost(id_post int, content string) {
	modificationDate := time.Now()
	stmt, err := db.Prepare("UPDATE Posts SET content = ?, modificationDate = ? WHERE id = ?")
	checkErr(err)
	_, err = stmt.Exec(content, modificationDate, id_post)
	checkErr(err)
}

func DeletePost(id_post int) {
	stmt, err := db.Prepare("UPDATE Posts SET Status = ? WHERE id = ?")
	checkErr(err)
	_, err = stmt.Exec("Supprimé", id_post)
	checkErr(err)
}

func CreateComment(nickname_users string, content string, post int) {
	creationDate := time.Now()
	stmt, err := db.Prepare("INSERT INTO Comments (content, creationDate, id_users, id_posts ) VALUES (?, ?, ?, ?)")
	checkErr(err)
	_, err = stmt.Exec(content, creationDate, nickname_users, post)
	checkErr(err)
}

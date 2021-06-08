package main

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"io/ioutil"
)

// La fonction insert prend une table et une interface de données pour insérer ces données dans la BDD
func insert(table Table, value interface{}) error {
	// call example: insert(USERS, Users{0, "Alecs", "alecs@ynov.com", "Admin", "I am a Pokemon fanboy", "shorturl.at/qtxIN", ""})
	var stmt *sql.Stmt
	var res sql.Result
	var err error

	db, err := sql.Open("sqlite3", "./forum.db")
	defer db.Close()
	checkErr(err)
	// insert
	switch table {
	case USERS:
		if fmt.Sprintf("%T", value) != "main.Users" {
			return errors.New("wrong table type")
		}
		stmt, err = db.Prepare(fmt.Sprintf("%s %s %s", "INSERT INTO ", table, "(nickname, email, role, biography, profileImage, status) values(?,?,?,?,?,?)"))
		checkErr(err)

		res, err = stmt.Exec(value.(Users).nickname, value.(Users).email, value.(Users).role, value.(Users).biography, value.(Users).profileImage, value.(Users).status)
		checkErr(err)
	case CATEGORY:
		if fmt.Sprintf("%T", value) != "main.Category" {
			return errors.New("wrong table type")
		}
		stmt, err = db.Prepare(fmt.Sprintf("%s %s %s", "INSERT INTO ", table, "(name, description) values(?,?)"))
		checkErr(err)

		res, err = stmt.Exec(value.(Category).name, value.(Category).description)
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
		stmt, err = db.Prepare(fmt.Sprintf("%s %s %s", "INSERT INTO ", table, "(title, creationDate, modificationDate, deleteDate, likes, dislikes, id_users) values(?,?,?,?,?,?,?)"))
		checkErr(err)

		res, err = stmt.Exec(value.(Posts).title, value.(Posts).creationDate, value.(Posts).modificationDate, value.(Posts).deleteDate, value.(Posts).likes, value.(Posts).dislikes, value.(Posts).id_users)
		checkErr(err)
	case POSTSCAT:
		if fmt.Sprintf("%T", value) != "main.PostsCat" {
			return errors.New("wrong table type")
		}
		stmt, err = db.Prepare(fmt.Sprintf("%s %s %s", "INSERT INTO ", table, "(id_posts, id_category) values(?,?)"))
		checkErr(err)

		res, err = stmt.Exec(value.(PostsCat).id_posts, value.(PostsCat).id_category)
		checkErr(err)
	case COMMENTS:
		if fmt.Sprintf("%T", value) != "main.Comments" {
			return errors.New("wrong table type")
		}
		stmt, err = db.Prepare(fmt.Sprintf("%s %s %s", "INSERT INTO ", table, "(creationDate, modificationDate, deleteDate, likes, dislikes, id_users, id_posts) values(?,?,?,?,?,?,?)"))
		checkErr(err)

		res, err = stmt.Exec(value.(Comments).creationDate, value.(Comments).modificationDate, value.(Comments).deleteDate, value.(Comments).likes, value.(Comments).dislikes, value.(Comments).id_users, value.(Comments).id_posts)
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

	db, err := sql.Open("sqlite3", "./forum.db")
	defer db.Close()
	checkErr(err)

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

	db, err := sql.Open("sqlite3", "./forum.db")
	checkErr(err)
	defer db.Close()

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

	dropTable, err :=  ioutil.ReadFile("data/sql/drop_tables.sql")
	checkErr(err)

	createTable, err :=  ioutil.ReadFile("data/sql/create_tables.sql")
	checkErr(err)

	basicdata, err := ioutil.ReadFile("data/sql/insert_data.sql")
	checkErr(err)

	tableList := string(dropTable) + string(createTable) + string(basicdata)
	return tableList
}

func registerBDD(nickname string, email string, hashedpwd string) {
	db, err := sql.Open("sqlite3", "./forum.db")
	defer db.Close()
	checkErr(err)

	stmt, err := db.Prepare("INSERT INTO Users (nickname, email, hashedpwd) VALUES (?, ?, ?)")
	checkErr(err)
	_, err = stmt.Exec(nickname, email, hashedpwd)
	checkErr(err)

}

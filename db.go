package main

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"time"
)

type Table string

const (
	USERS Table = "Users"
	CATEGORY Table = "Category"
	USERSCAT Table = "UsersCat"
	POSTS Table = "Posts"
	POSTSCAT Table = "PostsCat"
	COMMENTS Table = "Comments"
	BADGE Table = "Badge"
	USERSBADGE Table = "UsersBadge"
)

type Users struct {
	id int
    nickname string
	email string
	role string
	biography string
	profileImage string
	status string
}

type Category struct {
    id int
	name string
	description string
}

type UsersCat struct {
    id_users int
	id_category int
}

type Posts struct {
	id int
	title string
	creationDate time.Time
	modificationDate time.Time
	deleteDate time.Time
	likes int
	dislikes int
	id_users int
}

type PostsCat struct {
    id_posts int
	id_category int
}

type Comments struct {
	id int
	creationDate time.Time
	modificationDate time.Time
	deleteDate time.Time
	likes int 
	dislikes int
	id_users int
	id_posts int
}

type Badge struct {
	id int
	name string
	image string
	description string
}

type UsersBadge struct {
    id_users int
	id_badge int
}

// call example: insert(USERS, Users{0, "Alecs", "alecs@ynov.com", "Admin", "I am a Pokemon fanboy", "shorturl.at/qtxIN", ""})
func insert(table Table, value interface{}) error {
	var stmt *sql.Stmt
	var res sql.Result
	var err error

	db, err := sql.Open("sqlite3", "./data/forum.db")
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

func update(table Table, value interface{}, id int) {
	var stmt *sql.Stmt
	var res sql.Result

	db, err := sql.Open("sqlite3", "./data/forum.db")
	defer db.Close()
	checkErr(err)

	// update
	stmt, err = db.Prepare(fmt.Sprintf("%s %s %s", "UPDATE ", table, " SET username=? WHERE id=?"))
	checkErr(err)

	res, err = stmt.Exec("astaxieupdate", id)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)
}

func query(table Table) {
	db, err := sql.Open("sqlite3", "./data/forum.db")
	defer db.Close()
	checkErr(err)

	// query
	rows, err := db.Query(fmt.Sprintf("%s %s", "SELECT * FROM ", table))
	defer rows.Close()
	checkErr(err)
	// var uid int
	// var username string
	// var department string
	// var created time.Time

	// for rows.Next() {
	// 	err = rows.Scan(&uid, &username, &department, &created)
	// 	checkErr(err)
	// 	fmt.Println(uid)
	// 	fmt.Println(username)
	// 	fmt.Println(department)
	// 	fmt.Println(created)
	// }
}

func queryItem(table Table, id int) {
	db, err := sql.Open("sqlite3", "./data/forum.db")
	defer db.Close()
	checkErr(err)

	// query
	rows, err := db.Query(fmt.Sprintf("%s %s %s %d", "SELECT * FROM ", table, "WHERE id=", id))
	defer rows.Close()
	checkErr(err)
	// var uid int
	// var username string
	// var department string
	// var created time.Time

	// for rows.Next() {
	// 	err = rows.Scan(&uid, &username, &department, &created)
	// 	checkErr(err)
	// 	fmt.Println(uid)
	// 	fmt.Println(username)
	// 	fmt.Println(department)
	// 	fmt.Println(created)
	// }
}

func delete(table Table, id int) {
	var stmt *sql.Stmt
	var res sql.Result
	var affect int64

	db, err := sql.Open("sqlite3", "./data/forum.db")
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

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
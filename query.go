package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"time"
)

type Table string

const (
	USERS      Table = "Users"
	CATEGORY   Table = "Category"
	USERSCAT   Table = "UsersCat"
	POSTS      Table = "Posts"
	POSTSCAT   Table = "PostsCat"
	COMMENTS   Table = "Comments"
	BADGE      Table = "Badge"
	USERSBADGE Table = "UsersBadge"
)

type Users struct {
	id           int
	nickname     string
	email        string
	role         string
	biography    string
	profileImage string
	status       string
	password     string
}

type Category struct {
	id          int
	name        string
	description string
}

type UsersCat struct {
	id_users    int
	id_category int
}

type Posts struct {
	id               int
	title            string
	creationDate     time.Time
	modificationDate time.Time
	deleteDate       time.Time
	likes            int
	dislikes         int
	id_users         int
}

type PostsCat struct {
	id_posts    int
	id_category int
}

type Comments struct {
	id               int
	creationDate     time.Time
	modificationDate time.Time
	deleteDate       time.Time
	likes            int
	dislikes         int
	id_users         int
	id_posts         int
}

type Badge struct {
	id          int
	name        string
	image       string
	description string
}

type UsersBadge struct {
	id_users int
	id_badge int
}

// La fonction query prend une table en paramètre pour en afficher son contenu
func query(table Table) {
	db, err := sql.Open("sqlite3", "./data/forum.db")
	defer db.Close()
	checkErr(err)

	// query
	rows, err := db.Query(fmt.Sprintf("%s %s", "SELECT * FROM ", table))
	defer rows.Close()
	checkErr(err)
}

// La fonction queryItem prend une table et un id en paramètre pour faire afficher les informations correspondantes.
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

// La fonction queryEmail prend en paramètre un email et va vérifier si il existe un utilisateur avec cet email dans la base de données
func queryEmail(email string) bool {
	database, err := sql.Open("sqlite3", "./forum.db")
	checkErr(err)
	defer database.Close()
	verif := `SELECT email FROM Users WHERE email = ?`
	err = database.QueryRow(verif, email).Scan(&email)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Print(err)
		}
		return false
	}
	return true
}

// La fonction queryUname va vérifier si le nom d'utilisateur appartient déjà à un utilisateur enregistré dans la base de données
func queryUname(username string) bool {
	database, err := sql.Open("sqlite3", "./forum.db")
	checkErr(err)
	defer database.Close()
	verif := `SELECT username FROM Users WHERE username = ?`
	err = database.QueryRow(verif, username).Scan(&username)

	// vérificateur, si QueryRow.Scan retourne une erreur, c'est qu'il n'y a pas de ligne correspondant à la recherche SELECT
	if err != nil {
		if err != sql.ErrNoRows {
			log.Print(err)
		}
		return false
	}
	return true
}

// La fonction queryLogin va vérifier si le nom d'utilisateur et l'email appartient déjà à un utilisateur enregistré dans la base de données
func queryLogin(username string, email string) bool {
	database, err := sql.Open("sqlite3", "./forum.db")
	checkErr(err)
	defer database.Close()
	verif := `SELECT username, email FROM Users WHERE username = ? AND email = ?`
	err = database.QueryRow(verif, username, email).Scan(&username, &email)

	if err != nil {
		if err != sql.ErrNoRows {
			log.Print(err)
		}
		return false
	}
	return true
}

func queryPassword(email string) string {
	db, err := sql.Open("sqlite3", "./forum.db")
	checkErr(err)
	defer db.Close()
	verif := `SELECT password FROM Users WHERE email = ?`
	err = db.QueryRow(verif, email).Scan(&email)

	if err != nil {
		if err != sql.ErrNoRows {
			log.Print(err)
		}
		return verif
	}
	return verif

}

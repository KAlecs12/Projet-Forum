package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"time"

	_ "github.com/mattn/go-sqlite3"
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
	Id           int
	Nickname     string
	Email        string
	Role         string
	Biography    string
	ProfileImage string
	Status       string
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
	Id               int
	Title            string
	Content			 string
	CreationDate     time.Time
	ModificationDate time.Time
	DeleteDate       time.Time
	Likes            int
	Dislikes         int
	Id_users         int
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
	rows, err := db.Query(fmt.Sprintf("%s %s", "SELECT * FROM ", table))
	defer rows.Close()
	checkErr(err)
}

// La fonction queryItem prend une table et un id en paramètre pour faire afficher les informations correspondantes.
func queryItem(table Table, id int) {
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
func queryId(email string) int {
	query := "SELECT id FROM Users WHERE email = \"" + email + "\""
	result, err := db.Query(query)
	checkErr(err)
	var id int
	defer result.Close()
	for result.Next() {
		err = result.Scan(&id)
		checkErr(err)
		return id
	}
	return id
}

// La fonction queryUname va vérifier si le nom d'utilisateur appartient déjà à un utilisateur enregistré dans la base de données
func queryUname(email string) string {
	query := "SELECT nickname FROM Users WHERE email = \"" + email + "\""
	result, err := db.Query(query)
	checkErr(err)
	var nickname string
	defer result.Close()
	for result.Next() {
		err = result.Scan(&nickname)
		checkErr(err)
		return nickname
	}
	return nickname
}

// La fonction queryLogin va vérifier si le nom d'utilisateur et l'email appartient déjà à un utilisateur enregistré dans la base de données
func queryLogin(username string, email string) bool {
	var err error
	verif := `SELECT nickname, email FROM Users WHERE nickname = ? OR email = ?`
	err = db.QueryRow(verif, username, email).Scan(&username, &email)

	if err != nil {
		if err != sql.ErrNoRows {
			log.Print(err)
		}
		return false
	}
	return true
}

func queryPassword(email string) string {
	query := "SELECT hashedpwd FROM Users WHERE email = \"" + email + "\""
	result, err := db.Query(query)
	checkErr(err)
	var hashedpwd string
	defer result.Close()
	for result.Next() {
		err = result.Scan(&hashedpwd)
		checkErr(err)
		return hashedpwd
	}
	return hashedpwd
}

// La fonction InitDB permet de réinitialiser la base de données, puis de recréer les tables nécessaires, avec les informations par défaut
func InitDB() {
	var err error
	db, err = sql.Open("sqlite3", "./forum.db")
	checkErr(err)

	createBDD := CreateTables()
	_, err = db.Exec(createBDD)
	checkErr(err)
}

func infosU(id int) Users {
	query := "SELECT nickname, email, role, biography, profileImage, status FROM Users WHERE id = " + strconv.Itoa(id)
	result, err := db.Query(query)
	checkErr(err)
	var nicknameBDD, emailBDD, roleBDD, biographyBDD, profileImageBDD, statusBDD interface{}
	defer result.Close()
	User := Users{}
	for result.Next() {
		err = result.Scan(&nicknameBDD, &emailBDD, &roleBDD, &biographyBDD, &profileImageBDD, &statusBDD)
		checkErr(err)
		if nicknameBDD == nil {
			User.Nickname = ""
		} else {
			User.Nickname = fmt.Sprintf("%v", nicknameBDD)
		}
		if emailBDD == nil {
			User.Email = ""
		} else {
			User.Email = fmt.Sprintf("%v", emailBDD)
		}
		if roleBDD == nil {
			User.Role = ""
		} else {
			User.Role = fmt.Sprintf("%v", roleBDD)
		}
		if biographyBDD == nil {
			User.Biography = ""
		} else {
			User.Biography = fmt.Sprintf("%v", biographyBDD)
		}
		if profileImageBDD == nil {
			User.ProfileImage = ""
		} else {
			User.ProfileImage = fmt.Sprintf("%v", profileImageBDD)
		}
		if statusBDD == nil {
			User.Status = ""
		} else {
			User.Status = fmt.Sprintf("%v", statusBDD)
		}
	}
	return User
}

func infosPost(id int) Posts {
	query := "SELECT title, content, creationDate, modificationDate, deleteDate, likes, dislikes, id_users FROM Posts WHERE id = " + strconv.Itoa(id)
	result, err := db.Query(query)
	checkErr(err)
	var title, creationDate, modificationDate, deleteDate, content, likes, dislikes, id_users interface{}
	defer result.Close()
	Post := Posts{}
	for result.Next() {
		err = result.Scan(&title, &content, &creationDate, &modificationDate, &deleteDate, &likes, &dislikes, &id_users)
		checkErr(err)
		nilTime := time.Time{}
		if title == nil {
			Post.Title = ""
		} else {
			Post.Title = fmt.Sprintf("%v", title)
		}
		if content == nil {
			Post.Content = ""
		} else {
			Post.Content = fmt.Sprintf("%v", content)
		}
		if creationDate == nil {
			Post.CreationDate = nilTime
		} else {
			Post.CreationDate = creationDate.(time.Time)
		}

		if modificationDate == nil {
			Post.ModificationDate = nilTime
		} else {
			Post.ModificationDate = modificationDate.(time.Time)
		}
		if fmt.Sprintf("%T", deleteDate) != "time.Time" {
			Post.DeleteDate = nilTime
		} else {
			Post.DeleteDate = deleteDate.(time.Time)
		}
		if likes == nil {
			Post.Likes = 0
		} else {
			Post.Likes, err = strconv.Atoi(fmt.Sprintf("%v", likes))
			checkErr(err)
		}
		if dislikes == nil {
			Post.Dislikes = 0
		} else {
			Post.Dislikes, err = strconv.Atoi(fmt.Sprintf("%v", dislikes))
			checkErr(err)
		}
		if id_users == nil {
			Post.Id_users = 0
		} else {
			Post.Id_users, err = strconv.Atoi(fmt.Sprintf("%v", id_users))
		}
	}
	return Post
}
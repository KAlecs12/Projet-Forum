package main

import "C"
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
	Id          int
	Name        string
	Description string
}

type UsersCat struct {
	id_users    int
	id_category int
}

type Posts struct {
	Id               int
	Title            string
	Content          string
	CreationDate     string
	ModificationDate string
	DeleteDate       string
	Likes            int
	Dislikes         int
	Nickname_users   string
	Category         string
	Status           string
}

type Comments struct {
	Id               int
	Content          string
	CreationDate     string
	ModificationDate string
	DeleteDate       string
	Likes            int
	Dislikes         int
	NicknameUsers    string
	IdPosts          int
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
	query := "SELECT id, nickname, email, role, biography, profileImage, status FROM Users WHERE id = " + strconv.Itoa(id)
	result, err := db.Query(query)
	checkErr(err)
	var id_User, nicknameBDD, emailBDD, roleBDD, biographyBDD, profileImageBDD, statusBDD interface{}
	defer result.Close()
	User := Users{}
	for result.Next() {
		err = result.Scan(&id_User, &nicknameBDD, &emailBDD, &roleBDD, &biographyBDD, &profileImageBDD, &statusBDD)
		checkErr(err)

		User.Id, err = strconv.Atoi(fmt.Sprintf("%v", id_User))
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
	query := "SELECT id, title, content, creationDate, modificationDate, deleteDate, likes, dislikes, Nickname_users, category, status FROM Posts WHERE id = " + strconv.Itoa(id)
	result, err := db.Query(query)
	checkErr(err)
	var id_Post, title, creationDate, modificationDate, deleteDate, content, likes, dislikes, nicknameUsers, category, status interface{}
	defer result.Close()
	Post := Posts{}
	for result.Next() {
		err = result.Scan(&id_Post, &title, &content, &creationDate, &modificationDate, &deleteDate, &likes, &dislikes, &nicknameUsers, &category, &status)
		checkErr(err)

		Post.Id, err = strconv.Atoi(fmt.Sprintf("%v", id_Post))
		Post.Status = fmt.Sprintf("%v", status)
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
			Post.CreationDate = ""
		} else {
			t := creationDate.(time.Time)
			Post.CreationDate = t.Format("2006-01-02 15:04:05")
		}
		if modificationDate == nil {
			Post.ModificationDate = ""
		} else {
			t := modificationDate.(time.Time)
			Post.ModificationDate = t.Format("2006-01-02 15:04:05")
		}
		if fmt.Sprintf("%T", deleteDate) != "time.Time" {
			Post.DeleteDate = ""
		} else {
			t := deleteDate.(time.Time)
			Post.DeleteDate = t.Format("2006-01-02 15:04:05")
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
		if nicknameUsers == nil {
			Post.Nickname_users = ""
		} else {
			Post.Nickname_users = fmt.Sprintf("%v", nicknameUsers)
		}
		if category == nil {
			Post.Category = ""
		} else {
			Post.Category = fmt.Sprintf("%v", category)
		}
	}
	return Post
}

func infosPosts() []Posts {
	var table []Posts
	query := "SELECT * FROM Posts"
	result, err := db.Query(query)
	checkErr(err)
	var id_Post, title, creationDate, modificationDate, deleteDate, content, likes, dislikes, nicknameUsers, category, status interface{}
	defer result.Close()
	Post := Posts{}
	for result.Next() {
		err = result.Scan(&id_Post, &title, &content, &creationDate, &modificationDate, &deleteDate, &likes, &dislikes, &nicknameUsers, &category, &status)
		checkErr(err)

		Post.Id, err = strconv.Atoi(fmt.Sprintf("%v", id_Post))
		Post.Status = fmt.Sprintf("%v", status)
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
			Post.CreationDate = ""
		} else {
			t := creationDate.(time.Time)
			Post.CreationDate = t.Format("2006-01-02 15:04:05")
		}
		if modificationDate == nil {
			Post.ModificationDate = ""
		} else {
			t := modificationDate.(time.Time)
			Post.ModificationDate = t.Format("2006-01-02 15:04:05")
		}
		if fmt.Sprintf("%T", deleteDate) != "time.Time" {
			Post.DeleteDate = ""
		} else {
			t := deleteDate.(time.Time)
			Post.DeleteDate = t.Format("2006-01-02 15:04:05")
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
		if nicknameUsers == nil {
			Post.Nickname_users = ""
		} else {
			Post.Nickname_users = fmt.Sprintf("%v", nicknameUsers)
		}
		if category == nil {
			Post.Category = ""
		} else {
			Post.Category = fmt.Sprintf("%v", category)
		}
		table = append(table, Post)
	}
	return table
}

func infosCat() []Category {
	var table []Category

	query := "SELECT * FROM Category"
	result, err := db.Query(query)
	checkErr(err)
	var id, name, description interface{}
	defer result.Close()
	Category := Category{}
	for result.Next() {
		err = result.Scan(&id, &name, &description)
		checkErr(err)
		Category.Id, err = strconv.Atoi(fmt.Sprintf("%v", id))
		if name == nil {
			Category.Name = ""
		} else {
			Category.Name = fmt.Sprintf("%v", name)
		}
		if description == nil {
			Category.Description = ""
		} else {
			Category.Description = fmt.Sprintf("%v", description)
		}
		table = append(table, Category)
	}
	return table
}

func infosComments() []Comments {
	var table []Comments
	query := "SELECT * FROM Comments"
	result, err := db.Query(query)
	checkErr(err)
	var id_comments, content, creationDate, modificationDate, deleteDate, likes, dislikes, nicknameUsers, idPosts interface{}
	defer result.Close()
	Comments := Comments{}
	for result.Next() {
		err = result.Scan(&id_comments, &content, &creationDate, &modificationDate, &deleteDate, &likes, &dislikes, &nicknameUsers, &idPosts)
		checkErr(err)

		Comments.Id, err = strconv.Atoi(fmt.Sprintf("%v", id_comments))
		if content == nil {
			Comments.Content = ""
		} else {
			Comments.Content = fmt.Sprintf("%v", content)
		}
		if creationDate == nil {
			Comments.CreationDate = ""
		} else {
			t := creationDate.(time.Time)
			Comments.CreationDate = t.Format("2006-01-02 15:04:05")
		}
		if modificationDate == nil {
			Comments.ModificationDate = ""
		} else {
			t := modificationDate.(time.Time)
			Comments.ModificationDate = t.Format("2006-01-02 15:04:05")
		}
		if fmt.Sprintf("%T", deleteDate) != "time.Time" {
			Comments.DeleteDate = ""
		} else {
			t := deleteDate.(time.Time)
			Comments.DeleteDate = t.Format("2006-01-02 15:04:05")
		}
		if likes == nil {
			Comments.Likes = 0
		} else {
			Comments.Likes, err = strconv.Atoi(fmt.Sprintf("%v", likes))
			checkErr(err)
		}
		if dislikes == nil {
			Comments.Dislikes = 0
		} else {
			Comments.Dislikes, err = strconv.Atoi(fmt.Sprintf("%v", dislikes))
			checkErr(err)
		}
		if nicknameUsers == nil {
			Comments.NicknameUsers = ""
		} else {
			Comments.NicknameUsers = fmt.Sprintf("%v", nicknameUsers)
		}
		if idPosts == nil {
			Comments.IdPosts = 0
		} else {
			Comments.IdPosts, err = strconv.Atoi(fmt.Sprintf("%v", idPosts))
			checkErr(err)
		}
		table = append(table, Comments)
	}
	return table
}

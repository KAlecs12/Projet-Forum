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

////////// STRUCTURES /////////////

type Table string

const (
	USERS      Table = "Users"
	CATEGORY   Table = "Category"
	USERSCAT   Table = "UsersCat"
	POSTS      Table = "Posts"
	COMMENTS   Table = "Comments"
	BADGE      Table = "Badge"
	USERSBADGE Table = "UsersBadge"
)

type Users struct {
	Id           int
	Nickname     string
	Email        string
	Squestion    string
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
	CommentCount     int
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

// La fonction InitDB permet de réinitialiser la base de données, puis de recréer les tables nécessaires, avec les informations par défaut
func InitDB() {
	var err error
	db, err = sql.Open("sqlite3", "./forum.db")
	checkErr(err)

	createBDD := CreateTables()
	_, err = db.Exec(createBDD)
	checkErr(err)
}

////////// USERS /////////////

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

// La fonction queryEmail permet de vérifier si un utilisateur possède un email correspond à celui envoyé en paramètre
func queryEmail(email string) bool {
	var err error
	verif := `SELECT email FROM Users WHERE email = ?`
	err = db.QueryRow(verif, email).Scan(&email)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Print(err)
		}
		return false
	}
	return true
}

// La fonction queryPassword permet de récupérer le mot de passe haché en fonction de l'email
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

// La fonction infosU permet de récupérer une structure de Users en fonction de l'Id de celui-ci
func infosU(id int) Users {
	query := "SELECT id, nickname, email, squestion, role, biography, profileImage, status FROM Users WHERE id = " + strconv.Itoa(id)
	result, err := db.Query(query)
	checkErr(err)
	var id_User, nicknameBDD, emailBDD, roleBDD, squestion, biographyBDD, profileImageBDD, statusBDD interface{}
	defer result.Close()
	User := Users{}
	for result.Next() {
		err = result.Scan(&id_User, &nicknameBDD, &emailBDD, &squestion, &roleBDD, &biographyBDD, &profileImageBDD, &statusBDD)
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
		if squestion == nil {
			User.Squestion = ""
		} else {
			User.Squestion = fmt.Sprintf("%v", squestion)
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

// La fonction infosU2 permet de récupérer une structure de Users en fonction de l'email de celui-ci
func infosU2(email string) Users {
	query := "SELECT id, nickname, email, squestion, role, biography, profileImage, status FROM Users WHERE email = \"" + email + "\""
	result, err := db.Query(query)
	checkErr(err)
	var id_User, nicknameBDD, emailBDD, roleBDD, squestion, biographyBDD, profileImageBDD, statusBDD interface{}
	defer result.Close()
	User := Users{}
	for result.Next() {
		err = result.Scan(&id_User, &nicknameBDD, &emailBDD, &squestion, &roleBDD, &biographyBDD, &profileImageBDD, &statusBDD)
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
		if squestion == nil {
			User.Squestion = ""
		} else {
			User.Squestion = fmt.Sprintf("%v", squestion)
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

///////// LIKES DISLIKES /////////

// LikesUpdate permet d'augmenter ou de diminuer le nombre de likes sur un post ou un comment en BDD
func LikesUpdate(id_post int, table string, ToUp bool) {
	queryLikeBDD := "SELECT likes FROM " + table + " WHERE id = \"" + strconv.Itoa(id_post) + "\""
	result, err := db.Query(queryLikeBDD)
	checkErr(err)
	var nblikes int
	defer result.Close()
	for result.Next() {
		err = result.Scan(&nblikes)
		checkErr(err)
	}
	if ToUp {
		nblikes++
	} else {
		nblikes--
	}
	query := "UPDATE " + table + " SET likes = ? WHERE id = ?"
	updatelikes, err := db.Prepare(query)
	checkErr(err)
	_, err = updatelikes.Exec(nblikes, id_post)
	checkErr(err)
}

// DislikesUpdate permet d'augmenter ou de diminuer le nombre de dislikes sur un post ou un comment en BDD
func DislikesUpdate(id_post int, table string, ToUp bool) {
	queryDislikeBDD := "SELECT dislikes FROM " + table + " WHERE id = \"" + strconv.Itoa(id_post) + "\""
	result, err := db.Query(queryDislikeBDD)
	checkErr(err)
	var nbdislikes int
	defer result.Close()
	for result.Next() {
		err = result.Scan(&nbdislikes)
		checkErr(err)
	}
	if ToUp {
		nbdislikes++
	} else {
		nbdislikes--
	}
	query := "UPDATE " + table + " SET dislikes = ? WHERE id = ?"
	updatedislikes, err := db.Prepare(query)
	checkErr(err)
	_, err = updatedislikes.Exec(nbdislikes, id_post)
	checkErr(err)
}

// Like permet de voir s'il existe une ligne en BDD possédant l'id du post et de l'utilisateur entrées en paramètres
// Si c'est le cas, on va la supprimer, si ce n'est pas le cas, alors on va la créer
func Like(id_post int, id_user int, table string) {
	liketype := "Like"
	var err error
	verif := tableSelect(table, liketype)
	err = db.QueryRow(verif, id_post, id_user).Scan(&id_post, &id_user)
	if err == sql.ErrNoRows {
		insert := tableInsert(table, liketype)
		stmt, err := db.Prepare(insert)
		checkErr(err)
		_, err = stmt.Exec(id_post, id_user)
		checkErr(err)
		LikesUpdate(id_post, table, true)
	} else if err == nil {
		delete := tableDelete(table, liketype)
		stmt, err := db.Prepare(delete)
		checkErr(err)
		_, err = stmt.Exec(id_post, id_user)
		checkErr(err)
		LikesUpdate(id_post, table, false)
	} else {
		log.Fatal(err)
	}
}

// Dislike permet, de la même manière, de vérifier si une ligne existe en BDD possédant l'id du post et de l'utilisateur entrées en paramètres
func Dislike(id_post int, id_user int, table string) {
	liketype := "Dislike"
	var err error
	verif := tableSelect(table, liketype)
	err = db.QueryRow(verif, id_post, id_user).Scan(&id_post, &id_user)
	if err == sql.ErrNoRows {
		insert := tableInsert(table, liketype)
		stmt, err := db.Prepare(insert)
		checkErr(err)
		_, err = stmt.Exec(id_post, id_user)
		checkErr(err)
		DislikesUpdate(id_post, table, true)
	} else if err == nil {
		delete := tableDelete(table, liketype)
		stmt, err := db.Prepare(delete)
		checkErr(err)
		_, err = stmt.Exec(id_post, id_user)
		checkErr(err)
		DislikesUpdate(id_post, table, false)
	} else {
		log.Fatal(err)
	}
}

// tableSelect permet de renvoyer la phrase de requête de sélection en fonction de la table et du type de like que c'est (like ou dislike)
func tableSelect(table string, liketype string) string{
	if table == "Posts" {
		if liketype == "Like" {
			return `SELECT id_post, id_user FROM LikesPosts WHERE id_post = ? AND id_user = ?`
		} else {
			return `SELECT id_post, id_user FROM DislikesPosts WHERE id_post = ? AND id_user = ?`
		}
	} else if table == "Comments" {
		if liketype == "Like" {
			return `SELECT id_post, id_user FROM LikesComments WHERE id_post = ? AND id_user = ?`
		} else {
			return `SELECT id_post, id_user FROM DislikesComments WHERE id_post = ? AND id_user = ?`
		}
	}
	return ""
}

// tableInsert permet de renvoyer la phrase de requête d'insertion en fonction de la table et du type de like que c'est (like ou dislike)
func tableInsert(table string, liketype string) string{
	if table == "Posts" {
		if liketype == "Like" {
			return "INSERT INTO LikesPosts (id_post, id_user) VALUES (?, ?)"
		} else {
			return "INSERT INTO DislikesPosts (id_post, id_user) VALUES (?, ?)"
		}
	} else if table == "Comments" {
		if liketype == "Like" {
			return "INSERT INTO LikesComments (id_post, id_user) VALUES (?, ?)"
		} else {
			return "INSERT INTO DislikesComments (id_post, id_user) VALUES (?, ?)"
		}
	}
	return ""
}

// tableDelete permet de renvoyer la phrase de requête de suppression en fonction de la table et du type de like que c'est (like ou dislike)
func tableDelete(table string, liketype string) string{
	if table == "Posts" {
		if liketype == "Like" {
			return "DELETE FROM LikesPosts WHERE id_post = ? AND id_user = ?"
		} else {
			return "DELETE FROM DislikesPosts WHERE id_post = ? AND id_user = ?"
		}
	} else if table == "Comments" {
		if liketype == "Like" {
			return "DELETE FROM LikesComments WHERE id_post = ? AND id_user = ?"
		} else {
			return "DELETE FROM DislikesComments WHERE id_post = ? AND id_user = ?"
		}
	}
	return ""
}

////////// CATEGORY /////////////

// infosCat permet de renvoyer un table de structure de Category
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

////////// SESSION /////////////

// getIdSession permet de récupérer l'Id de l'utilisateur via le cookie
func getIdSession(uuid string) int {
	queryid := "SELECT id_user FROM SessionControl WHERE uuid = \"" + uuid + "\""
	result, err := db.Query(queryid)
	checkErr(err)
	var id_user interface{}
	defer result.Close()
	for result.Next() {
		err = result.Scan(&id_user)
		checkErr(err)
		if id_user == nil {
			return 0
		} else {
			intId, err := strconv.Atoi(fmt.Sprintf("%v", id_user))
			checkErr(err)
			return intId
		}
	}
	return 0
}

////////// POSTS /////////////

// infosPost permet de récupérer une structure de Posts de l'Id correspond dans la BDD
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
			Post.CreationDate = t.Format("02 January 2006 15:04")
		}
		if modificationDate == nil {
			Post.ModificationDate = ""
		} else {
			t := modificationDate.(time.Time)
			Post.ModificationDate = t.Format("02 January 2006 15:04")
		}
		if fmt.Sprintf("%T", deleteDate) != "time.Time" {
			Post.DeleteDate = ""
		} else {
			t := deleteDate.(time.Time)
			Post.DeleteDate = t.Format("02 January 2006 15:04")
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

// infosPosts permet de récupérer tous les Posts sous la forme d'un tableau de structure Posts
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
			Post.CreationDate = t.Format("02 January 2006 15:04")
		}
		if modificationDate == nil {
			Post.ModificationDate = ""
		} else {
			t := modificationDate.(time.Time)
			Post.ModificationDate = t.Format("02 January 2006 15:04")
		}
		if fmt.Sprintf("%T", deleteDate) != "time.Time" {
			Post.DeleteDate = ""
		} else {
			t := deleteDate.(time.Time)
			Post.DeleteDate = t.Format("02 January 2006 15:04")
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
		Post.CommentCount = countComments(Post.Id)
		table = append(table, Post)
	}
	return table
}

////////// COMMENTS /////////////

// infosComments permet de récupérer tous les Comments sous la forme d'un tableau de structure Comments
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
			Comments.CreationDate = t.Format("02 January 2006 15:04")
		}
		if modificationDate == nil {
			Comments.ModificationDate = ""
		} else {
			t := modificationDate.(time.Time)
			Comments.ModificationDate = t.Format("02 January 2006 15:04")
		}
		if fmt.Sprintf("%T", deleteDate) != "time.Time" {
			Comments.DeleteDate = ""
		} else {
			t := deleteDate.(time.Time)
			Comments.DeleteDate = t.Format("02 January 2006 15:04")
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

// countComments permet de récupérer le nombre de Comments liés à un Posts
func countComments(id_post int) int {
	queryComments := "SELECT id FROM Comments WHERE id_posts = \"" + strconv.Itoa(id_post) + "\""
	result, err := db.Query(queryComments)
	checkErr(err)
	var id int
	var count int
	defer result.Close()
	for result.Next() {
		err = result.Scan(&id)
		checkErr(err)
		count++
	}
	return count
}
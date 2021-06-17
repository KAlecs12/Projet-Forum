package main

import (
	_ "github.com/mattn/go-sqlite3"
	"io/ioutil"
	"time"
)

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

// CreateSession permet de créer une session à partir de l'id de l'utilisateur connecté et de son uuid présent dans son cookie
func CreateSession(id_user int, uuid string) {
	stmt, err := db.Prepare("INSERT INTO SessionControl (uuid, id_user) VALUES (?, ?)")
	checkErr(err)
	_, err = stmt.Exec(uuid, id_user)
	checkErr(err)
}

//////////////////////// CATEGORY ////////////////////////

// CreateCat permet de créer une Category
func CreateCat(name string) {
	if name != "" {
		stmt, err := db.Prepare("INSERT INTO Category (name) VALUES (?)")
		checkErr(err)
		_, err = stmt.Exec(name)
		checkErr(err)
	}
}

// ModifiedCat permet de modifier le nom d'une Category
func ModifiedCat(name string, newname string) {
	stmt, err := db.Prepare("UPDATE Category SET name = ? WHERE name = ?")
	checkErr(err)
	_, err = stmt.Exec(newname, name)
	checkErr(err)
}

// DeleteCat permet de supprimer une Category
func DeleteCat(name string) {
	stmt, err := db.Prepare("DELETE FROM Category WHERE name = ?")
	checkErr(err)
	_, err = stmt.Exec(name)
	checkErr(err)
}

//////////////////////// POSTS ////////////////////////

// CreatePost permet de créer un Posts à partir des informations de bases
func CreatePost(nickname_users string, bio_user string, title string, content string, category string) {
	creationDate := time.Now()
	stmt, err := db.Prepare("INSERT INTO Posts (title, content, creationDate, nickname_users, bio_users, category, status, likes, dislikes) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)")
	checkErr(err)
	_, err = stmt.Exec(title, content, creationDate, nickname_users, bio_user, category, "Actif", 0, 0)
	checkErr(err)
}

// ModifiedPost permet de modifier le contenu d'un Posts
func ModifiedPost(id_post int, content string) {
	modificationDate := time.Now()
	stmt, err := db.Prepare("UPDATE Posts SET content = ?, modificationDate = ? WHERE id = ?")
	checkErr(err)
	_, err = stmt.Exec(content, modificationDate, id_post)
	checkErr(err)
}

// DeletePost permet de supprimer un Posts de l'affichage, mais pas de la base de données
func DeletePost(id_post int) {
	stmt, err := db.Prepare("UPDATE Posts SET Status = ? WHERE id = ?")
	checkErr(err)
	_, err = stmt.Exec("Supprimé", id_post)
	checkErr(err)
}

// CreateComment permet de créer un Comments à partir des informations entrées en paramètre
func CreateComment(nickname_users string, biouser string, content string, post int) {
	creationDate := time.Now()
	stmt, err := db.Prepare("INSERT INTO Comments (content, creationDate, id_users, bio_users, id_posts ) VALUES (?, ?, ?, ?, ?)")
	checkErr(err)
	_, err = stmt.Exec(content, creationDate, nickname_users, biouser, post)
	checkErr(err)
}

////////////////////////  ACCOUNT  ////////////////////////

// Modifypseudo permet de changer le nickname d'un utilisateur dans la BDD
func Modifypseudo(newpseudo string, pseudo string) {
	stmt, err := db.Prepare("UPDATE Users SET nickname = ? WHERE nickname = ?")
	checkErr(err)
	_, err = stmt.Exec(newpseudo, pseudo)
	checkErr(err)
}

// Modifyemail permet de changer l'email d'un utilisateur dans la BDD
func Modifyemail(newemail string, email string) {
	stmt, err := db.Prepare("UPDATE Users SET email = ? WHERE email = ?")
	checkErr(err)
	_, err = stmt.Exec(newemail, email)
	checkErr(err)
}
func Modifybio(bio string, email string) {
	stmt, err := db.Prepare("UPDATE Users SET biography = ? WHERE email = ?")
	checkErr(err)
	_, err = stmt.Exec(bio, email)
	checkErr(err)
}

// registerBDD permet de créer un utilisateur en BDD à partir des informations récupérées sur la page
func registerBDD(nickname string, email string, hashedpwd string, question string) {
	stmt, err := db.Prepare("INSERT INTO Users (nickname, email, hashedpwd, squestion, role, status) VALUES (?, ?, ?, ?, ?, ?)")
	checkErr(err)
	_, err = stmt.Exec(nickname, email, hashedpwd, question, "User", "Actif")
	checkErr(err)
}

// ChangepwdBDD prend un mot de passe haché et un email pour changer le mot de passe d'un utilisateur
func ChangepwdBDD(hashedpwd string, email string) {
	stmt, err := db.Prepare("UPDATE Users SET hashedpwd = ? WHERE email = ?")
	checkErr(err)
	_, err = stmt.Exec(hashedpwd, email)
	checkErr(err)
}

func DeleteAcc(id int) {
	stmt, err := db.Prepare("DELETE FROM Users WHERE id = ?")
	checkErr(err)
	_, err = stmt.Exec(id)
	checkErr(err)
}

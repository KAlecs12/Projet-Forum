package BackToBDD

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"strconv"
)

// La fonction insert prend une table et une interface de données pour insérer ces données dans la BDD
func insert(table Table, value interface{}) error {
	// call example: insert(USERS, Users{0, "Alecs", "alecs@ynov.com", "Admin", "I am a Pokemon fanboy", "shorturl.at/qtxIN", ""})
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

// Encore à modifier, mais pour l'instant permet de changer un username en un autre, en prenant une table spécifique
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

// La fonction delete prend une table et un id en paramètre pour retirer les informations correspondantes.
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

// La fonction CreateTables permet de réinitialiser les tables en les supprimant et en les recréant, seules les tables précisées dans le code sont touchées
func CreateTables() string {
	tableList := `
	DROP TABLE if EXISTS Users;
	DROP TABLE if EXISTS Category;
	DROP TABLE if EXISTS Comments;
	DROP TABLE if EXISTS Posts;
	DROP TABLE if EXISTS Cookie;
	DROP TABLE if EXISTS Badge;
	DROP TABLE if EXISTS Posseder;
	DROP TABLE if EXISTS Modifier;
	
	CREATE TABLE if NOT EXISTS Users(
	ID_User INTEGER PRIMARY KEY autoincrement, 
	Nickname TEXT, 
	Email TEXT, 
	Password TEXT, 
	Biography TEXT, 
	profileImage BLOB, 
	Status TEXT
	);
	
	CREATE TABLE if NOT EXISTS Category (
	ID_Category INTEGER PRIMARY KEY autoincrement,
	categoryName TEXT,
	categoryDescription TEXT
	);
	
	CREATE TABLE if NOT EXISTS Posts (
	ID_Post INTEGER PRIMARY KEY autoincrement,
	Title TEXT,
	postContent TEXT,
	creationDate DATE,
	modificationDate DATE,
	deleteDate DATE,
	likesCounter INTEGER,
	dislikesCounter INTEGER,
	ID_User INTEGER,
	ID_Category INTEGER,
	FOREIGN KEY(ID_User) REFERENCES Users(ID_User),
	FOREIGN KEY(ID_Category) REFERENCES Category(ID_Category)
	);
	
	CREATE TABLE if NOT EXISTS Comments(
	ID_Comment INTEGER PRIMARY KEY autoincrement,
	commentContent TEXT,
	creationDate DATE,
	modificationDate DATE,
	deleteDate DATE,
	likesCounter INTEGER,
	dislikesCounter INTEGER,
	ID_User INTEGER,
	ID_Post INTEGER,
	FOREIGN KEY(ID_User) REFERENCES Users(ID_User),
	FOREIGN KEY(ID_Post) REFERENCES Posts(ID_Post)
	);
	
	CREATE TABLE if NOT EXISTS Badge(
	ID_Badge INTEGER PRIMARY KEY autoincrement,
	badgeName TEXT,
	badgeImage BLOB,
	badgeDescription TEXT
	);
	
	CREATE TABLE if NOT EXISTS Cookie(
	ID_Cookie INTEGER PRIMARY KEY autoincrement,
	creationDate DATE,
	ID_User INTEGER,
	FOREIGN KEY(ID_User) REFERENCES Users(ID_User)
	);
	
	CREATE TABLE if NOT EXISTS Modifier(
	ID_User INTEGER,
	ID_Category INTEGER,
	FOREIGN KEY(ID_User) REFERENCES Users(ID_User),
	FOREIGN KEY(ID_Category) REFERENCES Category(ID_Category),
	PRIMARY KEY(ID_User, ID_Category)
	);
	
	CREATE TABLE if NOT EXISTS Posseder(
	ID_Badge INTEGER,
	ID_User INTEGER,
	FOREIGN KEY(ID_Badge) REFERENCES Badge(ID_Badge),
	FOREIGN KEY(ID_User) REFERENCES Users(ID_User),
	PRIMARY KEY(ID_Badge, ID_User)
	);
	`
	return tableList
}

// ! Attention, je ne certifie pas qu'elle fonctionne
// La fonction InsertUtoU correspond à Insert User to Users, donc ajouter la configuration d'un User à la table Users
func InsertUtoU(Nickname string, Email string, Password string, Biography string, Status string) string {
	User := `
	INSERT INTO Users
	(Nickname, Email, Password, Biography, Status) 
	VALUES(`
	User += "\"" + Nickname + "\" ,\"" + Email + "\" ,\"" + Password + "\" ,\"" + Biography + "\" ,\"" + Status + "\");"
	return User
}

// ! Attention, je ne certifie pas qu'elle fonctionne
// La fonction InsertCatToCat correspond à Insert Category to Category, donc ajouter la configuration d'une Catégorie à la table Category
func InsertCatToCat(categoryName string, categoryDescription string) string {
	Category := `
	INSERT INTO Category
	(categoryName, categoryDescription) 
	VALUES(`
	Category += "\"" + categoryName + "\" ,\"" + categoryDescription + "\");"
	return Category
}

// ! Attention, je ne certifie pas qu'elle fonctionne
// La fonction InsertPtoP correspond à Insert Post to Post, donc ajouter la configuration d'un Post à la table Post
func InsertPtoP(Title string, postContent string, creationDate int, modificationDate int, deleteDate int, likesCounter int, dislikesCounter int, ID_User int, ID_Category int) string {
	Post := `
	INSERT INTO Posts
	(Title, postContent, creationDate, modificationDate, deleteDate, likesCounter, dislikesCounter, ID_User, ID_Category)
	VALUES(`
	Post += "\"" + Title + "\" ,\"" + postContent + "\" , " + strconv.Itoa(creationDate) + ", " + strconv.Itoa(modificationDate) + ", " + strconv.Itoa(deleteDate) + " ,\"" + strconv.Itoa(likesCounter) + "\" ,\"" + strconv.Itoa(dislikesCounter) + "\" ,\"" + strconv.Itoa(ID_User) + "\" ,\"" + strconv.Itoa(ID_Category) + "\");"
	return Post
}

// ! Attention, je ne certifie pas qu'elle fonctionne
// La fonction InsertComToCom correspond à Insert Comment to Comments, donc ajouter la configuration d'un Comment à la table Comments
func InsertComToCom(commentContent string, creationDate int, modificationDate int, deleteDate int, likesCounter int, dislikesCounter int, ID_User int, ID_Post int) string {
	Comment := `
	INSERT INTO Comments
	( commentContent, creationDate, modificationDate, deleteDate, likesCounter, dislikesCounter, ID_User, ID_Post)
	VALUES(`
	Comment += "\"" + commentContent + "\" , " + strconv.Itoa(creationDate) + ", " + strconv.Itoa(modificationDate) + ", " + strconv.Itoa(deleteDate) + ", " + strconv.Itoa(likesCounter) + "\" ,\"" + strconv.Itoa(dislikesCounter) + "\" ,\"" + strconv.Itoa(ID_User) + "\" ,\"" + strconv.Itoa(ID_Post) + "\");"
	return Comment
}

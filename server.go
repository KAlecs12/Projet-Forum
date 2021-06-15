package main

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type popup struct {
	Title   string
	Content string
	Need    string
}

type Homecontent struct {
	Users    Users
	Infos    []Posts
	Category []Category
}

type allPost struct {
	User    Users
	Posts   Posts
	Comment []Comments
}

var db *sql.DB
var tpl *template.Template
var cookie *http.Cookie

var idpost int

func init() {
	tpl = template.Must(template.ParseGlob("tmpl/*.html"))
}

func main() {
	InitDB()
	fs := http.FileServer(http.Dir("CSS"))
	http.Handle("/CSS/", http.StripPrefix("/CSS/", fs))

	http.HandleFunc("/", homehandler)
	http.HandleFunc("/signin", signinhandler)
	http.HandleFunc("/login", loginhandler)
	http.HandleFunc("/account", accounthandler)
	http.HandleFunc("/post", posthandler)
	http.HandleFunc("/postcreation", postcreation)
	http.HandleFunc("/modifcat", modifcat)

	http.HandleFunc("/logout", logout)
	http.HandleFunc("/like", likehandler)
	http.HandleFunc("/dislike", dislikehandler)


	http.Handle("/static/",
		http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	fmt.Printf("Starting server at port 8080\n")
	err := http.ListenAndServeTLS(":8080", "https-server.crt", "https-server.key", nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}

func homehandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		//http.Error(w, "404 not found.", http.StatusNotFound)

		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		t, err := template.ParseFiles("./static/errorpage.html")

		if err != nil {
			fmt.Fprint(w, "Unable to load page.")
			log.Fatal(err)
		}

		content := ""

		err = t.Execute(w, content)
		if err != nil {
			fmt.Fprint(w, "Unable to load page.")
			log.Fatal(err)
		}

		return
	}

	if r.Method == "POST" {
		LoginToBDD(w, r)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// sur la page d'accueil, on récupère le template index.html

	t, err := template.ParseFiles("./static/index.html", "./tmpl/header.html", "./tmpl/footer.html")

	if err != nil {
		fmt.Fprint(w, "Unable to load page.")
		log.Fatal(err)
	}

	id := getUserSession(w, r)

	Post := Homecontent{Users: infosU(id), Infos: infosPosts(), Category: infosCat()}

	err = t.Execute(w, Post)
	if err != nil {
		fmt.Fprint(w, "Unable to load page.")
		log.Fatal(err)
	}

}

func signinhandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/signin" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	if r.Method == "POST" {
		registerhandler(w, r)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// sur la page d'accueil, on récupère le template index.html
	t, err := template.ParseFiles("./static/signin.html", "./tmpl/footer.html")
	if err != nil {
		fmt.Fprint(w, "Unable to load page.")
		log.Fatal(err)
	}
	content := ""
	err = t.Execute(w, content)
	if err != nil {
		fmt.Fprint(w, "Unable to load page.")
		log.Fatal(err)
	}

}

func loginhandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/login" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	if r.Method == "POST" {
		LoginToBDD(w, r)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// sur la page d'accueil, on récupère le template index.html
	t, err := template.ParseFiles("./static/login.html", "./tmpl/footer.html")
	if err != nil {
		fmt.Fprint(w, "Unable to load page.")
		log.Fatal(err)
	}
	content := ""
	err = t.Execute(w, content)
	if err != nil {
		fmt.Fprint(w, "Unable to load page.")
		log.Fatal(err)
	}

}

func accounthandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/account" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method == "POST" {

		if r.FormValue("pseudo") != "" {
			pseudoModifBDD(w, r)
		}
		if r.FormValue("email") != "" {
			emailModifBDD(w, r)
		}

		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// sur la page d'accueil, on récupère le template index.html
	t, err := template.ParseFiles("./static/account.html", "./tmpl/footer.html")
	if err != nil {
		fmt.Fprint(w, "Unable to load page.")
		log.Fatal(err)
	}

	id := getUserSession(w, r)

	Users := infosU(id)
	err = t.Execute(w, Users)
	if err != nil {
		fmt.Fprint(w, "Unable to load page.")
		log.Fatal(err)
	}

}

type ModifCat struct {
	Users    Users
	Category []Category
}

func modifcat(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/modifcat" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	if r.Method == "POST" {

		if r.FormValue("cat") != "" {
			catToBDD(w, r)
		} else if r.FormValue("newname") != "" {
			catModifBDD(w, r)
		} else {
			catdelBDD(w, r)
		}

		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// sur la page d'accueil, on récupère le template index.html
	t, err := template.ParseFiles("./static/modifcat.html", "./tmpl/footer.html")
	if err != nil {
		fmt.Fprint(w, "Unable to load page.")
		log.Fatal(err)
	}

	id := getUserSession(w, r)

	err = t.Execute(w, ModifCat{Users: infosU(id), Category: infosCat()})
	if err != nil {
		fmt.Fprint(w, "Unable to load page.")
		log.Fatal(err)
	}

}

////////////////      INFOS POSTS       ////////////////

func posthandler(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	idpost := r.Form.Get("id")
	idpostint, err := strconv.Atoi(idpost)

	if r.URL.Path != "/post" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method == "POST" {
		commentToBDD(w, r, idpostint)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// sur la page d'accueil, on récupère le template index.html
	t, err := template.ParseFiles("./static/post.html", "./tmpl/header.html", "./tmpl/footer.html")
	if err != nil {
		fmt.Fprint(w, "Unable to load page.")
		log.Fatal(err)
	}

	id := getUserSession(w, r)

	fmt.Println(infosU(id))

	err = t.Execute(w, allPost{User: infosU(id), Posts: infosPost(idpostint), Comment: infosComments()})
	if err != nil {
		fmt.Fprint(w, "Unable to load page.")
		log.Fatal(err)
	}

}

func postcreation(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/postcreation" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	if r.Method == "POST" {
		postToBDD(w, r)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// sur la page d'accueil, on récupère le template index.html
	t, err := template.ParseFiles("./static/createpost.html", "./tmpl/footer.html")
	if err != nil {
		fmt.Fprint(w, "Unable to load page.")
		log.Fatal(err)
	}
	content := infosCat()
	err = t.Execute(w, content)
	if err != nil {
		fmt.Fprint(w, "Unable to load page.")
		log.Fatal(err)
	}

}

///////////      FUNCTION CONNEXION       ///////////////

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func registerhandler(w http.ResponseWriter, r *http.Request) {

	// TmpPopup := popup{Title: "Déjà existant", Content: "Nom ou Email déjà existant en base de données", Need: "false" }

	nickname := r.FormValue("pseudo")
	email := r.FormValue("email")
	password := r.FormValue("password")

	hashedpwd, err := HashPassword(password)
	if err != nil {
		fmt.Fprint(w, "Unable to hash password.")
		log.Fatal(err)
	}

	check := queryLogin(nickname, email)
	if check {
		log.Println("oui")
		//w.Header().Set("Content-Type", "text/html; charset=utf-8")
		//t, err := template.ParseFiles("./static/signin.html")
		//UnableLoad(w, err)

	} else {
		registerBDD(nickname, email, hashedpwd)

		http.Redirect(w, r, "/login", 302)
	}
}

func LoginToBDD(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")

	id := getUserSession(w, r)

	if CheckPasswordHash(password, queryPassword(email)) {

		key := uuid.NewString()

		// A la place de value, je veux qu'un UUID soit créer et mit
		cookie = &http.Cookie{Name: "id", Value: key, MaxAge: 1800}
		http.SetCookie(w, cookie)

		// Une fois que le cookie est créer, et qu'il est envoyé, je veux récupérer l'ID de l'utilisateur
		id = queryId(email)

		// Pour ensuite le stocker avec le UUID dans une table session, les deux donc ensemble
		CreateSession(id, key)

		http.Redirect(w, r, "/", 302)
		// Maintenant que ça sera créer, on pourra le récupérer où on veut, en sachant que la valeur de Value
		// Sera notre UUID, et donc que l'utilisateur aura toujours dans son cookie une des variables de recherche
	} else {
		log.Println("Le mdp est incorrect")
	}
}

func logout(w http.ResponseWriter, r *http.Request) {
	cookie = &http.Cookie{
		Name:   "id",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/", 302)
}

func postToBDD(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	content := r.FormValue("content")
	category := r.FormValue("category")

	id := getUserSession(w, r)

	nickname := infosU(id)

	CreatePost(nickname.Nickname, title, content, category)

	http.Redirect(w, r, "/", 302)
}

func pseudoModifBDD(w http.ResponseWriter, r *http.Request) {

	newpseudo := r.FormValue("pseudo")

	id := getUserSession(w, r)
	user := infosU(id)

	Modifypseudo(newpseudo, user.Nickname)

	http.Redirect(w, r, "/account", 302)
}

func emailModifBDD(w http.ResponseWriter, r *http.Request) {

	newemail := r.FormValue("email")

	id := getUserSession(w, r)
	user := infosU(id)

	Modifyemail(newemail, user.Email)

	http.Redirect(w, r, "/account", 302)
}

func catToBDD(w http.ResponseWriter, r *http.Request) {

	cat := r.FormValue("cat")

	CreateCat(cat)

	http.Redirect(w, r, "/", 302)
}

func catModifBDD(w http.ResponseWriter, r *http.Request) {

	cat := r.FormValue("name")
	newcat := r.FormValue("newname")

	ModifiedCat(cat, newcat)

	http.Redirect(w, r, "/", 302)
}

func catdelBDD(w http.ResponseWriter, r *http.Request) {

	cat := r.FormValue("delete")

	DeleteCat(cat)

	http.Redirect(w, r, "/", 302)
}

func commentToBDD(w http.ResponseWriter, r *http.Request, idpost int) {

	id := getUserSession(w, r)

	content := r.FormValue("content")
	nickname := infosU(id)

	CreateComment(nickname.Nickname, content, idpost)

	http.Redirect(w, r, "/", 302)
}

func getUserSession(w http.ResponseWriter, r *http.Request) int {
	cookie, err := r.Cookie("id")
	if err != nil || cookie == nil {
		return 0
	} else {
		uuid := cookie.Value
		return getIdSession(uuid)
	}
}

func likehandler(w http.ResponseWriter, r *http.Request) {

	id := getUserSession(w, r)

	if r.Method == "POST" {
		Like(1, id, "Posts")
		http.Redirect(w, r, "/post?id=id_post", 302) // Ici pour la redirection, c'est en gros la page du post
	}
}

func dislikehandler(w http.ResponseWriter, r *http.Request) {

	id := getUserSession(w, r)

	if r.Method == "POST" {
		Dislike(1, id, "Posts")
		http.Redirect(w, r, "/post?id=1", 302) // Ici pour la redirection, c'est en gros la page du post
	}
}

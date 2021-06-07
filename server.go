package main

import (
	"fmt"
	//"golang.org/x/crypto/bcrypt"
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("tmpl/*.html"))
}

func main() {
	fs := http.FileServer(http.Dir("CSS"))
	http.Handle("/CSS/", http.StripPrefix("/CSS/", fs))

	http.HandleFunc("/", homehandler)
	http.HandleFunc("/signin", signinhandler)
	http.HandleFunc("/login", loginhandler)
	http.HandleFunc("/account", accounthandler)
	http.HandleFunc("/post", posthandler)
	http.HandleFunc("/postcreation", postcreation)

	http.Handle("/static/",
		http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}

func homehandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// sur la page d'accueil, on récupère le template index.html
	t, err := template.ParseFiles("./static/index.html", "./tmpl/header.html")

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

	//username := "testeur n1"
	//
	//C := http.Cookie{
	//	Name:  "username",
	//	Value: username,
	//}
	//
	//// cookie : username=testeur n1
	//r.AddCookie(&C)
	//
	//// r.Cookie("username") --> retourne une erreur ou non

}

func signinhandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/signin" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// sur la page d'accueil, on récupère le template index.html
	t, err := template.ParseFiles("./static/signin.html")
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

	//nickname := r.FormValue("pseudo")
	//email := r.FormValue("email")
	//password := r.FormValue("password")

}

func loginhandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/login" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// sur la page d'accueil, on récupère le template index.html
	t, err := template.ParseFiles("./static/login.html")
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

	//email := r.FormValue("email")
	//password := r.FormValue("password")
	//
	//if CheckPasswordHash(password, queryPassword(email)) {
	//	// Si c'est bon, on lui associe le cookie
	//} else {
	//	// Si c'est pas bon, on lui renvoie une erreur
	//}

}

func accounthandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/account" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// sur la page d'accueil, on récupère le template index.html
	t, err := template.ParseFiles("./static/account.html")
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

func posthandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/post" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// sur la page d'accueil, on récupère le template index.html
	t, err := template.ParseFiles("./static/post.html")
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

func postcreation(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/postcreation" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// sur la page d'accueil, on récupère le template index.html
	t, err := template.ParseFiles("./static/createpost.html")
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

//func HashPassword(password string) (string, error) {
//	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
//	return string(bytes), err
//}
//
//func CheckPasswordHash(password, hash string) bool {
//	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
//	return err == nil
//}

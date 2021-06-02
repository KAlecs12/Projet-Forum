package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

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
	t, err := template.ParseFiles("./static/index.html")
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

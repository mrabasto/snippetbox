package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

const (
	Html  = "./ui/html"
	Pages = "./ui/html/pages"
)

func html(pageName string) string {
	return fmt.Sprintf("%s/%s", Html, pageName)
}
func page(pageName string) string {
	return fmt.Sprintf("%s/%s", Pages, pageName)
}

func Home(w http.ResponseWriter, r *http.Request) {
	files := []string{
		html("base.html"),
		page("home.html"),
	}

	tmpl, err := template.ParseFiles(files...)

	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, "base", nil)

	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func SnippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Creating a snippet"))
}

func SnippetView(w http.ResponseWriter, r *http.Request) {
	id, error := strconv.Atoi(r.PathValue("id"))

	if error != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "You are viewing snippet #%d", id)
}

func SnippetCreatePost(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Saved a new snippet"))
}

package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

const (
	Html     = "./ui/html"
	Pages    = "./ui/html/pages"
	Partials = "./ui/html/partials"
)

func html(pageName string) string {
	return fmt.Sprintf("%s/%s", Html, pageName)
}

func page(pageName string) string {
	return fmt.Sprintf("%s/%s", Pages, pageName)
}

func partials(pageName string) string {
	return fmt.Sprintf("%s/%s", Partials, pageName)
}

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	files := []string{
		html("base.html"),
		page("home.html"),
		partials("nav.html"),
	}

	tmpl, err := template.ParseFiles(files...)

	if err != nil {
		app.serverError(w, r, err)
		return
	}

	err = tmpl.ExecuteTemplate(w, "base", nil)

	if err != nil {
		app.serverError(w, r, err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Creating a snippet"))
}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	id, error := strconv.Atoi(r.PathValue("id"))

	if error != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "You are viewing snippet #%d", id)
}

func (app *application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Saved a new snippet"))
}

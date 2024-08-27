package main

import (
    "fmt"
    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
    "html/template"
    "log"
    "net/http"
    "path/filepath"
)

func executeTemplate(w http.ResponseWriter, filepath string) {
    w.Header().Set("Content-Type", "text/html; charset=utf-8")
    tpl, err := template.ParseFiles(filepath)
    if err != nil {
        log.Printf("parsing template: %v", err)
        http.Error(w, "There was an error parsing the template.", http.StatusInternalServerError)
        return
    }
    err = tpl.Execute(w, nil)
    if err != nil {
        log.Printf("executing template: %v", err)
        http.Error(w, "There was an error executing the template.", http.StatusInternalServerError)
        return
    }
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
    tplPath := filepath.Join("templates", "home.gohtml")
    executeTemplate(w, tplPath)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
    tplPath := filepath.Join("templates", "contact.gohtml")
    executeTemplate(w, tplPath)
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
    tplPath := filepath.Join("templates", "faq.gohtml")
    executeTemplate(w, tplPath)
}

func galleryHandler(w http.ResponseWriter, r *http.Request) {
    galleryID := chi.URLParam(r, "id")
    w.Header().Set("Content-Type", "text/html; charset=utf-8")
    fmt.Fprintf(w, "<h1>The gallery id is: %s</h1>", galleryID)
}

func main() {
    r := chi.NewRouter()
    r.Get("/", homeHandler)
    r.Get("/contact", contactHandler)
    r.Get("/faq", faqHandler)

    r.Group(func(r chi.Router) {
        r.Use(middleware.Logger)
        r.Get("/galleries/{id}", galleryHandler)
    })

    r.NotFound(func(w http.ResponseWriter, r *http.Request) {
        http.Error(w, "Page not found", http.StatusNotFound)
    })
    fmt.Println("Starting the server on :3000...")
    http.ListenAndServe(":3000", r)
}

package main

import (
    "fmt"
    "github.com/Spartan09/lenslocked/controllers"
    "github.com/Spartan09/lenslocked/views"
    "github.com/go-chi/chi/v5"
    "net/http"
    "path/filepath"
)

func main() {
    r := chi.NewRouter()

    tpl, err := views.Parse(filepath.Join("templates", "home.gohtml"))
    if err != nil {
        panic(err)
    }
    r.Get("/", controllers.StaticHandler(tpl))

    tpl, err = views.Parse(filepath.Join("templates", "contact.gohtml"))
    if err != nil {
        panic(err)
    }
    r.Get("/contact", controllers.StaticHandler(tpl))

    tpl, err = views.Parse(filepath.Join("templates", "faq.gohtml"))
    if err != nil {
        panic(err)
    }
    r.Get("/faq", controllers.StaticHandler(tpl))

    r.NotFound(func(w http.ResponseWriter, r *http.Request) {
        http.Error(w, "Page not found", http.StatusNotFound)
    })
    fmt.Println("Starting the server on :3000...")
    http.ListenAndServe(":3000", r)
}

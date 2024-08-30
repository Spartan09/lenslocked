package main

import (
    "fmt"
    "github.com/Spartan09/lenslocked/controllers"
    "github.com/Spartan09/lenslocked/templates"
    "github.com/Spartan09/lenslocked/views"
    "github.com/go-chi/chi/v5"
    "net/http"
)

func main() {
    r := chi.NewRouter()

    r.Get("/", controllers.StaticHandler(views.Must(
        views.ParseFS(templates.FS, "home.gohtml", "tailwind.gohtml"))))
    r.Get("/contact", controllers.StaticHandler(views.Must(
        views.ParseFS(templates.FS, "contact.gohtml", "tailwind.gohtml"))))
    r.Get("/faq", controllers.FAQ(
        views.Must(views.ParseFS(templates.FS, "faq.gohtml", "tailwind.gohtml"))))

    var usersC controllers.Users
    usersC.Templates.New = views.Must(views.ParseFS(templates.FS, "signup.gohtml", "tailwind.gohtml"))
    r.Get("/signup", usersC.New)

    r.NotFound(func(w http.ResponseWriter, r *http.Request) {
        http.Error(w, "Page not found", http.StatusNotFound)
    })
    fmt.Println("Starting the server on :3000...")
    http.ListenAndServe(":3000", r)
}

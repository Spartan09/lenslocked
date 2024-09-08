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

    cfg := models.DefaultPostgresConfig()
    db, err := models.Open(cfg)
    if err != nil {
        panic(err)
    }

    defer db.Close()
    userService := models.UserService{
        DB: db,
    }
    usersC := controllers.Users{
        UserService: &userService,
    }
    usersC.Templates.New = views.Must(views.ParseFS(templates.FS, "signup.gohtml", "tailwind.gohtml"))
    r.Get("/users/new", usersC.New)
    r.Post("/users", usersC.Create)

    r.NotFound(func(w http.ResponseWriter, r *http.Request) {
        http.Error(w, "Page not found", http.StatusNotFound)
    })
    fmt.Println("Starting the server on :3000...")
    http.ListenAndServe(":3000", r)
}

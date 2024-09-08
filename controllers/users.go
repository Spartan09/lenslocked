package controllers

import (
    "fmt"
    "github.com/Spartan09/lenslocked/models"
    "net/http"
)

type Users struct {
    Templates struct {
        New    Template
        SignIn Template
    }
    UserService *models.UserService
}

func (u Users) New(w http.ResponseWriter, r *http.Request) {
    var data struct {
        Email string
    }
    data.Email = r.FormValue("email")
    u.Templates.New.Execute(w, data)
}

func (u Users) Create(w http.ResponseWriter, r *http.Request) {
    email := r.FormValue("email")
    password := r.FormValue("password")
    user, err := u.UserService.Create(email, password)
    if err != nil {
        fmt.Println(err)
        http.Error(w, "Something went wrong.", http.StatusInternalServerError)
        return
    }
    fmt.Fprintf(w, "User created: %+v", user)
}

func (u Users) SignIn(w http.ResponseWriter, r *http.Request) {
    var data struct {
        Email string
    }
    data.Email = r.FormValue("email")
    u.Templates.SignIn.Execute(w, data)
}

func (u Users) ProcessSignIn(w http.ResponseWriter, r *http.Request) {
    var data struct {
        Email    string
        Password string
    }

    data.Email = r.FormValue("email")
    data.Password = r.FormValue("password")

    user, err := u.UserService.Authenticate(data.Email, data.Password)
    if err != nil {
        fmt.Println(err)
        http.Error(w, "Something went wrong.", http.StatusInternalServerError)
        return
    }

    cookie := http.Cookie{
        Name:     "email",
        Value:    user.Email,
        Path:     "/",
        HttpOnly: true,
    }
    http.SetCookie(w, &cookie)

    fmt.Fprintf(w, "User authenticated: %+v", user)
}

func (u Users) CurrentUser(w http.ResponseWriter, r *http.Request) {
    email, err := r.Cookie("email")
    if err != nil {
        fmt.Fprint(w, "The email cookie could not be read.")
        return
    }
    fmt.Fprintf(w, "Email cookie: %s\n", email.Value)
    fmt.Fprintf(w, "Headers: %+v\n", r.Header)
}

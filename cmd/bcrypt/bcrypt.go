package main

import (
    "fmt"
    "golang.org/x/crypto/bcrypt"
    "os"
)

func main() {
    switch os.Args[1] {
    case "hash":
        hash(os.Args[2])
    case "compare":
        compare(os.Args[2], os.Args[3])
    default:
        fmt.Printf("Invalid command: %v\n", os.Args[1])
    }
}

func hash(password string) {
    hashedBytes, err := bcrypt.GenerateFromPassword(
        []byte(password), bcrypt.DefaultCost)
    if err != nil {
        fmt.Printf("error handling: %v\n", err)
        return
    }
    hash := string(hashedBytes)
    fmt.Print(hash)
}

func compare(password, hash string) {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    if err != nil {
        fmt.Printf("Passsword is invalid: %v\n", password)
        return
    }
    fmt.Println("Password is correct!")
}

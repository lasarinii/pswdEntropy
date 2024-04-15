package main

import (
	"errors"
	"log"
	"net/http"
	"os"
    "text/template"
    validator "github.com/wagslane/go-password-validator"
)

type Entropy struct {
    Password string
    Value float64
    Message string
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
    log.Println(r.Method, r.RequestURI)

    t, err := template.ParseFiles("./index.tmpl")
    if err != nil {
        log.Println("Error parsing template:", err)
        os.Exit(1)
    }

    pswd := r.PathValue("password")
    if pswd == "" {
        pswd = "password"
    }

    err = validator.Validate(pswd, 60)

    var msg string
    if err != nil {
        msg = err.Error()
    } else {
        msg = "Suficient entropy"
    }

    result := Entropy{
        Password: pswd,
        Value: validator.GetEntropy(pswd),
        Message: msg,
    }

    w.Header().Add("Content-Type", "text/html")

    t.ExecuteTemplate(w, "index", map[string]interface{}{
        "Content": result,
    })
}

func main() {
    log.Println("Initializing server...")

    s := http.NewServeMux()

    s.HandleFunc("GET /", handleIndex)
    s.HandleFunc("GET /{password}", handleIndex)

    err := http.ListenAndServe(":8000", s); if err != nil {
        if errors.Is(err, http.ErrServerClosed) {
            log.Println("Server closed.")
        } else {
            log.Println("Error starting server:", err)
            os.Exit(1)
        }
    }
}

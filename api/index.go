package handler

import(
    "log"
    "errors"
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

func Handler(w http.ResponseWriter, r *http.Request) {
    t, err := template.ParseFiles("../index.tmpl")
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

func Main() {
    log.Println("Initializing server...")

    s := http.NewServeMux()

    s.HandleFunc("GET /", Handler)
    s.HandleFunc("GET /{password}", Handler)

    err := http.ListenAndServe(":8000", s); if err != nil {
        if errors.Is(err, http.ErrServerClosed) {
            log.Println("Server closed.")
        } else {
            log.Println("Error starting server:", err)
            os.Exit(1)
        }
    }
}

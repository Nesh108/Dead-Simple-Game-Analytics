package main

import (
	"net/http"
    "log"
	"fmt"
	"runtime"
    "os"

    "github.com/gorilla/mux"
    "github.com/Nesh108/Dead-Simple-Game-Analytics/pkg/db"
    "github.com/Nesh108/Dead-Simple-Game-Analytics/pkg/controllers"
    "github.com/Nesh108/Dead-Simple-Game-Analytics/pkg/middleware"
)

func main() {
    username := os.Getenv("AUTH_USERNAME")
    password := os.Getenv("AUTH_PASSWORD")

    if username == "" {
        log.Fatal("Basic auth username must be provided.")
        return;
    }

    if password == "" {
        log.Fatal("Basic auth password must be provided.")
        return;
    }

	// Use all available processing power
	runtime.GOMAXPROCS(runtime.NumCPU() - 1)

    DB := db.Init()
    c := controllers.New(DB)
    router := mux.NewRouter()

    router.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "{\"message\" :\"Hello!\"}")
    })

    router.HandleFunc("/events", c.GetEvents).Methods(http.MethodGet)
    router.HandleFunc("/events", c.CreateEvent).Methods(http.MethodPost)
    router.HandleFunc("/export", c.ExportEvents).Methods(http.MethodGet)

    router.Use(middleware.BasicAuth)

    log.Println("API is running!")
    fmt.Println(http.ListenAndServe(":" + os.Getenv("PORT"), router))
}

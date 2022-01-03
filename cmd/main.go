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
)

func main() {
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

    log.Println("API is running!")
    http.ListenAndServe(":" + os.Getenv("PORT"), router)
}

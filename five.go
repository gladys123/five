package main

import (
    "fmt"
    "github.com/julienschmidt/httprouter"
    "net/http"
    "log"
)

func Five(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    for i := 0; i < 5; i++ {
        fmt.Fprintf(w, "%d\n", 5)
    }
}

func main() { 
    router := httprouter.New()
    router.GET("/api/five", Five)

    log.Println("Listening on 127.0.0.1:5555...")
    log.Fatal(http.ListenAndServe(":5555", router)) 
}

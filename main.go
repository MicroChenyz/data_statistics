package main

import (
	"fmt"
	"log"
	"net/http"
	"wxcloudrun-golang/db"
	"wxcloudrun-golang/service"
)

func main() {
	if err := db.Init(); err != nil {
		panic(fmt.Sprintf("mysql init failed with %+v", err))
	}
	http.HandleFunc("/api/tm_user", service.TmUserHandler)
	http.HandleFunc("/api/user", service.UserHandler)
	http.HandleFunc("/api/client", service.ClientHandler)
	http.HandleFunc("/api/stemming", service.StemmingHandler)

	log.Fatal(http.ListenAndServe(":80", nil))

}

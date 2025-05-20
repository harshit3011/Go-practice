package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/harshit3011/students-api/internal/config"
)

 func main(){
	cfg:= config.MustLoad()
	router:= http.NewServeMux()

	router.HandleFunc("GET /",func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("Welcome to students api"))
	})

	server:=http.Server{
		Addr: cfg.Addr,
		Handler: router,
	}

	err:= server.ListenAndServe()
	if err!=nil{
		log.Fatal("Failed to start the server: ",err.Error())
	}
	fmt.Println("server started")
 }
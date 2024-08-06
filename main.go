package main

import (
	"log"
	"net/http"
	"webSiteProject/app/controller"

	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New()
	routes(r)
	err := http.ListenAndServe("localhost:7777", r)
	if err != nil {
		log.Fatal(err)
	}
}

func routes(r *httprouter.Router) {
	r.ServeFiles("/public/*filepath", http.Dir("public"))
	r.GET("/", controller.StartPage)

}

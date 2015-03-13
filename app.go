package main

import (
	"net/http"
	"os"

	"./controllers"

	"github.com/go-zoo/bone"
)

func main() {
	SERVER_PORT := ":" + os.Getenv("PORT")

	postController := new(controllers.PostController)

	router := bone.New()

	router.Get("/posts/:id", http.HandlerFunc(postController.ServePost))
	router.Get("/", http.HandlerFunc(postController.ServeIndex))
	router.Get("/:any", http.HandlerFunc(postController.ServeError))

	http.ListenAndServe(SERVER_PORT, router)
}

package controllers

import (
	"net/http"

	"github.com/go-zoo/bone"
)

type PostController struct {
}

func (controller *PostController) ServeIndex(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("Hello"))
}

func (controller *PostController) ServePost(rw http.ResponseWriter, req *http.Request) {

	postId := bone.GetValue(req, "id")

	rw.Write([]byte(postId))
}

func (controller *PostController) ServeError(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("Error"))
}

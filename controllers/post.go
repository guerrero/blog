package controllers

import (
	"net/http"

	"../models"
)

type PostController struct {
}

func (controller *PostController) ServeIndex(rw http.ResponseWriter, req *http.Request) {
	postList := models.GetIndex()

	var postNames string

	for _, post := range postList.Posts {
		postNames += post.Name
	}

	rw.Write([]byte(postNames))
}

func (controller *PostController) ServePost(rw http.ResponseWriter, req *http.Request) {
	models.GetPost(rw, req)
}

func (controller *PostController) ServeError(rw http.ResponseWriter, req *http.Request) {
	models.GetError(rw, req)
}

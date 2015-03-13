package controllers

import (
	"net/http"

	"../models"
)

type PostController struct {
}

func (controller *PostController) ServeIndex(rw http.ResponseWriter, req *http.Request) {
	models.GetIndex(rw, req)
}

func (controller *PostController) ServePost(rw http.ResponseWriter, req *http.Request) {
	models.GetPost(rw, req)
}

func (controller *PostController) ServeError(rw http.ResponseWriter, req *http.Request) {
	models.GetError(rw, req)
}

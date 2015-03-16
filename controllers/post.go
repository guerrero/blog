package controllers

import (
	"net/http"
	"text/template"

	"../models"
)

type PostController struct {
}

func (controller *PostController) ServeIndex(rw http.ResponseWriter, req *http.Request) {
	postList := models.GetIndex()

	renderView(rw, "home", "./views/home.html", postList)
}

func (controller *PostController) ServePost(rw http.ResponseWriter, req *http.Request) {
	post := models.GetPost(req)

	if post.Name != "" {
		renderView(rw, "post", "./views/post.html", post)
	} else {
		controller.ServeError(rw, req)
	}
}

func (controller *PostController) ServeError(rw http.ResponseWriter, req *http.Request) {

	error := models.GetError(req)

	renderView(rw, "error", "./views/error.html", error)

}

func renderView(rw http.ResponseWriter, name string, file string, data interface{}) {

	t := template.New(name)

	t, err := t.ParseFiles("./views/base.html",
		"./views/head.html",
		"./views/footer.html",
		file)

	t = template.Must(t, err)

	t.ExecuteTemplate(rw, "base", data)
}

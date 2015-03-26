package controllers

import (
	"net/http"
	"text/template"

	"github.com/go-zoo/bone"

	"../models"
)

type PostController struct {
}

func (controller *PostController) ServeIndex(rw http.ResponseWriter, req *http.Request) {
	postList := models.GetIndex()

	renderView(rw, "home", postList)
}

func (controller *PostController) ServePost(rw http.ResponseWriter, req *http.Request) {
	postRequested := bone.GetValue(req, "id")

	post := models.GetPost(postRequested)

	if post.Name != "" {
		renderView(rw, "post", post)
	} else {
		controller.ServeError(rw, req)
	}
}

func (controller *PostController) ServeError(rw http.ResponseWriter, req *http.Request) {

	pathRequested := req.URL.Path

	error := models.GetError(pathRequested)

	renderView(rw, "error", error)

}

func renderView(rw http.ResponseWriter, view string, data interface{}) {

	t := template.New(view)

	viewNames := []string{
		"base",
		"head",
		"header",
		"footer",
		view}

	viewFiles := make([]string, 0)

	for _, viewName := range viewNames {
		viewFile := "./views/" + viewName + ".html"
		viewFiles = append(viewFiles, viewFile)
	}

	t, err := t.ParseFiles(viewFiles...)

	t = template.Must(t, err)

	t.ExecuteTemplate(rw, "base", data)
}

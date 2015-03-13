package models

import (
	"net/http"

	"github.com/go-zoo/bone"
)

func GetIndex(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("Hello"))
}

func GetPost(rw http.ResponseWriter, req *http.Request) {
	postId := bone.GetValue(req, "id")

	rw.Write([]byte(postId))
}

func GetError(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("Error"))
}

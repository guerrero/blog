package models

import (
	"io/ioutil"
	"net/http"

	"github.com/go-zoo/bone"
	"github.com/russross/blackfriday"
)

type PostItem struct {
	Name    string
	Content []byte
}

type PostList struct {
	Posts []PostItem
}

func GetIndex() PostList {
	// rw.Write([]byte("Hello"))

	dirFiles, _ := ioutil.ReadDir("./content")

	posts := make([]PostItem, len(dirFiles))
	postList := PostList{Posts: posts}

	index := 0

	for _, file := range dirFiles {
		if file.Name() != ".DS_Store" {

			fileContent, _ := ioutil.ReadFile(file.Name())
			renderedOutput := blackfriday.MarkdownCommon(fileContent)

			postList.Posts[index] = PostItem{Name: file.Name(), Content: renderedOutput}

			index += 1
		}
	}

	return postList
}

func GetPost(rw http.ResponseWriter, req *http.Request) {
	postId := bone.GetValue(req, "id")

	rw.Write([]byte(postId))
}

func GetError(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("Error"))
}

package models

import (
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/go-zoo/bone"
	"github.com/russross/blackfriday"
)

type PostItem struct {
	Name     string
	Filename string
	Content  []byte
}

type PostList struct {
	Posts []PostItem
}

func GetIndex() PostList {

	dirFiles, _ := ioutil.ReadDir("./content")

	posts := make([]PostItem, len(dirFiles))
	postList := PostList{Posts: posts}

	index := 0

	for _, file := range dirFiles {
		if file.Name() != ".DS_Store" {

			filename := strings.Replace(file.Name(), ".md", "", -1)
			name := setPostName(filename)

			fileContent, _ := ioutil.ReadFile(file.Name())
			renderedOutput := blackfriday.MarkdownCommon(fileContent)

			postList.Posts[index] = PostItem{Name: name, Filename: filename, Content: renderedOutput}

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

func setPostName(filename string) string {
	nameWithoutExt := strings.Replace(filename, ".md", "", -1)
	return strings.Replace(nameWithoutExt, "-", " ", -1)
}

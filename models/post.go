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
	Content  string
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

			fileContent, _ := ioutil.ReadFile("content/" + file.Name())
			renderedOutput := string(blackfriday.MarkdownCommon(fileContent))

			postList.Posts[index] = PostItem{Name: name, Filename: filename, Content: renderedOutput}

			index += 1
		}
	}

	return postList
}

func GetPost(req *http.Request) PostItem {
	query := bone.GetValue(req, "id")
	postReq := strings.Replace(query, " ", "-", -1)

	postList := GetIndex()

	for _, post := range postList.Posts {
		if post.Filename == postReq {
			return post
		}
	}

	return PostItem{Name: ""}
}

func GetError(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("Error"))
}

func setPostName(filename string) string {
	nameWithoutExt := strings.Replace(filename, ".md", "", -1)
	return strings.Replace(nameWithoutExt, "-", " ", -1)
}

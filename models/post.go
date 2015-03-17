package models

import (
	"io/ioutil"
	"strings"

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

	posts := make([]PostItem, 0)
	postList := PostList{Posts: posts}

	index := 0

	for _, file := range dirFiles {
		if file.Name() != ".DS_Store" {

			filename := strings.Replace(file.Name(), ".md", "", -1)
			name := setPostName(filename)

			fileContent, _ := ioutil.ReadFile("content/" + file.Name())
			renderedOutput := string(blackfriday.MarkdownCommon(fileContent))

			postList.Posts = append(postList.Posts, PostItem{Name: name, Filename: filename, Content: renderedOutput})

			index += 1
		}
	}

	return postList
}

func GetPost(postRequested string) PostItem {

	postRequested = strings.Replace(postRequested, " ", "-", -1)

	postList := GetIndex()

	for _, post := range postList.Posts {
		if post.Filename == postRequested {
			return post
		}
	}

	return PostItem{Name: ""}
}

func GetError(pathRequested string) string {

	var errData string

	if strings.HasPrefix(pathRequested, "/posts/") {
		postRequested := strings.Replace(pathRequested, "/posts/", "", 1)
		postRequested = strings.Replace(postRequested, "-", " ", -1)

		errData = "The post \"" + postRequested + "\" doesn't exist or has been removed"
	} else {
		errData = "There's no page associated to " + pathRequested
	}

	return errData
}

func setPostName(filename string) string {
	nameWithoutExt := strings.Replace(filename, ".md", "", -1)
	return strings.Replace(nameWithoutExt, "-", " ", -1)
}

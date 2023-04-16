package main

import (
	"fmt"

	"github.com/sebastian009w/go_run/posts"
)

func main() {

	post := posts.Post{
		Title:       "title post",
		Description: "Description post",
		Published:   true,
	}
	fmt.Println(post.Title)

}

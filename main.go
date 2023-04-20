package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/sebastian009w/go_run/message"
	post "github.com/sebastian009w/go_run/posts"
)

func main() {

	file, err := os.OpenFile("posts.json", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var posts []post.Post

	info, err := file.Stat()

	if err != nil {
		panic(err)
	}

	if info.Size() != 0 {
		bytes, err := io.ReadAll(file)

		if err != nil {
			panic(err)
		}

		err = json.Unmarshal(bytes, &posts)
		if err != nil {
			panic(err)
		}
	} else {
		posts = []post.Post{}
	}

	if len(os.Args) < 2 {
		message.Message_init()
	}

	switch os.Args[1] {
	case "--help":
		message.Help_me()
	case "--posts":
		post.List_posts(posts)
	case "--create":
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Cual es tu tarea ?: ")
		title, _ := reader.ReadString('\n')
		title = strings.TrimSpace(title)

		posts = post.Add_post(posts, title)
		post.Add_JSON(file, posts)
	case "--delete":
		fmt.Println("delete")
	}
}

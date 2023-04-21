package post

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type Post struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func List_posts(posts []Post) {
	if len(posts) == 0 {
		fmt.Println("No hay Ninguna tarea disponibles")
		return
	}

	for _, post := range posts {
		status := " "
		if post.Completed {
			status = "✓"
		}
		fmt.Printf("[%s] %d %s \n", status, post.ID, post.Title)
	}
}

func Add_post(posts []Post, title string) []Post {
	post := Post{
		ID:        Gen_ID(posts),
		Title:     title,
		Completed: false,
	}
	return append(posts, post)
}

func Add_JSON(file *os.File, posts []Post) {
	bytes, err := json.Marshal(posts)

	if err != nil {
		panic(err)
	}
	_, err = file.Seek(0, 0)
	if err != nil {
		panic(err)
	}

	err = file.Truncate(0)
	if err != nil {
		panic(err)
	}
	writter := bufio.NewWriter(file)
	_, err = writter.Write(bytes)

	if err != nil {
		panic(err)
	}
	err = writter.Flush()
	if err != nil {
		panic(err)
	}
}

func Gen_ID(posts []Post) int {
	if len(posts) == 0 {
		return 1
	}
	return posts[len(posts)-1].ID + 1
}

func Delete_post(posts []Post, id int) []Post {
	for i, post := range posts {
		if post.ID == id {
			return append(posts[:i], posts[i+1:]...)
		}
	}
	return posts
}

func Completed_post(posts []Post, id int) []Post {
	for i, post := range posts {
		if post.ID == id {
			posts[i].Completed = true
			break
		}
	}
	return posts
}

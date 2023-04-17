package post

import "fmt"

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

	for i, post := range posts {

		fmt.Printf("%d %s \n", i, post.Title)
	}
}

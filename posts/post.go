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
		ID:        4,
		Title:     "TITLE FOR",
		Completed: false,
	}
	return append(posts, post)
}

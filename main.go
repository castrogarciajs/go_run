package main

import "os"

func main() {

	file, err := os.OpenFile("tasks.json", os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()
}

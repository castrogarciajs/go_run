package message

import "fmt"

func Message_init() {
	fmt.Println("Hello Welcome to CLI Go. Utiliza el comando --help para ver opciones")
}

func Help_me() {
	fmt.Println("Comandos disponibles: --posts | --create | --delete |--completed")
}

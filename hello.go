package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var str string = "Hello, world!"
	fmt.Println(str)
	fmt.Println(utf8.DecodeRuneInString(str)) //72

	// Извлечение пятой руны
	pos := 6 // Индекс 6 соответствует 7 символу (индексация начинается с 0)
	_, size := utf8.DecodeRuneInString(str)
	for i := 0; i < pos; i++ {
		_, s := utf8.DecodeRuneInString(str[size:])
		size += s
	}
	fifthRune, _ := utf8.DecodeRuneInString(str[size:])

	// Вывод пятого символа
	fmt.Println(pos, "character:", string(fifthRune)) // Вывод: Fifth character: o

	developerName := "Developer (first name): Ivan"

	fmt.Println(developerName)

	// определите переменные ver, id, pi
	// ...

	ver := "v0.0.1"
	id := 0
	pi := 3.1415

	fmt.Println("ver =", ver, "id =", id, "pi =", pi)

	const (
		Black = iota
		Gray
		White
	)

	fmt.Println(Black, Gray, White)
}

//Необходимо вызывать, для создания файла "go.mod" в корневой директории проекта:
//go mod init github.com/username/helloworld

package main

/* Создать приложение, которое сначала выдает меню:
- 1. Посмотреть закладки
- 2. Добавить закладку
- 3. Удалить закладку
- 4. Выйти из программы
При выборе 1 программа должна выводить все закладки из мапы.
При выборе 2 программа должна запрашивать название и URL и добавлять его в мапу.
При выборе 3 программа должна запрашивать название и удалять его из мапы.
При выборе 4 программа должна выходить из программы.
*/

import (
	"fmt"
)

type stringMap = map[string]string

func main() {
	var bookmarks = stringMap{}
	fmt.Println("--- Добро пожаловать в приложение для управления закладками ---")
Menu:
	for {
		printMenu()
		choice := getMenuChoice()
		switch choice {
		case 1:
			printBookmarks(bookmarks)
		case 2:
			addBookmark(bookmarks)
		case 3:
			deleteBookmark(bookmarks)
		case 4:
			break Menu
		}
	}
}

func printMenu() {
	fmt.Println("--- Меню ---")
	fmt.Println("1. Посмотреть закладки")
	fmt.Println("2. Добавить закладку")
	fmt.Println("3. Удалить закладку")
	fmt.Println("4. Выйти из программы")
}

func getMenuChoice() int {
	var choice int
	fmt.Scan(&choice)
	return choice
}

func addBookmark(bookmarks stringMap) {
	var name, url string
	fmt.Println("Введите название закладки:")
	fmt.Scan(&name)
	fmt.Println("Введите URL закладки:")
	fmt.Scan(&url)
	bookmarks[name] = url
}

func printBookmarks(bookmarks stringMap) {
	if len(bookmarks) == 0 {
		fmt.Println("Закладок нет")
		return
	}
	fmt.Println("--- Закладки ---")
	for name, url := range bookmarks {
		fmt.Println(name, ":", url)
	}
}

func deleteBookmark(bookmarks stringMap) {
	var name string
	fmt.Println("Введите название закладки:")
	fmt.Scan(&name)
	delete(bookmarks, name)
}

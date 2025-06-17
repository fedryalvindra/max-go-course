package main

import (
	"bufio"
	"fmt"
	"go-struct-practice-1/note"
	"go-struct-practice-1/todo"
	"os"
	"strings"
)

type saver interface {
	Save() error
}

type outputtable interface {
	saver
	Display()
}

func main() {
	result := add(1, 2)
	fmt.Println(result)

	title, content := getNoteData()
	todoText := getUserInput("Todo text: ")

	todo, err := todo.New(&todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	printSomething(todo)

	userNote, err := note.New(&title, &content)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(&todo)
	if err != nil {
		return
	}

	outputData(&userNote)
}

func add[T int | float64 | string](a, b T) T {
	return a + b
}

func printSomething(value interface{}) {
	intVal, ok := value.(int)
	if ok {
		fmt.Println("Integer:", intVal)
		return
	}

	floatVal, ok := value.(float64)
	if ok {
		fmt.Println("Float:", floatVal)
		return
	}

	stringVal, ok := value.(string)
	if ok {
		fmt.Println("String:", stringVal)
		return
	}
	// switch value.(type) {
	// case int:
	// 	fmt.Println("Integer:", value)
	// case float64:
	// 	fmt.Println("Float:", value)
	// case string:
	// 	fmt.Println(value)
	// }
}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}

func saveData(data saver) error {
	err := data.Save()
	if err != nil {
		fmt.Println("Saving the note failed.")
		return err
	}
	fmt.Println("Saving the note suceeded!")
	return nil
}

func getNoteData() (string, string) {
	title := getUserInput("Note Tile:")
	content := getUserInput("Note Content:")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}

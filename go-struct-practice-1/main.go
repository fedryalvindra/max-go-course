package main

import (
	"bufio"
	"fmt"
	"go-struct-practice-1/note"
	"os"
	"strings"
)

func main() {
	title, content := getNoteData()

	userNote, err := note.New(&title, &content)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote.Display()
	err = userNote.Save()

	if err != nil {
		fmt.Println("Saving the note failed.")
		return
	}
	fmt.Println("Saving the note suceeded!")
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

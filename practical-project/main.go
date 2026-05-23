package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"

	"example.com/note/note"
)

func main() {
	title, context := getNoteData()

	 userNote, err := note.New(title, context)

	if err != nil {
		fmt.Print(err)
		return
	}

	userNote.Display()

}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")

	context := getUserInput("Note content:")

	return title, context
}

func getUserInput(prompt string) string {
	fmt.Printf("%v", prompt)
	
	reader :=  bufio.NewReader(os.Stdin)

	text, err :=  reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}

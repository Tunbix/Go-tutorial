package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/change"
	"example.com/note/note"
)

type saver interface {
	Save() error
}

type outputter interface {
	Display()
	saver
}

func main() {
	title, context := getNoteData()
	todoText := getUserInput("Todo text:")

	todo, err := change.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote, err := note.New(title, context)

	if err != nil {
		fmt.Print(err)
		return
	}

	err = outputData(todo)

	if err != nil {
		return
	}

	outputData(userNote)

}

func outputData(data outputter) error {
	data.Display()
	return saveData(data)
}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Saving the todo failed.")
		return err
	}
	fmt.Println("Saving the todo succeeded!")
	return nil
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")

	context := getUserInput("Note content:")

	return title, context
}

func getUserInput(prompt string) string {
	fmt.Printf("%v", prompt)

	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}

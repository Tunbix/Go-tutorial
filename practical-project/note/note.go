package note

import (
	"errors"
	"fmt"
	"time"
)

type Note struct {
	title string
	content string
	createdAt time.Time
}

func (note Note) Display()  {
	fmt.Printf("Your note title %v has the following content:\n\n%v\n\n", note.title, note.content)
}

func New(title, context string) (Note, error) {
	if title == "" || context == "" {
		return Note{}, errors.New("Invalid input")
	}

	return Note{
		title: title,
		content: context,
		createdAt: time.Now(),
	}, nil
}
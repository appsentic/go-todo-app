package utils

import (
	"bufio"
	"context"
	"fmt"
	"strings"
)

var Options = map[string]string{
	"l": "List Tasks",
	"c": "Create a task",
	"d": "Delete a task",
	"h": "Get Help",
	"q": "Quit",
}

func CaptureUserInput(ctx context.Context, reader *bufio.Reader, question string) (string, error) {
	fmt.Printf("%s: ", question)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	value := strings.TrimSpace(input)

	return value, nil
}

func ShowSystemOptions() {

	fmt.Println("")
	for key, value := range Options {
		fmt.Printf("%s - %s\n", key, value)
	}
	fmt.Println("")
}

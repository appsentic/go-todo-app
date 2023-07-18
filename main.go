package main

import (
	"bufio"
	"context"
	"fmt"
	"os"

	"github.com/appsentic/go-todo-app/src/application/usecases"
	"github.com/appsentic/go-todo-app/src/application/utils"
)

func startTodoApp(ctx context.Context, reader *bufio.Reader) {
	tasks := usecases.GetTasks()
	if len(tasks) == 0 {
		fmt.Println("Hello there! \nYou do not have any pending tasks today!")
		utils.ShowSystemOptions()
		input, err := utils.CaptureUserInput(ctx, reader, "\nSelect an option above to continue")
		if err != nil {
			fmt.Println(err)
			utils.ShowSystemOptions()
		}
		handleOptionSelection(ctx, reader, input)
	} else {
		fmt.Println("Hello there! Here are your tasks for today!")
		for _, task := range tasks {
			fmt.Println(task.Format(ctx))
		}
	}
}

func handleRecursiveSelectionCalls(ctx context.Context, reader *bufio.Reader, q string) {

	fmt.Println(q)
	utils.ShowSystemOptions()
	input, err := utils.CaptureUserInput(ctx, reader, "\nSelect an option above to continue")
	if err != nil {
		fmt.Println(err)
		startTodoApp(ctx, reader)
	}
	handleOptionSelection(ctx, reader, input)
}

func handleOptionSelection(ctx context.Context, reader *bufio.Reader, selection string) {
	_, ok := utils.Options[selection]
	if !ok {
		handleRecursiveSelectionCalls(ctx, reader, "Invalid option!")
	}

	switch selection {
	case "c":
		item, err := utils.CaptureUserInput(ctx, reader, "\nEnter task name")
		if err != nil {
			fmt.Println(err)
			startTodoApp(ctx, reader)
		}

		task := usecases.CreateNewTask()
		task.AddTaskItem(ctx, item)
		handleRecursiveSelectionCalls(ctx, reader, "Task added successfully")
	case "d":
	case "l":
		startTodoApp(ctx, reader)
	case "h":
	default:
	}

}

func main() {
	reader := bufio.NewReader(os.Stdin)
	startTodoApp(context.TODO(), reader)
}

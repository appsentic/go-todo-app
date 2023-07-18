package dao

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

var DB = []*Task{}
var ErrorTaskNotFound = errors.New("Task was not found")

type Task struct {
	ID          string     `json:"id"`
	Items       []TaskItem `json:"items"`
	DateCreated string     `json:"date_created"`
	DateUpdated string     `json:"date_updated"`
}

type TaskItem struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Done        bool   `json:"done"`
	DateCreated string `json:"date_created"`
	DateUpdated string `json:"date_updated"`
}

func CreateTask(ctx context.Context) *Task {
	task := &Task{
		ID:          uuid.NewString(),
		DateCreated: time.Now().String(),
		DateUpdated: time.Now().String(),
		Items:       []TaskItem{},
	}
	DB = append(DB, task)
	return task
}

func GetTasks(ctx context.Context) []*Task {
	return DB
}

func GetTask(ctx context.Context, id string) (*Task, error) {
	for _, task := range DB {
		if task.ID == id {
			return task, nil
		}
	}
	return nil, ErrorTaskNotFound
}

func (t *Task) AddTaskItem(ctx context.Context, name string) *Task {
	task := &TaskItem{
		ID:          uuid.NewString(),
		Name:        name,
		Done:        false,
		DateCreated: time.Now().String(),
		DateUpdated: time.Now().String(),
	}

	t.Items = append(t.Items, *task)
	return t
}

func (t *Task) UpdateTaskItem(ctx context.Context, id string, done bool) *Task {
	for i, item := range t.Items {
		if item.ID == id {
			item.Done = done
			item.DateUpdated = time.Now().String()
			t.Items = append(t.Items[:i], item)
		} else {
			t.Items = append(t.Items[:i], t.Items[i+1:]...)
		}

	}
	return t
}

func (t *Task) RemoveTaskItem(ctx context.Context, id string) *Task {
	for i, item := range t.Items {
		if item.ID != id {
			t.Items = append(t.Items[:i], t.Items[i+1:]...)
		}
	}
	return t
}

func (t Task) Format(ctx context.Context) string {
	format := fmt.Sprintln("---------------------------------------------------------------")
	format += fmt.Sprintf("Task %s :\n\n", t.ID)

	for _, item := range t.Items {
		format += fmt.Sprintf("\n-- %s", item.Name)
		format += fmt.Sprintf("\n-- done %v", item.Done)
		format += fmt.Sprintf("\n-- created on %s", item.DateCreated)
	}

	format += fmt.Sprintln("---------------------------------------------------------------")
	return format
}

package usecases
import (
	"context"

	"github.com/appsentic/go-todo-app/src/domain/dao"
)

type TaskUsecase interface {
	CreateNewTask() *dao.Task
	GetTask(id string) (*dao.Task, error)
	GetTasks() []*dao.Task
}

func CreateNewTask() *dao.Task {
	return dao.CreateTask(context.TODO())
}

func GetTask(id string) (*dao.Task, error) {
	return dao.GetTask(context.TODO(), id)
}

func GetTasks() []*dao.Task {
	return dao.GetTasks(context.TODO())
}

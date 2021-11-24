package task

import (
	"errors"
	"github.com/jinzhu/gorm"
)

type TaskRepository struct {
	database *gorm.DB
}

func (repository *TaskRepository) FindAll() []Task {
	var tasks []Task
	repository.database.Find(&tasks)
	return tasks
}

func (repository *TaskRepository) Find(id int) (Task, error) {
	var task Task
	err := repository.database.Find(&task, id).Error
	if task.Name == "" {
		err = errors.New("Task not found")
	}
	return task, err
}

func (repository *TaskRepository) Create(task Task) (Task, error) {
	err := repository.database.Create(&task).Error
	return task, err
}

func (repository *TaskRepository) Save(task Task) (Task, error) {
	err := repository.database.Save(task).Error
	return task, err
}

func (repository *TaskRepository) Delete(id int) int64 {
	count := repository.database.Delete(&Task{}, id).RowsAffected
	return count
}

func NewTaskRepository(database *gorm.DB) *TaskRepository {
	return &TaskRepository{
		database: database,
	}
}

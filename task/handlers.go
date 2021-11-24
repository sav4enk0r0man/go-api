package task

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	"strconv"
)

type TaskHandler struct {
	repository *TaskRepository
}

func (handler *TaskHandler) GetAll(c *fiber.Ctx) error {
	var tasks []Task = handler.repository.FindAll()
	return c.JSON(tasks)
}

func (handler *TaskHandler) Get(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	task, err := handler.repository.Find(id)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status": 404,
			"error":  err,
		})
	}

	return c.JSON(task)
}

func (handler *TaskHandler) Create(c *fiber.Ctx) error {
	data := new(Task)

	if err := c.BodyParser(data); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error", 
			"message": "Invalid input data", 
			"error": err
		})
	}

	task, err := handler.repository.Create(*data)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": 400,
			"message": "Failed creating task",
			"error": err,
		})
	}

	return c.JSON(task)
}

func (handler *TaskHandler) Update(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  400,
			"message": "Task id not valid",
			"error":   err,
		})
	}

	task, err := handler.repository.Find(id)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Task not found",
		})
	}

	taskData := new(Task)

	if err := c.BodyParser(taskData); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "error", 
			"message": "Invalid input data", 
			"data": err})
	}

	task.Name = taskData.Name
	task.Description = taskData.Description
	task.Status = taskData.Status

	task, err := handler.repository.Save(task)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Error updating task",
			"error": err,
		})
	}

	return c.JSON(task)
}

func (handler *TaskHandler) Delete(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": 400,
			"message": "Task id not valid",
			"err": err,
		})
	}
	RowsAffected := handler.repository.Delete(id)
	statusCode := 204
	if RowsAffected == 0 {
		statusCode = 400
	}
	return c.Status(statusCode).JSON(nil)
}

func NewTaskHandler(repository *TaskRepository) *TaskHandler {
	return &TaskHandler{
		repository: repository,
	}
}

func Register(router fiber.Router, database *gorm.DB) {
	database.AutoMigrate(&Task{})
	taskRepository := NewTaskRepository(database)
	taskHandler := NewTaskHandler(taskRepository)

	taskRouter := router.Group("/task")
	taskRouter.Get("/", taskHandler.GetAll)
	taskRouter.Get("/:id", taskHandler.Get)
	taskRouter.Put("/:id", taskHandler.Update)
	taskRouter.Post("/", taskHandler.Create)
	taskRouter.Delete("/:id", taskHandler.Delete)
}

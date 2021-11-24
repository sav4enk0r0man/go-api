package task

import "github.com/jinzhu/gorm"

const (
	PENDING  = "pending"
	PROGRESS = "in_progress"
	DONE     = "done"
)

type Task struct {
	gorm.Model
	Name        string `gorm:"Not Null" json:"name"`
	Description string `json:"description"`
	Status      string `gorm:"Not Null" json:"status"`
}

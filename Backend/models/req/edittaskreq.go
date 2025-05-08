package req

import "time"

type EditTaskReq struct {
	TaskID         int       `json:"task_id" validate:"required" gorm:"primaryKey"`
	TaskName       string    `json:"task_name" validate:"required"`
	CourseID       int       `json:"course_id" validate:"required"`
	TaskDifficulty int       `json:"task_difficulty" validate:"required"`
	TaskWeight     int       `json:"task_weight" validate:"required"`
	TaskDueDate    time.Time `json:"task_due_date" validate:"required" gorm:"type:date"`
}

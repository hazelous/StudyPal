package entity

import "time"

type Tasks struct {
	TaskID         int       `json:"task_id" gorm:"primaryKey"`
	TaskName       string    `json:"task_name"`
	CourseID       int       `json:"course_id"`
	TaskDifficulty int       `json:"task_difficulty"`
	TaskWeight     int       `json:"task_weight"`
	TaskDueDate    time.Time `json:"task_due_date" gorm:"type:date"`
	TaskStatus     string    `json:"task_status"`
}

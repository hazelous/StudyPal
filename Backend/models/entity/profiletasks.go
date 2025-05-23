package entity

type ProfileTasks struct {
	ProfileID  int    `json:"profile_id" gorm:"primaryKey"`
	TaskID     int    `json:"task_id" gorm:"primaryKey"`
	TaskStatus string `json:"task_status"`
}

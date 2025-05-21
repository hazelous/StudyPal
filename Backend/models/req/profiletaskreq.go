package req

type ProfileTaskReq struct {
	ProfileID  int    `json:"profile_id" validate:"required"`
	TaskID     int    `json:"task_id" validate:"required"`
	TaskStatus string `json:"task_status" validate:"required"`
}

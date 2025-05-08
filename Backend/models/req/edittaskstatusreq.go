package req

type EditTaskStatusReq struct {
	TaskID     int    `json:"task_id" validate:"required"`
	TaskStatus string `json:"task_status" validate:"required"`
}

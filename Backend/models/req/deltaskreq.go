package req

type DelTaskReq struct {
	TaskID    int    `json:"task_id" gorm:"primaryKey"`
}
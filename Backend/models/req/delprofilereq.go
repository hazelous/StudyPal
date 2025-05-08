package req

type DelProfileReq struct {
	ProfileID    int    `json:"profile_id" gorm:"primaryKey"`
}
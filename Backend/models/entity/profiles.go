package entity

type Profiles struct {
	ProfileID       int    `json:"profile_id" gorm:"primaryKey"`
	ProfileName     string `json:"profile_name" gorm:"uniqueIndex"`
	ProfilePassword string `json:"profile_password"`
	ProfileImage    string `json:"profile_image"`
}

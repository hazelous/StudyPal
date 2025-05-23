package entity

type ProfileCourses struct {
	ProfileID int `json:"profile_id" gorm:"primaryKey"`
	CourseID  int `json:"course_id" gorm:"primaryKey"`
}

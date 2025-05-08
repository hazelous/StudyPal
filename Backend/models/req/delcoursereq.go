package req

type DelCourseReq struct {
	CourseID int `json:"course_id" gorm:"primaryKey"`
}

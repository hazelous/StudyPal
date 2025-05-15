package req

type ProfileCourseReq struct {
	ProfileID int `json:"profile_id" validate:"required"`
	CourseID  int `json:"course_id" validate:"required"`
}

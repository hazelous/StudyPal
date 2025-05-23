package req

type EditCourseReq struct {
	CourseID    int    `json:"course_id" validate:"required"`
	CourseName  string `json:"course_name" validate:"required"`
	CourseImage string `json:"course_image" validate:"required"`
}

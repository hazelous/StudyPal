package req

type CourseReq struct {
	CourseName  string `json:"course_name" validate:"required"`
	CourseImage string `json:"course_image" validate:"required"`
}

package entity

type Courses struct {
	CourseID    int    `json:"course_id" gorm:"primaryKey"`
	CourseName  string `json:"course_name"`
	CourseImage string `json:"course_image"`
}

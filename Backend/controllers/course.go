package controllers

import (
	"studypal/database"
	"studypal/models/entity"
	"studypal/models/req"

	validator "github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func ShowAllCourses(c *fiber.Ctx) error {
	var courses []entity.Courses
	if err := database.DB.Find(&courses).Error; err != nil {
		return err
	}
	return c.JSON(fiber.Map{
		"data": courses,
	})
}

func ShowCourseIdForProfile(c *fiber.Ctx) error {
	ProfileID := c.Params("id") //extract id from the url
	var courseIDs []int

	if err := database.DB.
		Table("profile_courses").
		Select("course_id").
		Where("profile_id = ?", ProfileID).
		Scan(&courseIDs).Error; err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"data": courseIDs,
	})
}

func AddCourse(c *fiber.Ctx) error {
	course := new(req.CourseReq)
	if err := c.BodyParser(course); err != nil {
		return err
	}
	validate := validator.New()
	if err := validate.Struct(course); err != nil {
		return err
	}
	newCourse := entity.Courses{
		CourseName:  course.CourseName,
		CourseImage: course.CourseImage,
	}
	if err := database.DB.Create(&newCourse).Error; err != nil {
		return err
	}
	return c.JSON(fiber.Map{
		"message": "Successfully created course",
	})
}

func DeleteCourse(c *fiber.Ctx) error {
	course := new(req.DelCourseReq)
	if err := c.BodyParser(course); err != nil {
		return err
	}
	validate := validator.New()
	if err := validate.Struct(course); err != nil {
		return err
	}
	delcourse := entity.Courses{
		CourseID: course.CourseID,
	}
	if err := database.DB.Delete(&delcourse).Error; err != nil {
		return err
	}
	return c.JSON(fiber.Map{
		"message": "Successfully Deleted Course",
	})
}

func EditCourse(c *fiber.Ctx) error {
	course := new(req.EditCourseReq)
	if err := c.BodyParser(course); err != nil {
		return err
	}
	validate := validator.New()
	if err := validate.Struct(course); err != nil {
		return err
	}
	editcourse := entity.Courses{
		CourseID:    course.CourseID,
		CourseName:  course.CourseName,
		CourseImage: course.CourseImage,
	}
	if err := database.DB.Model(&editcourse).Updates(map[string]interface{}{
		"course_name":  editcourse.CourseName,
		"course_image": editcourse.CourseImage,
	}).Error; err != nil {
		return err
	}
	return c.JSON(fiber.Map{
		"message": "Successfully Edited Course",
	})
}

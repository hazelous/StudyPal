package controllers

import (
	"studypal/database"
	"studypal/models/entity"
	"studypal/models/req"

	validator "github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func ShowAllProfiles(c *fiber.Ctx) error {
	var profiles []entity.Profiles
	if err := database.DB.Find(&profiles).Error; err != nil {
		return err
	}
	return c.JSON(fiber.Map{
		"data": profiles,
	})
}

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

func ShowAllTasks(c *fiber.Ctx) error {
	var tasks []entity.Tasks
	if err := database.DB.Find(&tasks).Error; err != nil {
		return err
	}
	return c.JSON(fiber.Map{
		"data": tasks,
	})
}

func ShowTasksForProfile(c *fiber.Ctx) error {
	ProfileID := c.Params("id")
	var tasks []entity.Tasks
	if err := database.DB.Table("tasks").
		Select("tasks.*").
		Joins("JOIN profile_courses ON tasks.course_id = profile_courses.course_id").
		Where("profile_id = ?", ProfileID).
		Find(&tasks).Error; err != nil {
		return err
	}
	return c.JSON(fiber.Map{
		"data": tasks,
	})
}

func AddProfile(c *fiber.Ctx) error {
	profile := new(req.ProfileReq)
	if err := c.BodyParser(profile); err != nil {
		return err
	}
	validate := validator.New()
	if err := validate.Struct(profile); err != nil {
		return err
	}
	newProfile := entity.Profiles{
		ProfileName:  profile.ProfileName,
		ProfileImage: profile.ProfileImage,
	}
	if err := database.DB.Create(&newProfile).Error; err != nil {
		return err
	}
	return c.JSON(fiber.Map{
		"message": "Successfully created profile",
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

func AddProfileCourse(c *fiber.Ctx) error {
	ProfileCourse := new(req.ProfileCourseReq)
	if err := c.BodyParser(ProfileCourse); err != nil {
		return err
	}
	validate := validator.New()
	if err := validate.Struct(ProfileCourse); err != nil {
		return err
	}
	newProfileCourse := entity.ProfileCourses{
		ProfileID: ProfileCourse.ProfileID,
		CourseID:  ProfileCourse.CourseID,
	}
	if err := database.DB.Create(&newProfileCourse).Error; err != nil {
		return err
	}
	return c.JSON(fiber.Map{
		"message": "successfully created ProfileCourse",
	})
}

func AddTask(c *fiber.Ctx) error {
	task := new(req.TaskReq)
	if err := c.BodyParser(task); err != nil {
		return err
	}
	validate := validator.New()
	if err := validate.Struct(task); err != nil {
		return err
	}
	newTask := entity.Tasks{
		TaskName:       task.TaskName,
		CourseID:       task.CourseID,
		TaskDifficulty: task.TaskDifficulty,
		TaskWeight:     task.TaskWeight,
		TaskDueDate:    task.TaskDueDate,
		TaskStatus:     task.TaskStatus,
	}
	if err := database.DB.Create(&newTask).Error; err != nil {
		return err
	}
	return c.JSON(fiber.Map{
		"message": "Successfully Created Task",
	})
}

func DeleteProfile(c *fiber.Ctx) error {
	profile := new(req.DelProfileReq)
	if err := c.BodyParser(profile); err != nil {
		return err
	}
	validate := validator.New()
	if err := validate.Struct(profile); err != nil {
		return err
	}
	delprofile := entity.Profiles{
		ProfileID: profile.ProfileID,
	}
	if err := database.DB.Delete(&delprofile).Error; err != nil {
		return err
	}
	return c.JSON(fiber.Map{
		"message": "Successfully Deleted Profile",
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

func DeleteProfileCourse(c *fiber.Ctx) error {
	ProfileCourse := new(req.ProfileCourseReq)
	if err := c.BodyParser(ProfileCourse); err != nil {
		return err
	}
	validate := validator.New()
	if err := validate.Struct(ProfileCourse); err != nil {
		return err
	}
	delprofilecourse := entity.ProfileCourses{
		ProfileID: ProfileCourse.ProfileID,
		CourseID:  ProfileCourse.CourseID,
	}

	if err := database.DB.Delete(&delprofilecourse).Error; err != nil {
		return err
	}
	return c.JSON(fiber.Map{
		"message": "Successfully Deleted ProfileCourse",
	})
}

func DeleteTask(c *fiber.Ctx) error {
	task := new(req.DelTaskReq)
	if err := c.BodyParser(task); err != nil {
		return err
	}
	validate := validator.New()
	if err := validate.Struct(task); err != nil {
		return err
	}
	deltask := entity.Tasks{
		TaskID: task.TaskID,
	}
	if err := database.DB.Delete(&deltask).Error; err != nil {
		return err
	}
	return c.JSON(fiber.Map{
		"message": "Successfully Deleted Task",
	})
}

func EditProfile(c *fiber.Ctx) error {
	profile := new(req.EditProfileReq)
	if err := c.BodyParser(profile); err != nil {
		return err
	}
	validate := validator.New()
	if err := validate.Struct(profile); err != nil {
		return err
	}
	editprofile := entity.Profiles{
		ProfileID:    profile.ProfileID,
		ProfileName:  profile.ProfileName,
		ProfileImage: profile.ProfileImage,
	}
	if err := database.DB.Model(&editprofile).Update("profile_name", editprofile.ProfileName).Error; err != nil {
		return err
	}
	if err := database.DB.Model(&editprofile).Update("profile_image", editprofile.ProfileImage).Error; err != nil {
		return err
	}
	return c.JSON(fiber.Map{
		"message": "Successfully Edited Profile",
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

func EditTask(c *fiber.Ctx) error {
	task := new(req.EditTaskReq)
	if err := c.BodyParser(task); err != nil {
		return err
	}
	validate := validator.New()
	if err := validate.Struct(task); err != nil {
		return err
	}
	edittask := entity.Tasks{
		TaskID:         task.TaskID,
		TaskName:       task.TaskName,
		CourseID:       task.CourseID,
		TaskDifficulty: task.TaskDifficulty,
		TaskWeight:     task.TaskWeight,
		TaskDueDate:    task.TaskDueDate,
	}
	if err := database.DB.Model(&edittask).Updates(map[string]interface{}{
		"task_name":       task.TaskName,
		"course_id":       task.CourseID,
		"task_difficulty": task.TaskDifficulty,
		"task_weight":     task.TaskWeight,
		"task_due_date":   task.TaskDueDate,
	}).Error; err != nil {
		return err
	}
	return c.JSON(fiber.Map{
		"message": "Successfully Edited Task",
	})
}

func EditTaskStatus(c *fiber.Ctx) error {
	task := new(req.EditTaskStatusReq)
	if err := c.BodyParser(task); err != nil {
		return err
	}
	validate := validator.New()
	if err := validate.Struct(task); err != nil {
		return err
	}
	edittask := entity.Tasks{
		TaskID:     task.TaskID,
		TaskStatus: task.TaskStatus,
	}
	if err := database.DB.Model(&edittask).Update("task_status", edittask.TaskStatus).Error; err != nil {
		return err
	}
	return c.JSON(fiber.Map{
		"message": "Successfully Edited Status",
	})
}

package controllers

import (
	"studypal/database"
	"studypal/models/entity"
	"studypal/models/req"

	validator "github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

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

func ShowTaskStatusForProfile(c *fiber.Ctx) error {
	ProfileID := c.Params("id")
	var ProfileTasks []entity.ProfileTasks
	if err := database.DB.Where("profile_id = ?", ProfileID).Find(&ProfileTasks).Error; err != nil {
		return err
	}
	return c.JSON(fiber.Map{
		"data": ProfileTasks,
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
	}
	if err := database.DB.Create(&newTask).Error; err != nil {
		return err
	}
	return c.JSON(fiber.Map{
		"message": "Successfully Created Task",
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

//func EditTaskStatus(c *fiber.Ctx) error {
//task := new(req.EditTaskStatusReq)
//if err := c.BodyParser(task); err != nil {
//	return err
//}
//validate := validator.New()
//if err := validate.Struct(task); err != nil {
//	return err
//}
//edittask := entity.Tasks{
//	TaskID:     task.TaskID,
//	TaskStatus: task.TaskStatus,
//}
//if err := database.DB.Model(&edittask).Update("task_status", edittask.TaskStatus).Error; err != nil {
//	return err
//}
//return c.JSON(fiber.Map{
//	"message": "Successfully Edited Status",
//	})
//}
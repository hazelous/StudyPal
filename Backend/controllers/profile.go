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

func ShowProfileByID(c *fiber.Ctx) error {
	ProfileID := c.Params("id")
	var profile entity.Profiles

	if err := database.DB.Where("profile_id = ?", ProfileID).Find(&profile).Error; err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"data": profile,
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

func UpdateTaskStatus(c *fiber.Ctx) error {
	ProfileTask := new(req.ProfileTaskReq)
	if err := c.BodyParser(ProfileTask); err != nil {
		return err
	}
	validate := validator.New()
	if err := validate.Struct(ProfileTask); err != nil {
		return err
	}

	editprofiletask := entity.ProfileTasks{
		ProfileID:  ProfileTask.ProfileID,
		TaskID:     ProfileTask.TaskID,
		TaskStatus: ProfileTask.TaskStatus,
	}

	res := database.DB.
		Model(&editprofiletask).
		Where("profile_id = ? AND task_id = ?", ProfileTask.ProfileID, ProfileTask.TaskID).
		Update("task_status", ProfileTask.TaskStatus)

	if res.Error != nil {
		return res.Error
	}

	// 4) If no row was updated, insert a new one
	if res.RowsAffected == 0 {
		if err := database.DB.Create(&editprofiletask).Error; err != nil {
			return err
		}
	}

	// 5) Return success
	return c.JSON(fiber.Map{
		"message": "status saved",
	})
}
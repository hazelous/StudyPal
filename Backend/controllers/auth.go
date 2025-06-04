package controllers

import (
	"studypal/database"
	"studypal/models/entity"
	"github.com/gofiber/fiber/v2"
)

func LoginProfile(c *fiber.Ctx) error {
	// extract name and password from url
	ProfileName := c.Params("name")
	ProfilePassword := c.Params("password")

	var profile entity.Profiles

	if err := database.DB.
	Where("profile_name = ?", ProfileName).
	Where("profile_password = ?", ProfilePassword).
	Find(&profile).Error; err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"data": profile,
	})
}
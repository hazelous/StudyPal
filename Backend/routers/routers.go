package routers

import (
	"studypal/controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func RouterApp(c *fiber.App) {
	c.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))
	c.Get("/api/showallprofiles", controllers.ShowAllProfiles)
	c.Get("/api/showallcourses", controllers.ShowAllCourses)
	c.Get("/api/showalltasks", controllers.ShowAllTasks)
	c.Post("/api/addprofile", controllers.AddProfile)
	c.Post("/api/addcourse", controllers.AddCourse)
	c.Post("/api/addtask", controllers.AddTask)
	c.Post("/api/deleteprofile", controllers.DeleteProfile)
	c.Post("/api/deletecourse", controllers.DeleteCourse)
	c.Post("/api/deletetask", controllers.DeleteTask)
	c.Post("/api/editprofile", controllers.EditProfile)
	c.Post("/api/editcourse", controllers.EditCourse)
	c.Post("/api/edittask", controllers.EditTask)
	c.Post("/api/editstatus", controllers.EditTaskStatus)
}

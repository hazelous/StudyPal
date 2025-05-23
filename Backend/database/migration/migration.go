package migration

import (
	"fmt"

	"studypal/database"
	"studypal/models/entity"
)

func Migrate() {
	err := database.DB.AutoMigrate(
		&entity.Profiles{},
		&entity.Courses{},
		&entity.ProfileCourses{},
		&entity.Tasks{},
		&entity.ProfileTasks{},
	)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Migration Successful")
	}
}

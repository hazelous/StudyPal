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
	)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Migration Successful")
	}
}

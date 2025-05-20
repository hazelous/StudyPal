package migration

import (
	"fmt"
	"time"

	"studypal/models/entity"

	"gorm.io/gorm"
)

func DBSeed(db *gorm.DB) error {
	if err := ProfileSeed(db); err != nil {
		return err
	}
	if err := CourseSeed(db); err != nil {
		return err
	}
	if err := ProfileCourseSeed(db); err != nil {
		return err
	}
	if err := TaskSeed(db); err != nil {
		return err
	}
	// if err := ProfileTaskSeed(db);  err != nil { return err }
	return nil
}

func ProfileSeed(db *gorm.DB) error {
	var count int64
	if err := db.Model(&entity.Profiles{}).Count(&count).Error; err != nil {
		return err
	}

	if count == 0 {
		Profiles := []entity.Profiles{
			{ProfileName: "Guest", ProfileImage: "https://media.istockphoto.com/id/1495088043/vector/user-profile-icon-avatar-or-person-icon-profile-picture-portrait-symbol-default-portrait.jpg?s=612x612&w=0&k=20&c=dhV2p1JwmloBTOaGAtaA3AW1KSnjsdMt7-U_3EZElZ0="},
			{ProfileName: "Test", ProfileImage: "Test"},
		}

		// Transaction means that if theres an error in adding one of the profiles and cancels halfway through,
		// The function will rollback all the changes and stop.
		if err := db.Transaction(func(tx *gorm.DB) error {
			if err := tx.Create(&Profiles).Error; err != nil {
				return err
			}
			fmt.Println("Seeding profiles success")
			return nil
		}); err != nil {
			return err
		}
	} else {
		fmt.Println("Profiles already seeded")
	}
	return nil
}

func CourseSeed(db *gorm.DB) error {
	var count int64
	if err := db.Model(&entity.Courses{}).Count(&count).Error; err != nil {
		return err
	}

	if count == 0 {
		Courses := []entity.Courses{
			{CourseName: "SIT120 - Responsive Web Apps", CourseImage: "https://i.imgur.com/4NNunpg.png"},
			{CourseName: "SIT111 - Computer Systems", CourseImage: "https://i.imgur.com/dz3bDfI.png"},
			{CourseName: "Test", CourseImage: "Test"},
		}

		if err := db.Transaction(func(tx *gorm.DB) error {
			if err := tx.Create(&Courses).Error; err != nil {
				return err
			}
			fmt.Println("Seeding courses success")
			return nil
		}); err != nil {
			return err
		}
	} else {
		fmt.Println("Courses already seeded")
	}
	return nil
}

func ProfileCourseSeed(db *gorm.DB) error {
	var count int64
	if err := db.Model(&entity.ProfileCourses{}).Count(&count).Error; err != nil {
		return err
	}

	if count == 0 {
		ProfileCourses := []entity.ProfileCourses{
			{ProfileID: 1, CourseID: 1},
			{ProfileID: 1, CourseID: 2},
			{ProfileID: 2, CourseID: 1},
		}

		if err := db.Transaction(func(tx *gorm.DB) error {
			if err := tx.Create(&ProfileCourses).Error; err != nil {
				return err
			}
			fmt.Println("Seeding profilecourses success")
			return nil
		}); err != nil {
			return err
		}
	} else {
		fmt.Println("Profilecourses already seeded")
	}
	return nil
}

func TaskSeed(db *gorm.DB) error {
	var count int64
	if err := db.Model(&entity.Tasks{}).Count(&count).Error; err != nil {
		return err
	}

	if count == 0 {
		Tasks := []entity.Tasks{
			{
				TaskName:       "1.1P Learning Plan",
				CourseID:       2,
				TaskDifficulty: 1,
				TaskWeight:     1,
				TaskDueDate:    time.Date(2025, time.October, 2, 0, 0, 0, 0, time.UTC),
				TaskStatus:     "Not Started",
			},
			{
				TaskName:       "Project 1",
				CourseID:       1,
				TaskDifficulty: 10,
				TaskWeight:     30,
				TaskDueDate:    time.Date(2025, time.January, 15, 0, 0, 0, 0, time.UTC),
				TaskStatus:     "Completed",
			},
			{
				TaskName:       "Test",
				CourseID:       3,
				TaskDifficulty: 10,
				TaskWeight:     100,
				TaskDueDate:    time.Date(2025, time.May, 21, 0, 0, 0, 0, time.UTC),
				TaskStatus:     "Not Started",
			},
		}

		if err := db.Transaction(func(tx *gorm.DB) error {
			if err := tx.Create(&Tasks).Error; err != nil {
				return err
			}
			fmt.Println("Seeding Tasks success")
			return nil
		}); err != nil {
			return err
		}
	} else {
		fmt.Println("Courses already seeded")
	}
	return nil
}

//func ProfileTaskSeed(db *gorm.DB) error {
//
//}

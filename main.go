package main

import (
	"log"
	"oddshub/domain"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	course := domain.Course{
		Name:        "CCO",
		Description: "Learn to deliver a high quality software from start to end",
		Capacity:    16,
		Price:       100000,
		Trainer: domain.Trainer{
			FirstName: "Peerawat",
			LastName:  "Poombua",
			Email:     "pong@odds.team",
		},
	}

	_ = course

	db.AutoMigrate(&domain.Trainer{}, &domain.Course{})

	saveCourseFunc := func(course domain.Course) error {
		return db.Create(&course).Error
	}
	course.With(saveCourseFunc)
	course.Save()
}

package repository

import (
	"fmt"
	"microservice_go/types"
)

var courses = []types.Course{
	{Id: 1, Title: "Course 1"},
	{Id: 2, Title: "Course 2"},
}

func GetStages() []types.Course {
	return courses
}

func FindStage(id int) (*types.Course, error) {
	item := types.Course{}
	idx := func() int {
		for i := 0; i < len(courses); i++ {
			if courses[i].Id == id {
				return i
			}
		}
		return -1
	}()

	if idx == -1 {
		return &item, fmt.Errorf("курс с таким id %d не найден", id)
	}

	item = courses[idx]
	return &item, nil
}

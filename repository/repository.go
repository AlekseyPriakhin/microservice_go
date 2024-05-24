package repository

import (
	"fmt"
	"microservice_go/types"
	"microservice_go/utils"
)

type CourseResDto struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

type CourseReqDto struct {
	Title  string `json:"title"`
	Status string `json:"status"`
}

var courses = []types.Course{}
var index = 0

func init() {
	courses = append(courses,
		types.Course{Id: 1, Title: "Course 1", Status: types.Active},
		types.Course{Id: 2, Title: "Course 2", Status: types.Archived},
	)
	index = len(courses)
}

func GetCourse() []CourseResDto {
	res := utils.Map(courses, func(c types.Course) CourseResDto {
		return CourseResDto{
			Id:     c.Id,
			Title:  c.Title,
			Status: c.Status.String(),
		}
	})

	return res
}

func FindCourse(id int) (CourseResDto, error) {
	item := CourseResDto{}
	idx := func() int {
		for i := 0; i < len(courses); i++ {
			if courses[i].Id == id {
				return i
			}
		}
		return -1
	}()

	if idx == -1 {
		return item, fmt.Errorf("курс с таким id %d не найден", id)
	}

	item = CourseResDto{
		Id:     courses[idx].Id,
		Title:  courses[idx].Title,
		Status: courses[idx].Status.String(),
	}
	return item, nil
}

func AddCourse(c CourseReqDto) (CourseResDto, error) {

	index++
	id := index
	status, err := types.MapToCourseStatus(c.Status)

	if err != nil {
		return CourseResDto{}, err
	}

	course := types.Course{
		Id:     id,
		Title:  c.Title,
		Status: types.Status(status),
	}

	courses = append(courses, course)
	return FindCourse(id)
}

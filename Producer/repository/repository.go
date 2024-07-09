package repository

import (
	"fmt"
	"microservice_go/types"
	"microservice_go/utils"
)

type CourseReqDto struct {
	Title  string        `json:"title"`
	Status string        `json:"status"`
	Stages []StageReqDto `json:"stages"`
}

type StageReqDto struct {
	Title string `json:"title"`
	Order int    `json:"order"`
	Desc  string `json:"desc"`
	Type  string `json:"type"`
}

var Users = []string{
	"User 1",
	"User 2",
	"User 3"}

var courses = []types.Course{}
var courseIndex = 0
var stageIndex = 0

func init() {
	courses = append(courses,
		types.Course{Id: 1, Title: "Course 1", Status: types.Active, Stages: []types.Stage{
			{Id: 1, Title: "Stage 1", Type: types.Document, Order: 1, Desc: "First stage"},
			{Id: 2, Title: "Stage 2", Type: types.Presentation, Order: 2, Desc: "Second stage"},
			{Id: 3, Title: "Stage 3", Type: types.Video, Order: 3, Desc: "Third stage"},
			{Id: 4, Title: "Stage 4", Type: types.Test, Order: 4, Desc: "Fourth stage"},
		}},
		types.Course{Id: 2, Title: "Course 2", Status: types.Archived, Stages: []types.Stage{}},
	)
	courseIndex = len(courses)
	for i := 0; i < len(courses); i++ {
		stageIndex += len(courses[i].Stages)
	}
}

func GetCourse() []types.Course {
	return courses
}

func FindCourse(id int) (types.Course, error) {
	idx := func() int {
		for i := 0; i < len(courses); i++ {
			if courses[i].Id == id {
				return i
			}
		}
		return -1
	}()

	if idx == -1 {
		return types.Course{}, fmt.Errorf("курс с таким id %d не найден", id)
	}

	return courses[idx], nil
}

func AddCourse(c CourseReqDto) (types.Course, error) {

	courseIndex++

	stages := utils.Map(c.Stages, func(s StageReqDto) types.Stage {
		return createStage(s)
	})

	status, err := types.MapToCourseStatus(c.Status)

	if err != nil {
		return types.Course{}, err
	}

	if !types.ValidateStages(stages) {
		status = int(types.Draft)
	}

	course := types.Course{
		Id:     courseIndex,
		Title:  c.Title,
		Status: types.Status(status),
		Stages: utils.Map(c.Stages, func(s StageReqDto) types.Stage {
			return createStage(s)
		}),
	}

	courses = append(courses, course)
	return FindCourse(course.Id)
}

func createStage(s StageReqDto) types.Stage {

	status, _ := types.MapToStageType(s.Type)

	stageIndex++
	return types.Stage{
		Id:    stageIndex,
		Title: s.Title,
		Type:  types.StageType(status),
		Order: s.Order,
		Desc:  s.Desc,
	}
}

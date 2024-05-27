package repository

import (
	"fmt"
	"microservice_go/types"
	"microservice_go/utils"
)

type CourseResDto struct {
	Id     int           `json:"id"`
	Title  string        `json:"title"`
	Status string        `json:"status"`
	Stages []StageResDto `json:"stages"`
}

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

type StageResDto struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Type  string `json:"type"`
	Order int    `json:"order"`
	Desc  string `json:"desc"`
}

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

func GetCourse() []CourseResDto {
	res := utils.Map(courses, func(c types.Course) CourseResDto {
		return CourseResDto{
			Id:     c.Id,
			Title:  c.Title,
			Status: c.Status.String(),
			Stages: utils.Map(c.Stages, func(s types.Stage) StageResDto {

				return StageResDto{
					Id:    s.Id,
					Title: s.Title,
					Type:  s.Type.String(),
					Order: s.Order,
					Desc:  s.Desc,
				}
			}),
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
		Stages: utils.Map(courses[idx].Stages, func(s types.Stage) StageResDto {
			return StageResDto{
				Id:    s.Id,
				Title: s.Title,
				Type:  s.Type.String(),
				Order: s.Order,
				Desc:  s.Desc,
			}
		}),
	}
	return item, nil
}

func AddCourse(c CourseReqDto) (CourseResDto, error) {

	courseIndex++

	stages := utils.Map(c.Stages, func(s StageReqDto) types.Stage {
		return createStage(s)
	})

	status, err := types.MapToCourseStatus(c.Status)

	if err != nil {
		return CourseResDto{}, err
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

package mapper

import (
	"microservice_go/types"
	"microservice_go/utils"
)

func ToCourseResDto(course types.Course) CourseResDto {
	return CourseResDto{
		Id:     course.Id,
		Title:  course.Title,
		Status: course.Status.String(),
		Stages: utils.Map(course.Stages, func(s types.Stage) StageResDto {
			return StageResDto{
				Id:    s.Id,
				Title: s.Title,
				Type:  s.Type.String(),
				Order: s.Order,
				Desc:  s.Desc,
			}
		}),
	}
}

func ToCourseResDtoList(courses []types.Course) []CourseResDto {
	return utils.Map(courses, func(c types.Course) CourseResDto {
		return ToCourseResDto(c)
	})
}

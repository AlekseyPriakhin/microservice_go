package types

import "fmt"

type Status int

const (
	Active Status = iota
	Draft
	Archived
)

func (s Status) String() string {
	return [...]string{"Active", "Draft", "Archived"}[s]
}

func MapToCourseStatus(s string) (int, error) {

	switch s {
	case "Active":
		{
			return int(Active), nil
		}
	case "Draft":
		{
			return int(Draft), nil
		}
	case "Archived":
		{
			return int(Archived), nil
		}
	default:
		{
			return -1, fmt.Errorf("не удалось преобразовать %s в тип int", s)
		}
	}
}

type Course struct {
	Id     int     `json:"id"`
	Title  string  `json:"title"`
	Status Status  `json:"status"`
	Stages []Stage `json:"stages"`
}

func ValidateStages(s []Stage) bool {
	isValid := false

	if len(s) == 0 {
		return isValid
	}

	for i := 0; i < len(s); i++ {
		if s[i].Type == Test {
			isValid = true
		} else {
			isValid = false
		}
	}

	return isValid
}

type StageType int

const (
	Test StageType = iota
	Video
	Document
	Presentation
)

func (s StageType) String() string {
	return [...]string{"Test", "Video", "Document", "Presentation"}[s]
}

func MapToStageType(s string) (int, error) {
	switch s {
	case "Test":
		{
			return int(Test), nil
		}
	case "Video":
		{
			return int(Video), nil
		}
	case "Document":
		{
			return int(Document), nil
		}
	case "Presentation":
		{
			return int(Presentation), nil
		}
	default:
		{
			return -1, fmt.Errorf("не удалось преобразовать %s в тип int", s)
		}
	}
}

type Stage struct {
	Id    int       `json:"id"`
	Title string    `json:"title"`
	Type  StageType `json:"type"`
	Order int       `json:"order"`
	Desc  string    `json:"desc"`
}

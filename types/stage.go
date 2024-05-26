package types

import "fmt"

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

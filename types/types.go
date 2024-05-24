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
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Status Status `json:"status"`
}

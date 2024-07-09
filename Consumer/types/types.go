package types

type Status int

const (
	Active Status = iota
	Draft
	Archived
)

type Course struct {
	Id     int     `json:"id"`
	Title  string  `json:"title"`
	Status Status  `json:"status"`
	Stages []Stage `json:"stages"`
}

type StageType int

const (
	Test StageType = iota
	Video
	Document
	Presentation
)

type Stage struct {
	Id    int       `json:"id"`
	Title string    `json:"title"`
	Type  StageType `json:"type"`
	Order int       `json:"order"`
	Desc  string    `json:"desc"`
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

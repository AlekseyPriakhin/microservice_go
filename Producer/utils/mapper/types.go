package mapper

type CourseResDto struct {
	Id     int           `json:"id"`
	Title  string        `json:"title"`
	Status string        `json:"status"`
	Stages []StageResDto `json:"stages"`
}

type StageResDto struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Type  string `json:"type"`
	Order int    `json:"order"`
	Desc  string `json:"desc"`
}

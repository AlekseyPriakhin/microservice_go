package types

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

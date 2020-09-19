package rlog

type Entry struct {
	Title string `json:"title"`
	Slug  string `json:"slug"`
	Text  string `json:"text"`
}

func (e *Entry) String() string {
	return e.Title
}

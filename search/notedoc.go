package search

type NoteDoc struct {
	Path     string   `json:"path"`
	Headings []string `json:"headings"`
	Tags     []string `json:"tags"`
	Body     string   `json:"body"`
	Type     string   `json:"type"`
}

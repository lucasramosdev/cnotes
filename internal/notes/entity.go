package notes

type Note struct {
	ID          int64    `json:"id"`
	Category    string   `json:"category"`
	Theme       string   `json:"theme"`
	Title       string   `json:"title"`
	Summary     string   `json:"summary"`
	Keywords    []string `json:"Keywords"`
	Annotations []string `json:"annotations"`
}

type BasicNote struct {
	ID       int64  `json:"id"`
	Title    string `json:"title"`
	Theme    string `json:"theme"`
	Category string `json:"category"`
}

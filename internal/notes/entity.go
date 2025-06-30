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

type CreateNote struct {
	Category    string   `json:"category" binding:"required"`
	Theme       string   `json:"theme" binding:"required"`
	Title       string   `json:"title" binding:"required"`
	Summary     string   `json:"summary" binding:"required"`
	Keywords    []string `json:"Keywords"`
	Annotations []string `json:"annotations"`
}

type BasicNote struct {
	ID       int64  `json:"id"`
	Title    string `json:"title"`
	Theme    string `json:"theme"`
	Category string `json:"category"`
}

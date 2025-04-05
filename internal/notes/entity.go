package notes

type Note struct {
	ID       int64  `json:"id"`
	Category string `json:"category"`
	Theme    string `json:"theme"`
	Title    string `json:"title"`
	Summary  string `json:"summary"`
}

type BasicNote struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
}

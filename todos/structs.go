package todos

type RequestBody struct {
	Action string `json:"action"`
	Body   `json:"responsebody"`
}

type Body struct {
	Org        string `json:"org"`
	Title      string `json:"title"`
	Repo       string `json:"repo"`
	MarkDone   bool   `json:"markdone"`
	Identifier string `json:"identifier"`
}

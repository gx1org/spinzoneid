package app

type SpinPayload struct {
	Name      string `json:"name" binding:"required"`
	Options   string `json:"options" binding:"required"`
	InputHint string `json:"input_hint" binding:"required"`
	Comment   string `json:"comment" binding:""`
	Password  string `json:"password" binding:"required"`
	IsDelete  bool   `json:"is_delete" binding:""`
}

type HtmlResponse struct {
	Spin   Spin
	Result Result
}

package task

type SimpleTaskStruct struct {
	Title string `json:"title"`
	Body  string `json:"body"`
	Receivers string `json:"receivers"`
	DelayTime int `json:"delayTime"`
}

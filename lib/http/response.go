package http

type Response struct {
	Code uint32      `json:"c"`
	Msg  string      `json:"m"`
	Data interface{} `json:"d"`
}

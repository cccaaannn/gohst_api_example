package result

import (
	"encoding/json"

	"github.com/cccaaannn/gohst"
)

type Message struct {
	Message string `json:"message,omitempty"`
}

func (message Message) Json() string {
	json, _ := json.Marshal(message)
	return string(json)
}

func To404(res *gohst.Response) {
	res.StatusCode = 404
	res.Body = Message{Message: "Not Found"}.Json()
}

func To400(res *gohst.Response, message string) {
	res.StatusCode = 400
	res.Body = Message{Message: message}.Json()
}

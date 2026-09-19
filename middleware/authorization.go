package middleware

import (
	"encoding/json"
	"gohst_api_example/result"

	"github.com/cccaaannn/gohst"
)

func Authorization(next gohst.HandlerFunc) gohst.HandlerFunc {
	return func(req *gohst.Request, res *gohst.Response) {

		var token string = req.Context["token"].(string)

		if token != "banana" {
			res.StatusCode = 403
			json, _ := json.Marshal(result.Message{Message: "Forbidden"})
			res.Body = string(json)
			return
		}

		next(req, res)
	}
}

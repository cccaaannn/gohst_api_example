package middleware

import (
	"encoding/json"
	"gohst_api_example/result"

	"github.com/cccaaannn/gohst"
)

func Authentication(next gohst.HandlerFunc) gohst.HandlerFunc {
	return func(req *gohst.Request, res *gohst.Response) {

		var raw string = req.Headers["Authorization"]

		var bearer string = ""
		var token string = ""
		if len(raw) >= 7 {
			bearer = raw[:6]
			token = raw[7:]
		}

		if bearer != "Bearer" || token == "" {
			res.StatusCode = 401
			json, _ := json.Marshal(result.Message{Message: "Unauthorized"})
			res.Body = string(json)
			return
		}

		req.Context["token"] = token

		next(req, res)
	}
}

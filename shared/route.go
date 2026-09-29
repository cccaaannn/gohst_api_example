package shared

import (
	"gohst_api_example/result"

	"github.com/cccaaannn/gohst"
)

func CreateNotFoundRoute(server *gohst.Server) {
	server.AddHandler("/*", func(req *gohst.Request, res *gohst.Response) {
		result.To404(res)
		res.StatusCode = 404
	})
}

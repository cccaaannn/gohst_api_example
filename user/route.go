package user

import "github.com/cccaaannn/gohst"

func CreateUserRoutes(server *gohst.Server) {
	server.AddHandler("GET /users", Get)
	server.AddHandler("GET /users/:id", GetById)
	server.AddHandler("POST /users", Add)
	server.AddHandler("PUT /users/:id", Update)
	server.AddHandler("DELETE /users/:id", Delete)
}

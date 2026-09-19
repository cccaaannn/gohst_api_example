package user

import (
	"encoding/json"
	"gohst_api_example/result"
	"strconv"

	"github.com/cccaaannn/gohst"
)

func Get(req *gohst.Request, res *gohst.Response) {
	search := req.Query["search"]
	users := GetUsers(search)
	json, _ := json.Marshal(users)
	res.Body = string(json)
}

func GetById(req *gohst.Request, res *gohst.Response) {
	id := req.Params["id"]
	userId, err := strconv.Atoi(id)
	if err != nil {
		result.To400(res, "Invalid user id")
		return
	}

	user := GetUserById(userId)

	json, _ := json.Marshal(user)
	res.Body = string(json)
}

func Add(req *gohst.Request, res *gohst.Response) {
	user := User{}
	json.Unmarshal([]byte(req.Body), &user)
	user = AddUser(user)

	json, _ := json.Marshal(user)
	res.Body = string(json)
	res.StatusCode = 201
}

func Update(req *gohst.Request, res *gohst.Response) {
	id := req.Params["id"]
	userId, err := strconv.Atoi(id)
	if err != nil {
		result.To400(res, "Invalid user id")
		return
	}

	user := User{}
	json.Unmarshal([]byte(req.Body), &user)
	user.Id = userId

	user = UpdateUser(user)

	json, _ := json.Marshal(user)
	res.Body = string(json)
}

func Delete(req *gohst.Request, res *gohst.Response) {
	id := req.Params["id"]
	userId, err := strconv.Atoi(id)
	if err != nil {
		result.To400(res, "Invalid user id")
		return
	}

	DeleteUser(userId)

	res.Body = result.Message{Message: "User deleted"}.Json()
}

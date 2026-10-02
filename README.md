# gohst_api_example

Example of a simple REST API built with Go using the gohst library.

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/cccaaannn/gohst_api_example) ![GitHub top language](https://img.shields.io/github/languages/top/cccaaannn/gohst_api_example?color=008B8B&style=flat-square) ![GitHub repo size](https://img.shields.io/github/repo-size/cccaaannn/gohst_api_example?color=FF6347&style=flat-square) [![GitHub](https://img.shields.io/github/license/cccaaannn/gohst_api_example?color=green&style=flat-square)](https://github.com/cccaaannn/gohst_api_example/blob/master/LICENSE)

---

## Usage

```bash
go run main.go
```

## API

### Authoriazation is simulated with simple text.
- `Authorization: Bearer banana`

`GET /users`
```json
[
    {
        "id": 1,
        "name": "Can kurt",
        "age": 30
    },
    {
        "id": 2,
        "name": "Banana king",
        "age": 25
    }
]
```
`GET /users?search=melon`
```json
[
    {
        "id": 3,
        "name": "melon lover",
        "age": 12
    }
]
```
`GET /users/{id}`
```json
{
	"id": 1,
	"name": "Can kurt",
	"age": 30
}
```
`POST /users`
```json
{
	"name": "New User",
	"age": 22
}
```
`PUT /users/{id}`
```json
{
	"name": "Updated User",
	"age": 35
}
```
`DELETE /users/{id}`
```json
{
    "message": "User deleted"
}
```

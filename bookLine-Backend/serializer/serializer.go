package serializer

import "github.com/golang-jwt/jwt"

type Task struct {
	Title      string
	Message    string
	Status     int
	CreateTime int64
	EndTime    int64
}

type TaskRequest struct {
	Title      string
	Message    string
	Status     int
	CreateTime int64
	EndTime    int64
	Id         int
}

type TaskReturn struct {
	Items []TaskRequest `json:"items"`
	Total int           `json:"total" example:"0"`
}

type User struct {
	Id       string `json:"id" example:"-1"`
	Username string `json:"username" example:""`
	Password string `json:"password" example:""`
	IsAdmin  string `json:"isadmin" example:"0"`
}

type Novel struct {
	Id        string `json:"id"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	Status    int    `json:"status"`
	FilePath  string `json:"filePath"`
	CoverPath string `json:"coverPath"`
	Profile   string `json:"profile"`
	Category  string `json:"category"`
	Update    string `json:"update"`
	User      string `json:"user"`
}

type NovelFav struct {
	Item    Novel  `json:"item"`
	Favtime string `json:"favtime"`
}

type NovelReturn struct {
	Items interface{} `json:"items"`
	Total int         `json:"total" example:"0"`
}

type JWTClaims struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	IsAdmin  string `json:"isadmin" example:"0"`
	jwt.StandardClaims
}

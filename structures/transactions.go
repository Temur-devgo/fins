package structures

type Transactions struct {
	Id       int    `json:"id"`
	UserId   int    `json:"userId"`
	AuthorId int    `json:"authorId"`
	Type     int    `json:"type"`
	Sum      int    `json:"sum"`
	Date     string `json:"date"`
}

type PostForm struct {
	Sum      int `json:"sum"`
	ToUserId int `json:"toUserId"`
	UserId   int `json:"userId"`
}

type Response struct {
	Data      []Transactions    `json:"data"`
	ErrorText map[string]string `json:"errorText"`
}

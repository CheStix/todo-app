package todo

type User struct {
	Id       int `json:"-"`
	Name     int `json:"name"`
	Username int `json:"username"`
	Password int `json:"password"`
}

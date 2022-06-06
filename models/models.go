package models

type User struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Username string `json:"username"`
}

type UpdateBody struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Username string `json:"username,omitempty"`
}

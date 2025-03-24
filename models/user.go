package models


type User struct {
	ID   int    `json:"id" orm:"auto"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}


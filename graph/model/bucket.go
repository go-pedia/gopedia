package model

//Bucket Struct
type Bucket struct {
	ID      string  `json:"id"`
	User    User    `json:"users"`
	Product Product `json:"product"`
}

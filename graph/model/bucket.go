package model

//Bucket Struct
type Bucket struct {
	ID      string  `json:"id"`
	User    User    `json:"user"`
	Product Product `json:"product"`
}

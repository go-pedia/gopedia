package model

//Bucket Struct
type Bucket struct {
	ID      string `json:"id"`
	User    string `json:"users"`
	Product string `json:"product"`
}

//IsOwnerB this is to check are bucket user or not
func (b *Bucket) IsBucketOwner(user *User) bool {
	return b.User == user.ID
}

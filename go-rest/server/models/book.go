package models

//Book database model
type Book struct {
	Id          uint64 `json:"id" gorm:"primary_key"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Publisher   string `json:"publisher"`
	PublishDate string `json:"publish date"` //Format: YYYY-MM-DD
	Rating      uint8  `json:"rating"`       //1, 2 or 3
	CheckedOut  bool   `json:"checked out"`
}

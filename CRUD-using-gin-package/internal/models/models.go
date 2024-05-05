package internal 

type Details struct {
	Id int `gorm:"primarykey"`
	Name  string `json:"name"`
	Email string `json:"email" binding:"required"`
}
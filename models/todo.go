package models

type Todo struct {
	Id          int64  `gorm:"primaryKey" json:"id"`
	Title       string `gorm:"type:varchar(50)" json:"title"`
	Description string `gorm:"type:text" json:"description"`
	DueDate     string `gorm:"type:date" json:"due_date"`
}

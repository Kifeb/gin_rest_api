package models

type ProducModel struct {
	Id          int    `gorm:"PrimaryKey" json:"id"`
	Name        string `gorm:"type:varchar(100)" json:"name"`
	Description string `gorm:"type:text" json:"description"`
}

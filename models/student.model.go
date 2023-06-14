package models

type Student struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Name     string `json:"name"`
	Major_Id uint   `json:"major_id"`
	Majors   Majors `json:"majors" gorm:"foreignKey:major_id"`
}

type StudentInput struct {
	Name     string `json:"name"`
	Major_Id uint   `json:"major_id"`
}

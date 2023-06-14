package models

type Courses struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Name     string `json:"name"`
	Major_Id uint   `json:"major_id"`
	Majors   Majors `json:"majors" gorm:"foreignKey:major_id"`
}

type CourseInput struct {
	Name     string `json:"name"`
	Major_Id uint   `json:"major_id"`
}

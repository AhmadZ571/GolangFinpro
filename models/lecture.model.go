package models

type Lecture struct {
	ID        uint    `json:"id" gorm:"primaryKey"`
	Name      string  `json:"name"`
	Course_Id uint    `json:"course_id"`
	Courses   Courses `json:"courses" gorm:"foreignKey:course_id"`
}

type LectureInput struct {
	Name      string `json:"name"`
	Course_Id uint   `json:"course_id"`
}

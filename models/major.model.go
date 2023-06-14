package models

type Majors struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}

type MajorInput struct {
	Name string `json:"name"`
}

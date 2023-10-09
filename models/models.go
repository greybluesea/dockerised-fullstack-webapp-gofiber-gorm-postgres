package models

import "gorm.io/gorm"

type Fact struct {
	gorm.Model
	Question string `jason:"question" gorm:"text;not null;default:null"`
	Answer   string `jason:"answer"   gorm:"text;not null;default:null" `
}

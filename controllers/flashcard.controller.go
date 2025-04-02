package controllers

import (
	"gorm.io/gorm"
)

type FlashcardController struct {
	DB *gorm.DB
}

func NewFlashcardController(DB *gorm.DB) FlashcardController {
	return FlashcardController{DB}
}

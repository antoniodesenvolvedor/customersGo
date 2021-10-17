package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Customer struct {
	gorm.Model
	ID     uuid.UUID  `json:"id"`
	Name  string  `json:"name"`
	CPF string  `json:"cpf"`
	BirthDate  time.Time `json:"birthDate"`
}

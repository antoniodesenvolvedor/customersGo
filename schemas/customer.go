package schemas

import (
	"github.com/google/uuid"
	"time"
)

type Customer struct {
	Name  string  `json:"name"`
	CPF string  `json:"cpf"`
	BirthDate  string `json:"birthDate"`
}

type CustomerResponse struct {
	ID     uuid.UUID  `json:"ID"`
	Name  string  `json:"name"`
	CPF string  `json:"cpf"`
	BirthDate  time.Time `json:"birthDate"`
}

package pkg

import "github.com/google/uuid"

func OrderNumbering() string {
	s := uuid.New().String()
	return s[:8]
}

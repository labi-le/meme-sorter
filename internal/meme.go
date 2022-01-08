package internal

import (
	"database/sql"
	"gorm.io/gorm"
)

type DeletedAt sql.NullTime

type Meme struct {
	gorm.Model

	Fun   bool   `json:"Fun"`
	Image []byte `json:"Image"`
}

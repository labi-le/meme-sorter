package structures

import (
	"gorm.io/gorm"
)

type Meme struct {
	gorm.Model

	Fun  bool   `json:"fun"`
	Path string `json:"path"`
}

package internal

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"meme-sorter/internal/structures"
)

type DB struct {
	db *gorm.DB
}

func NewDB(c structures.Config) *DB {
	conn, err := gorm.Open(sqlite.Open(c.Dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return &DB{db: conn}
}

func (conn *DB) Create(item *structures.Meme) error {
	return conn.db.Create(&item).Error
}

func (conn *DB) Take(id uint, item *structures.Meme) error {
	return conn.db.Take(&item, id).Error
}
func (conn *DB) Update(item *structures.Meme) error {
	return conn.db.Updates(&item).Error
}
func (conn *DB) Delete(id uint) error {
	return conn.db.Delete(&structures.Meme{}, id).Error
}

func (conn *DB) Migrate() error {
	return conn.db.AutoMigrate(structures.Meme{})
}

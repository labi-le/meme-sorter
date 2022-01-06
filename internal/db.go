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

func (conn *DB) Create(item *structures.Meme) *gorm.DB {
	return conn.db.Create(&item)
}

func (conn *DB) Take(id uint) structures.Meme {
	var Item structures.Meme
	conn.db.Take(&Item, id)

	return Item
}
func (conn *DB) Update(item *structures.Meme) *gorm.DB {
	return conn.db.Updates(&item)
}
func (conn *DB) Delete(id uint) *gorm.DB {
	return conn.db.Delete(&structures.Meme{}, id)
}

func (conn *DB) Migrate() error {
	return conn.db.AutoMigrate(structures.Meme{})
}

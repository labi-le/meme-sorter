package internal

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DB struct {
	db *gorm.DB
}

func NewDB(c Config) *DB {
	conn, err := gorm.Open(sqlite.Open(c.Dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return &DB{db: conn}
}

func (conn *DB) Create(item *Meme) error {
	return conn.db.Create(&item).Error
}

func (conn *DB) Take(id uint, item *Meme) error {
	return conn.db.Take(&item, id).Error
}
func (conn *DB) Update(item *Meme) error {
	return conn.db.Updates(&item).Error
}
func (conn *DB) Delete(id uint) error {
	return conn.db.Delete(&Meme{}, id).Error
}

func (conn *DB) Migrate() error {
	return conn.db.AutoMigrate(Meme{})
}

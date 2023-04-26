package appdb

import (
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	UserID    int    `gorm:"primaryKey;autoIncrement"`
	LoginID   string `gorm:"unique"`
	Name      string
	Role      string
	Passwd    string
	Country   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func GetConnection() (*gorm.DB, error) {

	return gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
}

func InitDB() {
	db, err := GetConnection()
	if err == nil {
		db.AutoMigrate(&User{})
	}

}

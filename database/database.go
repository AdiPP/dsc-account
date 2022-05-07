package database

import (
	"github.com/AdiPP/dsc-account/entity"
	"github.com/AdiPP/dsc-account/mock"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Database struct {
	DB *gorm.DB
}

func NewDatabase() Database {
	db, err := gorm.Open(sqlite.Open("account.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		panic("failed to connect database")
	}

	return Database{
		DB: db,
	}
}

func (d *Database) Init() {
	d.Migrate()
	d.Seed()
}

func (d *Database) Migrate() {
	// Run migrations
	d.DB.AutoMigrate(&entity.User{}, &entity.Role{})
}

func (d *Database) Seed() {
	// Run seeders
	roles := mock.Roles

	for _, r := range roles {
		result := d.DB.First(&entity.Role{}, "id = ?", r.ID)

		if result.RowsAffected == 0 {
			d.DB.Create(&r)
		}
	}

	users := mock.Users

	for _, u := range users {
		result := d.DB.First(&entity.User{}, "id = ?", u.ID)

		if result.RowsAffected == 0 {
			d.DB.Create(&u)
		}
	}
}

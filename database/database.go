package database

import (
	"fmt"

	"github.com/AdiPP/dsc-account/entity"
	"github.com/AdiPP/dsc-account/mock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type PostgresSqlDatabase struct {
	DB *gorm.DB
}

func NewPostgresSqlDatabase() PostgresSqlDatabase {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn),
	})

	if err != nil {
		panic("failed to connect database")
	}

	return PostgresSqlDatabase{
		DB: db,
	}
}

func (d *PostgresSqlDatabase) Init() {
	d.Migrate()
	d.Seed()
}

func (d *PostgresSqlDatabase) Migrate() {
	// Run migrations
	d.DB.AutoMigrate(&entity.User{}, &entity.Role{})
}

func (d *PostgresSqlDatabase) Seed() {
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

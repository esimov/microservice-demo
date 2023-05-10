package model

import (
	"fmt"
	"time"

	"github.com/esimov/xm/config"
	"gorm.io/gorm"
)

// ID (uuid) required
// •  Name (15 characters) required - unique
// •  Description (3000 characters) optional
// •  Amount of Employees (int) required
// •  Registered (boolean) required
// •  Type (Corporations | NonProfit | Cooperative | Sole Proprietorship) required

type Company struct {
	gorm.Model
	ID          string    `gorm:"type:uuid;primaryKey"`
	Name        string    `gorm:"size:15;unique;not null"`
	Description string    `gorm:"size:3000"`
	Employees   int       `gorm:"not null"`
	Registered  bool      `gorm:"not null"`
	Type        OrgType   `gorm:"type:varchar(20);not null"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

type OrgType string

const (
	Corporations       OrgType = "Corporations"
	NonProfit          OrgType = "NonProfit"
	Cooperative        OrgType = "Cooperative"
	SoleProprietorship OrgType = "Sole Proprietorship"
)

func Create() {
	config := config.GetConfig(".env")
	fmt.Println(config)
	// viper.SetConfigName(".env") // name of config file
	// viper.AddConfigPath("/")    // path to look for the config file in
	// err := viper.ReadInConfig() // Find and read the config file
	// if err != nil {             // Handle errors reading the config file
	// 	panic(fmt.Errorf("fatal error config file: %w", err))
	// }

	// dbName, ok := viper.Get("MYSQL_DATABASE")
	// if !ok {
	// 	log.Fatalf("Wrong DB name")
	// }

	// userName, ok := viper.Get("MYSQL_USERNAME")
	// if !ok {
	// 	log.Fatalf("Wrong user name")
	// }

	// password, ok := viper.Get("MYSQL_PASSWORD")
	// if !ok {
	// 	log.Fatalf("Wrong password")
	// }

	// host, ok := viper.Get("MYSQL_HOSTNAME")
	// if !ok {
	// 	log.Fatalf("Wrong hostname")
	// }

	// port, ok := viper.Get("MYSQL_PORT")
	// if !ok {
	// 	log.Fatalf("Wrong port")
	// }

	// dsn := fmt.Sprint("user=%s password=%s dbname=%s host=%s port=%s", userName, password, dbName, host, port)

	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// if err != nil {
	// 	panic("Failed to connect to database")
	// }

	// AutoMigrate will create the table if it doesn't exist already
	//db.AutoMigrate(&Company{})
}

// func (c *Company) Create() {
// 	result := db.Create(&c)
// 	if result.Error != nil {
// 		panic("Failed to create table")
// 	}
// 	log.Println("Table created successfully!")
// }

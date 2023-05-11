package model

import (
	"errors"
	"fmt"
	"time"

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
	ID          string  `gorm:"primaryKey"`
	Name        string  `gorm:"size:15;unique;not null"`
	Description string  `gorm:"size:3000"`
	Employees   int     `gorm:"not null"`
	Registered  bool    `gorm:"not null"`
	Type        OrgType `gorm:"type:varchar(20);not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type OrgType string

const (
	Corporations       OrgType = "Corporations"
	NonProfit          OrgType = "NonProfit"
	Cooperative        OrgType = "Cooperative"
	SoleProprietorship OrgType = "Sole Proprietorship"
)

func Load(db *gorm.DB) error {
	//AutoMigrate will create the table if it doesn't exist already
	err := db.Debug().AutoMigrate(&Company{})
	if err != nil {
		return errors.New(fmt.Sprintf("cannot migrate table: %v", err))
	}
	return nil
}

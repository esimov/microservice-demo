package models

import (
	"time"

	"github.com/go-playground/validator/v10"
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
	ID          string    `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"size:15;unique;not null" json:"name"`
	Description string    `gorm:"size:3000" json:"description"`
	Employees   int       `gorm:"not null" json:"employees"`
	Registered  bool      `gorm:"not null" json:"registered"`
	Type        OrgType   `gorm:"type:varchar(20);not null" json:"type"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"created_at"`
}

type OrgType string

const (
	Corporations       OrgType = "Corporations"
	NonProfit          OrgType = "NonProfit"
	Cooperative        OrgType = "Cooperative"
	SoleProprietorship OrgType = "Sole Proprietorship"
)

func (c *Company) Save(db *gorm.DB) error {
	// Create a new validator instance
	validate := validator.New()

	// Validate the struct fields
	if err := validate.Struct(c); err != nil {
		return err
	}

	// Save the company to the database
	result := db.Create(c)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (c *Company) FindById(db *gorm.DB, pid uint64) error {
	var err error
	err = db.Debug().Model(&Company{}).Where("id = ?", pid).Take(&c).Error
	if err != nil {
		return err
	}
	return nil
}

func (c *Company) Update(db *gorm.DB) error {
	var err error
	err = db.Debug().Model(&Company{}).Where("id = ?", c.ID).Updates(
		Company{
			Name:        c.Name,
			Description: c.Description,
			Employees:   c.Employees,
			Registered:  c.Registered,
			Type:        c.Type,
			UpdatedAt:   time.Now(),
		}).Error
	if err != nil {
		return err
	}
	return nil
}

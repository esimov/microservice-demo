package models

import (
	"errors"
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
	ID          uint32    `gorm:"primaryKey;auto_increment" json:"id"`
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

func (c *Company) FindAll(db *gorm.DB) (*[]Company, error) {
	var err error

	companies := []Company{}
	err = db.Debug().Model(&Company{}).Limit(100).Find(&companies).Error
	if err != nil {
		return &[]Company{}, err
	}
	return &companies, err
}

func (c *Company) FindById(db *gorm.DB, cid uint64) (*Company, error) {
	var err error
	err = db.Debug().Model(&Company{}).Where("id = ?", cid).Take(&c).Error
	if err != nil {
		return nil, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("Company Not Found")
	}

	return c, nil
}

func (c *Company) Update(db *gorm.DB, cid uint64) error {
	var err error
	err = db.Debug().Model(&Company{}).Where("id = ?", cid).Updates(
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

func (c *Company) Delete(db *gorm.DB, cid uint64) (int64, error) {
	db = db.Debug().Model(&Company{}).Unscoped().Where("id = ?", cid).Take(&Company{}).Delete(&Company{})
	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

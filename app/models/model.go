package models

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

func Load(db *gorm.DB) error {
	//AutoMigrate will create the table if it doesn't exist already
	err := db.Debug().AutoMigrate(&Company{}, &User{})
	if err != nil {
		return errors.New(fmt.Sprintf("cannot migrate table: %v", err))
	}
	return nil
}

package models

import (
	"time"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Email     string    `gorm:"size:100;not null;unique" json:"email"`
	Password  string    `gorm:"size:100;not null;" json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func (u *User) VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (u *User) BeforeSave() error {
	hashedPassword, err := Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) Save(db *gorm.DB) error {
	// Create a new validator instance
	validate := validator.New()

	// Validate the struct fields
	if err := validate.Struct(u); err != nil {
		return err
	}

	// Save the user to the database
	result := db.Create(u)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (u *User) FindById(db *gorm.DB, pid uint64) error {
	var err error
	err = db.Debug().Model(&User{}).Where("id = ?", pid).Take(&u).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *User) FindByEmail(db *gorm.DB, email string) error {
	var err error
	err = db.Debug().Model(&User{}).Where("email = ?", email).Take(&u).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *User) Update(db *gorm.DB) error {
	var err error
	err = db.Debug().Model(&User{}).Where("id = ?", u.ID).Updates(
		User{
			Email:     u.Email,
			Password:  u.Password,
			UpdatedAt: time.Now(),
		}).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *User) Delete(db *gorm.DB, uid uint32) (int64, error) {
	db = db.Debug().Model(&User{}).Where("id = ?", u.ID).Take(&User{}).Delete(&User{})
	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

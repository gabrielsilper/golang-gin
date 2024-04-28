package models

import "gopkg.in/validator.v2"

type User struct {
	ID    int    `gorm:"primaryKey" json:"id"`
	Name  string `gorm:"not null" json:"name" validate:"nonzero,min=3"`
	Email string `gorm:"not null;unique" json:"email" validate:"nonzero,regexp=^\\w+([-+.']\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*$"`
}

func ValidateUser(user *User) error {
	err := validator.Validate(user)
	return err
}

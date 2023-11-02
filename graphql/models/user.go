package models

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func (s *Conn) Signup(input NewUser) (NewUser, error) {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return NewUser{}, fmt.Errorf("generating password hash: %w", err)
	}
	u := NewUser{
		Name:     input.Name,
		Email:    input.Email,
		Password: string(hashedPass),
	}
	err = s.db.Create(&u).Error
	if err != nil {
		return NewUser{}, err
	}

	// Successfully created the record, return the user.
	return u, nil

}

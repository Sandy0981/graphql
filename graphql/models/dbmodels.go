package models

import (
	"gorm.io/gorm"
)

type NewCompany struct {
	gorm.Model
	CompanyName string `json:"company_name"`
	FoundedYear string `json:"founded_year"`
	Location    string `json:"location"`
}

type NewJob struct {
	gorm.Model
	Title              string `json:"title"`
	ExperienceRequired string `json:"experience_required"`
	CompanyID          string `json:"company_id"`
}

type Company struct {
	ID          string `json:"ID"`
	CompanyName string `json:"company_name"`
	FoundedYear string `json:"founded_year"`
	Location    string `json:"location"`
	Jobs        []*Job `json:"jobs"`
}

type Job struct {
	ID                 string `json:"ID"`
	Title              string `json:"title"`
	ExperienceRequired string `json:"experience_required"`
	CompanyID          string `json:"company_id"`
}

type NewUser struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	ID           string `json:"ID"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
}

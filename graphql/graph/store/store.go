package store

import (
	"graphql/graph/model"
)

type Storer interface {
	CreateCompany(company model.NewCompany) (*model.Company, error)
	CreateJob(job model.NewJob) (*model.Job, error)
	ViewAllCompanies() ([]*model.Company, error)
	ViewAllJobs() ([]*model.Job, error)
	FindCompanyByID(companyID string) (*model.Company, error)
	FindJobByJobID(jobID string) (*model.Job, error)
	FindJobByCompanyID(companyID string) ([]*model.Job, error)
	Signup(input model.NewUser) (*model.User, error)
}

type Store struct {
	Storer
}

func NewStore(storer Storer) Store {
	return Store{Storer: storer}
}

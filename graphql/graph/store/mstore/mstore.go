package mstore

import (
	"fmt"
	"graphql/graph/model"
	"graphql/models"
)

type Service struct {
	C models.Conn
}

func NewService(ms *models.Conn) Service {
	return Service{
		C: *ms,
	}
}

func (s *Service) CreateCompany(cd model.NewCompany) (*model.Company, error) {
	fmt.Println("creating company in database") // Displaying a message to signify performing database manipulation
	compData := models.NewCompany{
		CompanyName: cd.CompanyName,
		FoundedYear: cd.FoundedYear,
		Location:    cd.Location,
	}
	comp, err := s.C.CreateCompany(compData)
	if err != nil {
		return nil, err
	}
	return &model.Company{
		CompanyName: comp.CompanyName,
		FoundedYear: comp.FoundedYear,
		Location:    comp.Location,
		Jobs:        nil,
	}, nil
}

func (s *Service) CreateJob(ni model.NewJob) (*model.Job, error) {
	fmt.Println("creating job in database")
	jobData := models.NewJob{
		Title:              ni.Title,
		ExperienceRequired: ni.ExperienceRequired,
		CompanyID:          ni.CompanyID,
	}

	job, err := s.C.CreateJob(jobData)
	if err != nil {
		return nil, err
	}
	return &model.Job{
		Title:              job.Title,
		ExperienceRequired: job.ExperienceRequired,
		CompanyID:          job.CompanyID,
	}, nil
}

func (s *Service) ViewAllCompanies() ([]*model.Company, error) {
	cmp, err := s.C.ViewAllCompanies()
	if err != nil {
		return nil, err
	}
	var result []*model.Company

	for _, c := range cmp {
		result = append(result, &model.Company{
			ID:          c.ID,
			CompanyName: c.CompanyName,
			FoundedYear: c.FoundedYear,
			Location:    c.Location,
		})
	}
	return result, nil
}

func (s *Service) ViewAllJobs() ([]*model.Job, error) {
	job, err := s.C.ViewAllJobs()
	if err != nil {
		return nil, err
	}
	var result []*model.Job
	for _, c := range job {
		result = append(result, &model.Job{
			ID:                 c.ID,
			Title:              c.Title,
			ExperienceRequired: c.ExperienceRequired,
			CompanyID:          c.CompanyID,
		})
	}
	return result, nil
}

func (s *Service) FindCompanyByID(companyID string) (*model.Company, error) {
	cmp, err := s.C.FindCompanyByID(companyID)
	if err != nil {
		return nil, err
	}
	return &model.Company{
		ID:          cmp.ID,
		CompanyName: cmp.CompanyName,
		FoundedYear: cmp.FoundedYear,
		Location:    cmp.Location,
	}, nil
}

func (s *Service) FindJobByJobID(jobID string) (*model.Job, error) {
	job, err := s.C.FindJobByJobID(jobID)
	if err != nil {
		return nil, err
	}
	return &model.Job{
		ID:                 job.ID,
		Title:              job.Title,
		ExperienceRequired: job.ExperienceRequired,
		CompanyID:          job.CompanyID,
	}, nil
}

func (s *Service) FindJobByCompanyID(companyID string) ([]*model.Job, error) {
	job, err := s.C.FindJobByCompanyID(companyID)
	if err != nil {
		return nil, err
	}
	var result []*model.Job

	for _, c := range job {
		result = append(result, &model.Job{
			ID:                 c.ID,
			Title:              c.Title,
			ExperienceRequired: c.ExperienceRequired,
			CompanyID:          c.CompanyID,
		})
	}
	return result, nil
}

func (s *Service) Signup(input model.NewUser) (*model.User, error) {
	fmt.Println("signup is in progress")
	userData := models.NewUser{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
	}
	user, err := s.C.Signup(userData)
	if err != nil {
		return nil, err
	}
	return &model.User{
		Name:         user.Name,
		Email:        user.Email,
		PasswordHash: user.Password,
	}, nil
}

package models

func (s *Conn) AutoMigrate() error {
	err := s.db.Migrator().AutoMigrate(&NewCompany{}, &NewJob{}, &NewUser{})
	if err != nil {
		// If there is an error while migrating, log the error message and stop the program
		return err
	}
	s.db.AutoMigrate(&NewCompany{}, &NewJob{}, &NewUser{})

	// Add foreign key constraint

	return nil
}

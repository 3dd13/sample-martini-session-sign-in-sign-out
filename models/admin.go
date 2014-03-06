package models

import (
	"github.com/eaigner/hood"
)

// Admin can be any struct that represents a admin in my system

type Admin struct {
  Id            hood.Id `sql:"pk" form:"-"`
	Username      string  `form:"username"`
	Password      string  `form:"password"`
	Authenticated bool    `sql:"-"`

  Created       hood.Created
  Updated       hood.Updated
}

// Login will preform any actions that are required to make a admin model
// officially authenticated.
func (admin *Admin) Login() {
	// Update last login time
	// Add to logged-in admin's list
	// etc ...
	admin.Authenticated = true
}

// Logout will preform any actions that are required to completely
// logout a admin.
func (admin *Admin) Logout() {
	// Remove from logged-in admin's list
	// etc ...
	admin.Authenticated = false
}

func (admin Admin) IsAuthenticated() bool {
	return admin.Authenticated
}

func (admin Admin) GetId() hood.Id {
	return admin.Id
}

// GetById will populate a admin object from a database model with
// a matching id.
func (admin *Admin) GetById(id interface{}) error {
	var results []Admin
	err := GetHood().Where("id", "=", id).Limit(1).Find(&results)
	if err != nil {
		return err
	} else if len(results) == 1 {
		admin.Id = results[0].Id
		admin.Username = results[0].Username
		admin.Created = results[0].Created
		admin.Updated = results[0].Updated
	}

	return nil
}

func (admin *Admin) FindByUsernameAndPassword(username string, password string) error {
	var results []Admin
	err := GetHood().Where("username", "=", username).And("password", "=", password).Limit(1).Find(&results)

	if err != nil {
		return err
	} else if len(results) == 1 {
		admin.Id = results[0].Id
		admin.Username = results[0].Username
		admin.Created = results[0].Created
		admin.Updated = results[0].Updated
	}

	return nil
}

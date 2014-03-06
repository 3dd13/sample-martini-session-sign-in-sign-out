package models

import (
	"github.com/eaigner/hood"
)

// User can be any struct that represents a user in my system
type Student struct {
	Id            hood.Id  `sql:"pk" form:"-"`
	Username      string `form:"username"`
	Password      string `form:"password"`
	Authenticated bool   `sql:"-"`

	Created       hood.Created
	Updated       hood.Updated
}

// Login will preform any actions that are required to make a user model
// officially authenticated.
func (student *Student) Login() {
	// Update last login time
	// Add to logged-in user's list
	// etc ...
	student.Authenticated = true
}

// Logout will preform any actions that are required to completely
// logout a user.
func (student *Student) Logout() {
	// Remove from logged-in user's list
	// etc ...
	student.Authenticated = false
}

func (student Student) IsAuthenticated() bool {
	return student.Authenticated
}

func (student Student) GetId() hood.Id {
	return student.Id
}

// GetById will populate a student object from a database model with
// a matching id.
func (student Student) GetById(id interface{}) error {
	var results []Student
	err := GetHood().Where("id", "=", id).Limit(1).Find(&results)
	if err != nil {
		return err
	} else if len(results) == 1 {
		student.Id = results[0].Id
		student.Username = results[0].Username
		student.Created = results[0].Created
		student.Updated = results[0].Updated
	}

	return nil
}

func (student *Student) FindByUsernameAndPassword(username string, password string) error {
	var results []Student
	err := GetHood().Where("username", "=", username).And("password", "=", password).Limit(1).Find(&results)
	if err != nil {
		return err
	} else if len(results) == 1 {
		student.Id = results[0].Id
		student.Username = results[0].Username
		student.Created = results[0].Created
		student.Updated = results[0].Updated
	}

	return nil
}

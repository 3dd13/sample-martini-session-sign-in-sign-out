package db

import (
	"github.com/eaigner/hood"
)

type Student struct {
	Id       hood.Id
	Username string
	Password string
	Created  hood.Created
	Updated  hood.Updated
}

func (table *Student) Indexes(indexes *hood.Indexes) {
	indexes.AddUnique("student_unique_username", "username")
}

type Admin struct {
	Id       hood.Id
	Username string
	Password string
	Created  hood.Created
	Updated  hood.Updated
}

func (table *Admin) Indexes(indexes *hood.Indexes) {
	indexes.AddUnique("admin_unique_username", "username")
}

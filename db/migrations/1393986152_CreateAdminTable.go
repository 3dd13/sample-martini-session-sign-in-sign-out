package main

import (
	"github.com/eaigner/hood"
)

type Admin struct {
	Id            hood.Id
	Username      string
	Password      string

	Created       hood.Created
	Updated       hood.Updated
}

func (m *M) CreateAdminsTable_1393986152_Up(hd *hood.Hood) {
	hd.CreateTable(&Admin{})
}

func (m *M) CreateAdminTable_1393986152_Down(hd *hood.Hood) {
	hd.DropTable(&Admin{})
}

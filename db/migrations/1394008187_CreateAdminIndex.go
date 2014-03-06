package main

import (
	"github.com/eaigner/hood"
)

func (m *M) CreateAdminIndex_1394008187_Up(hd *hood.Hood) {
	hd.CreateIndex("admin", "admin_unique_username", true, "username")
}

func (m *M) CreateAdminIndex_1394008187_Down(hd *hood.Hood) {
	// TODO: implement
	hd.DropIndex("admin", "admin_unique_username")
}

package main

import (
	"github.com/eaigner/hood"
)

type Student struct {
	Id            hood.Id
	Username      string
	Password      string

	Created       hood.Created
	Updated       hood.Updated
}

func (m *M) CreateStudentTable_1393986146_Up(hd *hood.Hood) {
  hd.CreateTable(&Student{})
}

func (m *M) CreateStudentTable_1393986146_Down(hd *hood.Hood) {
	hd.DropTable(&Student{})
}

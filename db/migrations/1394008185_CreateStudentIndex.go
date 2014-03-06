package main

import (
	"github.com/eaigner/hood"
)

func (m *M) CreateStudentIndex_1394008185_Up(hd *hood.Hood) {
  hd.CreateIndex("student", "student_unique_username", true, "username")
}

func (m *M) CreateStudentIndex_1394008185_Down(hd *hood.Hood) {
	// TODO: implement
  hd.DropIndex("student", "student_unique_username")
}

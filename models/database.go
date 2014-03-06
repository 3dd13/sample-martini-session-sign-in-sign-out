package models

import (
  "github.com/eaigner/hood"
  "log"
)

import (
  _ "github.com/lib/pq"
)

func InitDb() {
/*  hd := GetHood()

  InsertDefaultStudent(hd)
  InsertDefaultAdmin(hd)*/
}

func GetHood() *hood.Hood {
  hd, err := hood.Open("postgres", "dbname=first_go_dev host=localhost sslmode=disable")
  if err != nil {
    log.Fatalln("Fail to connect postgresql", err)
  }

  return hd
}

func InsertDefaultStudent(hd *hood.Hood) {
  student := Student{Username: "testuser", Password: "password"}
  _, err := hd.Save(&student)
  if err != nil {
    log.Fatalln("Could not insert test student", err)
  }
}

func InsertDefaultAdmin(hd *hood.Hood) {
  admin := Admin{Username: "testadmin", Password: "password!"}
  _, err := hd.Save(&admin)
  if err != nil {
    log.Fatalln("Could not insert test admin", err)
  }
}

package sessionauth

import (
	"github.com/codegangsta/martini"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
	"github.com/eaigner/hood"

  "log"
	"net/http"
	"fmt"
	"../models"
)

var (
	RedirectStudentUrl = "/students/sign_in"
	RedirectStudentParams = "redirect-to"
	studentSessionKey = "student-session"
)

type StudentUser interface {
	IsAuthenticated() bool
	GetById(id interface{}) error
	Login()
	Logout()
	GetId() hood.Id
}

func GenerateAnonymousStudentUser() StudentUser {
	return &(models.Student{})
}

func SessionStudent() martini.Handler {
	return func(s sessions.Session, c martini.Context, l *log.Logger) {
		studentId := s.Get(studentSessionKey)
		student := GenerateAnonymousStudentUser()
		if studentId != nil {
			err := student.GetById(studentId)
			if err != nil {
				l.Printf("Login Error: %v\n", err)
			} else {
				student.Login()
			}
		}

		c.MapTo(student, (*StudentUser)(nil))
	}
}

func StudentLogin(s sessions.Session, student StudentUser) error {
	student.Login()
	return updateStudent(s, student)
}

func StudentLogout(s sessions.Session, student StudentUser) {
	student.Logout()
	s.Delete(studentSessionKey)
}

func StudentLoginRequired(r render.Render, s sessions.Session, student StudentUser, req *http.Request) {
	if student.IsAuthenticated() == false {
		path := fmt.Sprintf("%s?%s=%s", RedirectStudentUrl, RedirectStudentParams, req.URL.Path)
		r.Redirect(path, 302)
	}
}

func updateStudent(s sessions.Session, student StudentUser) error {
	s.Set(studentSessionKey, int64(student.GetId()))
	return nil
}

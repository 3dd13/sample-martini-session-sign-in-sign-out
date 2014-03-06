package sessionauth

import (
  "fmt"
  "github.com/codegangsta/martini"
  "github.com/martini-contrib/render"
  "github.com/martini-contrib/sessions"
  "github.com/eaigner/hood"

  "log"
  "net/http"
  "../models"
)

var (
  RedirectAdminUrl = "/admins/sign_in"
  RedirectAdminParams = "redirect-to"
  adminSessionKey = "admin-session"
)

type AdminUser interface {
  IsAuthenticated() bool
  GetById(id interface{}) error
  Login()
  Logout()
  GetId() hood.Id
}

func GenerateAnonymousAdminUser() AdminUser {
  return &models.Admin{}
}

func SessionAdmin() martini.Handler {
  return func(s sessions.Session, c martini.Context, l *log.Logger) {
    adminId := s.Get(adminSessionKey)
    admin := GenerateAnonymousAdminUser()
    if adminId != nil {
      err := admin.GetById(adminId)
      if err != nil {
        l.Printf("Login Error: %v\n", err)
      } else {
        admin.Login()
      }
    }

    c.MapTo(admin, (*AdminUser)(nil))
  }
}

func AdminLogin(s sessions.Session, admin AdminUser) error {
  admin.Login()
  return updateAdmin(s, admin)
}

func AdminLogout(s sessions.Session, admin AdminUser) {
  admin.Logout()
  s.Delete(adminSessionKey)
}

func AdminLoginRequired(r render.Render, admin AdminUser, req *http.Request) {
  if admin.IsAuthenticated() == false {
    path := fmt.Sprintf("%s?%s=%s", RedirectAdminUrl, RedirectAdminParams, req.URL.Path)
    r.Redirect(path, 302)
  }
}

func updateAdmin(s sessions.Session, admin AdminUser) error {
  s.Set(adminSessionKey, int64(admin.GetId()))
  return nil
}

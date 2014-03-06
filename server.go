package main

import(
  "github.com/codegangsta/martini"

  "github.com/martini-contrib/render"
/*  "github.com/martini-contrib/gzip"*/
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/sessions"

	"net/http"

  "./models"
  "./sessionauth"
)

func main() {
  store := sessions.NewCookieStore([]byte("secret-session-12345678"))
	models.InitDb()

	store.Options(sessions.Options{
		MaxAge: 0,
	})

  m := martini.Classic()
/*  m.Use(gzip.All())*/
  m.Use(render.Renderer(render.Options{
    // Directory: "templates", // Specify what path to load the templates from.
    Layout: "layout", // Specify a layout template. Layouts can call {{ yield }} to render the current template.
    // Extensions: []string{".tmpl", ".html"}, // Specify extensions to load for templates.
    // Funcs: []template.FuncMap{AppHelpers}, // Specify helper function maps for templates to access.
    // Delims: render.Delims{"{[{", "}]}"}, // Sets delimiters to the specified strings.
    Charset: "UTF-8", // Sets encoding for json and html content-types. Default is "UTF-8".
    IndentJSON: true, // Output human readable JSON
  }))
	m.Use(sessions.Sessions("my_session", store))
	m.Use(sessionauth.SessionStudent())
  m.Use(sessionauth.SessionAdmin())

  m.Get("/", func(r render.Render) {
    r.HTML(200, "home/index", nil)
  })

  setupAdminRoutes(m)
  setupStudentRoutes(m)

  m.Get("/api", func(r render.Render) {
    r.JSON(200, map[string]interface{}{"hello": "world"})
  })

  m.Run()
}

func setupAdminRoutes(m *martini.ClassicMartini) {
  m.Get("/admins/sign_in", func(r render.Render) {
    r.HTML(200, "admins/sessions/new", nil)
  })
  m.Post("/admins/sign_in", binding.Bind(models.Admin{}), func(session sessions.Session, postedAdmin models.Admin, r render.Render, req *http.Request) {
    admin := &models.Admin{}
    err := admin.FindByUsernameAndPassword(postedAdmin.Username, postedAdmin.Password)
    if err != nil {
      r.Redirect(sessionauth.RedirectAdminUrl)
      return
    } else {
      err := sessionauth.AdminLogin(session, admin)
      if err != nil {
        r.JSON(500, err)
      } else {

      }

      params := req.URL.Query()
      redirect := params.Get(sessionauth.RedirectAdminParams)
      r.Redirect(redirect)
      return
    }
  })
  m.Get("/admins/sign_out", sessionauth.AdminLoginRequired, func(session sessions.Session, admin sessionauth.AdminUser, r render.Render) {
    sessionauth.AdminLogout(session, admin)
    r.Redirect("/")
  })
  m.Get("/admins", sessionauth.AdminLoginRequired, func(r render.Render, admin sessionauth.AdminUser) {
    r.HTML(200, "admins/dashboard/index", admin.(*models.Admin))
  })
}

func setupStudentRoutes(m *martini.ClassicMartini) {
  m.Get("/students/sign_in", func(r render.Render) {
    r.HTML(200, "students/sessions/new", nil)
  })
  m.Post("/students/sign_in", binding.Bind(models.Student{}), func(session sessions.Session, postedStudent models.Student, r render.Render, req *http.Request) {
    student := &models.Student{}
    err := student.FindByUsernameAndPassword(postedStudent.Username, postedStudent.Password)
    if err != nil {
      r.Redirect(sessionauth.RedirectStudentUrl)
      return
    } else {
      err := sessionauth.StudentLogin(session, student)
      if err != nil {
        r.JSON(500, err)
      }

      params := req.URL.Query()
      redirect := params.Get(sessionauth.RedirectStudentParams)
      r.Redirect(redirect)
      return
    }
  })
  m.Get("/students/sign_out", sessionauth.StudentLoginRequired, func(session sessions.Session, student sessionauth.StudentUser, r render.Render) {
    sessionauth.StudentLogout(session, student)
    r.Redirect("/")
  })
  m.Get("/students", sessionauth.StudentLoginRequired, func(r render.Render, student sessionauth.StudentUser) {
    r.HTML(200, "students/dashboard/index", student.(*models.Student))
  })
}

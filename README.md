Sample Project for Go with Martini
====================================

was trying to have two separated routes protected by two different session model:

* /students is protected by student session
* /admins is protected by admin session


Reference / Credit
======================

was started with [https://github.com/martini-contrib/sessionauth](https://github.com/martini-contrib/sessionauth)
but couldn't get the multi-scope session working because of the c.MapTo(user, (*User)(nil))

So, I copied most of the code from sessionauth and make the injection independent from sessionauth.User.
Will think about refactoring the duplicated code in admin.go and student.go if more people find it useful.


Setting up the database for the sample
======================

using [hood](https://github.com/eaigner/hood) to handle the ORM and database versioning.
If you are familiar with Rails, you should be able to understand this easily:

    hood db:migrate


Starting the Go server
======================

you can always start with go run:

    go run server.go

but I prefer fresh:

    fresh

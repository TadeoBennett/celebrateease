package main

import (
	"errors"
	"log"
	"net/http"
	"regexp"
	"strings"
	"tadeobennett/celebrateease/pkg/models"
	"tadeobennett/celebrateease/templates"
	"text/template"
	"unicode/utf8"

	"github.com/justinas/nosurf"
)

func (app *Application) signup(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.ClientError(w, http.StatusBadRequest)
		return
	}
	firstname := strings.TrimSpace(r.PostForm.Get("firstname"))
	lastname := strings.TrimSpace(r.PostForm.Get("lastname"))
	email := strings.TrimSpace(r.PostForm.Get("email"))
	password := strings.TrimSpace(r.PostForm.Get("password"))
	repeatedPassword := strings.TrimSpace(r.PostForm.Get("repeatedPassword"))

	errors_user := make(map[string]string)

	//validating the firstname
	if firstname == "" {
		errors_user["firstname"] = "This field cannot be left blank"
	} else if utf8.RuneCountInString(firstname) > 30 {
		errors_user["firstname"] = "This field is too long"
	} else if utf8.RuneCountInString(firstname) < 3 {
		errors_user["firstname"] = "This field is too short"
	}

	//validating the lastname
	if lastname == "" {
		errors_user["lastname"] = "This field cannot be left blank"
	} else if utf8.RuneCountInString(firstname) > 30 {
		errors_user["lastname"] = "This field is too long"
	} else if utf8.RuneCountInString(firstname) < 3 {
		errors_user["lastname"] = "This field is too short"
	}

	//validating the email
	if strings.TrimSpace(email) == "" {
		errors_user["email"] = "This field cannot be left blank"
	} else if utf8.RuneCountInString(email) > 70 {
		errors_user["email"] = "This field is too long"
	}
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(email) {
		errors_user["email"] = "Invalid Email"
	}

	//validating the password
	if password == "" {
		errors_user["password"] = "This field cannot be left blank"
	} else if utf8.RuneCountInString(password) > 60 {
		errors_user["password"] = "This field is too long"
	} else if utf8.RuneCountInString(password) < 5 {
		errors_user["password"] = "This field is too short"
	}

	//checking the repeated password
	if password != repeatedPassword {
		errors_user["repeatedPassword"] = "The repeated password does not match"
	}

	if len(errors_user) > 0 { //an error exists
		ts, err := template.ParseFiles("../../views/pages/signup.tmpl",
		"../../views/partials/test_header.tmpl", 
		"../../views/partials/scripts.tmpl", 
		) //go back to signup page

		if err != nil { //error loading the template
			log.Println(err.Error())
			app.ServerError(w, err)
			return
		}

		err = ts.Execute(w, &templates.TemplateData{
			ErrorsFromForm:  errors_user,
			FormData:        r.PostForm,
			IsAuthenticated: app.IsAuthenticated(r),
			CSRFToken:       nosurf.Token(r),
		})
		if err != nil {
			panic(err.Error())
		}
		return
	}

	err = app.users.Insert(firstname, lastname, email, password)
	if err != nil {
		if errors.Is(err, models.ErrDuplicateEmail) {
			errors_user["email"] = "Email already exists in the system"
			ts, err := template.ParseFiles("../../views/pages/signup.tmpl",
			"../../views/partials/test_header.tmpl",
			"../../views/partials/scripts.tmpl",
			)
			if err != nil {
				app.ServerError(w, err)
				return
			}
			err = ts.Execute(w, &templates.TemplateData{
				ErrorsFromForm:  errors_user,
				FormData:        r.PostForm,
				IsAuthenticated: app.IsAuthenticated(r),
				CSRFToken:       nosurf.Token(r),
			})
			if err != nil {
				app.ServerError(w, err)
			}
			return
		} else {
			app.ServerError(w, err)
			return
		}
	}

	//set some Session data after a quote is added
	app.Session.Put(r, "flash", "User added. You can Login.")

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func (app *Application) signupForm(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("../../views/pages/signup.tmpl",
		"../../views/partials/test_header.tmpl",
		"../../views/partials/scripts.tmpl",
	)

	if err != nil {
		app.ServerError(w, err)
		return
	}

	err = ts.Execute(w, &templates.TemplateData{
		IsAuthenticated: app.IsAuthenticated(r),
		CSRFToken:       nosurf.Token(r),
	})

	if err != nil {
		app.ServerError(w, err)
	}
}

func (app *Application) login(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.ClientError(w, http.StatusBadRequest)
		return
	}
	//get these values from the form
	email := r.PostForm.Get("email")
	password := r.PostForm.Get("password")

	errors_user := make(map[string]string)
	id, err := app.users.Authenticate(email, password)
	if err != nil {
		log.Print("could not authenticate")
		if errors.Is(err, models.ErrInvalidCredentials) {
			errors_user["default"] = "Email or Password is Incorrect"
			ts, err := template.ParseFiles("../../views/pages/login.tmpl",
			"../../views/partials/test_header.tmpl", 
			"../../views/partials/scripts.tmpl", 
			) //load the template file
			if err != nil {
				log.Println(err.Error())
				app.ServerError(w, err)
				return
			}
			err = ts.Execute(w, &templates.TemplateData{
				ErrorsFromForm:  errors_user,
				FormData:        r.PostForm,
				IsAuthenticated: app.IsAuthenticated(r),
				CSRFToken:       nosurf.Token(r),
			})

			if err != nil {
				log.Panicln(err.Error())
				app.ServerError(w, err)
				return
			}
			return
		}
		return
	}
	app.Session.Put(r, "authenticatedUserId", id)
	app.Session.Put(r, "flash", "Welcome to CelebrateEase")
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (app *Application) loginForm(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("../../views/pages/login.tmpl",
		"../../views/partials/test_header.tmpl",
		"../../views/partials/scripts.tmpl",
	)
	if err != nil {
		app.ServerError(w, err)
		return
	}

	flash := app.Session.PopString(r, "flash")

	err = ts.Execute(w, &templates.TemplateData{
		Flash:           flash,
		IsAuthenticated: app.IsAuthenticated(r),
		CSRFToken:       nosurf.Token(r),
	})

	if err != nil {
		app.ServerError(w, err)
	}
}

func (app *Application) logout(w http.ResponseWriter, r *http.Request) {
	app.Session.Remove(r, "authenticatedUserId")
	log.Println("logged out user")
	app.Session.Put(r, "flash", "You have logged out")
	http.Redirect(w, r, "/login", http.StatusSeeOther) //go to home when logged out
}

package controllers

import (
	"fmt"
	"net/http"

	"lenslocked.com/views"
)

// NewUsers are cool
func NewUsers() *Users {
	return &Users{
		NewView: views.NewView("bootstrap", "users/new"),
	}
}

// Users is cool
type Users struct {
	NewView *views.View
}

// New is used to render the form where a user can
// create a new user account.
//
// GET /signup
func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	if err := u.NewView.Render(w, nil); err != nil {
		panic(err)
	}
}

// SignupForm rules!
type SignupForm struct {
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

// Create new user. sumbit the signup form. /POST signup
func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	var form SignupForm
	if err := parseForm(r, &form); err != nil {
		panic(err)
	}
	fmt.Fprintln(w, "Email is", form.Email)
	fmt.Fprintln(w, "Password is", form.Password)
}

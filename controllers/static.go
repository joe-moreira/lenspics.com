package controllers

import "lenslocked.com/views"

//NewStatic rules
func NewStatic() *Static {
	return &Static{
		Home: views.NewView(
			"bootstrap", "views/static/home.gohtml"),
		Contact: views.NewView(
			"bootstrap", "views/static/contact.gohtml"),
	}
}

// Static is a new static
type Static struct {
	Home    *views.View
	Contact *views.View
}

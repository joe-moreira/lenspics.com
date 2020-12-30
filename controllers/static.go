package controllers

import "lenslocked.com/views"

//NewStatic rules
func NewStatic() *Static {
	return &Static{
		Home: views.NewView(
			"bootstrap", "static/home"),
		Contact: views.NewView(
			"bootstrap", "static/contact"),
	}
}

// Static is a new static
type Static struct {
	Home    *views.View
	Contact *views.View
}

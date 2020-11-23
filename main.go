package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"lenslocked.com/views"
)

var homeView *views.View
var contactView *views.View

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := homeView.Template.Execute(w, nil)
	if err != nil {
		panic(err)
	}
}
func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := contactView.Template.Execute(w, nil)
	if err != nil {
		panic(err)
	}
}

/*func faq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "Any question?.")
}
*/
/*func notfound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "Page Not Found!")
}
*/
func main() {
	//	var h http.Handler = http.HandlerFunc(notfound)
	//var err error
	homeView = views.NewView("views/home.gohtml")
	contactView = views.NewView("views/contact.gohtml")

	/*homeTemplate, err = template.ParseFiles(
		"views/home.gohtml",
		"views/layouts/footer.gohtml")
	if err != nil {
		panic(err)
	}
	contactTemplate, err = template.ParseFiles(
		"views/contact.gohtml",
		"views/layouts/footer.gohtml")
	if err != nil {
		panic(err)
	}*/
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	//	r.HandleFunc("/faq", faq)
	//	r.NotFoundHandler = h
	http.ListenAndServe(":3000", r)
}

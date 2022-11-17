package main

import (
	"fmt"
	"html/template"

	//functionality for dealing with HTTP requests
	"net/http"
)

type Rsvp struct {
	Name, Email, Phone string
	WillAttend         bool
}

// collect Rsvp values together.
// make initializes a new slice, last two arguments are initial size and initial capacity 0= empty slice
var responses = make([]*Rsvp, 0, 10)
var templates = make(map[string]*template.Template, 3)

// will be responsible for loading the HTML files defined in earlier listings
// and processing them to create the *template.Template values that will be stored in the map
func loadTemplates() {
	templateNames := [5]string{"welcome", "form", "thanks", "sorry", "list"}
	for index, name := range templateNames {
		t, err := template.ParseFiles("layout.html", name+".html")
		if err == nil {
			templates[name] = t
			fmt.Println("Loaded template", index, name)
		} else {
			panic(err)
		}
	}
}

//HTTP Handlers and Server
//user requests the default URL
//pointer to an instance of the Request struct. Request being processed.

func welcomeHandler(writer http.ResponseWriter, request *http.Request) {
	templates["welcome"].Execute(writer, nil)
}

func listHandler(writer http.ResponseWriter, request *http.Request) {
	templates["list"].Execute(writer, responses)
}

type formData struct {
	*Rsvp
	Errors []string
}

// checks the value of the request.Method field, which returns the type of HTTP request that has been received.
// For GET requests, the form template is executed.
// There is no data to use when responding to GET requests,
// but there is a need to provide the template with the expected data structure.
func formHandler(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		templates["form"].Execute(writer, formData{
			Rsvp: &Rsvp{}, Errors: []string{},
		})
	}
}

// entry point of the app
func main() {

	loadTemplates()
	http.HandleFunc("/", welcomeHandler)
	http.HandleFunc("/list", listHandler)
	http.HandleFunc("/form", formHandler)

	//create and http server that listens for requests on port 5000
	// nil is to tell the server that requests should be processed using funcs registered with HandleFunc
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		fmt.Println(err)
	}
}

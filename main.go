package main

import (
	"net/http"
	"fmt"
	"time"
	"html/template"
)
//Struct, simlar to DTO, information to be displayed in our HTML file
type Welcome struct {
	FirstName string
	LastName string
	Time string
}
//Go application entrypoint
func main() {
	//Instantiate a Welcome struct obj & pass in some rand info.
	//get name of the uer as a query pram from the url
	welcome := Welcome{"No first name", "No last Name", time.Now().Format(time.Stamp)}

	//We tell Go where it can find our html file. We as Go to parse the html file.

	templates := template.Must(template.ParseFiles("templates/welcome-template.html"))

	//Our HTML comes with CSS that go needs to provide when we run the app. 
	//Here we tell go to create a handle that looks in the static dir, go then
	//uses the "/static/" as a url that our html can refer to when looking for our css & other files.
	
	
	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("static")))) //Go looks in the relative static directory first, 
			//then matches it to a url of our choice as shown in http.handle("/static/")
			//THis url is what we need when referencing our css files.
		//Note: the final url can be whatever we like, as long as we are consistent.


	//Method takes in the url path / and a fucction that takes in a reponse writeer, and a http request.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		//Takes the name from the URL query. ie if name is Barb then, will set welcome.Name to Martin (from our struct)
		if firstname := r.FormValue("firstname"); firstname != "" {
			welcome.FirstName = firstname;
		}
		if lastname := r.FormValue("lastname"); lastname != "" {
			welcome.LastName = lastname;
		}
		//If errors show an internal server error message.
		if err := templates.ExecuteTemplate(w, "welcome-template.html", welcome); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}) 

	fmt.Println("Server Running on PORT 8080");
		//starting the web server, set the port and listening on 8080. 
		//WIthout a path it assumes localhost, print any errors from starting the webserver using fmt.
	fmt.Println(http.ListenAndServe(":8080", nil));
}
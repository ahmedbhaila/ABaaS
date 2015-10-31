package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hello World")
	
	router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/", Index)
    router.HandleFunc("/ahmedbhaila", AhmedBhaila)
    router.HandleFunc("/ahmedbhaila/age", AhmedBhailaAge)
    router.HandleFunc("/ahmedbhaila/email", AhmedBhailaEmail)
    router.HandleFunc("/ahmedbhaila/birthday", AhmedBhailaBirthday)
    router.HandleFunc("/ahmedbhaila/married", AhmedBhailaMarried)
    router.HandleFunc("/ahmedbhaila/kids", AhmedBhailaKids)
    //router.HandleFunc("/todos/{todoId}", TodoShow)

    log.Fatal(http.ListenAndServe(":8090", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Welcome! Start by appending /ahmedbhaila to this url")
}

func AhmedBhaila(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "You made it this far, now find all about me by appending \n\n - /age\n - /email\n - /birthday\n - /married\n - /kids")
}

func AhmedBhailaAge(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "I am 35 years old")
}

func AhmedBhailaEmail(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Reach me at ahmed.bhaila at gmail.com")
}

func AhmedBhailaBirthday(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "My Birthday is 14th December")
}

func AhmedBhailaMarried(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Yes, sorry ladies and/or germs")
}

func AhmedBhailaKids(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Yes, One 3 year old munchkin")
}
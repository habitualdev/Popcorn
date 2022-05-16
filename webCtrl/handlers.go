package webCtrl

import (
	"Popcorn/auth"
	"Popcorn/download"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Message struct {
	Message string
	Data string
}

var MessageBuffer = make(chan Message, 100)

func index(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	if _, err := r.Cookie("auth"); err != nil {
		page, _ := ioutil.ReadFile("static/index.html")
		returnString := string(page)
		fmt.Fprintf(w, returnString)
		return
	}

	page, _ := ioutil.ReadFile("static/index_auth.html")
	returnString := string(page)
	fmt.Fprintf(w, returnString)

}

func add(w http.ResponseWriter, r *http.Request) {

	body, _ := ioutil.ReadAll(r.Body)
	if r.Method != "POST" {
		fmt.Fprintf(w, "Method not allowed")
		return
	}

	if _, err := r.Cookie("auth"); err != nil {
		fmt.Fprintf(w, "Not logged in")
		return
	}

	authCookie , _ := r.Cookie("auth")


	decodedString, _ := base64.StdEncoding.DecodeString(authCookie.Value)
	splitAuth := strings.Split(string(decodedString), ":")

	user := auth.User{Username: splitAuth[0], Password: splitAuth[1]}

	err := auth.CheckValue(user)
	if err != nil {
		fmt.Fprintf(w,err.Error())
		return
	}

	if len(body) == 0 {
		w.Write([]byte("empty post"))
		return
	}

	MessageBuffer <- Message{Message: "add", Data: string(body)}
	go download.GetVideo(string(body))
	http.Redirect(w, r, "/", http.StatusFound)
}

func addUser(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		fmt.Fprintf(w, "Method not allowed")
		return
	}

	body, _ := ioutil.ReadAll(r.Body)

	decodedString, err := base64.StdEncoding.DecodeString(string(body))
	if err != nil {
		fmt.Fprintf(w,err.Error())
		return
	}

	splitAuth := strings.Split(string(decodedString), ":")

	masterPass, _ := ioutil.ReadFile("masterpassword.txt")

	if string(masterPass) != splitAuth[2] {
		fmt.Fprintf(w,"Wrong master password")
		return
	}

	user := auth.User{Username: splitAuth[0], Password: splitAuth[1]}

	auth.AddUser(user)
}

func login(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		fmt.Fprintf(w, "Method not allowed")
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")



	http.SetCookie(w, &http.Cookie{
		Name: "auth",
		Value: base64.StdEncoding.EncodeToString([]byte(username + ":" + password)),
		Path: "/",
		MaxAge: 3600,
	})
}

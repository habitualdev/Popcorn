package webCtrl

import "net/http"

func ServeHttp(){
	http.HandleFunc("/", index)
	http.HandleFunc("/adduser", addUser)
	http.HandleFunc("/add", add)
	http.HandleFunc("/login", login)


	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets/"))))
	http.ListenAndServe(":8080", nil)
}
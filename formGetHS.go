//	Code for YouTube Go Tutorial course Titled:
//	Advanced Enterprise Platform Coding with Go in Depth Programming
//	https://youtu.be/whL0Toym5u4

package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	webS := http.FileServer(http.Dir("."))
	http.Handle("/", webS)
	http.HandleFunc("/xyzForm", formHandler)

	http.ListenAndServe(":80", nil)
}

func formHandler(xz http.ResponseWriter, yz *http.Request) {
	yz.ParseForm()
	email := yz.FormValue("email")
	subject := yz.FormValue("subject")
	message := yz.FormValue("message")

	appendFileBlob, _ := os.OpenFile("formGetHS.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	appendFileBlob.Write([]byte(string(email + "\n" + subject + "\n" + message)))
	fmt.Fprintf(xz, "Thanks for your feedback, your message has been recorded:\n")
}

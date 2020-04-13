package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net"
	"net/http"
	"net/smtp"
)

//Request struct
type Request struct {
	from    string
	to      []string
	subject string
	body    string
}

func NewRequest(to []string, subject, body string) *Request {
	return &Request{
		to:      to,
		subject: subject,
		body:    body,
	}
}

func (r *Request) SendEmail() (bool, error) {
	servername := "smtp.gmail.com:587"
	host, _, _ := net.SplitHostPort(servername)
	log.Println("HOST:", host)
	auth := smtp.PlainAuth("", "youremail", "yourpassword", host)
	mime := "MIME-version: 1.0;\nContent-Type: text/plain; charset=\"UTF-8\";\n\n"

	subject := "Subject: " + r.subject + "!\n"
	msg := []byte(subject + mime + "\n" + r.body)

	if err := smtp.SendMail(servername, auth, "youremail", r.to, msg); err != nil {
		return false, err
	}
	return true, nil
}

func (r *Request) ParseTemplate(templateFileName string, data interface{}) error {
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return err
	}
	r.body = buf.String()
	return nil
}

func sendEmailTest(w http.ResponseWriter, r *http.Request) {
	templateData := struct {
		Name string
		URL  string
	}{
		Name: "Laís Lima",
		URL:  "http://lais.dev",
	}

	if r.Method == "POST" {
		r := NewRequest([]string{"to@email.com"}, "Primeiro email", "Olá Laís este é o seu primeiro email")

		if err != nil {
			fmt.Println(err)
		}
		ok, error := r.SendEmail()
	}
}

func main() {
	log.Println("Server started on: http://localhost:8080")
	http.HandleFunc("/sendEmail", sendEmailTest)
	http.ListenAndServe(":8080", nil)
}

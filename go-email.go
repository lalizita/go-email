package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
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
	auth := smtp.PlainAuth("", "laislima98@hotmail.com", "xxxxx", "smtp-mail.outlook.com")
	mime := "MIME-version: 1.0;\nContent-Type: text/plain; charset=\"UTF-8\";\n\n"
	subject := "Subject: " + r.subject + "!\n"
	msg := []byte(subject + mime + "\n" + r.body)
	addr := "smtp-mail.outlook.com:587"

	if err := smtp.SendMail(addr, auth, "laislima98@hotmail.com", r.to, msg); err != nil {
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

func sendEmailTest(w http.ResponseWriter, requestParam *http.Request) {
	templateData := struct {
		Name string
		URL  string
	}{
		Name: "Laís Lima",
		URL:  "http://lais.dev",
	}

	r := NewRequest([]string{"lais.lima70@gmail.com"}, "Primeiro email", "Olá Laís este é o seu primeiro email")
	fmt.Println(r)
	err := r.ParseTemplate("template.html", templateData)
	if err != nil {
		ok, _ := r.SendEmail()
		fmt.Println(ok)
	}
}

func main() {
	http.HandleFunc("/sendEmail", sendEmailTest)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

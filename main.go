package main

import (
	"bytes"
	"html/template"
	"log"

	gomail "gopkg.in/gomail.v2"
)

type info struct {
	Name string
}

func (i info) sendMail() {

	t := template.New("template.html")

	var err error
	t, err = t.ParseFiles("template.html")
	if err != nil {
		log.Println(err)
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, i); err != nil {
		log.Println(err)
	}

	result := tpl.String()
	m := gomail.NewMessage()
	m.SetHeader("From", "")
	m.SetHeader("To", "")
	//m.SetAddressHeader("Cc", "<RECIPIENT CC>", "<RECIPIENT CC NAME>")
	m.SetHeader("Subject", "golang test")
	m.SetBody("text/html", result)
	m.Attach("template.html")// attach whatever you want

	d := gomail.NewDialer("smtp.mailtrap.io", 587, "", "")

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
func main() {

	d := info{"jack"}

	d.sendMail()
}
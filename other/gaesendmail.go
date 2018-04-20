package main

import (
	"fmt"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/mail"
)

func main() {
	//	https://project-123456.appspot.com/mail/send?to=123456@qq.com&subject=hi&body=bye
	http.HandleFunc("/mail/send", sendMail)
	appengine.Main()
}

//	App Engine>设置>电子邮件发件人>添加：sender@gmail.com
func sendMail(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		fmt.Fprintln(w, err.Error())
		return
	}

	to := r.Form.Get("to")
	subject := r.Form.Get("subject")
	body := r.Form.Get("body")
	msg := &mail.Message{
		Sender:  "sender@gmail.com",
		To:      []string{to},
		Subject: subject,
		Body:    body,
	}

	ctx := appengine.NewContext(r)
	if err := mail.Send(ctx, msg); err != nil {
		fmt.Fprintln(w, err.Error())
		return
	}
	fmt.Fprintln(w, "sent mail!")
}

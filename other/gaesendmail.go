package main

import (
	"fmt"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/mail"
)

func main() {
	//	https://project-123.appspot.com/mail/send?to=123@qq.com&subject=hi&body=bye
	http.HandleFunc("/mail/send", sendMail)
	appengine.Main()
}

//	App Engine>设置>电子邮件发件人>添加：当前谷歌云账号邮箱
func sendMail(w http.ResponseWriter, r *http.Request) {

	to := r.FormValue("to")
	subject := r.FormValue("subject")
	body := r.FormValue("body")
	msg := &mail.Message{
		Sender:  "123@gmail.com",
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

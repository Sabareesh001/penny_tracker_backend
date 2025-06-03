package email

import (
	"net/smtp"
	"os"
)

   func ComposeEmail(toEmail string,subject string,body string) (string,error){

       from := os.Getenv("EMAIL")
       password := os.Getenv("EMAIL_PASSWORD")
       to := []string{toEmail}
       smtpHost := os.Getenv("SMTP_HOST")
       smtpPort := "587"

       message := []byte("Subject: "+subject+"\r\n" +
          body+"\r\n")

       auth := smtp.PlainAuth("", from, password, smtpHost)

       err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
       if err != nil {
           return err.Error(),err
       }
       return "Email Sent Successfully!",err
   }
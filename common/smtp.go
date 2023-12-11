package common

import (
	"bytes"
	"fmt"
	"io"
	"net/smtp"
	"os"
	"strings"
	"text/template"
)

type EmailTemplateData struct {
	Name  string
	Token string
	Url   string

	// Add more fields as needed for your template
}

// Send email using smtp and html template
func SendEmail(to string, subject string, data EmailTemplateData, templateName string) error {
	templatePath := fmt.Sprintf("templates/email/%s.html", templateName)
	htmlTemplate, err := loadHTMLTemplate(templatePath)
	if err != nil {
		return err
	}
	if err != nil {
		return err
	}

	// Parse the HTML template
	tmpl, err := template.New("emailTemplate").Parse(htmlTemplate)
	if err != nil {
		return err
	}

	// Buffer to store the rendered template
	var tpl bytes.Buffer
	if err := tmpl.Execute(&tpl, data); err != nil {
		return err
	}

	// Replace the tokens in the HTML template
	emailBody := tpl.String()
	//TODO make populating template to be dynamic
	emailBody = strings.Replace(emailBody, "%Name%", data.Name, -1)
	emailBody = strings.Replace(emailBody, "%Token%", data.Token, -1)
	emailBody = strings.Replace(emailBody, "%Url%", data.Url, -1)

	// Set up SMTP authentication credentials
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := "587"
	auth := smtp.PlainAuth("", os.Getenv("SMTP_USER"), os.Getenv("SMTP_PASS"), smtpHost)

	// Set up email headers
	from := os.Getenv("FROM_EMAIL")

	// Compose the email message
	msg := []byte("To: " + to + "\r\n" +
		"From: " + from + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"MIME-version: 1.0;\r\n" +
		"Content-Type: text/html; charset=\"UTF-8\";\r\n\r\n" +
		emailBody + "\r\n")

	// Send the email using SMTP
	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, msg)
	if err != nil {
		return err
	}
}

func loadHTMLTemplate(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var buf bytes.Buffer
	_, err = io.Copy(&buf, file)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

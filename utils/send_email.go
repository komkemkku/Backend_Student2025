package utils

import (
	"crypto/tls"
	"fmt"
	"net/smtp"
	"os"
)

// SendEmail - ฟังก์ชันส่งอีเมลผ่าน Gmail
func SendEmail(to, subject, body string) error {
	// ดึงค่าอีเมลและรหัสผ่านจาก Environment Variables
	from := os.Getenv("EMAIL_USER")
	password := os.Getenv("EMAIL_PASSWORD")

	// ตั้งค่า SMTP Server
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// ตั้งค่า message
	message := fmt.Sprintf("From: %s\nTo: %s\nSubject: %s\n\n%s", from, to, subject, body)

	// ตั้งค่า Authentication
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// ใช้ `smtp.Dial()` เชื่อมต่อ SMTP Server แทน `tls.Dial()`
	client, err := smtp.Dial(smtpHost + ":" + smtpPort)
	if err != nil {
		return fmt.Errorf("failed to connect to SMTP server: %v", err)
	}
	defer client.Close()

	// เริ่ม StartTLS
	tlsConfig := &tls.Config{
		ServerName: smtpHost,
	}
	if err = client.StartTLS(tlsConfig); err != nil {
		return fmt.Errorf("failed to start TLS: %v", err)
	}

	// ตรวจสอบและตั้งค่า Authentication
	if err = client.Auth(auth); err != nil {
		return fmt.Errorf("failed to authenticate: %v", err)
	}

	// ตั้งค่า Sender และ Recipient
	if err = client.Mail(from); err != nil {
		return fmt.Errorf("failed to set sender: %v", err)
	}
	if err = client.Rcpt(to); err != nil {
		return fmt.Errorf("failed to set recipient: %v", err)
	}

	// เขียนข้อความอีเมลลงไปใน SMTP Server
	wc, err := client.Data()
	if err != nil {
		return fmt.Errorf("failed to open email body: %v", err)
	}
	defer wc.Close()

	_, err = wc.Write([]byte(message))
	if err != nil {
		return fmt.Errorf("failed to send email content: %v", err)
	}

	return nil
}

package services

import (
    "crypto/tls"
    "fmt"
    "net/smtp"

    "github.com/kadyrbayev2005/studysync/internal/utils"
)

// EmailConfig holds SMTP connection settings loaded from the environment in InitSMTP.
type EmailConfig struct {
    Host     string
    Port     string
    Username string
    Password string
    From     string
}

var emailConfig *EmailConfig

// InitSMTP reads SMTP_HOST, SMTP_PORT, SMTP_USERNAME, SMTP_PASSWORD, and SMTP_FROM for outbound mail.
func InitSMTP() {
    emailConfig = &EmailConfig{
        Host:     utils.GetEnv("SMTP_HOST", "smtp.gmail.com"),
        Port:     utils.GetEnv("SMTP_PORT", "587"),
        Username: utils.GetEnv("SMTP_USERNAME", ""),
        Password: utils.GetEnv("SMTP_PASSWORD", ""),
        From:     utils.GetEnv("SMTP_FROM", ""),
    }
}

// SendEmail delivers a plain-text message over SMTP with TLS (e.g. Gmail submission on port 587).
func SendEmail(to, subject, body string) error {
    if emailConfig == nil {
        return fmt.Errorf("SMTP not initialized")
    }

    msg := fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\r\n\r\n%s",
        emailConfig.From, to, subject, body)

    auth := smtp.PlainAuth("",
        emailConfig.Username,
        emailConfig.Password,
        emailConfig.Host,
    )

    addr := fmt.Sprintf("%s:%s", emailConfig.Host, emailConfig.Port)

    // For Gmail with TLS
    tlsConfig := &tls.Config{
        ServerName: emailConfig.Host,
    }

    conn, err := tls.Dial("tcp", addr, tlsConfig)
    if err != nil {
        return fmt.Errorf("failed to connect to SMTP server: %w", err)
    }
    defer conn.Close()

    client, err := smtp.NewClient(conn, emailConfig.Host)
    if err != nil {
        return fmt.Errorf("failed to create SMTP client: %w", err)
    }
    defer client.Close()

    if err = client.Auth(auth); err != nil {
        return fmt.Errorf("authentication failed: %w", err)
    }

    if err = client.Mail(emailConfig.From); err != nil {
        return fmt.Errorf("failed to set sender: %w", err)
    }

    if err = client.Rcpt(to); err != nil {
        return fmt.Errorf("failed to set recipient: %w", err)
    }

    w, err := client.Data()
    if err != nil {
        return fmt.Errorf("failed to get data writer: %w", err)
   }

    _, err = w.Write([]byte(msg))
    if err != nil {
        return fmt.Errorf("failed to write email: %w", err)
    }

    err = w.Close()
    if err != nil {
        return fmt.Errorf("failed to close writer: %w", err)
    }

    Info("Email sent successfully", "to", to, "subject", subject)
    return nil
}

// SendDeadlineReminder emails the student a fixed-template notice about an upcoming task due time.
func SendDeadlineReminder(userEmail, taskTitle string, dueDate string) error {
    subject := "StudySync: Upcoming Deadline Reminder"
    body := fmt.Sprintf(`
Dear Student,

You have an upcoming deadline:

Task: %s
Due Date: %s

Please complete your task on time!

Best regards,
StudySync Team
`, taskTitle, dueDate)

    return SendEmail(userEmail, subject, body)
}
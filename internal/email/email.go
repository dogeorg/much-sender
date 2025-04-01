package email

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/smtp"
	"strings"

	"github.com/dogeorg/much-sender/internal/config"
)

// Such EmailRequest struct for JSON payload
type EmailRequest struct {
	ReplyToEmail string `json:"reply_to_email"`
	ReplyToName  string `json:"reply_to_name"`
	ToEmail      string `json:"to_email"`
	ToName       string `json:"to_name"`
	Subject      string `json:"subject"`
	HTML         string `json:"html"`
}

// Such SendEmail, handles the email sending logic
func SendEmail(config config.Config, req EmailRequest) error {
	auth := smtp.PlainAuth("", config.SMTP.Username, config.SMTP.Password, config.SMTP.Server)
	addr := fmt.Sprintf("%s:%d", config.SMTP.Server, config.SMTP.Port)

	from := fmt.Sprintf("From: %s <%s>\r\n", req.ReplyToName, req.ReplyToEmail)
	to := fmt.Sprintf("To: %s <%s>\r\n", req.ToName, req.ToEmail)
	replyTo := fmt.Sprintf("Reply-To: %s <%s>\r\n", req.ReplyToName, req.ReplyToEmail)
	subject := fmt.Sprintf("Subject: %s\r\n", req.Subject)
	mime := "MIME-Version: 1.0\r\nContent-Type: text/html; charset=\"UTF-8\"\r\n\r\n"

	msg := []byte(from + to + replyTo + subject + mime + req.HTML)

	return smtp.SendMail(addr, auth, config.SMTP.Username, []string{req.ToEmail}, msg)
}

// Much Handler, creates HTTP handler for email sending
func Handler(config config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Much Sad! Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Much Sad! Missing or invalid Authorization header", http.StatusUnauthorized)
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token != config.Security.BearerToken {
			http.Error(w, "Much Sad! Invalid token", http.StatusUnauthorized)
			return
		}

		var emailReq EmailRequest
		err := json.NewDecoder(r.Body).Decode(&emailReq)
		if err != nil {
			http.Error(w, "Much Sad! Invalid JSON payload", http.StatusBadRequest)
			return
		}

		if emailReq.ToEmail == "" || emailReq.Subject == "" || emailReq.HTML == "" {
			http.Error(w, "Much Sad! Missing required fields", http.StatusBadRequest)
			return
		}

		err = SendEmail(config, emailReq)
		if err != nil {
			log.Printf("Much Sad! Error sending email: %v", err)
			http.Error(w, "Much Sad! Failed to send email", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Much Wow! Email sent successfully",
		})
	}
}

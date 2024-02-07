package middleware

import (
	"bytes"
	"foodcourt/app/api/request"
	"foodcourt/app/config"
	"html/template"
	"log"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v3"
	"gopkg.in/mail.v2"
)

func SendEmail(request request.EmailRequest, cfg *config.Config) error {
	templatePath := os.Getenv("TEMPLATE_PATH")
	fullPath := filepath.Join(templatePath, request.Template)
	t, err := template.ParseFiles(fullPath)
	if err != nil {
		log.Printf("Erreur lors de la lecture du template: %v", err)
		return err
	}

	var buf bytes.Buffer
	err = t.Execute(&buf, request.Data)
	if err != nil {
		log.Printf("Erreur lors de l'exécution du template: %v", err)
		return err
	}

	m := mail.NewMessage()
	m.SetHeader("From", cfg.EmailFrom)
	m.SetHeader("To", request.To)
	m.SetHeader("Subject", request.Subject)
	m.SetBody("text/html", buf.String())

	d := mail.NewDialer(cfg.SMTPHost, 587, cfg.SMTPUser, cfg.SMTPPass)
	if err := d.DialAndSend(m); err != nil {
		log.Printf("Erreur lors de l'envoi de l'email: %v", err)
		return err
	}

	return nil
}

func RegisterMiddleware(c fiber.Ctx) error {
	cfg := config.LoadConfig()

	req := c.Locals("registerRequest").(request.RegisterRequest)

	emailRequest := request.EmailRequest{
		To:       req.Email,
		Subject:  "Bienvenue sur notre site !",
		Template: "welcome_template.html",
		Data:     struct{ Username string }{Username: req.Username},
	}

	err := SendEmail(emailRequest, cfg)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Erreur lors de l'envoi de l'email")
	}

	return nil
}

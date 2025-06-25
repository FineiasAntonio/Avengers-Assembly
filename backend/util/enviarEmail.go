package util

import (
	"fmt"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
)

func EnviarEmail(mensagemSubject, mensagem, destinatario string) error {
	err := godotenv.Load("../../.env")

	if err != nil {
		return fmt.Errorf("erro ao carregar .env: %v", err)
	}

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"
	from := os.Getenv("GMAIL_CCTS_EMAIL")
	password := os.Getenv("GMAIL_CCTS_SENHA_APP")

	auth := smtp.PlainAuth("", from, password, smtpHost)

	message := []byte(fmt.Sprintf("Subject: %s\r\n", mensagemSubject) +
		"\r\n" +
		fmt.Sprintf("Mensagem: %s\r\n", mensagem))

	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{destinatario}, message)
	return err
}

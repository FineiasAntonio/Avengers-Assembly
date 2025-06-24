package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
)

func FormatarChatID(telefone string) string {
	re := regexp.MustCompile(`\D`)
	num := re.ReplaceAllString(telefone, "")
	return num + "@c.us"
}

func FormatarTelefone(telefone string) string {
	re := regexp.MustCompile(`\D`)
	num := re.ReplaceAllString(telefone, "")

	if len(num) == 11 {
		return fmt.Sprintf("(%s) %s-%s", num[0:2], num[2:7], num[7:])

	} else if len(num) == 10 {
		return fmt.Sprintf("(%s) %s-%s", num[0:2], num[2:6], num[6:])

	} else {
		return num
	}
}

func EnviarMensagemWaha(telefone, texto string) error {
	chatId := FormatarChatID(telefone)

	payload := map[string]string{
		"chatId": chatId,
		"text":   texto,
	}
	jsonData, _ := json.Marshal(payload)

	resp, err := http.Post("http://localhost:3000/sendText", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("status code: %d", resp.StatusCode)
	}

	return nil
}

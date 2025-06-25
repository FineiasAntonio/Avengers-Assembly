package scheduler

import (
	"log"

	"backend/service"

	"github.com/robfig/cron/v3"
)

func IniciarScheduler(srv *service.RequisicaoExameService) *cron.Cron {
	c := cron.New()

	_, err := c.AddFunc("0 12 * * *", func() {
		err := srv.ProcessarLembretes()
		if err != nil {
			log.Printf("Erro ao processar lembretes: %v", err)
		} else {
			log.Println("Lembretes processados com sucesso (scheduler)")
		}
	})
	if err != nil {
		log.Fatalf("Erro ao configurar cron: %v", err)
	}

	c.Start()
	return c
}

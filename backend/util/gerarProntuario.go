package util

import (
	"fmt"
	"math/rand"
	"time"
)

func GerarProntuario() string {
	now := time.Now().Format("20060102150405")
	rand.Seed(time.Now().UnixNano())
	suf := rand.Intn(1000)
	return fmt.Sprintf("PRT-%s-%03d", now, suf)
}

package util

import (
	"fmt"
	"math/rand"
	"time"
)

func GerarCodigo() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%06d", rand.Intn(1000000))
}

package util

import (
	"math/rand"
	"strconv"
	"time"
)

func GerarId(tamanho int) string {

	min := int64(pow10(tamanho - 1))
	max := int64(pow10(tamanho)) - 1

	rand.Seed(time.Now().UnixNano())
	num := min + rand.Int63n(max-min+1)

	id := strconv.FormatInt(num, 10)
	return id
}

func pow10(n int) int64 {
	result := int64(1)
	for i := 0; i < n; i++ {
		result *= 10
	}
	return result
}

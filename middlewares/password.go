package middlewares

import (
	"crypto/rand"
	"fmt"
)

func TokenGenerator() string {
	t := make([]byte, 23)
	rand.Read(t)
	return fmt.Sprintf("%x", t)
}

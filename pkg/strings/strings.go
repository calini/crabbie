package strings

import (
	"math/rand"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func GenerateCode(length int) string {
	charSet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var output strings.Builder

	for i := 0; i < length; i++ {
		random := rand.Intn(len(charSet))
		randomChar := charSet[random]
		output.WriteString(string(randomChar))
	}

	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return output.String()
}

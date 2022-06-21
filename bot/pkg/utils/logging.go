package utils

import (
	"fmt"
	"log"
)

func LogError(m string, e error) {
	log.Print(fmt.Sprintf("%s: %e", m, e))
}

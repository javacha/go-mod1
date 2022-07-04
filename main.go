package myLogger

import (
	"log"
)

func logInfo(message string) {
	log.Println("INFO -", message)
}
func logWarning(message string) {
	log.Println("WARN -", message)
}
func logError(message string) {
	log.Println("ERRO -", message)
}

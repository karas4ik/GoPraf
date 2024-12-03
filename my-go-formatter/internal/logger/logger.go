package logger

import (
	"fmt"
	"os"
	"time"
)

// Log записывает сообщение в лог-файл.
func Log(message string) {
	logFile, err := os.OpenFile("logs/formatter.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("Ошибка при открытии лог-файла:", err)
		return
	}
	defer logFile.Close()

	logger := fmt.Sprintf("%s: %s\n", time.Now().Format("2006-01-02 15:04:05"), message)
	logFile.WriteString(logger)
}

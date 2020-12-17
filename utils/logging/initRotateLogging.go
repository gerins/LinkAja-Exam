package logging

import (
	"log"
	"os"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
)

// Set log output to file every 24 hour
func InitRotateLogger() *rotatelogs.RotateLogs {
	currentDirectory, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	fileLocation := currentDirectory + "/log/"
	fileName := "server_log.%Y-%b-%d-%H-%M" + ".txt"

	rl, _ := rotatelogs.New(
		fileLocation+fileName,
		rotatelogs.WithMaxAge(30*24*time.Hour),    // Maximum time before deleting file log
		rotatelogs.WithRotationTime(24*time.Hour), // Time before creating new file
		rotatelogs.WithClock(rotatelogs.Local))

	return rl
}

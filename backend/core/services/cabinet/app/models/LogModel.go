package models

import (
	"log"
	"os"
)

func Logger(level uint8) *log.Logger {
	customLogger := log.New(os.Stdout, "STANDART: ", log.LstdFlags|log.Lshortfile)

	if level <= ConfigProcess().LogLevel {

		switch level {
		case 1:
			customLogger = log.New(os.Stdout, "CRITICAL: ", log.LstdFlags|log.Lshortfile)
			break

		case 2:
			customLogger = log.New(os.Stdout, "WARNING: ", log.LstdFlags|log.Lshortfile)
			break

		case 3:
			customLogger = log.New(os.Stdout, "INFO: ", log.LstdFlags|log.Lshortfile)
			break

		case 4:
			customLogger = log.New(os.Stdout, "DEBUG: ", log.LstdFlags|log.Lshortfile)
			break
		}
	}

	return customLogger
}

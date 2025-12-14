package main

import (
	"log"
	"os"
)

func main() {

	log.Println("This is a log message.")

	log.SetPrefix("INFO Prefix: ")
	log.Println("This is a info log.")

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Println("This is a log message with data")

	// using custom logger

	infoLogger.Println("This is an informative log")
	warnLogger.Println("This is just a warning message")
	errorLogger.Println("Something seems wrong")

}

var (
	infoLogger  = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	warnLogger  = log.New(os.Stdout, "WARN: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
)

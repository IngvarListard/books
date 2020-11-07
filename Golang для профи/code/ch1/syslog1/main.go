package main

import (
	"fmt"
	"log"
	"log/syslog"
	"os"
	"path/filepath"
)

func main() {
	programName := filepath.Base(os.Args[0])
	sysLog, err := syslog.New(syslog.LOG_INFO|syslog.LOG_LOCAL7, programName)
	log.Println("LOG_INFO + LOG_LOCAL7: Logging in Go!")

	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(sysLog)

	log.Println("SYSLOG MESSAGE")
	fmt.Println("aAAAAA")
}

package internal

import (
	"log"
	"time"
)

func writeLog(s string) {
	logTime := time.Now()
	log.Printf("%s %s\n", logTime.Format("2006-01-02|15-04-05"), s)
}

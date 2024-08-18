package main

import (
	"crontab-without-race-execution/util"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	pid := os.Getpid()

	log.Default().Println("Run with new PID ", pid)

	errPid := util.WritePID(strconv.Itoa(pid))

	if errPid != nil {
		panic(errPid)
	}

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	sleep, _ := strconv.Atoi(os.Getenv("SLEEP"))

	for {
		processUUID, errUuid := uuid.NewUUID()

		if errUuid != nil {
			panic(errUuid)
		}

		util.Log(processUUID.String(), "INFO", "Start new process")

		time.Sleep(time.Duration(sleep) * time.Second)
	}
}

package util

import (
	"log"
	"os"
)

func WritePID(pid string) error {
	f, err := os.Create("pid")

	if err != nil {
		return err
	}

	_, err = f.WriteString(pid)

	if err != nil {
		errClose := f.Close()

		if errClose != nil {
			return errClose
		}

		return err
	}

	err = f.Close()

	if err != nil {
		return err
	}

	return nil
}

func Log(processUUID string, level string, msg string) {
	log.Default().Println(level, processUUID+":", msg)
}

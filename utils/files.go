package utils

import (
	"bufio"
	"log"
	"os"
)

func Read(file string, lines *[]string) {

	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		value := scanner.Text()
		*lines = append(*lines, value)
	}
}

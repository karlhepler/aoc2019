package input

import (
	"bufio"
	"log"
	"os"
)

var cwd string

func init() {
	var err error
	cwd, err = os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
}

// Lines sends lines of the given relative path to the channel
func Lines(relpath string) <-chan string {
	file, err := os.Open(cwd + "/" + relpath)
	if err != nil {
		log.Fatal(err)
	}

	line := make(chan string)
	scanner := bufio.NewScanner(file)

	go func() {
		defer close(line)
		defer file.Close()

		for scanner.Scan() {
			line <- scanner.Text()
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}()

	return line
}

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

// Lines sends lines of the given filename to the channel
func Lines(filename string, test ...bool) <-chan string {
	dir := "/input/"
	if len(test) > 0 && test[0] == true {
		dir = "/../input/"
	}

	file, err := os.Open(cwd + dir + filename)
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

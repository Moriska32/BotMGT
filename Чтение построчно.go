package main

import (
	"bufio"

	"fmt"

	"os"
)

const logfile = "ROUTES"

var print = fmt.Println

func scanFile() string {

	f, _ := os.OpenFile(logfile, os.O_RDONLY, os.ModePerm)

	defer f.Close()

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		sc.Text()

	}

}

func main() {
	scanFile()
}

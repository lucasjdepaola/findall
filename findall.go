package main

import (
	"bufio"
	"fmt"
	fw "github.com/charlievieth/fastwalk"
	"io/fs"
	"os"
	"strings"
)

func handle(err error) {
	if err != nil {
		panic(err)
	}
}

func getExtension(fileName string) string {
	str := strings.Split(fileName, ".")
	return str[len(str)-1]
}

func grepFile(findString string, path string) bool {
	file, err := os.Open(path)
	handle(err)
	scanner := bufio.NewScanner(file)
	defer file.Close()
	for scanner.Scan() {
		tx := scanner.Text()
		if strings.Contains(tx, findString) {
			fmt.Println(path)
			return true
		}
	}
	return false
}

func recurseDir(extension string, findString string) {
	conf := fw.Config{
		Follow: false,
	}
	wd, err := os.Getwd()
	handle(err)
	er := fw.Walk(&conf, wd,
		func(path string, entry fs.DirEntry, err error) error {
			if entry.IsDir() {
			} else if getExtension(entry.Name()) == extension {
				// now cat further into the file, and find string within such file
				grepFile(findString, path)
			}
			return nil
		},
	)
	handle(er)
}

func main() {
	if len(os.Args) < 2 {
		panic("arguments need to be greater than 1")
	}
	extension := os.Args[1]
	findString := os.Args[2]

	recurseDir(extension, findString)
}

package main

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	//Read a file at once
	_, err := ioutil.ReadFile("Go-Concepts\\workers.go")
	if err != nil {
		log.Println(err.Error())
		return
	}
	//log.Println(string(data))

	f, err := os.Open("Go-Concepts\\workers.go")
	if err != nil {
		log.Println(err.Error())
		return
	}

	defer f.Close()

	// scanner := bufio.NewScanner(f)
	// scanner.Split(bufio.ScanLines)
	// var str strings.Builder
	// for scanner.Scan() {
	// 	text := scanner.Text()
	// 	str.WriteString(text)
	// }
	// log.Println(str.String())

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		log.Println(err.Error())
		return
	}
	log.Println(text)

}

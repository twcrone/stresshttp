package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func submissions(r int, data []string) {
	i := 0
	for {
		j := rand.Intn(len(data))
		submission(data[j])
		i++
		message := fmt.Sprintf("Routine %d - %d", r, i)
		if r == 0 {
			message = "***" + message + "***"
		}
		fmt.Println(message)
	}
}

func submission(text string) {
	resp, err := http.Post("http://localhost:8080/submission", "text/plain",
		bytes.NewBuffer([]byte(text)))
	if err != nil {
		print(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		print(err)
	}
	fmt.Println(string(body))
}

func readData() []string {
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return lines
}

func main() {
	data := readData()

	go submissions(1, data)
	go submissions(2, data)
	go submissions(3, data)
	go submissions(4, data)
	go submissions(5, data)
	go submissions(6, data)
	go submissions(7, data)
	go submissions(8, data)
	go submissions(9, data)
	go submissions(0, data)

	time.Sleep(10 * time.Second)
}

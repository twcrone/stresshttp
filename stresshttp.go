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
	for i := 0; i < 100; i++ {
		go submissions(i, data)
	}
	time.Sleep(10 * time.Minute)
}

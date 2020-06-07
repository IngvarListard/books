package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	contentChan := make(chan []byte)
	filename := "test2.html"
	for _, url := range os.Args[1:] {
		go fetch(url, ch, contentChan)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	for range os.Args[1:] {
		err := writeFile("test2.html", <-contentChan)
		if err != nil {
			fmt.Fprintf(os.Stderr, "При записи данных в файл %s произошла ошибка %v", filename, err)
		}
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string, contentChan chan<- []byte) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("While reading: %s: %v", url, err)
	}
	nbytes := len(bytes)

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%2.fs %7d %s", secs, nbytes, url)
	contentChan <- bytes
}

func writeFile(filename string, data []byte) error {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()
	if err != nil {
		fmt.Printf("Error open file: %s", filename)
		return err
	}

	f.Write(data)
	return err
}

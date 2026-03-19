package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

func downloadCutePic(link string) {
	sps := strings.Split(link, "/")
	filename := sps[len(sps)-1]

	file, err := os.OpenFile("cute/"+filename, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal("Error opening file:", err)
		return
	}
	defer file.Close()

	resp, err := http.Get(link)
	if err != nil {
		log.Fatal("Error on downloading:", link)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error on read response:", link)
		return
	}

	n, err := file.Write(body)
	if err != nil {
		log.Fatal("Error on write to file:", link)
		return
	}

	fmt.Printf("%s success %d bytes\n", filename, n)
}

func main() {
	resp, err := http.Get("https://forum.gamer.com.tw/C.php?bsn=74934&snA=10873")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Printf("%s\n", resp.Status)

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	//f, err := os.OpenFile("result.dat", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)

	if err != nil {
		panic(err)
	}

	var links []string
	doc.Find("table tbody a.photoswipe-image").Each(func(i int, s *goquery.Selection) {
		link, exist := s.Attr("href")
		if exist {
			links = append(links, link)
		}
	})

	for _, link := range links {
		fmt.Println(link)
	}

	// dispatch task
	wg := &sync.WaitGroup{}

	for i := 0; i < len(links); i += 10 {
		wg.Add(1)

		go func() {
			task_num := i
			for j := task_num; j < task_num+10 && j < len(links); j++ {
				downloadCutePic(links[j])
			}
			wg.Done()
		}()
	}

	wg.Wait()
}

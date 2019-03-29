package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"
)

func titulo(urls ...string) <-chan [2]string {
	c := make(chan [2]string)

	for _, url := range urls {
		go func(url string) {

			now := time.Now()

			resp, _ := http.Get(url)
			html, _ := ioutil.ReadAll(resp.Body)

			r, _ := regexp.Compile("<title(.*?)>(.*?)<\\/title>")
			c <- [2]string{r.FindStringSubmatch(string(html))[2], fmt.Sprint(time.Since(now))}

		}(url)
	}
	return c
}

func main() {
	// f1 := titulo("http://www.google.com",
	// 	"http://www.grupoltm.com.br",
	// 	"http://nasa.gov",
	// 	"http://youtube.com",
	// )

	t1 := titulo("https://www.nasa.gov", "https://www.google.com")
	t2 := titulo("https://www.uol.com.br", "https://www.youtube.com")

	var result1 = <-t1
	var result2 = <-t2

	fmt.Println("Primeiros: ", result1[0], " time: ", result1[1], "|", result2[0], " time: ", result2[1])

	result1 = <-t1
	result2 = <-t2

	fmt.Println("Segundos: ", result1[0], " time: ", result1[1], "|", result2[0], " time: ", result2[1])

}

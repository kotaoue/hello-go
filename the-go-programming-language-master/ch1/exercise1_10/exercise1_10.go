// e.g. go run exercise1_10.go http://ftp.arin.net/pub/stats/arin/delegated-arin-extended-latest;go run exercise1_10.go http://ftp.arin.net/pub/stats/arin/delegated-arin-extended-latest
// $ diff -y exercise1_10_1592921481.txt exercise1_10_1592921490.txt
// 7.94s 9994816 http://ftp.arin.net/pub/stats/arin/delegated-ar | 14.56s 9994816 http://ftp.arin.net/pub/stats/arin/delegated-a
// 7.94s elapsed                                                 | 14.56s elapsed
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}

	file, _ := os.Create(fmt.Sprintf("exercise1_10_%v.txt", time.Now().Unix()))
	for range os.Args[1:] {
		fmt.Fprintln(file, <-ch)
	}
	fmt.Fprintf(file, "%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}

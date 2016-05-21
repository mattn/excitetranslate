package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/yhat/scrape"
	"golang.org/x/net/html"
)

func main() {
	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	values := url.Values{}
	values.Add("before", string(b))
	values.Add("wb_lp", "ENJA")
	resp, err := http.PostForm("http://www.excite.co.jp/world/english/", values)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	root, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	after, ok := scrape.Find(root, scrape.ById("after"))
	if ok {
		fmt.Println(scrape.Text(after))
	}
}

package main

import (
	"flag"
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
	var mode string
	flag.StringVar(&mode, "mode", "ENJA", "translate mode ENJA/JAEN")
	flag.Parse()
	if mode != "ENJA" && mode != "JAEN" {
		flag.Usage()
		os.Exit(1)
	}
	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	values := url.Values{}
	values.Add("before", string(b))
	values.Add("wb_lp", mode)
	resp, err := http.PostForm("https://www.excite.co.jp/world/english_japanese/", values)
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

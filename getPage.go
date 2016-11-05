package getPage

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func GetHTTPCode(u string, num int, cresults chan<- int, client *http.Client) {
	GetHTTPCodeNoChannel(u, num, client)
	cresults <- 1
}

func GetHTTPCodeNoChannel(u string, num int, client *http.Client) {
	t := time.Now()
	res, err := client.Get(u)
	t2 := time.Now()

	if err != nil {
		fmt.Printf("%v,"+err.Error()+",%v\n", num, t2.Sub(t))
	}
	if err == nil {
		res.Body.Close()
		fmt.Printf("%v,"+res.Status+",%v\n", num, t2.Sub(t))
	}
}

func GetContent(u string) string {

	res, err := http.Get(u)

	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	s := string(body[:])
	return s
}

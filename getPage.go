package getPage

import (
	"io/ioutil"
	"net/http"
	"log"
)

func PrintMessage() string {
	return "Message from PrintMessage"
}

func PrintEcho(s string) string {
	return s
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


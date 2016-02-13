package getPage

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func PrintMessage() string {
	return "Message from PrintMessage"
}

func PrintEcho(s string) string {
	return s
}

func GetHTTPCode(u string, num int, cresults chan<- int) {

	client := &http.Client{}

	req, err := http.NewRequest("GET", u, nil)
	req.Header.Add("Authorization", "anJrDZQ9M5PYvmGuADqYAxbkdPxUkHsFvoj1EWLj4WFOKHO4TqUFOx8PmtuubuWO")

	t := time.Now()
	res, err := client.Do(req)
	t2 := time.Now()

	if err != nil {
		//log.Fatal(err)
		log.Println(err)
	}
	if err == nil {
		ioutil.ReadAll(res.Body)
		res.Body.Close()
		fmt.Printf("%v,"+res.Status+",%v\n", num, t2.Sub(t))
	}
	
	cresults <- 1
}


func GetHTTPCodeNoChannel(u string, num int) {

	client := &http.Client{}

	req, err := http.NewRequest("GET", u, nil)
	
	t := time.Now()
	res, err := client.Do(req)
	t2 := time.Now()

	if err != nil {
		//log.Fatal(err)
		log.Println(err)
	}
	if err == nil {
		ioutil.ReadAll(res.Body)
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


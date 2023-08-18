package main

import (
	"fmt"
	"net/http"
	"time"
)


func main() {
	links := []string{
		"http://facebook.com",
		"http://google.com",
		"http://stackoverflow.com",
		"http://golang.org",
	}

	c := make(chan string)

	for _, link := range links{
		go checkLink(link, c)
	}
	for l := range c {
		go func(link string){
			time.Sleep(3* time.Second)
			checkLink(link,c)
		}(l)
	}
}

func checkLink(link string, c chan string){
	time.Sleep(3*time.Second)
	_, err := http.Get(link)
	if err != nil{
		fmt.Println(link, " might be down!")
		c <- "MIght be down my shinobi"
		return
	}
	fmt.Println(link, " the site is ONE ONLINE")
	c <- link

}
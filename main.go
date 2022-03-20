package main

import "sync"

var wg sync.WaitGroup

func main() {

	liste := ReadWebSite("./sites.json")
	wg.Add(len(liste))

	for _, site := range liste {
		go Fetch(site)
	}
	wg.Wait()
}

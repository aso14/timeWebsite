package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func ReadWebSite(JsFile string) []string {
	bytesArray, err := ioutil.ReadFile(JsFile)
	if err != nil {
		panic(err)
	}
	var listWeb []string
	err = json.Unmarshal(bytesArray, &listWeb)
	if err != nil {
		fmt.Println("Error : ", err)
	}
	return listWeb
}

func Fetch(url string) {
	defer wg.Done()

	starTime := time.Now()
	resp, err := http.Get(url)
	endTime := time.Since(starTime)

	if err != nil {
		fmt.Println(err)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	size := len(body)
	defer resp.Body.Close()
	fmt.Printf("Site Web : %s | Taille %d | Temps d'execution : %d\n", url, size, endTime)

}

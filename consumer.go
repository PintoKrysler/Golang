package main

import(
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Payload struct {
	Profile Data
}

type Personal struct{
	First_name string
	Last_name string
	Age	int
	Email string 
}

type Data struct {
	Info Personal
	Languages Languages
}

type Languages map[string]int

func main() {
	url := "http://localhost:1336"
	resp,error := http.Get(url) 
	if error != nil{
		panic(error)
	}
	defer resp.Body.Close()
	body , error := ioutil.ReadAll(resp.Body)
	if error!= nil{
		panic(error)
	}

	var p Payload
	error = json.Unmarshal(body,&p)
	if error!=nil{
		panic(error)
	}

	fmt.Println(p.Profile.Languages, "\n" , p.Profile.Info)
}
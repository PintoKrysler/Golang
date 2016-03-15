package main 

import (
	"encoding/json"
	"fmt"
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

func serveRest(w http.ResponseWriter, r *http.Request){
	response, err := getJsonResponse()
	if err != nil {
		panic (err)
	}

	fmt.Fprintf(w, string(response))
}

func main(){
	http.HandleFunc("/",serveRest)
	http.ListenAndServe("localhost:1336",nil)
}

func getJsonResponse() ([]byte , error){
	k := Personal{"Krysler","Pinto",24,"pinto.krysler@gmail.com"}
	lang := make(map[string]int);
	lang["C"] = 100
	lang["JavaScript"] = 90
	lang["Java"] = 85
	lang["Haskell"] = 20
	lang["Prolog"] = 70
	lang["PHP"] = 95
	lang["Python"] = 50
	d := Data{k,lang}
	p := Payload{d}

	return json.MarshalIndent(p, "", "")
}
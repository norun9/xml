package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Page struct{
	Title string
	Body []byte
}

func (p *Page) save() (err error) {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page , error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request){
	
}

func main() {
	//p1 := &Page{Title: "test", Body:[]byte("This is a sample page.\t")}
	//p1.save()
	//
	//p2, _ := loadPage(p1.Title)
	//fmt.Println(string(p2.Body))

	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc('/view', viewHandler)
	server.ListenAndServe()
}
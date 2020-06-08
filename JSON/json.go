package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"fmt"
)

type Post struct {
	Id int `json:"id"`
	Content string `json:"content"`
	Author Author `json:"author"`
	Comments []Comment `json:"comments"`
}

type Author struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Comment struct {
	Id      string `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func main() {
	jsonFile, err := os.Open("post.json")
	if err != nil{
		fmt.Println("Error opening JSON file:", err)
		return
	}
	defer jsonFile.Close()
	jsonData, err := ioutil.ReadAll(jsonFile)
	if err != nil{
		panic(jsonData)
	}

	var post Post
	json.Unmarshal(jsonData, &post)
	fmt.Println(post)
}

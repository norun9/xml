package main

import (
	"encoding/json"
	"os"
)

type Post struct {
	Id int `json:"id"`
	Content string `json:"content"`
	Author Author `json:"author"`
	Comments []Comment `json:"comments"`
}

type Author struct {
	Id   int `json:"id"`
	Name string `json:"name"`
}

type Comment struct {
	Id      int `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func main() {
	post := Post{
		Id: 1,
		Content: "Hello World",
		Author: Author{
			Id: 2,
			Name: "Sau",
		},
		Comments: []Comment{
			Comment{
				Id: 3,
				Content: "How are you today?",
				Author: "Adam"
			},
			Comment{
				Id: 4,
				Content: "Have a great day!",
				Author: "Betty",
			},
		},
	}

	jsonFile, err := os.Create("post.json")
	if err != nil {
		panic(err)
	}
	encoder := json.NewEncoder(jsonFile)
	err = encoder.Encode(&post)
	if err != nil {
		panic(err)
	}
}


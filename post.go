package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

type Post struct {
	XMLName xml.Name `xml:"post"`
	Id string `xml:"id,attr"`
	Content string  `xml:"content"`
	Author string `xml:"author"`
	Xml string `xml:",innerxml"`
	Comments []Comment `xml:"comments>comment"`
}

type Comment struct {
	Id string `xml:"is,attr"`
	Content string `xml:"content"`
	Author Author `xml:"author"`
}

type Author struct {
	Id string `xml:"id",attr`
	Name string `xml:",chardata"`
}

func main() {
	xmlFile, err := os.Open("post.xml")
	if err != nil{
		fmt.Println("Error opening XML file:", err)
		return
	}
	defer xmlFile.Close()

	//xmlDate, err := ioutil.ReadAll(xmlFile)
	//if err != nil {
	//	fmt.Println("Error reading XML data:", err)
	//	return
	//}
	//
	//post := new(Post)
	//xml.Unmarshal(xmlDate, post)
	//fmt.Println(post)

	decorder := xml.NewDecoder(xmlFile)
	for {
		t, err := decoder.Token()
		if err == io.EOF {
			break
		}
		if err != nil{
			fmt.Println("Error decoding XML into tokens:", err)
			return
		}

		switch se := t.(type){
		case xml.StartElement:
			if se.Name.Local == "comment" {
				var comment Comment
				decorder.DecodeElement(&comment, &se)
			}
		}
	}
}
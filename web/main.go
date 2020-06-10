package web

import "io/ioutil"

type Page struct {
	Title string
	Body []byte
}

func (p *Page) save() (err error) {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page ,error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {

	}
}

func main() {
	p1 := &Post{Title: "test", Body:[]byte("This is a sample page.\t")}
	p1.save()
}
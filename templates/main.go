package main

import (
	"html/template"
	"os"
)

//Index handle the values to be subtituted in the template
//All start with uppercase as we have to export them
type Index struct {
	Content string
	Num     int
	Names   []string
}

func main() {
	//It is important the order in which we call the templates
	t := template.Must(template.ParseFiles("index.html", "header.html", "footer.html"))
	cont := Index{Content: "my example content",
		Num:   42,
		Names: []string{"field0", "field1", "field2"},
	}
	err := t.Execute(os.Stdout, cont)
	if err != nil {
		panic(err)
	}
}

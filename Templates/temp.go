package main

import (
	"html/template"
	"os"
)

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

func main() {
	data := TodoPageData{
		PageTitle: "My TODO list",
		Todos: []Todo{
			{Title: "Task 1", Done: false},
			{Title: "Task 2", Done: true},
			{Title: "Task 3", Done: true},
		},
	}

	tmpl := template.Must(template.New("todo").Parse(`
<h1>{{.PageTitle}}</h1>
<ul>
    {{range .Todos}}
        {{if .Done}}
            <li class="done">{{.Title}}</li>
        {{else}}
            <li>{{.Title}}</li>
        {{end}}
    {{end}}
</ul>`))

	tmpl.Execute(os.Stdout, data)

	tmpl, err := template.ParseFiles("layout.html")
	if err != nil {
		panic(err)
	}

	parsed_tmpl := template.Must(tmpl, err)
	parsed_tmpl.Execute(os.Stdout, data)
}

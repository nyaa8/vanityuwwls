package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var domain = flag.String("domain", "localhost", "Vanity domain")
var user = flag.String("user", "", "Your GitHub username")

type VanityData struct {
	Domain string
	User   string
	Path   string
}

const vanitypage = `<!DOCTYPE html>
<html lang="en">
<head>

<meta charset="UTF-8">
<title>{{ .Path }}</title>
<meta name="go-import" content="{{ .Domain }}{{ .Path }} git https://github.com/{{ .User }}{{ .Path }}">
<meta name="go-source" content="https://github.com/{{ .User }}{{ .Path }} _ https://github.com/{{ .User }}{{ .Path }}/tree/master{/dir} https://github.com/{{ .User }}{{ .Path }}/blob/master{/dir}/{file}#L{line}">

</head>
<body>

<a href="https://github.com/{{ .User }}{{ .Path }}" style="font-family: monospace">{{ .Domain }}{{ .Path }}</a>

</body>
</html>`

func main() {
	flag.Parse()

	if len(*user) < 1 {
		log.Fatal("Please set -user flag")
	}

	t := template.Must(template.New("vanitypage").Parse(vanitypage))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := t.Execute(w, VanityData{*domain, *user, r.URL.Path})
		if err != nil {
			fmt.Fprint(w, "whoops :(")
		}
	})

	http.ListenAndServe(":3001", nil)
}

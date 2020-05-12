package main

import (
	"log"
	"io"
	"io/ioutil"
	"strings"
	"net/http"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
)

func main() {
	http.HandleFunc("/", servePage)
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	  
	log.Println("Server is running on port 80")
	
	  if err := http.ListenAndServe(":80", nil); err != nil {
    	panic(err)
  	}
}

func servePage(writer http.ResponseWriter, request *http.Request) {
	filename := strings.TrimPrefix(request.URL.Path, "/")
	if (filename == "") {
		filename = "page/home"
	}

	result, err := loadMarkdown(filename)
	if (err != nil) {
		log.Println(request.RemoteAddr, "Error:", err)
	} else {
		log.Println(request.RemoteAddr, filename)
	}

	io.WriteString(writer, result)
}

func loadMarkdown(filename string) (string, error) {
	md, err := ioutil.ReadFile(filename + ".md")
	if err != nil {
        md, err = ioutil.ReadFile("page/error.md")
    }
    if err != nil {
        return "400 server error", err
    }
    flags := html.CommonFlags | html.CompletePage
	opts := html.RendererOptions{
		Flags: flags,
		Title: "Thalkz's blog",
		CSS: "/static/style.css",
		Icon: "/static/favicon.ico",
	}
	renderer := html.NewRenderer(opts)
    result := string(markdown.ToHTML(md, nil, renderer))
    return result, nil
}

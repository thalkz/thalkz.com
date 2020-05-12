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
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
  	log.Println("Server is running on port 80")
	if err := http.ListenAndServe(":80", nil); err != nil {
    	panic(err)
  	}
}

func servePage(writer http.ResponseWriter, request *http.Request) {
	page := strings.TrimPrefix(request.URL.Path, "/")
	if (page == "") {
		page = "index"
	} 
	html, err := loadPage(page)
	if (err != nil) {
		io.WriteString(writer, "400 server error")
		log.Println(err)
	}
	log.Println("Request from", request.RemoteAddr, "for", page)
	io.WriteString(writer, string(html))
}

func loadPage(title string) ([]byte, error) {
    filename := "public/" + title + ".md"
    md, err := ioutil.ReadFile(filename)
    if err != nil {
        return nil, err
    }
    flags := html.CommonFlags | html.CompletePage
	opts := html.RendererOptions{
		Flags: flags,
		Title: "Thalkz's blog",
		CSS: "../static/style.css",
		Icon: "../static/favicon.ico",
	}
	renderer := html.NewRenderer(opts)
    result := markdown.ToHTML(md, nil, renderer)
    return result, nil
}

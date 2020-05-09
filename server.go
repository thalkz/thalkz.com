package main

import (
	"net/http"
	"io"
	"strings"
	"io/ioutil"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	 http.HandleFunc("/blog/", servePage)
	if err := http.ListenAndServe(":80", nil); err != nil {
    	panic(err)
  	}
}

func servePage(writer http.ResponseWriter, request *http.Request) {
	page := request.URL.Path
	page = strings.TrimPrefix(page, "/blog/")
	html, err := loadPage(page)
	if (err != nil) {
		io.WriteString(writer, "404 Page not foud")
	}
	io.WriteString(writer, string(html))
}

func loadPage(title string) ([]byte, error) {
    filename := "blog/" + title + ".txt"
    md, err := ioutil.ReadFile(filename)
    if err != nil {
        return nil, err
    }
    flags := html.CommonFlags | html.CompletePage | html.HrefTargetBlank
	opts := html.RendererOptions{
		Flags: flags,
		CSS: "../static/style.css",
	}
	renderer := html.NewRenderer(opts)
    result := markdown.ToHTML(md, nil, renderer)
    return result, nil
}

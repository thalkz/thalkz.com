package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
)

var flags = html.CommonFlags | html.CompletePage

func getRenderer(title string) markdown.Renderer {
	var opts = html.RendererOptions{
		Flags: flags,
		Title: fmt.Sprintf("%v | Roland Lmd", title),
		CSS:   "/static/style.css",
		Icon:  "/static/favicon.ico",
		Head: []byte(`<meta name=\"viewport\" content=\"width=device-width, initial-scale=1\">
			<script defer data-domain="thalkz.com" src="https://analytics.thalkz.com/js/script.js"></script>`),
	}
	return html.NewRenderer(opts)
}

func loadMarkdown(filename string) (string, error) {
	file, err := os.ReadFile(filename + ".md")
	if err != nil {
		file, _ = os.ReadFile("pages/error.md")
	}

	title := toTitle(filename)
	renderer := getRenderer(title)
	result := string(markdown.ToHTML(file, nil, renderer))
	return result, err
}

func servePage(writer http.ResponseWriter, request *http.Request) {
	filename := strings.TrimPrefix(request.URL.Path, "/")
	if filename == "" {
		filename = "pages/home"
	}

	result, err := loadMarkdown(filename)
	if err != nil {
		log.Println(request.RemoteAddr, "Error:", err)
	} else {
		log.Println(request.RemoteAddr, filename)
	}

	io.WriteString(writer, result)
}

func serveVersion(writer http.ResponseWriter, request *http.Request) {
	version := os.Getenv("VERSION")
	text := fmt.Sprintf("version: %v", version)
	renderer := getRenderer("Version")
	page := string(markdown.ToHTML([]byte(text), nil, renderer))
	io.WriteString(writer, page)
}

func main() {
	http.HandleFunc("/version", serveVersion)
	http.HandleFunc("/", servePage)
	staticFs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", staticFs))
	imagesFs := http.FileServer(http.Dir("images"))
	http.Handle("/images/", http.StripPrefix("/images/", imagesFs))

	log.Println("Server is running on port 80")

	if err := http.ListenAndServe(":80", nil); err != nil {
		panic(err)
	}
}

package main

import (
	"fmt"
	"net/http"

	"github.com/feynmaz/gophercises/exercise_2_url_shortener/handler"
)

func main() {
	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := handler.MapHandler(pathsToUrls, mux)

	// 	yaml := `
	// - path: /urlshort
	//   url: https://github.com/gophercises/urlshort
	// - path: /urlshort-final
	//   url: https://github.com/gophercises/urlshort/tree/solution
	// `

	// 	yamlFile := flag.String("file", "", "YAML file to load redirect rules from")
	// 	flag.Parse()
	// 	if yamlFile != nil && *yamlFile != "" {
	// 		byteContent, err := os.ReadFile(*yamlFile)
	// 		if err != nil {
	// 			log.Fatal(err)
	// 		}
	// 		yaml = string(byteContent)
	// 	}

	// 	yamlHandler, err := handler.YAMLHandler([]byte(yaml), mapHandler)
	// 	if err != nil {
	// 		panic(err)
	// 	}

	jsonData := `
	[
		{
			"path":"/urlshort",
			"url":"https://github.com/gophercises/urlshort"
		},
		{
			"path":"/urlshort-final",
			"url":"https://github.com/gophercises/urlshort/tree/solution"
		}
	]`

	JSONHandler, err := handler.JSONHandler([]byte(jsonData), mapHandler)
	if err != nil {
		panic(err)
	}

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", JSONHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}

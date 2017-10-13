package main

import(
	"net/http"
	"io"
	"io/ioutil"
)

func index(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html")

	bs,_ := ioutil.ReadFile("index.html")
	html := string(bs)
	io.WriteString(res, html)
}

func main() {
	http.Handle(
		"/assets/", 
		http.StripPrefix(
			"/assets/",
			http.FileServer(http.Dir("assets")),
		),
	)
	http.HandleFunc("/", index)
	http.ListenAndServe(":9000", nil)
}

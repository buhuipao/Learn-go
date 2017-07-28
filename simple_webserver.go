package main

import (
	"io"
	"net/http"
)

const form = `
	<http><body>
			<form action="#" method="post" name="bar">
					<input type="text" name="in" />
					<input type="submit" value="submit"/>
			</form>
	</body></html>
`

func IndexHandler(w http.ResponseWriter, request *http.Request) {
	io.WriteString(w, "<h1>Hello world!</h1>")
}

func FormHandler(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	switch request.Method {
	case "GET":
		io.WriteString(w, form)
	case "POST":
		io.WriteString(w, request.FormValue("in"))
	}
}

func main() {
	http.HandleFunc("/index", IndexHandler)
	http.HandleFunc("/form", FormHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

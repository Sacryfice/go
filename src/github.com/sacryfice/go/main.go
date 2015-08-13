package main

import (
	"bytes"
	"io"
	"net/http"
	"github.com/eknkc/amber"
)

func readJadeFile(filename string) string {
	person := make(map[string]string)
	person["Name"] = "Evgeniy"

	temp, err := amber.CompileFile(filename+".jade", amber.DefaultOptions)
	if err != nil {
		return `<div class="error">` + err.Error() + `</div>`
	}

	buf := new(bytes.Buffer)
	temp.Execute(buf, person)
	html := buf.String()

	return html
}

func hello(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"text/html",
	)
	io.WriteString(
		res,
		readJadeFile("test"),
	)
}

func main() {
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":80", nil)
}

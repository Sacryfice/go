package main

import (
	"io"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
}

var mux map[string]func(http.ResponseWriter, *http.Request)

func main() {
	server := http.Server{
		Addr:    ":8000",
		Handler: &myHandler{},
	}

	mux = make(map[string]func(http.ResponseWriter, *http.Request))
	mux["/"] = hello

	server.ListenAndServe()
}

type myHandler struct{}

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h, ok := mux[r.URL.String()]; ok {
		h(w, r)
		return
	}

	io.WriteString(w, "My server: "+r.URL.String())
}


/*package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/eknkc/amber"
)

type Path struct {
	Path string
	File string
}

type Config struct {
	Paths []Path
	Port  int
}

var config = readConfig()

func main() {
	for _, path := range config.Paths {
		addHandle(path)
	}

	http.ListenAndServe(":"+strconv.Itoa(config.Port), nil)
}

func addHandle(path Path) {
	var html = readJadeFile(path.File)
	var handler = func(res http.ResponseWriter, req *http.Request) {		
		res.Header().Set(
			"Content-Type",
			"text/html",
		)
		io.WriteString(
			res,
			html,
		)
	}

	http.HandleFunc(path.Path, handler)
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

func readConfig() *Config {
	content, err := ioutil.ReadFile("server.json")
	if err != nil {
		fmt.Print("Error:", err)
	}
	var conf Config
	err = json.Unmarshal(content, &conf)
	if err != nil {
		fmt.Print("Error:", err)
	}

	return &conf
}*/

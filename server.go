package main

// imports
import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// port flag
var port int

func main() {
	flag.IntVar(&port, "port", 8080, "Specify the port to listen to.")
	flag.Parse()
	// http routing
	http.HandleFunc("/pixel.gif", pixel)
	http.HandleFunc("/echo/", echo)
	http.HandleFunc("/code/", code)
	http.HandleFunc("/", pong)
	// server setup
	listenAddr := fmt.Sprint("127.0.0.1:", port)
	fmt.Printf("Server started at port %s", listenAddr)
	log.Fatal(http.ListenAndServe(listenAddr, nil))
}

func pong(w http.ResponseWriter, r *http.Request) {
	nowTime := strconv.FormatInt(time.Now().UnixMilli(), 10)
	w.Header().Set("meta-recv-ms", nowTime)
	fmt.Fprint(w, "pong")
}

func code(w http.ResponseWriter, r *http.Request) {
	httpCode := strings.TrimPrefix(r.URL.Path, "/code/")
	httpCodeInt, err := strconv.Atoi(httpCode)
	if err != nil {
		http.Error(w, "Invalid HTTP code", 400)
		return
	}
	http.Error(w, httpCode, httpCodeInt)
}

func echo(w http.ResponseWriter, r *http.Request) {
	prompt := strings.TrimPrefix(r.URL.Path, "/echo/")
	fmt.Fprint(w, prompt)
}

func pixel(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/gif")
	w.Write([]byte(`GIF89a     !ù  ,       L ;`))
}

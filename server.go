package main

// imports
import (
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// port flag
var port int
var sslCertificate string
var sslKey string

const VERSION = "1.0.4"

func init() {
	flag.StringVar(&sslCertificate, "ssl-cert", "", "path to SSL server certificate")
	flag.StringVar(&sslKey, "ssl-key", "", "path to SSL private key")
	flag.IntVar(&port, "port", 8080, "Specify the port to listen to.")
	flag.Parse()
}

func main() {
	err := server()
	if err != nil {
		log.Fatalf("server error %s", err)
	}
}

func server() error {
	// http routing
	http.HandleFunc("/pixel.gif", pixel)
	http.HandleFunc("/echo/", echo)
	http.HandleFunc("/code/", code)
	http.HandleFunc("/ping", pong)
	http.HandleFunc("/version", vers)
	http.HandleFunc("/", pong)
	cfg := &tls.Config{
		MinVersion:               tls.VersionTLS12,
		CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
		PreferServerCipherSuites: true,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
			tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_RSA_WITH_AES_256_CBC_SHA,
		},
	}
	// server setup
	listenAddr := fmt.Sprint(":", port)
	fmt.Printf("mchangrh/ping v%s listening on port %s\n", VERSION, listenAddr)
	srv := &http.Server{
		Addr:         listenAddr,
		TLSConfig:    cfg,
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
	}
	if sslCertificate != "" && sslKey != "" {
		log.Print("listening on HTTPS")
		return srv.ListenAndServeTLS(sslCertificate, sslKey)
	}
	return http.ListenAndServe(listenAddr, nil)
}

func pong(w http.ResponseWriter, r *http.Request) {
	cors(&w)
	fmt.Fprint(w, "pong")
}

func code(w http.ResponseWriter, r *http.Request) {
	cors(&w)
	httpCode := strings.TrimPrefix(r.URL.Path, "/code/")
	httpCode = strings.Split(httpCode, "/")[0]
	httpCodeInt, err := strconv.Atoi(httpCode)
	if err != nil {
		http.Error(w, "Invalid HTTP code", 400)
		return
	}
	http.Error(w, httpCode, httpCodeInt)
}

func vers(w http.ResponseWriter, r *http.Request) {
	cors(&w)
	fmt.Fprint(w, VERSION)
}

func echo(w http.ResponseWriter, r *http.Request) {
	cors(&w)
	prompt := strings.TrimPrefix(r.URL.Path, "/echo/")
	fmt.Fprint(w, prompt)
}

func pixel(w http.ResponseWriter, r *http.Request) {
	cors(&w)
	w.Header().Set("Content-Type", "image/gif")
	w.Write([]byte(`GIF89a     !ù  ,       L ;`))
}

func cors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Timing-Allow-Origin", "*")
}

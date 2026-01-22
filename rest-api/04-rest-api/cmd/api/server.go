package main

import (
	"crypto/tls"
	"fmt"
	"rest-api/internal/api/middlewares"

	"log"
	"net/http"
	"strings"
)

type user struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	City string `json:"city"`
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Hello root route")
	fmt.Printf("%v request on %v route\n", r.Method, r.URL)
	w.Write([]byte(fmt.Sprintf("%v request on %v route\n", r.Method, r.URL)))
	switch r.Method {
	case http.MethodGet:
	case http.MethodPost:
	case http.MethodPut:
	case http.MethodPatch:
	case http.MethodDelete:
	default:
	}

}

func teachersHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Printf("%v request on %v route\n", r.Method, r.URL)
	w.Write([]byte(fmt.Sprintf("%v request on %v route\n", r.Method, r.URL)))

	// teachers/{id}
	// teachers/9
	// teachers/?key=value&query=value2&sortby=email&sortorder=ASC

	fmt.Println(r.URL.Path)
	path := strings.TrimPrefix(r.URL.Path, "/teachers/")
	userID := strings.TrimSuffix(path, "/")

	fmt.Println("The id is: ", userID)

	// fmt.Println("Query params:", r.URL.Query())
	queryParams := r.URL.Query()
	sortby := queryParams.Get("sortby")
	key := queryParams.Get("key")
	sortorder := queryParams.Get("sortorder")

	if sortorder == "" {
		sortorder = "DESC"
	}

	fmt.Println(sortby, key, sortorder)

	switch r.Method {
	case http.MethodPost:
	case http.MethodPut:
	case http.MethodPatch:
	case http.MethodDelete:
	default:
	}

}

func studentsHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Printf("%v request on %v route\n", r.Method, r.URL)
	w.Write([]byte(fmt.Sprintf("%v request on %v route\n", r.Method, r.URL)))

	switch r.Method {
	case http.MethodGet:
	case http.MethodPost:
	case http.MethodPut:
	case http.MethodPatch:
	case http.MethodDelete:
	default:
	}

}

func execsHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Printf("%v request on %v route\n", r.Method, r.URL)
	w.Write([]byte(fmt.Sprintf("%v request on %v route\n", r.Method, r.URL)))

	switch r.Method {
	case http.MethodGet:
	case http.MethodPost:
	case http.MethodPut:
	case http.MethodPatch:
	case http.MethodDelete:
	default:
	}

}

func main() {

	port := ":3000"

	cert := "cert.pem"
	key := "key.pem"

	mux := http.NewServeMux()

	mux.HandleFunc("/", rootHandler)

	mux.HandleFunc("/teachers/", teachersHandler)

	mux.HandleFunc("/students/", studentsHandler)

	mux.HandleFunc("/execs/", execsHandler)

	tlsConfg := tls.Config{
		MinVersion: tls.VersionTLS12,
	}

	server := &http.Server{
		Addr:      port,
		Handler:   middlewares.SecurityHeaders(middlewares.Cors(mux)),
		TLSConfig: &tlsConfg,
	}

	fmt.Printf("Server is running on port %v\n", port)
	err := server.ListenAndServeTLS(cert, key)
	if err != nil {
		log.Fatalln("Error starting the server", err)
	}
}

/*
	// processing a reuest and printing out its components

	// parse form data for x-www-form-urlencoded
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form.", http.StatusBadRequest)
		return
	}
	fmt.Println("Form:", r.Form)

	response := make(map[string]interface{})
	for k, v := range r.Form {
		response[k] = v[0]
	}

	fmt.Println("Processed response map:", response)

	// RAW body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error parsing form.", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	fmt.Println(string(body))

	// var response1 user
	response2 := make(map[string]interface{})
	err = json.Unmarshal(body, &response2)
	// err = json.Unmarshal(body, &response1)
	if err != nil {
		log.Fatal("error unmarshaling:", err)
		return
	}

	fmt.Printf("%+v\n", response2)

	// things present in request
	fmt.Println("Body:", r.Body)
	fmt.Println("Form:", r.Form)
	fmt.Println("Header:", r.Header)
	fmt.Println("Context:", r.Context())
	fmt.Println("Context Length:", r.ContentLength)
	fmt.Println("Host:", r.Host)
	fmt.Println("Method:", r.Method)
	fmt.Println("Proto:", r.Proto)
	fmt.Println("Proto Minor:", r.ProtoMinor)
	fmt.Println("Proto Major:", r.ProtoMajor)
	fmt.Println("Remote Addr", r.RemoteAddr)
	fmt.Println("Request URI", r.RequestURI)
	fmt.Println("TLS:", r.TLS)
	fmt.Println("Trailer:", r.Trailer)
	fmt.Println("URL:", r.URL)
	fmt.Println("Transfer encoding:", r.TransferEncoding)
	fmt.Println("UserAgent:", r.UserAgent())
	fmt.Println("Port:", r.URL.Port())
	fmt.Println("Url scheme:", r.URL.Scheme)


*/

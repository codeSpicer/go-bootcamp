package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/http2"
)

// curl -v -k https://localhost:3000/orders

func getHttpVersion(r *http.Request) {
	if r.ProtoMajor == 2 {
		fmt.Println("[Server] Handling order via HTTP/2")
	} else {
		fmt.Println("[Server] Handling order via HTTP/1.1")
	}
}

func getTlsVersion(version uint16) {
	switch version {
	case tls.VersionTLS10:
		fmt.Println("TLS version 10")
	case tls.VersionTLS11:
		fmt.Println("TLS version 11")
	case tls.VersionTLS12:
		fmt.Println("TLS version 12")
	case tls.VersionTLS13:
		fmt.Println("TLS version 13")
	}
}

func main() {
	port := 3000

	// Routes
	http.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
		// We can check if the request is actually HTTP/2
		getHttpVersion(r)
		if r.TLS != nil {
			getTlsVersion(r.TLS.Version)
		}
		fmt.Fprintf(w, "Handling incoming orders")
	})

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Handling Users")
	})

	// Certificate paths
	cert := "cert.pem"
	key := "key.pem"

	// Configure TLS
	// We set MinVersion to 1.2 because HTTP/2 requires at least TLS 1.2
	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
		// NextProtos is required for the browser to negotiate HTTP/2 (h2)
		NextProtos: []string{"h2", "http/1.1"},
	}

	server := &http.Server{
		Addr:      fmt.Sprintf(":%d", port),
		TLSConfig: tlsConfig,
	}

	// Enable HTTP/2 explicitly
	// This adds the HTTP/2 framing layer to your server
	err := http2.ConfigureServer(server, &http2.Server{})
	if err != nil {
		log.Fatalf("Failed to configure HTTP/2: %v", err)
	}

	fmt.Printf("Server is running on https://localhost:%d\n", port)
	fmt.Println("Note: You MUST use https:// and click 'Advanced' -> 'Proceed' in your browser.")

	// Start TLS Server
	err = server.ListenAndServeTLS(cert, key)
	if err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}

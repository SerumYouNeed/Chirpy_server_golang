package main

import (
	"log"
	"net/http"
)
// Handlers are responsible for carrying out your application logic and writing response headers and bodies. Whereas a servemux (also known as a _router_) stores a mapping between the predefined URL paths for your application and the corresponding handlers. Usually you have one servemux for your application containing all your routes.

func main() {
	const filepathRoot = "."
	const port = "8080"

	// tworzę pusty servemux
	mux := http.NewServeMux() 
	// http.FileServer zwraca wbudowany http.Handler
	// mux.Handle rejestruje servemux żeby zachowywał się jak handler
	// dla wszystkich requestów pod URL "/"
	mux.Handle("/", http.FileServer(http.Dir(filepathRoot)))

	// hhtp.server to struct konfigurujący ustawienia serwera
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Serving files from %s on port: %s\n", filepathRoot, port)
	log.Fatal(srv.ListenAndServe()) 
	// po wezwaniu ListenAndServe main() jest blokowane aż serwer nie zostanie zamknięty

}

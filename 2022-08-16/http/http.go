package main

import "net/http"

func main() {

	serveMux := http.NewServeMux()

	serveMux.HandleFunc("/test/name/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("/test/name"))
	})

	serveMux.HandleFunc("/test/name/sex/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("/test/name/sex"))
	})

	// http.ListenAndServe()

	(&http.Server{
		Addr:              ":8080",
		Handler:           serveMux,
		TLSConfig:         nil,
		ReadTimeout:       0,
		ReadHeaderTimeout: 0,
		WriteTimeout:      0,
		IdleTimeout:       0,
		MaxHeaderBytes:    0,
		TLSNextProto:      nil,
		ConnState:         nil,
		ErrorLog:          nil,
		BaseContext:       nil,
		ConnContext:       nil,
	}).ListenAndServe()

}

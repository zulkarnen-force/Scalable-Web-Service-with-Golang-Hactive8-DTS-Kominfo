package main

import (
	"fmt"
	"log"
	"net/http"
)


func greet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`hello world`))
}


func middleware1(next http.Handler) http.Handler {
	
	return http.HandlerFunc( func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Middleware 1")
		next.ServeHTTP(w, r)
	} )

}


func middleware2(next http.Handler) http.Handler {
	return http.HandlerFunc( func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Middleware 2")
		next.ServeHTTP(w, r)
	})
}


func main() {
	mux := http.NewServeMux()

	endpoint := http.HandlerFunc(greet)

	mux.Handle("/", middleware1(middleware2(endpoint)))

	err := http.ListenAndServe(":8080", mux)
	fmt.Println("serve on http://localhost:8080/")

	if err != nil {
		log.Fatal(err)
	}


}
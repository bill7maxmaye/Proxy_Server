package main

import "net/http"

  
func main() {
	//
	mux := http.NewServeMux()
	mux.HandleFunc("/", getDoc)
	http.ListenAndServe(":8081", mux)
}


//
func getDoc(w http.ResponseWriter, r *http.Request){
	http.Redirect(w, r, "https://go.dev/doc/", http.StatusSeeOther)

	w.Header().Set("Content-Type", "text/plain")

	// Write "Hello" to the response
	w.Write([]byte("Hello the server responded successfuly"))
}


//This code sets up a basic web server that listens on port 8081 and redirects all requests to https://go.dev/doc/.

//The getDoc function is a handler function that gets called when a request is made to the root path ("/").
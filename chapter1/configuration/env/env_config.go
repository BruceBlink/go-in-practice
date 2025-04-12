package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	var port string
	if port = os.Getenv("PORT"); port == "" {
		panic("environment variable PORT has not been set!")
	}

	http.HandleFunc("/", homePage)
	fmt.Printf("Server start on http://localhost:%s \n", port)
	http.ListenAndServe(":"+port, nil)

}

func homePage(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(res, req)
		return
	}
	fmt.Fprint(res, "The homepage.")
}

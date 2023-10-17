package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

const startupMessage = `                 ___________________
                /   ________ -[]--. \
               / ,-'         ` + "`" + `-.   \ \
              / (       o       )  _) \
             /   ` + "`" + `-._________,-'_ /_/-.\
            /  __ _   Phasik   " " "    \ 
           /_____________________________\
           "-=-------------------------=-"
		   phasik.tv started!`

func main() {
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "Hello! you've requested %s\n", r.URL.Path)
	// })
	fs := http.FileServer(http.Dir("/srv/www"))
	http.Handle("/", fs)

	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}

	for _, line := range strings.Split(startupMessage, "\n") {
		fmt.Println(line)
	}
	fmt.Printf("Server listening at :%s 🚀\n", port)

	err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	if err != nil {
		panic(err)
	}
}

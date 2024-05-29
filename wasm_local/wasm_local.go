package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	const port = ":8080"
	if len(os.Args) < 2 {
		log.Panicln("Folder name not specifed. Aborting now...")
		panic("Folder path not specified")
	}
	path := os.Args[1]
	http.Handle("/", http.FileServer(http.Dir(path)))
	fmt.Printf("Now hosting wasm application at http://localhost%s\n", port)
	fmt.Println("Press Ctrl+C to close the server...")
	log.Fatal(http.ListenAndServe(port, nil))
}

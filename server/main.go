package main

import "log"
import "net/http"
import "fmt"

func main() {
	fmt.Println("Server have started")
	log.Fatal(http.ListenAndServe(":8080", Handlers()))
}

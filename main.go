package main

	import (
		"log"
		"os"
		"fmt"
		"net/http"
	)

	func main() {
		port := os.Getenv("PORT")

		if port == "" {
			log.Fatal("$PORT must be set")
		} else {
			fmt.Println("go-jitshappens is starting!")
		}
		log.Fatal(http.ListenAndServe(":" + port, http.FileServer(http.Dir("./public"))))
	}

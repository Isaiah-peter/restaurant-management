package main

import "os"

func main() {
	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = ":4000"
	}

}

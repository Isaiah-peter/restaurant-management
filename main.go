package main

import (
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = ":4000"
	}

}

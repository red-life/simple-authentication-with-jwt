package cmd

import (
	"github.com/red-life/simple-authentication-with-jwt/internal/routes"
)

func main() {
	r := router.InitRouter()
	if err := r.Run(":5000"); err != nil {
		panic(err)
	}
}

package cmd

import (
	"github.com/red-life/simple-authentication-with-jwt/internal/router"
)

func main() {
	r := router.NewRouter()
	if err := r.Run(); err != nil {
		panic(err)
	}
}

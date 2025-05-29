package main

import (
	"fmt"

	"github.com/hadihalimm/sigizi-rsam/internal/api"
)

func main() {
	server := api.NewServer()
	err := server.ListenAndServeTLS("cert.crt", "key.key")
	// err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("Cannot start server: %s", err))
	}
}

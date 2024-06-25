package pkg

import (
	"log"
	"net/http"
)

func StartServer(port string) {
	log.Println("INFO StartServer: starting server on :" + port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("ERROR StartServer fatal error: %v", err)
	}

}

package health

import (
	"fmt"
	"log"
	"net/http"
)

func Start() {
	go func() {
		http.HandleFunc("/health", func(w http.ResponseWriter, _ *http.Request) {
			w.WriteHeader(http.StatusOK)
			fmt.Println(w, "OK")
		})
		log.Println("Health server listening on port 8080")
		log.Fatal(http.ListenAndServe(":8080", nil))
	}()
}

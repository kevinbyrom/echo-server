package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const (
	ENV_HOST          = "CONN_HOST"
	ENV_PORT          = "CONN_PORT"
	DEFAULT_CONN_HOST = "127.0.0.1"
	DEFAULT_CONN_PORT = "8080"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func main() {
	http.HandleFunc("/", helloWorld)

	conn := envOrDefault(ENV_HOST, DEFAULT_CONN_HOST) + ":" + envOrDefault(ENV_PORT, DEFAULT_CONN_PORT)
	log.Printf("Listening on %s...\n", conn)

	err := http.ListenAndServe(conn, nil)
	if err != nil {
		log.Fatal("error starting http server: ", err)
		return
	}
}

func envOrDefault(envName, defVal string) string {
	config, ex := os.LookupEnv(envName)
	if ex {
		return config
	}

	return defVal
}

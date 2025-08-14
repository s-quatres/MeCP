package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"

	"github.com/s-quatres/MeCP/pkg/ai"
	"github.com/s-quatres/MeCP/pkg/k8s"
	"github.com/s-quatres/MeCP/pkg/log"
)

func main() {
	// Load config from YAML file
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	// Create AI MCP server
	aiServer := ai.NewMCP()

	// Create K8s manifests endpoint
	k8sHandler := k8s.NewManifestsHandler()

	// Create log collector endpoint
	logHandler := log.NewLogCollector()

	// Create router
	r := mux.NewRouter()
	r.HandleFunc("/k8s/manifests", k8sHandler.GetManifests).Methods("GET")
	r.HandleFunc("/logs/loglines", logHandler.GetLogLines).Methods("GET")

	// Start server
	fmt.Println("Server listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
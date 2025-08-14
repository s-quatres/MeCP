package k8s

import (
	"encoding/json"
	"net/http"
)

type ManifestsHandler struct{}

func NewManifestsHandler() *ManifestsHandler {
	return &ManifestsHandler{}
}

func (m *ManifestsHandler) GetManifests(w http.ResponseWriter, r *http.Request) {
	// Load K8s manifests from YAML file
	manifests := []byte(`apiVersion: apps/v1
kind: Deployment
metadata:
  name: my-deployment
spec:
  replicas: 3
  selector:
    matchLabels:
      app: my-app
  template:
    metadata:
      labels:
        app: my-app
    spec:
      containers:
      - name: my-container
        image: my-image:latest`)

	// Return manifests as JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(manifests)
}
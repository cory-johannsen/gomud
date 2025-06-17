package swaggerui

import (
	"github.com/cory-johannsen/gomud/internal/config"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type SwaggerUI struct {
	cfg *config.Config
}

func NewSwaggerUI(cfg *config.Config) *SwaggerUI {
	return &SwaggerUI{
		cfg: cfg,
	}
}

func StartSwaggerUIServer(server *SwaggerUI) error {
	// Serve Swagger JSON
	mux := http.NewServeMux()
	mux.HandleFunc("/swagger/mud.swagger.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "generated/mud.swagger.json")
	})
	// Serve Swagger UI static files
	fs := http.FileServer(http.Dir("generated/swagger-ui"))
	mux.Handle("/swagger/", http.StripPrefix("/swagger/", fs))

	// Start swagger UI server on the configured port
	log.Printf("REST gateway listening at %s", server.cfg.SwaggerUIAddress)
	return http.ListenAndServe(server.cfg.SwaggerUIAddress, mux)
}

package handler

import (
	"net/http"

	. "github.com/tbxark/g4vercel"
	"github.com/tbxark/g4vercel-demo/api/src/routes"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	server := New()
	routes.Router(server)
	server.Handle(w, r)
}

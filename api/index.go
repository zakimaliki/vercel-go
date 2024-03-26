package handler

import (
	"net/http"

	"github.com/subosito/gotenv"
	. "github.com/tbxark/g4vercel"
	"github.com/tbxark/g4vercel-demo/api/src/config"
	"github.com/tbxark/g4vercel-demo/api/src/helper"
	"github.com/tbxark/g4vercel-demo/api/src/routes"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	server := New()
	gotenv.Load()
	config.InitDB()
	helper.Migration()
	defer config.DB.Close()
	routes.Router(server)
	server.Handle(w, r)
}

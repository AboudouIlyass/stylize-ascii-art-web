package routes

import (
	"net/http"
	"path"

	"artweb/helpers"
	"artweb/model"
)

func StaticsRouter(w http.ResponseWriter, r *http.Request) {
	if path.Ext(r.URL.Path) != ".css" && path.Ext(r.URL.Path) != ".jpg" {
		helpers.ParseAndExecute("templates/error.html", model.Data{Error: "403 Status Forbidden"}, http.StatusForbidden, w, false)
		return
	}
	http.StripPrefix("/statics/", http.FileServer(http.Dir("statics"))).ServeHTTP(w, r)
}

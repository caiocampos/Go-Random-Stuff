package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"go.db.restapi/config"
)

type controller struct {
	router *mux.Router
	user   userController
}

var ctrl *controller

// Init method boots the end-points for the server
func Init() {
	if ctrl == nil {
		ctrl = &controller{}
		ctrl.router = mux.NewRouter()
		ctrl.user = userController{}
		ctrl.user.init(ctrl.router)

		config.ReadTOML()
		port := strconv.Itoa(config.TOMLConfig.App.Port)
		if err := http.ListenAndServe(":"+port, ctrl.router); err != nil {
			log.Fatal(err)
		}
	}
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

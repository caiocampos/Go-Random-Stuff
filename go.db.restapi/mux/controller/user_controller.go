package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"

	"go.db.restapi/model"
	serv "go.db.restapi/service"
)

type userController struct{ serv *serv.UserService }

func (u *userController) init(router *mux.Router) {
	if u.serv == nil {
		u.serv = &serv.UserService{}
		router.HandleFunc("/user", u.findAll).Methods("GET")
		router.HandleFunc("/user/id/{id}", u.findByID).Methods("GET")
		router.HandleFunc("/user/name/{name}", u.findByName).Methods("GET")
		router.HandleFunc("/user", u.insert).Methods("POST")
		router.HandleFunc("/user", u.delete).Methods("DELETE")
		router.HandleFunc("/user", u.update).Methods("PUT")
	}
}

func (u *userController) findAll(w http.ResponseWriter, r *http.Request) {
	users, err := u.serv.FindAll()
	if err != nil {
		respondWithJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	if users == nil {
		users = []model.User{}
	}
	respondWithJSON(w, http.StatusOK, users)
}

func (u *userController) findByName(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	user, err := u.serv.FindByName(params["name"])
	if err != nil {
		respondWithJSON(w, http.StatusBadRequest, "Name not found")
		return
	}
	respondWithJSON(w, http.StatusOK, user)
}

func (u *userController) findByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	user, err := u.serv.FindByID(params["id"])
	if err != nil {
		respondWithJSON(w, http.StatusBadRequest, "ID not found")
		return
	}
	respondWithJSON(w, http.StatusOK, user)
}

func (u *userController) insert(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		respondWithJSON(w, http.StatusBadRequest, "Invalid request")
		return
	}
	user.ID = bson.NewObjectId()
	if err := u.serv.Insert(user); err != nil {
		respondWithJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusCreated, user)
}

func (u *userController) delete(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		respondWithJSON(w, http.StatusBadRequest, "Invalid request")
		return
	}
	if err := u.serv.Delete(user); err != nil {
		respondWithJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, "OK")
}

func (u *userController) update(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		respondWithJSON(w, http.StatusBadRequest, "Invalid request")
		return
	}
	if err := u.serv.Update(user); err != nil {
		respondWithJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, "OK")
}

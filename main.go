package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/globalsign/mgo/bson"
	"github.com/gorilla/mux"
	cfg "github.com/martinroddam/lists/config"
	dao "github.com/martinroddam/lists/dao"
	"github.com/martinroddam/lists/model"
	log "github.com/sirupsen/logrus"
)

var config = cfg.Config{}

var mongo = dao.ListsDAO{}

func CreateUserEndPoint(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	user.ID = bson.NewObjectId()

	if err := mongo.InsertUser(user); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusCreated, user)
}

func GetUserEndPoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	user, err := mongo.FindUserById(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid User ID")
		return
	}
	respondWithJSON(w, http.StatusOK, user)
}

func GetAllUsersEndPoint(w http.ResponseWriter, r *http.Request) {
	users, err := mongo.FindAllUsers()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, users)
}

func GetAllListsEndPoint(w http.ResponseWriter, r *http.Request) {
	lists, err := mongo.FindAllLists()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, lists)
}

func UpdateListsEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

func GetListEndPoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	list, err := mongo.FindListById(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid List ID")
		return
	}
	respondWithJSON(w, http.StatusOK, list)
}

func CreateListEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var list model.List
	if err := json.NewDecoder(r.Body).Decode(&list); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	list.ID = bson.NewObjectId()

	if err := mongo.InsertList(list); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusCreated, list)
}

func UpdateListEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

func DeleteListEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

func GetAllTasksForList(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

func CreateTaskEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

func GetTaskEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

func UpdateTaskEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

func CompleteTaskEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

func GetTaskHistoryEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/user", CreateUserEndPoint).Methods("POST")
	r.HandleFunc("/user/{id}", GetUserEndPoint).Methods("GET")
	r.HandleFunc("/users", GetAllUsersEndPoint).Methods("GET")

	r.HandleFunc("/lists", GetAllListsEndPoint).Methods("GET")
	r.HandleFunc("/lists", UpdateListsEndPoint).Methods("PUT")

	r.HandleFunc("/list/{id}", GetListEndPoint).Methods("GET")
	r.HandleFunc("/list", CreateListEndPoint).Methods("POST")
	r.HandleFunc("/list/{id}", UpdateListEndPoint).Methods("PUT")
	r.HandleFunc("/list/{id}", DeleteListEndPoint).Methods("DELETE")
	r.HandleFunc("/list/{id}/tasks", GetAllTasksForList).Methods("GET")

	r.HandleFunc("/task/", CreateTaskEndPoint).Methods("POST")
	r.HandleFunc("/task/{id}", GetTaskEndPoint).Methods("GET")
	r.HandleFunc("/task/{id}", UpdateTaskEndPoint).Methods("PUT")
	r.HandleFunc("/task/{id}/completed", CompleteTaskEndPoint).Methods("POST")
	r.HandleFunc("/task/{id}/history", GetTaskHistoryEndPoint).Methods("GET")

	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}

func respondWithError(w http.ResponseWriter, statusCode int, errorMsg string) {

	respondWithJSON(w, statusCode, model.ListsAPIError{ErrorCode: statusCode, ErrorMessage: errorMsg})
}

// RespondWithJSON creates a json response with a input message and code
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		log.WithError(err).Error("error during json marshalling of payload for response")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		returnErr := fmt.Sprintf(`{"code": %v, "error": "error during json marshalling of payload for response"}`, http.StatusInternalServerError)
		w.Write([]byte(returnErr))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func init() {
	config.Read()

	mongo.Server = config.Server
	mongo.Database = config.Database
	mongo.Connect()
}

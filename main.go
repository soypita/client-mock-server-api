package main

import (
	"github.com/gorilla/mux"
	. "github.com/soypita/client-mock-server-api/controllers"
	. "github.com/soypita/client-mock-server-api/dao"
	"log"
	"net/http"
	"os"
)

var controller = DataObjectController{}

func init() {
	dao := &ObjectDao{}
	dao.Database = os.Getenv("DATABASE")
	dao.Server = os.Getenv("MONGODB_URI")
	dao.Connect()
	controller.ObjectDao = dao
}

func main() {
	port := os.Getenv("PORT")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/resources/{groupName}/create", controller.CreateNewObjectInGroup).Methods("POST")
	router.HandleFunc("/api/resources/{groupName}/getById/{id}", controller.GetObjectInGroupById).Methods("GET")
	router.HandleFunc("/api/resources/{groupName}/getAll", controller.GetAllDataInGroups).Methods("GET")
	router.HandleFunc("/api/resources/{groupName}/deleteAll", controller.DeleteAllObjectsInGroup).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":"+port, router))
}

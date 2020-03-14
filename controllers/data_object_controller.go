package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	. "github.com/soypita/client-mock-server-api/dao"
	. "github.com/soypita/client-mock-server-api/models"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"net/url"
)

type DataObjectController struct {
	ObjectDao *ObjectDao
}

const (
	VehicleGroup = "vehicle"
	TrafficGroup = "traffic"
	HandleGroup = "handle"
)

func (d DataObjectController) CreateNewObjectInGroup(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	groupName := params["groupName"]
	defer r.Body.Close()

	switch groupName {
	case VehicleGroup:
		var vehicleList []VehiclesModel

		if err := json.NewDecoder(r.Body).Decode(&vehicleList); err != nil {
			respondWithError(w, http.StatusBadRequest, "Invalid Request")
			return
		}
		for i, _ := range vehicleList {
			vehicleList[i].ID = bson.NewObjectId()
		}
		if err := d.ObjectDao.InsertNewVehiclesObject(vehicleList); err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		respondWithJson(w, http.StatusOK, `{"status": "success"}`)

	case TrafficGroup:
		var trafficList []TrafficInfraModel

		if err := json.NewDecoder(r.Body).Decode(&trafficList); err != nil {
			respondWithError(w, http.StatusBadRequest, "Invalid Request")
			return
		}
		for i, _ := range trafficList {
			trafficList[i].ID = bson.NewObjectId()
		}
		if err := d.ObjectDao.InsertNewTrafficInfraObject(trafficList); err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		respondWithJson(w, http.StatusOK, `{"status": "success"}`)

	case HandleGroup:

		var trafficList []TrafficInfraModel

		if err := json.NewDecoder(r.Body).Decode(&trafficList); err != nil {
			respondWithError(w, http.StatusBadRequest, "Invalid Request")
			return
		}
		for i, _ := range trafficList {
			trafficList[i].ID = bson.NewObjectId()
		}
		if err := d.ObjectDao.InsertNewTrafficInfraObject(trafficList); err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		respondWithJson(w, http.StatusOK, `{"status": "success"}`)

	default:
		respondWithError(w, http.StatusBadRequest, "Invalid group name")
	}
}

func (d DataObjectController) GetObjectInGroupById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	groupName := params["groupName"]
	objectId := params["id"]

	switch groupName {
	case VehicleGroup:
		vehicle, err := d.ObjectDao.GetVehiclesByVin(objectId)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		respondWithJson(w, http.StatusOK, vehicle)
	case TrafficGroup:
		decodeId, err := url.QueryUnescape(objectId)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		trafficInfra, err := d.ObjectDao.GetTrafficInfraByRegistryNumber(decodeId)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		respondWithJson(w, http.StatusOK, trafficInfra)
	default:
		respondWithError(w, http.StatusBadRequest, "Invalid group name")
	}
}

func (d DataObjectController) GetAllDataInGroups(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	groupName := params["groupName"]

	switch groupName {
	case VehicleGroup:
		vehicles, err := d.ObjectDao.GetAllVehicles()
		if err != nil {
			respondWithError(w, http.StatusNotFound, err.Error())
			return
		}
		respondWithJson(w, http.StatusOK, vehicles)
	case TrafficGroup:
		trafficList, err := d.ObjectDao.GetAllTrafficInfra()
		if err != nil {
			respondWithError(w, http.StatusNotFound, err.Error())
			return
		}
		respondWithJson(w, http.StatusOK, trafficList)

	case HandleGroup:
		handleList, err := d.ObjectDao.GetAllHandles()
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		respondWithJson(w, http.StatusOK, handleList)

	default:
		respondWithError(w, http.StatusBadRequest, "Invalid group name")
	}
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
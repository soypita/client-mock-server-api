package dao

import (
	. "github.com/soypita/client-mock-server-api/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type ObjectDao struct {
	Server   string
	Database string
}

var db *mgo.Session

const (
	CollectionCargo          = "cargo"
	CollectionVehicles       = "vehicles"
	CollectionTransportInfra = "transport_infra"
	CollectionHandles        = "handles"
)

func (m *ObjectDao) Connect() {
	session, err := mgo.Dial(m.Server) // establish connection
	if err != nil {                    // if connection failed to establish
		log.Fatal(err) // print error log
	}
	db = session

	m.addIndexes()
}

func (m *ObjectDao) addIndexes() {
	var err error
	vehiclesIndex := mgo.Index{
		Key:        []string{"vin"},
		Unique:     true,
		Background: true,
		Sparse:     true,
	}
	handleIndex := mgo.Index{
		Key:        []string{"handleId"},
		Unique:     true,
		Background: true,
		Sparse:     true,
	}
	trafficInfraIndex := mgo.Index{
		Key:        []string{"registryNumber"},
		Unique:     true,
		Background: true,
		Sparse:     true,
	}

	session := db.Copy()
	defer session.Close()
	handleColl := session.DB(m.Database).C(CollectionHandles)
	vehicleColl := session.DB(m.Database).C(CollectionVehicles)
	transportInfra := session.DB(m.Database).C(CollectionTransportInfra)

	err = handleColl.EnsureIndex(handleIndex)
	if err != nil {
		log.Fatalf("[addIndex] %s\n", err)
	}
	err = vehicleColl.EnsureIndex(vehiclesIndex)
	if err != nil {
		log.Fatalf("[addIndex] %s\n", err)
	}
	err = transportInfra.EnsureIndex(trafficInfraIndex)
	if err != nil {
		log.Fatalf("[addIndex] %s\n", err)
	}
}

// Traffic Infrastructure Objects methods
func (m *ObjectDao) GetAllTrafficInfra() ([]TrafficInfraModel, error) {
	var trafficInfraList []TrafficInfraModel
	session := db.Clone()
	defer session.Close()

	err := session.DB(m.Database).C(CollectionTransportInfra).Find(bson.M{}).All(&trafficInfraList)
	return trafficInfraList, err
}

func (m *ObjectDao) GetTrafficInfraByRegistryNumber(number string) (TrafficInfraModel, error) {
	var trafficInfra TrafficInfraModel
	session := db.Clone()
	defer session.Close()

	err := session.DB(m.Database).C(CollectionTransportInfra).Find(bson.M{"registryNumber": number}).One(&trafficInfra)
	return trafficInfra, err
}

func (m *ObjectDao) InsertNewTrafficInfraObject(objList []TrafficInfraModel) error {
	session := db.Clone()
	defer session.Close()

	var insData []interface{}

	for _, val := range objList {
		insData = append(insData, val)
	}

	collection := session.DB(m.Database).C(CollectionTransportInfra)
	bulk := collection.Bulk()
	bulk.Unordered()
	bulk.Insert(insData...)
	_, err := bulk.Run()
	return err
}

func (m *ObjectDao) InsertNewCargoObject(objList []CargoModel) error {
	session := db.Clone()
	defer session.Close()

	var insData []interface{}

	for _, val := range objList {
		insData = append(insData, val)
	}

	collection := session.DB(m.Database).C(CollectionCargo)
	bulk := collection.Bulk()
	bulk.Unordered()
	bulk.Insert(insData...)
	_, err := bulk.Run()
	return err
}

func (m *ObjectDao) GetCargoByRegistryNumber(number string) (TrafficInfraModel, error) {
	var trafficInfra TrafficInfraModel
	session := db.Clone()
	defer session.Close()

	err := session.DB(m.Database).C(CollectionCargo).Find(bson.M{"registryNumber": number}).One(&trafficInfra)
	return trafficInfra, err
}

func (m *ObjectDao) DeleteAllTrafficInfraObjects() error {
	session := db.Clone()
	defer session.Close()

	_, err := session.DB(m.Database).C(CollectionTransportInfra).RemoveAll(bson.M{})

	return err
}

func (m *ObjectDao) DeleteAllCargoObjects() error {
	session := db.Clone()
	defer session.Close()

	_, err := session.DB(m.Database).C(CollectionCargo).RemoveAll(bson.M{})

	return err
}

// Vehicles Objects methods
func (m *ObjectDao) GetAllVehicles() ([]VehiclesModel, error) {
	session := db.Clone()
	defer session.Close()

	var vehiclesList []VehiclesModel
	err := session.DB(m.Database).C(CollectionVehicles).Find(bson.M{}).All(&vehiclesList)
	return vehiclesList, err
}

func (m *ObjectDao) GetVehiclesByVin(vin string) (VehiclesModel, error) {
	session := db.Clone()
	defer session.Close()

	var vehiclesModel VehiclesModel
	err := session.DB(m.Database).C(CollectionVehicles).Find(bson.M{"vin": vin}).One(&vehiclesModel)
	return vehiclesModel, err
}

func (m *ObjectDao) InsertNewVehiclesObject(objList []VehiclesModel) error {
	session := db.Clone()
	defer session.Close()

	var insData []interface{}

	for _, val := range objList {
		insData = append(insData, val)
	}

	collection := session.DB(m.Database).C(CollectionVehicles)
	bulk := collection.Bulk()
	bulk.Unordered()
	bulk.Insert(insData...)
	_, err := bulk.Run()
	return err
}

func (m *ObjectDao) DeleteAllVehicleObjects() error {
	session := db.Clone()
	defer session.Close()

	_, err := session.DB(m.Database).C(CollectionVehicles).RemoveAll(bson.M{})

	return err
}

// Handles methods
func (m *ObjectDao) GetAllHandles() ([]HandlesModel, error) {
	session := db.Clone()
	defer session.Close()

	var handlesList []HandlesModel
	err := session.DB(m.Database).C(CollectionHandles).Find(bson.M{}).All(&handlesList)
	return handlesList, err
}

func (m *ObjectDao) GetAllHandlesInGroup(group string) ([]HandlesModel, error) {
	session := db.Clone()
	defer session.Close()

	var handlesModel []HandlesModel
	err := session.DB(m.Database).C(CollectionHandles).Find(bson.M{"group": group}).All(&handlesModel)
	return handlesModel, err
}

func (m *ObjectDao) InsertNewHandles(objList []HandlesModel) error {
	session := db.Clone()
	defer session.Close()

	var insData []interface{}

	for _, val := range objList {
		insData = append(insData, val)
	}

	collection := session.DB(m.Database).C(CollectionHandles)
	bulk := collection.Bulk()
	bulk.Unordered()
	bulk.Insert(insData...)
	_, err := bulk.Run()
	return err
}

func (m *ObjectDao) DeleteAllHandleObjects() error {
	session := db.Clone()
	defer session.Close()

	_, err := session.DB(m.Database).C(CollectionHandles).RemoveAll(bson.M{})

	return err
}

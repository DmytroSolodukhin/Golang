package test

import (
	"Ports/serverApi/model"
	"context"
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

const (
	MongoDBHost = "mongodb://localhost:27017"
	DbName = "testPort"
)
func TestDBIntegration(t *testing.T)  {

	Convey("Mongo integrations test", t, func() {
		model.ConnectToDB(MongoDBHost, DbName)

		testObject1 := &model.Port{
			PortID: "TEST1",
			Name: "test",
			City: "test",
			Country: "test",
			Alias: []string{},
			Regions: []string{},
			Coordinates: []float64{55.5136433, 25.4052165},
			Province: "test",
			Timezone: "test",
			Unlocs: []string{"TEST"},
			Code: "52000",
		}

		testObject2 := &model.Port{
			PortID: "TEST2",
			Name: "test2",
			City: "test2",
			Country: "test2",
			Alias: []string{},
			Regions: []string{},
			Coordinates: []float64{55.4564, 75.4674},
			Province: "test2",
			Timezone: "test2",
			Unlocs: []string{"TEST2"},
			Code: "10000",
		}

		Convey("Save Port", func() {
			model.SavePort(testObject1)
			model.SavePort(testObject2)
		})
		Convey("Getting Port", func() {
			geting := model.GetPortFromDB("TEST1")
			So(geting, ShouldResemble, testObject1)
		})
		Convey("Getting Ports", func() {
			geting := model.GetPortsFromDB()
			So(geting, ShouldNotBeEmpty)
		})
		Convey("Remove Port", func() {
			geting := model.DeletePortFromDB("TEST1")
			So(geting, ShouldBeTrue)
			geting = model.DeletePortFromDB("TEST2")
			So(geting, ShouldBeTrue)
			geting = model.DeletePortFromDB("TEST3")
			So(geting, ShouldBeFalse)
		})
		collection := model.ClientDB.Database(dbName).Collection("ports")
		filter := bson.M{}
		_, _ = collection.DeleteMany(context.TODO(), filter)
	})
}

package db_manager

import (
	api "github.com/kazak/Golang/modules/grpcapi"
	"testing"
)

const (
	MongoDBHost = "mongodb://localhost:27017"
	DbName = "testPort"
)
func TestDBIntegration(t *testing.T)  {

	Convey("Mongo integrations test", t, func() {
		repo := ConnectToDB(MongoDBHost, DbName)

		testObject1 := &api.Port{
			PortId: "TEST1",
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

		testObject2 := &api.Port{
			PortId: "TEST2",
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
			geting := repo.Save(testObject1)
			So(geting, ShouldBeTrue)
			geting = repo.Save(testObject2)
			So(geting, ShouldBeTrue)
		})
		Convey("Getting Port", func() {
			geting := repo.GetByID("TEST1")
			So(geting, ShouldResemble, testObject1)
		})
		Convey("Getting Ports", func() {
			geting := repo.GetAll()
			So(geting, ShouldNotBeEmpty)
		})
		Convey("Remove Port", func() {
			geting := repo.Delete("TEST1")
			So(geting, ShouldBeTrue)
			geting = repo.Delete("TEST2")
			So(geting, ShouldBeTrue)
			geting = repo.Delete("TEST3")
			So(geting, ShouldBeFalse)
		})
	})
}

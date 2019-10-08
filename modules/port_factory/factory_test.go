package port_factory

import (
	"bufio"
	"bytes"
	"testing"
	api "github.com/kazak/Golang/modules/grpcapi"
	. "github.com/smartystreets/goconvey/convey"
)

// Test to parse request body
func TestScanToChank(t *testing.T) {
	t.Parallel()
	Convey("Line by request body should be getting corretly", t, func() {

		chankText := make(chan string)
		buf := bytes.NewBufferString("line 1\nline 2\nline 1")
		scanner := bufio.NewScanner(buf)
		go ScanRequestBodyToChank(scanner, chankText)

		Convey("Getting chanks", func() {
			exp :=  <-chankText
			So(exp, ShouldEqual, "line 2")
			_, done :=  <-chankText
			So(done, ShouldEqual, false)
		})
	})
}

//Test to build port object
func TestStartProdactionPort(t *testing.T) {
	t.Parallel()
	Convey("Object from chank and create chank with port object corretly", t, func() {
		chankText, chankPort := make(chan string), make(chan *api.Port)
		jsonPort := `"Test": {
			"name": "test",
			"city": "test",
			"country": "test",
			"alias": [],
			"regions": [],
			"coordinates": [
			  55.5136433,
			  25.4052165
			],
			"province": "test",
			"timezone": "test",
			"unlocs": [
			  "TEST"
			],
			"code": "52000"
		  },`
		objectPort := &api.Port{
			PortId: "Test",
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
		go StartProdactionPort(chankText, chankPort)

		Convey("Getting chanks", func() {
			go func() {
				chankText <- "{"
				chankText <- jsonPort
				exp :=  <-chankText

				So(objectPort, ShouldEqual, exp)
				close(chankText)
				_, done :=  <-chankText
				So(done, ShouldEqual, false)
			}()
		})
	})
}
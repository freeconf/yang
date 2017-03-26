package main

import (
	"github.com/c2stack/c2g/conf"
	"github.com/c2stack/c2g/examples/car"
	"github.com/c2stack/c2g/meta"
	"github.com/c2stack/c2g/restconf"
)

// Initialize and start our Car micro-service application with C2Stack for
// RESTful based management
//
// To run:
//    cd ./src/vendor/github.com/c2stack/c2g/examples/car/car-basic
//    go run ./main.go
//
// Then open web browser to
//   http://localhost:8080/
//
func main() {

	// Your application
	app := car.New()

	// Where to looks for yang files, this tells library to use cwd
	ypath := &meta.FileStreamSource{Root: ".."}

	// Every management has a "device" container for
	device := conf.NewLocalDeviceWithUi(ypath, ypath)

	// Browser is the management handle to your entire application
	// see how we're connecting 3 things here
	//  1.) model - schema map for your application, nothing get's in/out with adhereing to model
	//  2.) car - this can be multiple objects, or often just one but that's up to you and how
	//      you designed your application.
	//  3.) carNode - your bridge between your model and your application
	device.Add("car", car.Node(app))

	// in our car app, we start off by running start.
	app.Start()

	// RESTful management.  Only required for RESTful based management.
	restconf.NewManagement(device, ":8080")

	select {}
}

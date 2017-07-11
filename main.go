//go:generate goagen bootstrap -d github.com/intervention-engine/ie/design

package main

import (
	"fmt"
	"log"

	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/intervention-engine/ie/app"
	"github.com/intervention-engine/ie/appt"
	"github.com/intervention-engine/ie/mongo"
	mgo "gopkg.in/mgo.v2"
)

func main() {
	cfg := ConfigurationInit()
	if cfg.PrintConfiguration {
		fmt.Printf("%#+v\n", cfg)
		return
	}
	//setupLogFile(cfg.LogDir)

	// Connect database
	session, err := mgo.Dial(cfg.MongoURL)
	if err != nil {
		log.Fatalln("dialing mongo failed for session at: ", cfg.MongoURL)
	}
	defer session.Close()

	svcFactory := mongo.NewServiceFactory(session.Copy(), "fhir")
	appt.Boot(svcFactory, cfg.HuddleConfigFiles)

	// Create service
	service := goa.New("api")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())
	service.Use(exposeHeaderField("Link"))
	service.Use(withMongoService(session))
	service.Use(withRiskServices(cfg.RiskServicesPath))
	service.Use(withHuddleConfig(cfg.HuddleConfigFiles))

	// Mount "patient" controller
	pc := NewPatientController(service)
	app.MountPatientController(service, pc)
	ct := NewCareTeamController(service)
	app.MountCareTeamController(service, ct)
	hc := NewHuddleController(service)
	app.MountHuddleController(service, hc)
	sc := NewSwaggerController(service)
	app.MountSwaggerController(service, sc)
	rac := NewRiskAssessmentController(service)
	app.MountRiskAssessmentController(service, rac)
	rcc := NewRiskCategoriesController(service)
	app.MountRiskCategoriesController(service, rcc)
	ec := NewEventController(service)
	app.MountEventController(service, ec)

	rsc := NewRiskServiceController(service)
	app.MountRiskServiceController(service, rsc)

	// Start service
	log.Println("serving api at: ", cfg.HostURL)
	if err := service.ListenAndServe(cfg.HostURL); err != nil {
		service.LogError("startup", "err", err)
	}
}

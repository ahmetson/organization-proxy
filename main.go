package main

import (
	"github.com/ahmetson/organization-proxy/handler"
	"github.com/ahmetson/organization-proxy/smartcontract"
	"github.com/ahmetson/service-lib/configuration"
	"github.com/ahmetson/service-lib/controller"
	"github.com/ahmetson/service-lib/log"
	"github.com/ahmetson/service-lib/proxy"
)

func main() {
	logger, err := log.New("organization", false)
	if err != nil {
		log.Fatal("log.New", "error", err)
	}
	appConfig, err := configuration.New(logger)
	if err != nil {
		logger.Fatal("configuration.New", "error", err)
	}

	//
	// Proxy service itself
	//
	err = smartcontract.Init(appConfig)
	if err != nil {
		logger.Fatal("smartcontract.Validate", "error", err)
	}

	sourceController, err := controller.NewReplier(logger)
	if err != nil {
		logger.Fatal("failed to create a source as controller.NewReplier", "error", err)
	}
	err = sourceController.AddRoute(handler.InsertRoute())
	if err != nil {
		logger.Fatal("controller.AddRoute", "error", err)
	}

	service := proxy.New(appConfig, logger)
	if err != nil {
		logger.Fatal("proxy.New", "error", err)
	}

	service.SetCustomSource(sourceController)

	service.Controller.RequireDestination(configuration.ReplierType)

	err = service.Prepare()
	if err != nil {
		logger.Fatal("service.Prepare", "error", err)
	}

	service.Run()
}

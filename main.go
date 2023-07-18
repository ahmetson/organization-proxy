package main

import (
	"github.com/ahmetson/service-lib/configuration"
	"github.com/ahmetson/service-lib/controller"
	"github.com/ahmetson/service-lib/log"
	"github.com/ahmetson/service-lib/proxy"
	"github.com/ahmetson/service-lib/remote"
)

var proxyHandler = func(messages []string, logger log.Logger, destinations []*proxy.DestinationClient, _ remote.Clients) ([]string, error) {
	//request, err := message.ParseRequest(messages)
	return messages, nil
}

func main() {
	logger, err := log.New("main", false)
	if err != nil {
		log.Fatal("log.New", "error", err)
	}
	appConfig, err := configuration.NewAppConfig(logger)
	if err != nil {
		logger.Fatal("configuration.NewAppConfig", "error", err)
	}
	if len(appConfig.Services) == 0 {
		logger.Fatal("missing service configuration in yaml")
	}

	//
	// Proxy service itself
	//

	sourceController, err := controller.NewReplier(logger)
	if err != nil {
		logger.Fatal("failed to create a source as controller.NewReplier", "error", err)
	}
	sourceConfig, err := appConfig.Services[0].GetController(proxy.SourceName)
	if err != nil {
		logger.Fatal("failed to get controller config",
			"service instance", appConfig.Services[0].Instance,
			"controller name", proxy.SourceName)
	}
	sourceController.AddConfig(sourceConfig)
	var source controller.Interface = sourceController
	proxy.SourceHandler(source)

	service, err := proxy.New(appConfig.Services[0], logger)
	if err != nil {
		logger.Fatal("proxy.New", "error", err)
	}

	err = service.AddSourceController(proxy.SourceName, sourceController)
	if err != nil {
		logger.Fatal("service.AddSourceController", "controller.name", proxy.SourceName, "error", err)
	}
	service.SetRequestHandler(proxyHandler)

	service.Run()
}

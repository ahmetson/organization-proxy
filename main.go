package main

import (
	"fmt"
	"github.com/ahmetson/common-lib/data_type/key_value"
	"github.com/ahmetson/common-lib/topic"
	"github.com/ahmetson/organization-proxy/smartcontract"
	"github.com/ahmetson/service-lib/communication/message"
	"github.com/ahmetson/service-lib/configuration"
	"github.com/ahmetson/service-lib/controller"
	"github.com/ahmetson/service-lib/extension/database"
	"github.com/ahmetson/service-lib/log"
	"github.com/ahmetson/service-lib/proxy"
	"github.com/ahmetson/service-lib/remote"
)

var smartcontractFileName = func(id topic.Topic) (string, error) {
	if !id.Has("org", "net", "name") {
		return "", fmt.Errorf("missing one of the parameters: 'org', 'net', 'name' in topic")
	}

	return fmt.Sprintf("%s.%s.%s.json", id.Organization, id.NetworkId, id.Name), nil
}

var configurationFileName = func(id topic.Id) (string, error) {
	parsed, err := id.Unmarshal()
	if !id.Has("org", "proj") {
		return "", fmt.Errorf("missing one of the parameters: 'org', 'net', 'name' in topic")
	}

	if err != nil {
		return "", fmt.Errorf("failed to umarshall id to topic: %w", err)
	}
	return fmt.Sprintf("%s.%s.json", parsed.Organization, parsed.Project), nil
}

var proxyRequestHandler = func(messages []string, logger log.Logger, destinations []*proxy.DestinationClient, _ remote.Clients) ([]string, error) {
	request, err := message.ParseRequest(messages)
	if err != nil {
		return nil, fmt.Errorf("message.ParseRequest: %w", err)
	}

	// we accept only database requests
	var requestParameters database.QueryRequest
	err = request.Parameters.Interface(&requestParameters)
	if err != nil {
		return nil, fmt.Errorf("failed to parse database request: %w", err)
	}

	// let's handle the abi
	if request.Command == database.Insert.String() {
		// abi insert
		if len(requestParameters.Tables) == 0 {
			return nil, fmt.Errorf("missing any tables: %w", err)
		}

		tableName := requestParameters.Tables[0]
		// insert abi
		// requires two fields: abi_id, body
		// requires two arguments: id, body
		if tableName == "abi" {
			abiId := requestParameters.Arguments[0].(string)
			body := requestParameters.Arguments[1].(string)

			requestMessage := database.QueryRequest{
				Fields:    []string{abiId + ".json"},
				Arguments: []interface{}{body},
			}.Request(database.Insert)

			requestString, _ := requestMessage.ToString()

			return []string{requestString}, nil
		} else if tableName == "smartcontract" {
			// insert smartcontract
			id, ok := requestParameters.Arguments[1].(topic.Topic)
			if !ok {
				return nil, fmt.Errorf("failed to decode first argument as topic.Topic")
			}

			fileName, err := smartcontractFileName(id)
			if err != nil {
				return nil, fmt.Errorf("failed to create smartcontract file name %s: %w", fileName, err)
			}

			arguments := key_value.Empty()
			for i := 1; i < len(requestParameters.Fields); i++ {
				field := requestParameters.Fields[i]
				argument := requestParameters.Arguments[i]
				arguments.Set(field, argument)
			}

			requestMessage := database.QueryRequest{
				Fields:    []string{fileName},
				Arguments: []interface{}{arguments},
			}.Request(database.Insert)

			requestString, _ := requestMessage.ToString()

			return []string{requestString}, nil
		} else if tableName == "configuration" {
			// insert configuration
			id, ok := requestParameters.Arguments[0].(topic.Id)
			if !ok {
				return nil, fmt.Errorf("failed to decode first argument as topic.Id")
			}

			mainFileName, err := configurationFileName(id)
			if err != nil {
				return nil, fmt.Errorf("failed to create configuration file name: %w", err)
			}

			smartcontracts, ok := requestParameters.Arguments[1].([]topic.Id)
			if !ok {
				return nil, fmt.Errorf("failed to decode the second argument as []topic.Id")
			}
			if len(smartcontracts) == 0 {
				return nil, fmt.Errorf("missing smartcontracts for the configuration")
			}
			smartcontractFileNames := make([]string, len(smartcontracts))
			for i := range smartcontracts {
				parsed, err := smartcontracts[i].Unmarshal()
				if err != nil {
					return nil, fmt.Errorf("failed to umarshall smartcontract id into topic")
				}

				fileName, err := smartcontractFileName(parsed)
				if err != nil {
					return nil, fmt.Errorf("failed to create a file name for smartcontract %s: %w", smartcontracts[i], err)
				}
				smartcontractFileNames[i] = fileName
			}

			requestMessage := database.QueryRequest{
				Fields:    []string{mainFileName},
				Arguments: []interface{}{smartcontractFileNames},
			}.Request(database.Insert)

			requestString, _ := requestMessage.ToString()

			return []string{requestString}, nil
		} else {
			return nil, fmt.Errorf("the %s table is not supported", tableName)
		}
	} else {
		return nil, fmt.Errorf("only database.Insert command supported for now")
	}
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
	err = smartcontract.Init(appConfig)
	if err != nil {
		logger.Fatal("smartcontract.Validate", "error", err)
	}
	logger.Fatal("finished the smartcontract loading")

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
	service.SetRequestHandler(proxyRequestHandler)

	service.Run()
}

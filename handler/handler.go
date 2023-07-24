package handler

import (
	"fmt"
	"github.com/ahmetson/common-lib/data_type/key_value"
	"github.com/ahmetson/common-lib/topic"
	"github.com/ahmetson/organization-proxy/converter"
	"github.com/ahmetson/service-lib/communication/command"
	"github.com/ahmetson/service-lib/communication/message"
	"github.com/ahmetson/service-lib/extension/database"
	"github.com/ahmetson/service-lib/log"
	"github.com/ahmetson/service-lib/proxy"
	"github.com/ahmetson/service-lib/remote"
)

var onInsert = func(request message.Request, _ *log.Logger, extensions ...*remote.ClientSocket) message.Reply {
	// we accept only database requests
	var requestParameters database.QueryRequest
	err := request.Parameters.Interface(&requestParameters)
	if err != nil {
		return message.Fail("failed to parse database request: %w" + err.Error())
	}

	// table to identify what kind of information we will store
	if len(requestParameters.Tables) == 0 {
		return message.Fail("missing tables in the request parameters")
	}

	proxyClient := remote.FindClient(extensions, proxy.ControllerName)

	tableName := requestParameters.Tables[0]

	var convertedParameters database.QueryRequest
	// insert abi
	// requires two fields: abi_id, body
	// requires two arguments: id, body
	if tableName == "abi" {
		abiId := requestParameters.Arguments[0].(string)
		body := requestParameters.Arguments[1].(string)

		convertedParameters = database.QueryRequest{
			Fields:    []string{abiId + ".json"},
			Arguments: []interface{}{body},
		}
	} else if tableName == "smartcontract" {
		// insert smartcontract
		id, ok := requestParameters.Arguments[1].(topic.Topic)
		if !ok {
			return message.Fail("failed to decode first argument as topic.Topic")
		}

		fileName, err := converter.SmartcontractFileName(id)
		if err != nil {
			return message.Fail("failed to create smartcontract file name: " + err.Error())
		}

		arguments := key_value.Empty()
		for i := 1; i < len(requestParameters.Fields); i++ {
			field := requestParameters.Fields[i]
			argument := requestParameters.Arguments[i]
			arguments.Set(field, argument)
		}

		convertedParameters = database.QueryRequest{
			Fields:    []string{fileName},
			Arguments: []interface{}{arguments},
		}
	} else if tableName == "configuration" {
		// insert configuration
		id, ok := requestParameters.Arguments[0].(topic.Id)
		if !ok {
			return message.Fail("failed to decode first argument as topic.Id")
		}

		mainFileName, err := converter.ConfigurationFileName(id)
		if err != nil {
			return message.Fail("failed to create configuration file name: " + err.Error())
		}

		smartcontracts, ok := requestParameters.Arguments[1].([]topic.Id)
		if !ok {
			return message.Fail("failed to decode the second argument as []topic.Id")
		}
		if len(smartcontracts) == 0 {
			return message.Fail("missing smartcontracts for the configuration")
		}
		smartcontractFileNames := make([]string, len(smartcontracts))
		for i := range smartcontracts {
			parsed, err := smartcontracts[i].Unmarshal()
			if err != nil {
				return message.Fail("failed to unmarshall smartcontract id into topic")
			}

			fileName, err := converter.SmartcontractFileName(parsed)
			if err != nil {
				return message.Fail(fmt.Sprintf("failed to create a file name for '%s' smartcontract: %v", smartcontracts[i], err))
			}
			smartcontractFileNames[i] = fileName
		}

		convertedParameters = database.QueryRequest{
			Fields:    []string{mainFileName},
			Arguments: []interface{}{smartcontractFileNames},
		}
	} else {
		return message.Fail(fmt.Sprintf("the %s table is not supported", tableName))
	}

	requestMessage := convertedParameters.Request(database.Insert)

	replyParameter, err := proxyClient.RequestRemoteService(&requestMessage)
	if err != nil {
		return message.Fail("proxyClient: %w" + err.Error())
	}

	var insertReply database.InsertReply
	err = replyParameter.Interface(&insertReply)
	if err != nil {
		return message.Fail("proxy reply is not database.InsertReply: " + err.Error())
	}
	reply, err := command.Reply(&insertReply)
	if err != nil {
		return message.Fail("failed to convert parameters to reply message: " + err.Error())
	}

	return reply
}

var InsertRoute = func() *command.Route {
	return command.NewRoute(database.Insert, onInsert)
}

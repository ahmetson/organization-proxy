// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package smartcontract

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// SmartcontractMetaData contains all meta data concerning the Smartcontract contract.
var SmartcontractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"org\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"fileName\",\"type\":\"string\"}],\"name\":\"DeleteConfiguration\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"org\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"fileName\",\"type\":\"string\"}],\"name\":\"DeleteSmartcontract\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"org\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"fileName\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"}],\"name\":\"InsertConfiguration\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"org\",\"type\":\"string\"}],\"name\":\"InsertOrganization\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"org\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"fileName\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"}],\"name\":\"InsertSmartcontract\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"fileName\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"}],\"name\":\"UpdateConfiguration\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"fileName\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"}],\"name\":\"UpdateSmartcontract\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"configurations\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"org\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"fileName\",\"type\":\"string\"}],\"name\":\"deleteConfiguration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"org\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"fileName\",\"type\":\"string\"}],\"name\":\"deleteSmartcontract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"org\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"fileName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"}],\"name\":\"insertConfiguration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"org\",\"type\":\"string\"}],\"name\":\"insertOrganization\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"org\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"fileName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"}],\"name\":\"insertSmartcontract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"organizationConfigurations\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"organizationSmartcontracts\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"organizations\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"smartcontracts\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"fileName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"}],\"name\":\"updateConfiguration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"fileName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"}],\"name\":\"updateSmartcontract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// SmartcontractABI is the input ABI used to generate the binding from.
// Deprecated: Use SmartcontractMetaData.ABI instead.
var SmartcontractABI = SmartcontractMetaData.ABI

// Smartcontract is an auto generated Go binding around an Ethereum contract.
type Smartcontract struct {
	SmartcontractCaller     // Read-only binding to the contract
	SmartcontractTransactor // Write-only binding to the contract
	SmartcontractFilterer   // Log filterer for contract events
}

// SmartcontractCaller is an auto generated read-only Go binding around an Ethereum contract.
type SmartcontractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SmartcontractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SmartcontractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SmartcontractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SmartcontractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SmartcontractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SmartcontractSession struct {
	Contract     *Smartcontract    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SmartcontractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SmartcontractCallerSession struct {
	Contract *SmartcontractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// SmartcontractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SmartcontractTransactorSession struct {
	Contract     *SmartcontractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// SmartcontractRaw is an auto generated low-level Go binding around an Ethereum contract.
type SmartcontractRaw struct {
	Contract *Smartcontract // Generic contract binding to access the raw methods on
}

// SmartcontractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SmartcontractCallerRaw struct {
	Contract *SmartcontractCaller // Generic read-only contract binding to access the raw methods on
}

// SmartcontractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SmartcontractTransactorRaw struct {
	Contract *SmartcontractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSmartcontract creates a new instance of Smartcontract, bound to a specific deployed contract.
func NewSmartcontract(address common.Address, backend bind.ContractBackend) (*Smartcontract, error) {
	contract, err := bindSmartcontract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Smartcontract{SmartcontractCaller: SmartcontractCaller{contract: contract}, SmartcontractTransactor: SmartcontractTransactor{contract: contract}, SmartcontractFilterer: SmartcontractFilterer{contract: contract}}, nil
}

// NewSmartcontractCaller creates a new read-only instance of Smartcontract, bound to a specific deployed contract.
func NewSmartcontractCaller(address common.Address, caller bind.ContractCaller) (*SmartcontractCaller, error) {
	contract, err := bindSmartcontract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SmartcontractCaller{contract: contract}, nil
}

// NewSmartcontractTransactor creates a new write-only instance of Smartcontract, bound to a specific deployed contract.
func NewSmartcontractTransactor(address common.Address, transactor bind.ContractTransactor) (*SmartcontractTransactor, error) {
	contract, err := bindSmartcontract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SmartcontractTransactor{contract: contract}, nil
}

// NewSmartcontractFilterer creates a new log filterer instance of Smartcontract, bound to a specific deployed contract.
func NewSmartcontractFilterer(address common.Address, filterer bind.ContractFilterer) (*SmartcontractFilterer, error) {
	contract, err := bindSmartcontract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SmartcontractFilterer{contract: contract}, nil
}

// bindSmartcontract binds a generic wrapper to an already deployed contract.
func bindSmartcontract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SmartcontractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Smartcontract *SmartcontractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Smartcontract.Contract.SmartcontractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Smartcontract *SmartcontractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Smartcontract.Contract.SmartcontractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Smartcontract *SmartcontractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Smartcontract.Contract.SmartcontractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Smartcontract *SmartcontractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Smartcontract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Smartcontract *SmartcontractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Smartcontract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Smartcontract *SmartcontractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Smartcontract.Contract.contract.Transact(opts, method, params...)
}

// Configurations is a free data retrieval call binding the contract method 0x1214dd58.
//
// Solidity: function configurations(string ) view returns(string)
func (_Smartcontract *SmartcontractCaller) Configurations(opts *bind.CallOpts, arg0 string) (string, error) {
	var out []interface{}
	err := _Smartcontract.contract.Call(opts, &out, "configurations", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Configurations is a free data retrieval call binding the contract method 0x1214dd58.
//
// Solidity: function configurations(string ) view returns(string)
func (_Smartcontract *SmartcontractSession) Configurations(arg0 string) (string, error) {
	return _Smartcontract.Contract.Configurations(&_Smartcontract.CallOpts, arg0)
}

// Configurations is a free data retrieval call binding the contract method 0x1214dd58.
//
// Solidity: function configurations(string ) view returns(string)
func (_Smartcontract *SmartcontractCallerSession) Configurations(arg0 string) (string, error) {
	return _Smartcontract.Contract.Configurations(&_Smartcontract.CallOpts, arg0)
}

// OrganizationConfigurations is a free data retrieval call binding the contract method 0x6f74463f.
//
// Solidity: function organizationConfigurations(string , uint256 ) view returns(string)
func (_Smartcontract *SmartcontractCaller) OrganizationConfigurations(opts *bind.CallOpts, arg0 string, arg1 *big.Int) (string, error) {
	var out []interface{}
	err := _Smartcontract.contract.Call(opts, &out, "organizationConfigurations", arg0, arg1)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// OrganizationConfigurations is a free data retrieval call binding the contract method 0x6f74463f.
//
// Solidity: function organizationConfigurations(string , uint256 ) view returns(string)
func (_Smartcontract *SmartcontractSession) OrganizationConfigurations(arg0 string, arg1 *big.Int) (string, error) {
	return _Smartcontract.Contract.OrganizationConfigurations(&_Smartcontract.CallOpts, arg0, arg1)
}

// OrganizationConfigurations is a free data retrieval call binding the contract method 0x6f74463f.
//
// Solidity: function organizationConfigurations(string , uint256 ) view returns(string)
func (_Smartcontract *SmartcontractCallerSession) OrganizationConfigurations(arg0 string, arg1 *big.Int) (string, error) {
	return _Smartcontract.Contract.OrganizationConfigurations(&_Smartcontract.CallOpts, arg0, arg1)
}

// OrganizationSmartcontracts is a free data retrieval call binding the contract method 0x8718dc72.
//
// Solidity: function organizationSmartcontracts(string , uint256 ) view returns(string)
func (_Smartcontract *SmartcontractCaller) OrganizationSmartcontracts(opts *bind.CallOpts, arg0 string, arg1 *big.Int) (string, error) {
	var out []interface{}
	err := _Smartcontract.contract.Call(opts, &out, "organizationSmartcontracts", arg0, arg1)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// OrganizationSmartcontracts is a free data retrieval call binding the contract method 0x8718dc72.
//
// Solidity: function organizationSmartcontracts(string , uint256 ) view returns(string)
func (_Smartcontract *SmartcontractSession) OrganizationSmartcontracts(arg0 string, arg1 *big.Int) (string, error) {
	return _Smartcontract.Contract.OrganizationSmartcontracts(&_Smartcontract.CallOpts, arg0, arg1)
}

// OrganizationSmartcontracts is a free data retrieval call binding the contract method 0x8718dc72.
//
// Solidity: function organizationSmartcontracts(string , uint256 ) view returns(string)
func (_Smartcontract *SmartcontractCallerSession) OrganizationSmartcontracts(arg0 string, arg1 *big.Int) (string, error) {
	return _Smartcontract.Contract.OrganizationSmartcontracts(&_Smartcontract.CallOpts, arg0, arg1)
}

// Organizations is a free data retrieval call binding the contract method 0xe311f43c.
//
// Solidity: function organizations(string ) view returns(bool)
func (_Smartcontract *SmartcontractCaller) Organizations(opts *bind.CallOpts, arg0 string) (bool, error) {
	var out []interface{}
	err := _Smartcontract.contract.Call(opts, &out, "organizations", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Organizations is a free data retrieval call binding the contract method 0xe311f43c.
//
// Solidity: function organizations(string ) view returns(bool)
func (_Smartcontract *SmartcontractSession) Organizations(arg0 string) (bool, error) {
	return _Smartcontract.Contract.Organizations(&_Smartcontract.CallOpts, arg0)
}

// Organizations is a free data retrieval call binding the contract method 0xe311f43c.
//
// Solidity: function organizations(string ) view returns(bool)
func (_Smartcontract *SmartcontractCallerSession) Organizations(arg0 string) (bool, error) {
	return _Smartcontract.Contract.Organizations(&_Smartcontract.CallOpts, arg0)
}

// Smartcontracts is a free data retrieval call binding the contract method 0xbcd8e820.
//
// Solidity: function smartcontracts(string ) view returns(string)
func (_Smartcontract *SmartcontractCaller) Smartcontracts(opts *bind.CallOpts, arg0 string) (string, error) {
	var out []interface{}
	err := _Smartcontract.contract.Call(opts, &out, "smartcontracts", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Smartcontracts is a free data retrieval call binding the contract method 0xbcd8e820.
//
// Solidity: function smartcontracts(string ) view returns(string)
func (_Smartcontract *SmartcontractSession) Smartcontracts(arg0 string) (string, error) {
	return _Smartcontract.Contract.Smartcontracts(&_Smartcontract.CallOpts, arg0)
}

// Smartcontracts is a free data retrieval call binding the contract method 0xbcd8e820.
//
// Solidity: function smartcontracts(string ) view returns(string)
func (_Smartcontract *SmartcontractCallerSession) Smartcontracts(arg0 string) (string, error) {
	return _Smartcontract.Contract.Smartcontracts(&_Smartcontract.CallOpts, arg0)
}

// DeleteConfiguration is a paid mutator transaction binding the contract method 0xb58f296a.
//
// Solidity: function deleteConfiguration(string org, string fileName) returns()
func (_Smartcontract *SmartcontractTransactor) DeleteConfiguration(opts *bind.TransactOpts, org string, fileName string) (*types.Transaction, error) {
	return _Smartcontract.contract.Transact(opts, "deleteConfiguration", org, fileName)
}

// DeleteConfiguration is a paid mutator transaction binding the contract method 0xb58f296a.
//
// Solidity: function deleteConfiguration(string org, string fileName) returns()
func (_Smartcontract *SmartcontractSession) DeleteConfiguration(org string, fileName string) (*types.Transaction, error) {
	return _Smartcontract.Contract.DeleteConfiguration(&_Smartcontract.TransactOpts, org, fileName)
}

// DeleteConfiguration is a paid mutator transaction binding the contract method 0xb58f296a.
//
// Solidity: function deleteConfiguration(string org, string fileName) returns()
func (_Smartcontract *SmartcontractTransactorSession) DeleteConfiguration(org string, fileName string) (*types.Transaction, error) {
	return _Smartcontract.Contract.DeleteConfiguration(&_Smartcontract.TransactOpts, org, fileName)
}

// DeleteSmartcontract is a paid mutator transaction binding the contract method 0x371961f8.
//
// Solidity: function deleteSmartcontract(string org, string fileName) returns()
func (_Smartcontract *SmartcontractTransactor) DeleteSmartcontract(opts *bind.TransactOpts, org string, fileName string) (*types.Transaction, error) {
	return _Smartcontract.contract.Transact(opts, "deleteSmartcontract", org, fileName)
}

// DeleteSmartcontract is a paid mutator transaction binding the contract method 0x371961f8.
//
// Solidity: function deleteSmartcontract(string org, string fileName) returns()
func (_Smartcontract *SmartcontractSession) DeleteSmartcontract(org string, fileName string) (*types.Transaction, error) {
	return _Smartcontract.Contract.DeleteSmartcontract(&_Smartcontract.TransactOpts, org, fileName)
}

// DeleteSmartcontract is a paid mutator transaction binding the contract method 0x371961f8.
//
// Solidity: function deleteSmartcontract(string org, string fileName) returns()
func (_Smartcontract *SmartcontractTransactorSession) DeleteSmartcontract(org string, fileName string) (*types.Transaction, error) {
	return _Smartcontract.Contract.DeleteSmartcontract(&_Smartcontract.TransactOpts, org, fileName)
}

// InsertConfiguration is a paid mutator transaction binding the contract method 0x9ea508f8.
//
// Solidity: function insertConfiguration(string org, string fileName, string cid) returns()
func (_Smartcontract *SmartcontractTransactor) InsertConfiguration(opts *bind.TransactOpts, org string, fileName string, cid string) (*types.Transaction, error) {
	return _Smartcontract.contract.Transact(opts, "insertConfiguration", org, fileName, cid)
}

// InsertConfiguration is a paid mutator transaction binding the contract method 0x9ea508f8.
//
// Solidity: function insertConfiguration(string org, string fileName, string cid) returns()
func (_Smartcontract *SmartcontractSession) InsertConfiguration(org string, fileName string, cid string) (*types.Transaction, error) {
	return _Smartcontract.Contract.InsertConfiguration(&_Smartcontract.TransactOpts, org, fileName, cid)
}

// InsertConfiguration is a paid mutator transaction binding the contract method 0x9ea508f8.
//
// Solidity: function insertConfiguration(string org, string fileName, string cid) returns()
func (_Smartcontract *SmartcontractTransactorSession) InsertConfiguration(org string, fileName string, cid string) (*types.Transaction, error) {
	return _Smartcontract.Contract.InsertConfiguration(&_Smartcontract.TransactOpts, org, fileName, cid)
}

// InsertOrganization is a paid mutator transaction binding the contract method 0x62389e43.
//
// Solidity: function insertOrganization(string org) returns()
func (_Smartcontract *SmartcontractTransactor) InsertOrganization(opts *bind.TransactOpts, org string) (*types.Transaction, error) {
	return _Smartcontract.contract.Transact(opts, "insertOrganization", org)
}

// InsertOrganization is a paid mutator transaction binding the contract method 0x62389e43.
//
// Solidity: function insertOrganization(string org) returns()
func (_Smartcontract *SmartcontractSession) InsertOrganization(org string) (*types.Transaction, error) {
	return _Smartcontract.Contract.InsertOrganization(&_Smartcontract.TransactOpts, org)
}

// InsertOrganization is a paid mutator transaction binding the contract method 0x62389e43.
//
// Solidity: function insertOrganization(string org) returns()
func (_Smartcontract *SmartcontractTransactorSession) InsertOrganization(org string) (*types.Transaction, error) {
	return _Smartcontract.Contract.InsertOrganization(&_Smartcontract.TransactOpts, org)
}

// InsertSmartcontract is a paid mutator transaction binding the contract method 0xad5f0707.
//
// Solidity: function insertSmartcontract(string org, string fileName, string cid) returns()
func (_Smartcontract *SmartcontractTransactor) InsertSmartcontract(opts *bind.TransactOpts, org string, fileName string, cid string) (*types.Transaction, error) {
	return _Smartcontract.contract.Transact(opts, "insertSmartcontract", org, fileName, cid)
}

// InsertSmartcontract is a paid mutator transaction binding the contract method 0xad5f0707.
//
// Solidity: function insertSmartcontract(string org, string fileName, string cid) returns()
func (_Smartcontract *SmartcontractSession) InsertSmartcontract(org string, fileName string, cid string) (*types.Transaction, error) {
	return _Smartcontract.Contract.InsertSmartcontract(&_Smartcontract.TransactOpts, org, fileName, cid)
}

// InsertSmartcontract is a paid mutator transaction binding the contract method 0xad5f0707.
//
// Solidity: function insertSmartcontract(string org, string fileName, string cid) returns()
func (_Smartcontract *SmartcontractTransactorSession) InsertSmartcontract(org string, fileName string, cid string) (*types.Transaction, error) {
	return _Smartcontract.Contract.InsertSmartcontract(&_Smartcontract.TransactOpts, org, fileName, cid)
}

// UpdateConfiguration is a paid mutator transaction binding the contract method 0x8084163e.
//
// Solidity: function updateConfiguration(string fileName, string cid) returns()
func (_Smartcontract *SmartcontractTransactor) UpdateConfiguration(opts *bind.TransactOpts, fileName string, cid string) (*types.Transaction, error) {
	return _Smartcontract.contract.Transact(opts, "updateConfiguration", fileName, cid)
}

// UpdateConfiguration is a paid mutator transaction binding the contract method 0x8084163e.
//
// Solidity: function updateConfiguration(string fileName, string cid) returns()
func (_Smartcontract *SmartcontractSession) UpdateConfiguration(fileName string, cid string) (*types.Transaction, error) {
	return _Smartcontract.Contract.UpdateConfiguration(&_Smartcontract.TransactOpts, fileName, cid)
}

// UpdateConfiguration is a paid mutator transaction binding the contract method 0x8084163e.
//
// Solidity: function updateConfiguration(string fileName, string cid) returns()
func (_Smartcontract *SmartcontractTransactorSession) UpdateConfiguration(fileName string, cid string) (*types.Transaction, error) {
	return _Smartcontract.Contract.UpdateConfiguration(&_Smartcontract.TransactOpts, fileName, cid)
}

// UpdateSmartcontract is a paid mutator transaction binding the contract method 0xf252b7ec.
//
// Solidity: function updateSmartcontract(string fileName, string cid) returns()
func (_Smartcontract *SmartcontractTransactor) UpdateSmartcontract(opts *bind.TransactOpts, fileName string, cid string) (*types.Transaction, error) {
	return _Smartcontract.contract.Transact(opts, "updateSmartcontract", fileName, cid)
}

// UpdateSmartcontract is a paid mutator transaction binding the contract method 0xf252b7ec.
//
// Solidity: function updateSmartcontract(string fileName, string cid) returns()
func (_Smartcontract *SmartcontractSession) UpdateSmartcontract(fileName string, cid string) (*types.Transaction, error) {
	return _Smartcontract.Contract.UpdateSmartcontract(&_Smartcontract.TransactOpts, fileName, cid)
}

// UpdateSmartcontract is a paid mutator transaction binding the contract method 0xf252b7ec.
//
// Solidity: function updateSmartcontract(string fileName, string cid) returns()
func (_Smartcontract *SmartcontractTransactorSession) UpdateSmartcontract(fileName string, cid string) (*types.Transaction, error) {
	return _Smartcontract.Contract.UpdateSmartcontract(&_Smartcontract.TransactOpts, fileName, cid)
}

// SmartcontractDeleteConfigurationIterator is returned from FilterDeleteConfiguration and is used to iterate over the raw logs and unpacked data for DeleteConfiguration events raised by the Smartcontract contract.
type SmartcontractDeleteConfigurationIterator struct {
	Event *SmartcontractDeleteConfiguration // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SmartcontractDeleteConfigurationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmartcontractDeleteConfiguration)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SmartcontractDeleteConfiguration)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SmartcontractDeleteConfigurationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmartcontractDeleteConfigurationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmartcontractDeleteConfiguration represents a DeleteConfiguration event raised by the Smartcontract contract.
type SmartcontractDeleteConfiguration struct {
	Org      common.Hash
	FileName common.Hash
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDeleteConfiguration is a free log retrieval operation binding the contract event 0x3645ce45e193d81395e9355613763738a298551d4b0e7ddd6e399707620fed9c.
//
// Solidity: event DeleteConfiguration(string indexed org, string indexed fileName)
func (_Smartcontract *SmartcontractFilterer) FilterDeleteConfiguration(opts *bind.FilterOpts, org []string, fileName []string) (*SmartcontractDeleteConfigurationIterator, error) {

	var orgRule []interface{}
	for _, orgItem := range org {
		orgRule = append(orgRule, orgItem)
	}
	var fileNameRule []interface{}
	for _, fileNameItem := range fileName {
		fileNameRule = append(fileNameRule, fileNameItem)
	}

	logs, sub, err := _Smartcontract.contract.FilterLogs(opts, "DeleteConfiguration", orgRule, fileNameRule)
	if err != nil {
		return nil, err
	}
	return &SmartcontractDeleteConfigurationIterator{contract: _Smartcontract.contract, event: "DeleteConfiguration", logs: logs, sub: sub}, nil
}

// WatchDeleteConfiguration is a free log subscription operation binding the contract event 0x3645ce45e193d81395e9355613763738a298551d4b0e7ddd6e399707620fed9c.
//
// Solidity: event DeleteConfiguration(string indexed org, string indexed fileName)
func (_Smartcontract *SmartcontractFilterer) WatchDeleteConfiguration(opts *bind.WatchOpts, sink chan<- *SmartcontractDeleteConfiguration, org []string, fileName []string) (event.Subscription, error) {

	var orgRule []interface{}
	for _, orgItem := range org {
		orgRule = append(orgRule, orgItem)
	}
	var fileNameRule []interface{}
	for _, fileNameItem := range fileName {
		fileNameRule = append(fileNameRule, fileNameItem)
	}

	logs, sub, err := _Smartcontract.contract.WatchLogs(opts, "DeleteConfiguration", orgRule, fileNameRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmartcontractDeleteConfiguration)
				if err := _Smartcontract.contract.UnpackLog(event, "DeleteConfiguration", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDeleteConfiguration is a log parse operation binding the contract event 0x3645ce45e193d81395e9355613763738a298551d4b0e7ddd6e399707620fed9c.
//
// Solidity: event DeleteConfiguration(string indexed org, string indexed fileName)
func (_Smartcontract *SmartcontractFilterer) ParseDeleteConfiguration(log types.Log) (*SmartcontractDeleteConfiguration, error) {
	event := new(SmartcontractDeleteConfiguration)
	if err := _Smartcontract.contract.UnpackLog(event, "DeleteConfiguration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SmartcontractDeleteSmartcontractIterator is returned from FilterDeleteSmartcontract and is used to iterate over the raw logs and unpacked data for DeleteSmartcontract events raised by the Smartcontract contract.
type SmartcontractDeleteSmartcontractIterator struct {
	Event *SmartcontractDeleteSmartcontract // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SmartcontractDeleteSmartcontractIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmartcontractDeleteSmartcontract)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SmartcontractDeleteSmartcontract)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SmartcontractDeleteSmartcontractIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmartcontractDeleteSmartcontractIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmartcontractDeleteSmartcontract represents a DeleteSmartcontract event raised by the Smartcontract contract.
type SmartcontractDeleteSmartcontract struct {
	Org      common.Hash
	FileName common.Hash
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDeleteSmartcontract is a free log retrieval operation binding the contract event 0x7d4d9b89988fba54319ca71cadca5b9a99595aa2abbba9dc3a53889512942fc8.
//
// Solidity: event DeleteSmartcontract(string indexed org, string indexed fileName)
func (_Smartcontract *SmartcontractFilterer) FilterDeleteSmartcontract(opts *bind.FilterOpts, org []string, fileName []string) (*SmartcontractDeleteSmartcontractIterator, error) {

	var orgRule []interface{}
	for _, orgItem := range org {
		orgRule = append(orgRule, orgItem)
	}
	var fileNameRule []interface{}
	for _, fileNameItem := range fileName {
		fileNameRule = append(fileNameRule, fileNameItem)
	}

	logs, sub, err := _Smartcontract.contract.FilterLogs(opts, "DeleteSmartcontract", orgRule, fileNameRule)
	if err != nil {
		return nil, err
	}
	return &SmartcontractDeleteSmartcontractIterator{contract: _Smartcontract.contract, event: "DeleteSmartcontract", logs: logs, sub: sub}, nil
}

// WatchDeleteSmartcontract is a free log subscription operation binding the contract event 0x7d4d9b89988fba54319ca71cadca5b9a99595aa2abbba9dc3a53889512942fc8.
//
// Solidity: event DeleteSmartcontract(string indexed org, string indexed fileName)
func (_Smartcontract *SmartcontractFilterer) WatchDeleteSmartcontract(opts *bind.WatchOpts, sink chan<- *SmartcontractDeleteSmartcontract, org []string, fileName []string) (event.Subscription, error) {

	var orgRule []interface{}
	for _, orgItem := range org {
		orgRule = append(orgRule, orgItem)
	}
	var fileNameRule []interface{}
	for _, fileNameItem := range fileName {
		fileNameRule = append(fileNameRule, fileNameItem)
	}

	logs, sub, err := _Smartcontract.contract.WatchLogs(opts, "DeleteSmartcontract", orgRule, fileNameRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmartcontractDeleteSmartcontract)
				if err := _Smartcontract.contract.UnpackLog(event, "DeleteSmartcontract", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDeleteSmartcontract is a log parse operation binding the contract event 0x7d4d9b89988fba54319ca71cadca5b9a99595aa2abbba9dc3a53889512942fc8.
//
// Solidity: event DeleteSmartcontract(string indexed org, string indexed fileName)
func (_Smartcontract *SmartcontractFilterer) ParseDeleteSmartcontract(log types.Log) (*SmartcontractDeleteSmartcontract, error) {
	event := new(SmartcontractDeleteSmartcontract)
	if err := _Smartcontract.contract.UnpackLog(event, "DeleteSmartcontract", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SmartcontractInsertConfigurationIterator is returned from FilterInsertConfiguration and is used to iterate over the raw logs and unpacked data for InsertConfiguration events raised by the Smartcontract contract.
type SmartcontractInsertConfigurationIterator struct {
	Event *SmartcontractInsertConfiguration // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SmartcontractInsertConfigurationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmartcontractInsertConfiguration)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SmartcontractInsertConfiguration)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SmartcontractInsertConfigurationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmartcontractInsertConfigurationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmartcontractInsertConfiguration represents a InsertConfiguration event raised by the Smartcontract contract.
type SmartcontractInsertConfiguration struct {
	Org      common.Hash
	FileName common.Hash
	Cid      common.Hash
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterInsertConfiguration is a free log retrieval operation binding the contract event 0xe4c71a24d7d86a4b9dcdf7bb43d15433ed42b5239d98cdd095cc8433545911e0.
//
// Solidity: event InsertConfiguration(string indexed org, string indexed fileName, string indexed cid)
func (_Smartcontract *SmartcontractFilterer) FilterInsertConfiguration(opts *bind.FilterOpts, org []string, fileName []string, cid []string) (*SmartcontractInsertConfigurationIterator, error) {

	var orgRule []interface{}
	for _, orgItem := range org {
		orgRule = append(orgRule, orgItem)
	}
	var fileNameRule []interface{}
	for _, fileNameItem := range fileName {
		fileNameRule = append(fileNameRule, fileNameItem)
	}
	var cidRule []interface{}
	for _, cidItem := range cid {
		cidRule = append(cidRule, cidItem)
	}

	logs, sub, err := _Smartcontract.contract.FilterLogs(opts, "InsertConfiguration", orgRule, fileNameRule, cidRule)
	if err != nil {
		return nil, err
	}
	return &SmartcontractInsertConfigurationIterator{contract: _Smartcontract.contract, event: "InsertConfiguration", logs: logs, sub: sub}, nil
}

// WatchInsertConfiguration is a free log subscription operation binding the contract event 0xe4c71a24d7d86a4b9dcdf7bb43d15433ed42b5239d98cdd095cc8433545911e0.
//
// Solidity: event InsertConfiguration(string indexed org, string indexed fileName, string indexed cid)
func (_Smartcontract *SmartcontractFilterer) WatchInsertConfiguration(opts *bind.WatchOpts, sink chan<- *SmartcontractInsertConfiguration, org []string, fileName []string, cid []string) (event.Subscription, error) {

	var orgRule []interface{}
	for _, orgItem := range org {
		orgRule = append(orgRule, orgItem)
	}
	var fileNameRule []interface{}
	for _, fileNameItem := range fileName {
		fileNameRule = append(fileNameRule, fileNameItem)
	}
	var cidRule []interface{}
	for _, cidItem := range cid {
		cidRule = append(cidRule, cidItem)
	}

	logs, sub, err := _Smartcontract.contract.WatchLogs(opts, "InsertConfiguration", orgRule, fileNameRule, cidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmartcontractInsertConfiguration)
				if err := _Smartcontract.contract.UnpackLog(event, "InsertConfiguration", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInsertConfiguration is a log parse operation binding the contract event 0xe4c71a24d7d86a4b9dcdf7bb43d15433ed42b5239d98cdd095cc8433545911e0.
//
// Solidity: event InsertConfiguration(string indexed org, string indexed fileName, string indexed cid)
func (_Smartcontract *SmartcontractFilterer) ParseInsertConfiguration(log types.Log) (*SmartcontractInsertConfiguration, error) {
	event := new(SmartcontractInsertConfiguration)
	if err := _Smartcontract.contract.UnpackLog(event, "InsertConfiguration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SmartcontractInsertOrganizationIterator is returned from FilterInsertOrganization and is used to iterate over the raw logs and unpacked data for InsertOrganization events raised by the Smartcontract contract.
type SmartcontractInsertOrganizationIterator struct {
	Event *SmartcontractInsertOrganization // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SmartcontractInsertOrganizationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmartcontractInsertOrganization)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SmartcontractInsertOrganization)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SmartcontractInsertOrganizationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmartcontractInsertOrganizationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmartcontractInsertOrganization represents a InsertOrganization event raised by the Smartcontract contract.
type SmartcontractInsertOrganization struct {
	Org common.Hash
	Raw types.Log // Blockchain specific contextual infos
}

// FilterInsertOrganization is a free log retrieval operation binding the contract event 0xc88b7686edda01cb5cb9cfc327a234fd55553b72ce44f51772a9ffd0be4f517a.
//
// Solidity: event InsertOrganization(string indexed org)
func (_Smartcontract *SmartcontractFilterer) FilterInsertOrganization(opts *bind.FilterOpts, org []string) (*SmartcontractInsertOrganizationIterator, error) {

	var orgRule []interface{}
	for _, orgItem := range org {
		orgRule = append(orgRule, orgItem)
	}

	logs, sub, err := _Smartcontract.contract.FilterLogs(opts, "InsertOrganization", orgRule)
	if err != nil {
		return nil, err
	}
	return &SmartcontractInsertOrganizationIterator{contract: _Smartcontract.contract, event: "InsertOrganization", logs: logs, sub: sub}, nil
}

// WatchInsertOrganization is a free log subscription operation binding the contract event 0xc88b7686edda01cb5cb9cfc327a234fd55553b72ce44f51772a9ffd0be4f517a.
//
// Solidity: event InsertOrganization(string indexed org)
func (_Smartcontract *SmartcontractFilterer) WatchInsertOrganization(opts *bind.WatchOpts, sink chan<- *SmartcontractInsertOrganization, org []string) (event.Subscription, error) {

	var orgRule []interface{}
	for _, orgItem := range org {
		orgRule = append(orgRule, orgItem)
	}

	logs, sub, err := _Smartcontract.contract.WatchLogs(opts, "InsertOrganization", orgRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmartcontractInsertOrganization)
				if err := _Smartcontract.contract.UnpackLog(event, "InsertOrganization", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInsertOrganization is a log parse operation binding the contract event 0xc88b7686edda01cb5cb9cfc327a234fd55553b72ce44f51772a9ffd0be4f517a.
//
// Solidity: event InsertOrganization(string indexed org)
func (_Smartcontract *SmartcontractFilterer) ParseInsertOrganization(log types.Log) (*SmartcontractInsertOrganization, error) {
	event := new(SmartcontractInsertOrganization)
	if err := _Smartcontract.contract.UnpackLog(event, "InsertOrganization", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SmartcontractInsertSmartcontractIterator is returned from FilterInsertSmartcontract and is used to iterate over the raw logs and unpacked data for InsertSmartcontract events raised by the Smartcontract contract.
type SmartcontractInsertSmartcontractIterator struct {
	Event *SmartcontractInsertSmartcontract // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SmartcontractInsertSmartcontractIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmartcontractInsertSmartcontract)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SmartcontractInsertSmartcontract)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SmartcontractInsertSmartcontractIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmartcontractInsertSmartcontractIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmartcontractInsertSmartcontract represents a InsertSmartcontract event raised by the Smartcontract contract.
type SmartcontractInsertSmartcontract struct {
	Org      common.Hash
	FileName common.Hash
	Cid      common.Hash
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterInsertSmartcontract is a free log retrieval operation binding the contract event 0x0e7ebd6fc52c2b4f8dfefdaaba6996ec0c5104c174da6e43d16a314a286fdd93.
//
// Solidity: event InsertSmartcontract(string indexed org, string indexed fileName, string indexed cid)
func (_Smartcontract *SmartcontractFilterer) FilterInsertSmartcontract(opts *bind.FilterOpts, org []string, fileName []string, cid []string) (*SmartcontractInsertSmartcontractIterator, error) {

	var orgRule []interface{}
	for _, orgItem := range org {
		orgRule = append(orgRule, orgItem)
	}
	var fileNameRule []interface{}
	for _, fileNameItem := range fileName {
		fileNameRule = append(fileNameRule, fileNameItem)
	}
	var cidRule []interface{}
	for _, cidItem := range cid {
		cidRule = append(cidRule, cidItem)
	}

	logs, sub, err := _Smartcontract.contract.FilterLogs(opts, "InsertSmartcontract", orgRule, fileNameRule, cidRule)
	if err != nil {
		return nil, err
	}
	return &SmartcontractInsertSmartcontractIterator{contract: _Smartcontract.contract, event: "InsertSmartcontract", logs: logs, sub: sub}, nil
}

// WatchInsertSmartcontract is a free log subscription operation binding the contract event 0x0e7ebd6fc52c2b4f8dfefdaaba6996ec0c5104c174da6e43d16a314a286fdd93.
//
// Solidity: event InsertSmartcontract(string indexed org, string indexed fileName, string indexed cid)
func (_Smartcontract *SmartcontractFilterer) WatchInsertSmartcontract(opts *bind.WatchOpts, sink chan<- *SmartcontractInsertSmartcontract, org []string, fileName []string, cid []string) (event.Subscription, error) {

	var orgRule []interface{}
	for _, orgItem := range org {
		orgRule = append(orgRule, orgItem)
	}
	var fileNameRule []interface{}
	for _, fileNameItem := range fileName {
		fileNameRule = append(fileNameRule, fileNameItem)
	}
	var cidRule []interface{}
	for _, cidItem := range cid {
		cidRule = append(cidRule, cidItem)
	}

	logs, sub, err := _Smartcontract.contract.WatchLogs(opts, "InsertSmartcontract", orgRule, fileNameRule, cidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmartcontractInsertSmartcontract)
				if err := _Smartcontract.contract.UnpackLog(event, "InsertSmartcontract", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInsertSmartcontract is a log parse operation binding the contract event 0x0e7ebd6fc52c2b4f8dfefdaaba6996ec0c5104c174da6e43d16a314a286fdd93.
//
// Solidity: event InsertSmartcontract(string indexed org, string indexed fileName, string indexed cid)
func (_Smartcontract *SmartcontractFilterer) ParseInsertSmartcontract(log types.Log) (*SmartcontractInsertSmartcontract, error) {
	event := new(SmartcontractInsertSmartcontract)
	if err := _Smartcontract.contract.UnpackLog(event, "InsertSmartcontract", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SmartcontractUpdateConfigurationIterator is returned from FilterUpdateConfiguration and is used to iterate over the raw logs and unpacked data for UpdateConfiguration events raised by the Smartcontract contract.
type SmartcontractUpdateConfigurationIterator struct {
	Event *SmartcontractUpdateConfiguration // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SmartcontractUpdateConfigurationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmartcontractUpdateConfiguration)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SmartcontractUpdateConfiguration)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SmartcontractUpdateConfigurationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmartcontractUpdateConfigurationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmartcontractUpdateConfiguration represents a UpdateConfiguration event raised by the Smartcontract contract.
type SmartcontractUpdateConfiguration struct {
	FileName common.Hash
	Cid      common.Hash
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUpdateConfiguration is a free log retrieval operation binding the contract event 0xa183a608b7c8ba72d4b45247ef3b431ceb231e76761f3709c6679c7d1fbf9749.
//
// Solidity: event UpdateConfiguration(string indexed fileName, string indexed cid)
func (_Smartcontract *SmartcontractFilterer) FilterUpdateConfiguration(opts *bind.FilterOpts, fileName []string, cid []string) (*SmartcontractUpdateConfigurationIterator, error) {

	var fileNameRule []interface{}
	for _, fileNameItem := range fileName {
		fileNameRule = append(fileNameRule, fileNameItem)
	}
	var cidRule []interface{}
	for _, cidItem := range cid {
		cidRule = append(cidRule, cidItem)
	}

	logs, sub, err := _Smartcontract.contract.FilterLogs(opts, "UpdateConfiguration", fileNameRule, cidRule)
	if err != nil {
		return nil, err
	}
	return &SmartcontractUpdateConfigurationIterator{contract: _Smartcontract.contract, event: "UpdateConfiguration", logs: logs, sub: sub}, nil
}

// WatchUpdateConfiguration is a free log subscription operation binding the contract event 0xa183a608b7c8ba72d4b45247ef3b431ceb231e76761f3709c6679c7d1fbf9749.
//
// Solidity: event UpdateConfiguration(string indexed fileName, string indexed cid)
func (_Smartcontract *SmartcontractFilterer) WatchUpdateConfiguration(opts *bind.WatchOpts, sink chan<- *SmartcontractUpdateConfiguration, fileName []string, cid []string) (event.Subscription, error) {

	var fileNameRule []interface{}
	for _, fileNameItem := range fileName {
		fileNameRule = append(fileNameRule, fileNameItem)
	}
	var cidRule []interface{}
	for _, cidItem := range cid {
		cidRule = append(cidRule, cidItem)
	}

	logs, sub, err := _Smartcontract.contract.WatchLogs(opts, "UpdateConfiguration", fileNameRule, cidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmartcontractUpdateConfiguration)
				if err := _Smartcontract.contract.UnpackLog(event, "UpdateConfiguration", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdateConfiguration is a log parse operation binding the contract event 0xa183a608b7c8ba72d4b45247ef3b431ceb231e76761f3709c6679c7d1fbf9749.
//
// Solidity: event UpdateConfiguration(string indexed fileName, string indexed cid)
func (_Smartcontract *SmartcontractFilterer) ParseUpdateConfiguration(log types.Log) (*SmartcontractUpdateConfiguration, error) {
	event := new(SmartcontractUpdateConfiguration)
	if err := _Smartcontract.contract.UnpackLog(event, "UpdateConfiguration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SmartcontractUpdateSmartcontractIterator is returned from FilterUpdateSmartcontract and is used to iterate over the raw logs and unpacked data for UpdateSmartcontract events raised by the Smartcontract contract.
type SmartcontractUpdateSmartcontractIterator struct {
	Event *SmartcontractUpdateSmartcontract // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SmartcontractUpdateSmartcontractIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmartcontractUpdateSmartcontract)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SmartcontractUpdateSmartcontract)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SmartcontractUpdateSmartcontractIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmartcontractUpdateSmartcontractIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmartcontractUpdateSmartcontract represents a UpdateSmartcontract event raised by the Smartcontract contract.
type SmartcontractUpdateSmartcontract struct {
	FileName common.Hash
	Cid      common.Hash
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUpdateSmartcontract is a free log retrieval operation binding the contract event 0xcdf23f587b224317852b246c67b694bc63fc30d22be0c20193cbe5bd69d9b7c2.
//
// Solidity: event UpdateSmartcontract(string indexed fileName, string indexed cid)
func (_Smartcontract *SmartcontractFilterer) FilterUpdateSmartcontract(opts *bind.FilterOpts, fileName []string, cid []string) (*SmartcontractUpdateSmartcontractIterator, error) {

	var fileNameRule []interface{}
	for _, fileNameItem := range fileName {
		fileNameRule = append(fileNameRule, fileNameItem)
	}
	var cidRule []interface{}
	for _, cidItem := range cid {
		cidRule = append(cidRule, cidItem)
	}

	logs, sub, err := _Smartcontract.contract.FilterLogs(opts, "UpdateSmartcontract", fileNameRule, cidRule)
	if err != nil {
		return nil, err
	}
	return &SmartcontractUpdateSmartcontractIterator{contract: _Smartcontract.contract, event: "UpdateSmartcontract", logs: logs, sub: sub}, nil
}

// WatchUpdateSmartcontract is a free log subscription operation binding the contract event 0xcdf23f587b224317852b246c67b694bc63fc30d22be0c20193cbe5bd69d9b7c2.
//
// Solidity: event UpdateSmartcontract(string indexed fileName, string indexed cid)
func (_Smartcontract *SmartcontractFilterer) WatchUpdateSmartcontract(opts *bind.WatchOpts, sink chan<- *SmartcontractUpdateSmartcontract, fileName []string, cid []string) (event.Subscription, error) {

	var fileNameRule []interface{}
	for _, fileNameItem := range fileName {
		fileNameRule = append(fileNameRule, fileNameItem)
	}
	var cidRule []interface{}
	for _, cidItem := range cid {
		cidRule = append(cidRule, cidItem)
	}

	logs, sub, err := _Smartcontract.contract.WatchLogs(opts, "UpdateSmartcontract", fileNameRule, cidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmartcontractUpdateSmartcontract)
				if err := _Smartcontract.contract.UnpackLog(event, "UpdateSmartcontract", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdateSmartcontract is a log parse operation binding the contract event 0xcdf23f587b224317852b246c67b694bc63fc30d22be0c20193cbe5bd69d9b7c2.
//
// Solidity: event UpdateSmartcontract(string indexed fileName, string indexed cid)
func (_Smartcontract *SmartcontractFilterer) ParseUpdateSmartcontract(log types.Log) (*SmartcontractUpdateSmartcontract, error) {
	event := new(SmartcontractUpdateSmartcontract)
	if err := _Smartcontract.contract.UnpackLog(event, "UpdateSmartcontract", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

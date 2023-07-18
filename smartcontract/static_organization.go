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

// StaticOrganizationMetaData contains all meta data concerning the StaticOrganization contract.
var StaticOrganizationMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"org\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"fileName\",\"type\":\"string\"}],\"name\":\"DeleteConfiguration\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"org\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"fileName\",\"type\":\"string\"}],\"name\":\"DeleteSmartcontract\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"org\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"fileName\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"}],\"name\":\"InsertConfiguration\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"org\",\"type\":\"string\"}],\"name\":\"InsertOrganization\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"org\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"fileName\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"}],\"name\":\"InsertSmartcontract\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"fileName\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"}],\"name\":\"UpdateConfiguration\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"fileName\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"}],\"name\":\"UpdateSmartcontract\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"configurations\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"org\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"fileName\",\"type\":\"string\"}],\"name\":\"deleteConfiguration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"org\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"fileName\",\"type\":\"string\"}],\"name\":\"deleteSmartcontract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"org\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"fileName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"}],\"name\":\"insertConfiguration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"org\",\"type\":\"string\"}],\"name\":\"insertOrganization\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"org\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"fileName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"}],\"name\":\"insertSmartcontract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"organizationConfigurations\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"organizationSmartcontracts\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"organizations\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"smartcontracts\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"fileName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"}],\"name\":\"updateConfiguration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"fileName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"}],\"name\":\"updateSmartcontract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// StaticOrganizationABI is the input ABI used to generate the binding from.
// Deprecated: Use StaticOrganizationMetaData.ABI instead.
var StaticOrganizationABI = StaticOrganizationMetaData.ABI

// StaticOrganization is an auto generated Go binding around an Ethereum contract.
type StaticOrganization struct {
	StaticOrganizationCaller     // Read-only binding to the contract
	StaticOrganizationTransactor // Write-only binding to the contract
	StaticOrganizationFilterer   // Log filterer for contract events
}

// StaticOrganizationCaller is an auto generated read-only Go binding around an Ethereum contract.
type StaticOrganizationCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StaticOrganizationTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StaticOrganizationTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StaticOrganizationFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StaticOrganizationFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StaticOrganizationSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StaticOrganizationSession struct {
	Contract     *StaticOrganization // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// StaticOrganizationCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StaticOrganizationCallerSession struct {
	Contract *StaticOrganizationCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// StaticOrganizationTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StaticOrganizationTransactorSession struct {
	Contract     *StaticOrganizationTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// StaticOrganizationRaw is an auto generated low-level Go binding around an Ethereum contract.
type StaticOrganizationRaw struct {
	Contract *StaticOrganization // Generic contract binding to access the raw methods on
}

// StaticOrganizationCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StaticOrganizationCallerRaw struct {
	Contract *StaticOrganizationCaller // Generic read-only contract binding to access the raw methods on
}

// StaticOrganizationTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StaticOrganizationTransactorRaw struct {
	Contract *StaticOrganizationTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStaticOrganization creates a new instance of StaticOrganization, bound to a specific deployed contract.
func NewStaticOrganization(address common.Address, backend bind.ContractBackend) (*StaticOrganization, error) {
	contract, err := bindStaticOrganization(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StaticOrganization{StaticOrganizationCaller: StaticOrganizationCaller{contract: contract}, StaticOrganizationTransactor: StaticOrganizationTransactor{contract: contract}, StaticOrganizationFilterer: StaticOrganizationFilterer{contract: contract}}, nil
}

// NewStaticOrganizationCaller creates a new read-only instance of StaticOrganization, bound to a specific deployed contract.
func NewStaticOrganizationCaller(address common.Address, caller bind.ContractCaller) (*StaticOrganizationCaller, error) {
	contract, err := bindStaticOrganization(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StaticOrganizationCaller{contract: contract}, nil
}

// NewStaticOrganizationTransactor creates a new write-only instance of StaticOrganization, bound to a specific deployed contract.
func NewStaticOrganizationTransactor(address common.Address, transactor bind.ContractTransactor) (*StaticOrganizationTransactor, error) {
	contract, err := bindStaticOrganization(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StaticOrganizationTransactor{contract: contract}, nil
}

// NewStaticOrganizationFilterer creates a new log filterer instance of StaticOrganization, bound to a specific deployed contract.
func NewStaticOrganizationFilterer(address common.Address, filterer bind.ContractFilterer) (*StaticOrganizationFilterer, error) {
	contract, err := bindStaticOrganization(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StaticOrganizationFilterer{contract: contract}, nil
}

// bindStaticOrganization binds a generic wrapper to an already deployed contract.
func bindStaticOrganization(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StaticOrganizationMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StaticOrganization *StaticOrganizationRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StaticOrganization.Contract.StaticOrganizationCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StaticOrganization *StaticOrganizationRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StaticOrganization.Contract.StaticOrganizationTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StaticOrganization *StaticOrganizationRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StaticOrganization.Contract.StaticOrganizationTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StaticOrganization *StaticOrganizationCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StaticOrganization.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StaticOrganization *StaticOrganizationTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StaticOrganization.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StaticOrganization *StaticOrganizationTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StaticOrganization.Contract.contract.Transact(opts, method, params...)
}

// Configurations is a free data retrieval call binding the contract method 0x1214dd58.
//
// Solidity: function configurations(string ) view returns(string)
func (_StaticOrganization *StaticOrganizationCaller) Configurations(opts *bind.CallOpts, arg0 string) (string, error) {
	var out []interface{}
	err := _StaticOrganization.contract.Call(opts, &out, "configurations", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Configurations is a free data retrieval call binding the contract method 0x1214dd58.
//
// Solidity: function configurations(string ) view returns(string)
func (_StaticOrganization *StaticOrganizationSession) Configurations(arg0 string) (string, error) {
	return _StaticOrganization.Contract.Configurations(&_StaticOrganization.CallOpts, arg0)
}

// Configurations is a free data retrieval call binding the contract method 0x1214dd58.
//
// Solidity: function configurations(string ) view returns(string)
func (_StaticOrganization *StaticOrganizationCallerSession) Configurations(arg0 string) (string, error) {
	return _StaticOrganization.Contract.Configurations(&_StaticOrganization.CallOpts, arg0)
}

// OrganizationConfigurations is a free data retrieval call binding the contract method 0x6f74463f.
//
// Solidity: function organizationConfigurations(string , uint256 ) view returns(string)
func (_StaticOrganization *StaticOrganizationCaller) OrganizationConfigurations(opts *bind.CallOpts, arg0 string, arg1 *big.Int) (string, error) {
	var out []interface{}
	err := _StaticOrganization.contract.Call(opts, &out, "organizationConfigurations", arg0, arg1)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// OrganizationConfigurations is a free data retrieval call binding the contract method 0x6f74463f.
//
// Solidity: function organizationConfigurations(string , uint256 ) view returns(string)
func (_StaticOrganization *StaticOrganizationSession) OrganizationConfigurations(arg0 string, arg1 *big.Int) (string, error) {
	return _StaticOrganization.Contract.OrganizationConfigurations(&_StaticOrganization.CallOpts, arg0, arg1)
}

// OrganizationConfigurations is a free data retrieval call binding the contract method 0x6f74463f.
//
// Solidity: function organizationConfigurations(string , uint256 ) view returns(string)
func (_StaticOrganization *StaticOrganizationCallerSession) OrganizationConfigurations(arg0 string, arg1 *big.Int) (string, error) {
	return _StaticOrganization.Contract.OrganizationConfigurations(&_StaticOrganization.CallOpts, arg0, arg1)
}

// OrganizationSmartcontracts is a free data retrieval call binding the contract method 0x8718dc72.
//
// Solidity: function organizationSmartcontracts(string , uint256 ) view returns(string)
func (_StaticOrganization *StaticOrganizationCaller) OrganizationSmartcontracts(opts *bind.CallOpts, arg0 string, arg1 *big.Int) (string, error) {
	var out []interface{}
	err := _StaticOrganization.contract.Call(opts, &out, "organizationSmartcontracts", arg0, arg1)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// OrganizationSmartcontracts is a free data retrieval call binding the contract method 0x8718dc72.
//
// Solidity: function organizationSmartcontracts(string , uint256 ) view returns(string)
func (_StaticOrganization *StaticOrganizationSession) OrganizationSmartcontracts(arg0 string, arg1 *big.Int) (string, error) {
	return _StaticOrganization.Contract.OrganizationSmartcontracts(&_StaticOrganization.CallOpts, arg0, arg1)
}

// OrganizationSmartcontracts is a free data retrieval call binding the contract method 0x8718dc72.
//
// Solidity: function organizationSmartcontracts(string , uint256 ) view returns(string)
func (_StaticOrganization *StaticOrganizationCallerSession) OrganizationSmartcontracts(arg0 string, arg1 *big.Int) (string, error) {
	return _StaticOrganization.Contract.OrganizationSmartcontracts(&_StaticOrganization.CallOpts, arg0, arg1)
}

// Organizations is a free data retrieval call binding the contract method 0xe311f43c.
//
// Solidity: function organizations(string ) view returns(bool)
func (_StaticOrganization *StaticOrganizationCaller) Organizations(opts *bind.CallOpts, arg0 string) (bool, error) {
	var out []interface{}
	err := _StaticOrganization.contract.Call(opts, &out, "organizations", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Organizations is a free data retrieval call binding the contract method 0xe311f43c.
//
// Solidity: function organizations(string ) view returns(bool)
func (_StaticOrganization *StaticOrganizationSession) Organizations(arg0 string) (bool, error) {
	return _StaticOrganization.Contract.Organizations(&_StaticOrganization.CallOpts, arg0)
}

// Organizations is a free data retrieval call binding the contract method 0xe311f43c.
//
// Solidity: function organizations(string ) view returns(bool)
func (_StaticOrganization *StaticOrganizationCallerSession) Organizations(arg0 string) (bool, error) {
	return _StaticOrganization.Contract.Organizations(&_StaticOrganization.CallOpts, arg0)
}

// Smartcontracts is a free data retrieval call binding the contract method 0xbcd8e820.
//
// Solidity: function smartcontracts(string ) view returns(string)
func (_StaticOrganization *StaticOrganizationCaller) Smartcontracts(opts *bind.CallOpts, arg0 string) (string, error) {
	var out []interface{}
	err := _StaticOrganization.contract.Call(opts, &out, "smartcontracts", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Smartcontracts is a free data retrieval call binding the contract method 0xbcd8e820.
//
// Solidity: function smartcontracts(string ) view returns(string)
func (_StaticOrganization *StaticOrganizationSession) Smartcontracts(arg0 string) (string, error) {
	return _StaticOrganization.Contract.Smartcontracts(&_StaticOrganization.CallOpts, arg0)
}

// Smartcontracts is a free data retrieval call binding the contract method 0xbcd8e820.
//
// Solidity: function smartcontracts(string ) view returns(string)
func (_StaticOrganization *StaticOrganizationCallerSession) Smartcontracts(arg0 string) (string, error) {
	return _StaticOrganization.Contract.Smartcontracts(&_StaticOrganization.CallOpts, arg0)
}

// DeleteConfiguration is a paid mutator transaction binding the contract method 0xb58f296a.
//
// Solidity: function deleteConfiguration(string org, string fileName) returns()
func (_StaticOrganization *StaticOrganizationTransactor) DeleteConfiguration(opts *bind.TransactOpts, org string, fileName string) (*types.Transaction, error) {
	return _StaticOrganization.contract.Transact(opts, "deleteConfiguration", org, fileName)
}

// DeleteConfiguration is a paid mutator transaction binding the contract method 0xb58f296a.
//
// Solidity: function deleteConfiguration(string org, string fileName) returns()
func (_StaticOrganization *StaticOrganizationSession) DeleteConfiguration(org string, fileName string) (*types.Transaction, error) {
	return _StaticOrganization.Contract.DeleteConfiguration(&_StaticOrganization.TransactOpts, org, fileName)
}

// DeleteConfiguration is a paid mutator transaction binding the contract method 0xb58f296a.
//
// Solidity: function deleteConfiguration(string org, string fileName) returns()
func (_StaticOrganization *StaticOrganizationTransactorSession) DeleteConfiguration(org string, fileName string) (*types.Transaction, error) {
	return _StaticOrganization.Contract.DeleteConfiguration(&_StaticOrganization.TransactOpts, org, fileName)
}

// DeleteSmartcontract is a paid mutator transaction binding the contract method 0x371961f8.
//
// Solidity: function deleteSmartcontract(string org, string fileName) returns()
func (_StaticOrganization *StaticOrganizationTransactor) DeleteSmartcontract(opts *bind.TransactOpts, org string, fileName string) (*types.Transaction, error) {
	return _StaticOrganization.contract.Transact(opts, "deleteSmartcontract", org, fileName)
}

// DeleteSmartcontract is a paid mutator transaction binding the contract method 0x371961f8.
//
// Solidity: function deleteSmartcontract(string org, string fileName) returns()
func (_StaticOrganization *StaticOrganizationSession) DeleteSmartcontract(org string, fileName string) (*types.Transaction, error) {
	return _StaticOrganization.Contract.DeleteSmartcontract(&_StaticOrganization.TransactOpts, org, fileName)
}

// DeleteSmartcontract is a paid mutator transaction binding the contract method 0x371961f8.
//
// Solidity: function deleteSmartcontract(string org, string fileName) returns()
func (_StaticOrganization *StaticOrganizationTransactorSession) DeleteSmartcontract(org string, fileName string) (*types.Transaction, error) {
	return _StaticOrganization.Contract.DeleteSmartcontract(&_StaticOrganization.TransactOpts, org, fileName)
}

// InsertConfiguration is a paid mutator transaction binding the contract method 0x9ea508f8.
//
// Solidity: function insertConfiguration(string org, string fileName, string cid) returns()
func (_StaticOrganization *StaticOrganizationTransactor) InsertConfiguration(opts *bind.TransactOpts, org string, fileName string, cid string) (*types.Transaction, error) {
	return _StaticOrganization.contract.Transact(opts, "insertConfiguration", org, fileName, cid)
}

// InsertConfiguration is a paid mutator transaction binding the contract method 0x9ea508f8.
//
// Solidity: function insertConfiguration(string org, string fileName, string cid) returns()
func (_StaticOrganization *StaticOrganizationSession) InsertConfiguration(org string, fileName string, cid string) (*types.Transaction, error) {
	return _StaticOrganization.Contract.InsertConfiguration(&_StaticOrganization.TransactOpts, org, fileName, cid)
}

// InsertConfiguration is a paid mutator transaction binding the contract method 0x9ea508f8.
//
// Solidity: function insertConfiguration(string org, string fileName, string cid) returns()
func (_StaticOrganization *StaticOrganizationTransactorSession) InsertConfiguration(org string, fileName string, cid string) (*types.Transaction, error) {
	return _StaticOrganization.Contract.InsertConfiguration(&_StaticOrganization.TransactOpts, org, fileName, cid)
}

// InsertOrganization is a paid mutator transaction binding the contract method 0x62389e43.
//
// Solidity: function insertOrganization(string org) returns()
func (_StaticOrganization *StaticOrganizationTransactor) InsertOrganization(opts *bind.TransactOpts, org string) (*types.Transaction, error) {
	return _StaticOrganization.contract.Transact(opts, "insertOrganization", org)
}

// InsertOrganization is a paid mutator transaction binding the contract method 0x62389e43.
//
// Solidity: function insertOrganization(string org) returns()
func (_StaticOrganization *StaticOrganizationSession) InsertOrganization(org string) (*types.Transaction, error) {
	return _StaticOrganization.Contract.InsertOrganization(&_StaticOrganization.TransactOpts, org)
}

// InsertOrganization is a paid mutator transaction binding the contract method 0x62389e43.
//
// Solidity: function insertOrganization(string org) returns()
func (_StaticOrganization *StaticOrganizationTransactorSession) InsertOrganization(org string) (*types.Transaction, error) {
	return _StaticOrganization.Contract.InsertOrganization(&_StaticOrganization.TransactOpts, org)
}

// InsertSmartcontract is a paid mutator transaction binding the contract method 0xad5f0707.
//
// Solidity: function insertSmartcontract(string org, string fileName, string cid) returns()
func (_StaticOrganization *StaticOrganizationTransactor) InsertSmartcontract(opts *bind.TransactOpts, org string, fileName string, cid string) (*types.Transaction, error) {
	return _StaticOrganization.contract.Transact(opts, "insertSmartcontract", org, fileName, cid)
}

// InsertSmartcontract is a paid mutator transaction binding the contract method 0xad5f0707.
//
// Solidity: function insertSmartcontract(string org, string fileName, string cid) returns()
func (_StaticOrganization *StaticOrganizationSession) InsertSmartcontract(org string, fileName string, cid string) (*types.Transaction, error) {
	return _StaticOrganization.Contract.InsertSmartcontract(&_StaticOrganization.TransactOpts, org, fileName, cid)
}

// InsertSmartcontract is a paid mutator transaction binding the contract method 0xad5f0707.
//
// Solidity: function insertSmartcontract(string org, string fileName, string cid) returns()
func (_StaticOrganization *StaticOrganizationTransactorSession) InsertSmartcontract(org string, fileName string, cid string) (*types.Transaction, error) {
	return _StaticOrganization.Contract.InsertSmartcontract(&_StaticOrganization.TransactOpts, org, fileName, cid)
}

// UpdateConfiguration is a paid mutator transaction binding the contract method 0x8084163e.
//
// Solidity: function updateConfiguration(string fileName, string cid) returns()
func (_StaticOrganization *StaticOrganizationTransactor) UpdateConfiguration(opts *bind.TransactOpts, fileName string, cid string) (*types.Transaction, error) {
	return _StaticOrganization.contract.Transact(opts, "updateConfiguration", fileName, cid)
}

// UpdateConfiguration is a paid mutator transaction binding the contract method 0x8084163e.
//
// Solidity: function updateConfiguration(string fileName, string cid) returns()
func (_StaticOrganization *StaticOrganizationSession) UpdateConfiguration(fileName string, cid string) (*types.Transaction, error) {
	return _StaticOrganization.Contract.UpdateConfiguration(&_StaticOrganization.TransactOpts, fileName, cid)
}

// UpdateConfiguration is a paid mutator transaction binding the contract method 0x8084163e.
//
// Solidity: function updateConfiguration(string fileName, string cid) returns()
func (_StaticOrganization *StaticOrganizationTransactorSession) UpdateConfiguration(fileName string, cid string) (*types.Transaction, error) {
	return _StaticOrganization.Contract.UpdateConfiguration(&_StaticOrganization.TransactOpts, fileName, cid)
}

// UpdateSmartcontract is a paid mutator transaction binding the contract method 0xf252b7ec.
//
// Solidity: function updateSmartcontract(string fileName, string cid) returns()
func (_StaticOrganization *StaticOrganizationTransactor) UpdateSmartcontract(opts *bind.TransactOpts, fileName string, cid string) (*types.Transaction, error) {
	return _StaticOrganization.contract.Transact(opts, "updateSmartcontract", fileName, cid)
}

// UpdateSmartcontract is a paid mutator transaction binding the contract method 0xf252b7ec.
//
// Solidity: function updateSmartcontract(string fileName, string cid) returns()
func (_StaticOrganization *StaticOrganizationSession) UpdateSmartcontract(fileName string, cid string) (*types.Transaction, error) {
	return _StaticOrganization.Contract.UpdateSmartcontract(&_StaticOrganization.TransactOpts, fileName, cid)
}

// UpdateSmartcontract is a paid mutator transaction binding the contract method 0xf252b7ec.
//
// Solidity: function updateSmartcontract(string fileName, string cid) returns()
func (_StaticOrganization *StaticOrganizationTransactorSession) UpdateSmartcontract(fileName string, cid string) (*types.Transaction, error) {
	return _StaticOrganization.Contract.UpdateSmartcontract(&_StaticOrganization.TransactOpts, fileName, cid)
}

// StaticOrganizationDeleteConfigurationIterator is returned from FilterDeleteConfiguration and is used to iterate over the raw logs and unpacked data for DeleteConfiguration events raised by the StaticOrganization contract.
type StaticOrganizationDeleteConfigurationIterator struct {
	Event *StaticOrganizationDeleteConfiguration // Event containing the contract specifics and raw log

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
func (it *StaticOrganizationDeleteConfigurationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StaticOrganizationDeleteConfiguration)
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
		it.Event = new(StaticOrganizationDeleteConfiguration)
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
func (it *StaticOrganizationDeleteConfigurationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StaticOrganizationDeleteConfigurationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StaticOrganizationDeleteConfiguration represents a DeleteConfiguration event raised by the StaticOrganization contract.
type StaticOrganizationDeleteConfiguration struct {
	Org      common.Hash
	FileName common.Hash
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDeleteConfiguration is a free log retrieval operation binding the contract event 0x3645ce45e193d81395e9355613763738a298551d4b0e7ddd6e399707620fed9c.
//
// Solidity: event DeleteConfiguration(string indexed org, string indexed fileName)
func (_StaticOrganization *StaticOrganizationFilterer) FilterDeleteConfiguration(opts *bind.FilterOpts, org []string, fileName []string) (*StaticOrganizationDeleteConfigurationIterator, error) {

	var orgRule []interface{}
	for _, orgItem := range org {
		orgRule = append(orgRule, orgItem)
	}
	var fileNameRule []interface{}
	for _, fileNameItem := range fileName {
		fileNameRule = append(fileNameRule, fileNameItem)
	}

	logs, sub, err := _StaticOrganization.contract.FilterLogs(opts, "DeleteConfiguration", orgRule, fileNameRule)
	if err != nil {
		return nil, err
	}
	return &StaticOrganizationDeleteConfigurationIterator{contract: _StaticOrganization.contract, event: "DeleteConfiguration", logs: logs, sub: sub}, nil
}

// WatchDeleteConfiguration is a free log subscription operation binding the contract event 0x3645ce45e193d81395e9355613763738a298551d4b0e7ddd6e399707620fed9c.
//
// Solidity: event DeleteConfiguration(string indexed org, string indexed fileName)
func (_StaticOrganization *StaticOrganizationFilterer) WatchDeleteConfiguration(opts *bind.WatchOpts, sink chan<- *StaticOrganizationDeleteConfiguration, org []string, fileName []string) (event.Subscription, error) {

	var orgRule []interface{}
	for _, orgItem := range org {
		orgRule = append(orgRule, orgItem)
	}
	var fileNameRule []interface{}
	for _, fileNameItem := range fileName {
		fileNameRule = append(fileNameRule, fileNameItem)
	}

	logs, sub, err := _StaticOrganization.contract.WatchLogs(opts, "DeleteConfiguration", orgRule, fileNameRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StaticOrganizationDeleteConfiguration)
				if err := _StaticOrganization.contract.UnpackLog(event, "DeleteConfiguration", log); err != nil {
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
func (_StaticOrganization *StaticOrganizationFilterer) ParseDeleteConfiguration(log types.Log) (*StaticOrganizationDeleteConfiguration, error) {
	event := new(StaticOrganizationDeleteConfiguration)
	if err := _StaticOrganization.contract.UnpackLog(event, "DeleteConfiguration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StaticOrganizationDeleteSmartcontractIterator is returned from FilterDeleteSmartcontract and is used to iterate over the raw logs and unpacked data for DeleteSmartcontract events raised by the StaticOrganization contract.
type StaticOrganizationDeleteSmartcontractIterator struct {
	Event *StaticOrganizationDeleteSmartcontract // Event containing the contract specifics and raw log

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
func (it *StaticOrganizationDeleteSmartcontractIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StaticOrganizationDeleteSmartcontract)
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
		it.Event = new(StaticOrganizationDeleteSmartcontract)
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
func (it *StaticOrganizationDeleteSmartcontractIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StaticOrganizationDeleteSmartcontractIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StaticOrganizationDeleteSmartcontract represents a DeleteSmartcontract event raised by the StaticOrganization contract.
type StaticOrganizationDeleteSmartcontract struct {
	Org      common.Hash
	FileName common.Hash
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDeleteSmartcontract is a free log retrieval operation binding the contract event 0x7d4d9b89988fba54319ca71cadca5b9a99595aa2abbba9dc3a53889512942fc8.
//
// Solidity: event DeleteSmartcontract(string indexed org, string indexed fileName)
func (_StaticOrganization *StaticOrganizationFilterer) FilterDeleteSmartcontract(opts *bind.FilterOpts, org []string, fileName []string) (*StaticOrganizationDeleteSmartcontractIterator, error) {

	var orgRule []interface{}
	for _, orgItem := range org {
		orgRule = append(orgRule, orgItem)
	}
	var fileNameRule []interface{}
	for _, fileNameItem := range fileName {
		fileNameRule = append(fileNameRule, fileNameItem)
	}

	logs, sub, err := _StaticOrganization.contract.FilterLogs(opts, "DeleteSmartcontract", orgRule, fileNameRule)
	if err != nil {
		return nil, err
	}
	return &StaticOrganizationDeleteSmartcontractIterator{contract: _StaticOrganization.contract, event: "DeleteSmartcontract", logs: logs, sub: sub}, nil
}

// WatchDeleteSmartcontract is a free log subscription operation binding the contract event 0x7d4d9b89988fba54319ca71cadca5b9a99595aa2abbba9dc3a53889512942fc8.
//
// Solidity: event DeleteSmartcontract(string indexed org, string indexed fileName)
func (_StaticOrganization *StaticOrganizationFilterer) WatchDeleteSmartcontract(opts *bind.WatchOpts, sink chan<- *StaticOrganizationDeleteSmartcontract, org []string, fileName []string) (event.Subscription, error) {

	var orgRule []interface{}
	for _, orgItem := range org {
		orgRule = append(orgRule, orgItem)
	}
	var fileNameRule []interface{}
	for _, fileNameItem := range fileName {
		fileNameRule = append(fileNameRule, fileNameItem)
	}

	logs, sub, err := _StaticOrganization.contract.WatchLogs(opts, "DeleteSmartcontract", orgRule, fileNameRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StaticOrganizationDeleteSmartcontract)
				if err := _StaticOrganization.contract.UnpackLog(event, "DeleteSmartcontract", log); err != nil {
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
func (_StaticOrganization *StaticOrganizationFilterer) ParseDeleteSmartcontract(log types.Log) (*StaticOrganizationDeleteSmartcontract, error) {
	event := new(StaticOrganizationDeleteSmartcontract)
	if err := _StaticOrganization.contract.UnpackLog(event, "DeleteSmartcontract", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StaticOrganizationInsertConfigurationIterator is returned from FilterInsertConfiguration and is used to iterate over the raw logs and unpacked data for InsertConfiguration events raised by the StaticOrganization contract.
type StaticOrganizationInsertConfigurationIterator struct {
	Event *StaticOrganizationInsertConfiguration // Event containing the contract specifics and raw log

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
func (it *StaticOrganizationInsertConfigurationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StaticOrganizationInsertConfiguration)
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
		it.Event = new(StaticOrganizationInsertConfiguration)
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
func (it *StaticOrganizationInsertConfigurationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StaticOrganizationInsertConfigurationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StaticOrganizationInsertConfiguration represents a InsertConfiguration event raised by the StaticOrganization contract.
type StaticOrganizationInsertConfiguration struct {
	Org      common.Hash
	FileName common.Hash
	Cid      common.Hash
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterInsertConfiguration is a free log retrieval operation binding the contract event 0xe4c71a24d7d86a4b9dcdf7bb43d15433ed42b5239d98cdd095cc8433545911e0.
//
// Solidity: event InsertConfiguration(string indexed org, string indexed fileName, string indexed cid)
func (_StaticOrganization *StaticOrganizationFilterer) FilterInsertConfiguration(opts *bind.FilterOpts, org []string, fileName []string, cid []string) (*StaticOrganizationInsertConfigurationIterator, error) {

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

	logs, sub, err := _StaticOrganization.contract.FilterLogs(opts, "InsertConfiguration", orgRule, fileNameRule, cidRule)
	if err != nil {
		return nil, err
	}
	return &StaticOrganizationInsertConfigurationIterator{contract: _StaticOrganization.contract, event: "InsertConfiguration", logs: logs, sub: sub}, nil
}

// WatchInsertConfiguration is a free log subscription operation binding the contract event 0xe4c71a24d7d86a4b9dcdf7bb43d15433ed42b5239d98cdd095cc8433545911e0.
//
// Solidity: event InsertConfiguration(string indexed org, string indexed fileName, string indexed cid)
func (_StaticOrganization *StaticOrganizationFilterer) WatchInsertConfiguration(opts *bind.WatchOpts, sink chan<- *StaticOrganizationInsertConfiguration, org []string, fileName []string, cid []string) (event.Subscription, error) {

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

	logs, sub, err := _StaticOrganization.contract.WatchLogs(opts, "InsertConfiguration", orgRule, fileNameRule, cidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StaticOrganizationInsertConfiguration)
				if err := _StaticOrganization.contract.UnpackLog(event, "InsertConfiguration", log); err != nil {
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
func (_StaticOrganization *StaticOrganizationFilterer) ParseInsertConfiguration(log types.Log) (*StaticOrganizationInsertConfiguration, error) {
	event := new(StaticOrganizationInsertConfiguration)
	if err := _StaticOrganization.contract.UnpackLog(event, "InsertConfiguration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StaticOrganizationInsertOrganizationIterator is returned from FilterInsertOrganization and is used to iterate over the raw logs and unpacked data for InsertOrganization events raised by the StaticOrganization contract.
type StaticOrganizationInsertOrganizationIterator struct {
	Event *StaticOrganizationInsertOrganization // Event containing the contract specifics and raw log

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
func (it *StaticOrganizationInsertOrganizationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StaticOrganizationInsertOrganization)
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
		it.Event = new(StaticOrganizationInsertOrganization)
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
func (it *StaticOrganizationInsertOrganizationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StaticOrganizationInsertOrganizationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StaticOrganizationInsertOrganization represents a InsertOrganization event raised by the StaticOrganization contract.
type StaticOrganizationInsertOrganization struct {
	Org common.Hash
	Raw types.Log // Blockchain specific contextual infos
}

// FilterInsertOrganization is a free log retrieval operation binding the contract event 0xc88b7686edda01cb5cb9cfc327a234fd55553b72ce44f51772a9ffd0be4f517a.
//
// Solidity: event InsertOrganization(string indexed org)
func (_StaticOrganization *StaticOrganizationFilterer) FilterInsertOrganization(opts *bind.FilterOpts, org []string) (*StaticOrganizationInsertOrganizationIterator, error) {

	var orgRule []interface{}
	for _, orgItem := range org {
		orgRule = append(orgRule, orgItem)
	}

	logs, sub, err := _StaticOrganization.contract.FilterLogs(opts, "InsertOrganization", orgRule)
	if err != nil {
		return nil, err
	}
	return &StaticOrganizationInsertOrganizationIterator{contract: _StaticOrganization.contract, event: "InsertOrganization", logs: logs, sub: sub}, nil
}

// WatchInsertOrganization is a free log subscription operation binding the contract event 0xc88b7686edda01cb5cb9cfc327a234fd55553b72ce44f51772a9ffd0be4f517a.
//
// Solidity: event InsertOrganization(string indexed org)
func (_StaticOrganization *StaticOrganizationFilterer) WatchInsertOrganization(opts *bind.WatchOpts, sink chan<- *StaticOrganizationInsertOrganization, org []string) (event.Subscription, error) {

	var orgRule []interface{}
	for _, orgItem := range org {
		orgRule = append(orgRule, orgItem)
	}

	logs, sub, err := _StaticOrganization.contract.WatchLogs(opts, "InsertOrganization", orgRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StaticOrganizationInsertOrganization)
				if err := _StaticOrganization.contract.UnpackLog(event, "InsertOrganization", log); err != nil {
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
func (_StaticOrganization *StaticOrganizationFilterer) ParseInsertOrganization(log types.Log) (*StaticOrganizationInsertOrganization, error) {
	event := new(StaticOrganizationInsertOrganization)
	if err := _StaticOrganization.contract.UnpackLog(event, "InsertOrganization", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StaticOrganizationInsertSmartcontractIterator is returned from FilterInsertSmartcontract and is used to iterate over the raw logs and unpacked data for InsertSmartcontract events raised by the StaticOrganization contract.
type StaticOrganizationInsertSmartcontractIterator struct {
	Event *StaticOrganizationInsertSmartcontract // Event containing the contract specifics and raw log

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
func (it *StaticOrganizationInsertSmartcontractIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StaticOrganizationInsertSmartcontract)
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
		it.Event = new(StaticOrganizationInsertSmartcontract)
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
func (it *StaticOrganizationInsertSmartcontractIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StaticOrganizationInsertSmartcontractIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StaticOrganizationInsertSmartcontract represents a InsertSmartcontract event raised by the StaticOrganization contract.
type StaticOrganizationInsertSmartcontract struct {
	Org      common.Hash
	FileName common.Hash
	Cid      common.Hash
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterInsertSmartcontract is a free log retrieval operation binding the contract event 0x0e7ebd6fc52c2b4f8dfefdaaba6996ec0c5104c174da6e43d16a314a286fdd93.
//
// Solidity: event InsertSmartcontract(string indexed org, string indexed fileName, string indexed cid)
func (_StaticOrganization *StaticOrganizationFilterer) FilterInsertSmartcontract(opts *bind.FilterOpts, org []string, fileName []string, cid []string) (*StaticOrganizationInsertSmartcontractIterator, error) {

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

	logs, sub, err := _StaticOrganization.contract.FilterLogs(opts, "InsertSmartcontract", orgRule, fileNameRule, cidRule)
	if err != nil {
		return nil, err
	}
	return &StaticOrganizationInsertSmartcontractIterator{contract: _StaticOrganization.contract, event: "InsertSmartcontract", logs: logs, sub: sub}, nil
}

// WatchInsertSmartcontract is a free log subscription operation binding the contract event 0x0e7ebd6fc52c2b4f8dfefdaaba6996ec0c5104c174da6e43d16a314a286fdd93.
//
// Solidity: event InsertSmartcontract(string indexed org, string indexed fileName, string indexed cid)
func (_StaticOrganization *StaticOrganizationFilterer) WatchInsertSmartcontract(opts *bind.WatchOpts, sink chan<- *StaticOrganizationInsertSmartcontract, org []string, fileName []string, cid []string) (event.Subscription, error) {

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

	logs, sub, err := _StaticOrganization.contract.WatchLogs(opts, "InsertSmartcontract", orgRule, fileNameRule, cidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StaticOrganizationInsertSmartcontract)
				if err := _StaticOrganization.contract.UnpackLog(event, "InsertSmartcontract", log); err != nil {
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
func (_StaticOrganization *StaticOrganizationFilterer) ParseInsertSmartcontract(log types.Log) (*StaticOrganizationInsertSmartcontract, error) {
	event := new(StaticOrganizationInsertSmartcontract)
	if err := _StaticOrganization.contract.UnpackLog(event, "InsertSmartcontract", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StaticOrganizationUpdateConfigurationIterator is returned from FilterUpdateConfiguration and is used to iterate over the raw logs and unpacked data for UpdateConfiguration events raised by the StaticOrganization contract.
type StaticOrganizationUpdateConfigurationIterator struct {
	Event *StaticOrganizationUpdateConfiguration // Event containing the contract specifics and raw log

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
func (it *StaticOrganizationUpdateConfigurationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StaticOrganizationUpdateConfiguration)
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
		it.Event = new(StaticOrganizationUpdateConfiguration)
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
func (it *StaticOrganizationUpdateConfigurationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StaticOrganizationUpdateConfigurationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StaticOrganizationUpdateConfiguration represents a UpdateConfiguration event raised by the StaticOrganization contract.
type StaticOrganizationUpdateConfiguration struct {
	FileName common.Hash
	Cid      common.Hash
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUpdateConfiguration is a free log retrieval operation binding the contract event 0xa183a608b7c8ba72d4b45247ef3b431ceb231e76761f3709c6679c7d1fbf9749.
//
// Solidity: event UpdateConfiguration(string indexed fileName, string indexed cid)
func (_StaticOrganization *StaticOrganizationFilterer) FilterUpdateConfiguration(opts *bind.FilterOpts, fileName []string, cid []string) (*StaticOrganizationUpdateConfigurationIterator, error) {

	var fileNameRule []interface{}
	for _, fileNameItem := range fileName {
		fileNameRule = append(fileNameRule, fileNameItem)
	}
	var cidRule []interface{}
	for _, cidItem := range cid {
		cidRule = append(cidRule, cidItem)
	}

	logs, sub, err := _StaticOrganization.contract.FilterLogs(opts, "UpdateConfiguration", fileNameRule, cidRule)
	if err != nil {
		return nil, err
	}
	return &StaticOrganizationUpdateConfigurationIterator{contract: _StaticOrganization.contract, event: "UpdateConfiguration", logs: logs, sub: sub}, nil
}

// WatchUpdateConfiguration is a free log subscription operation binding the contract event 0xa183a608b7c8ba72d4b45247ef3b431ceb231e76761f3709c6679c7d1fbf9749.
//
// Solidity: event UpdateConfiguration(string indexed fileName, string indexed cid)
func (_StaticOrganization *StaticOrganizationFilterer) WatchUpdateConfiguration(opts *bind.WatchOpts, sink chan<- *StaticOrganizationUpdateConfiguration, fileName []string, cid []string) (event.Subscription, error) {

	var fileNameRule []interface{}
	for _, fileNameItem := range fileName {
		fileNameRule = append(fileNameRule, fileNameItem)
	}
	var cidRule []interface{}
	for _, cidItem := range cid {
		cidRule = append(cidRule, cidItem)
	}

	logs, sub, err := _StaticOrganization.contract.WatchLogs(opts, "UpdateConfiguration", fileNameRule, cidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StaticOrganizationUpdateConfiguration)
				if err := _StaticOrganization.contract.UnpackLog(event, "UpdateConfiguration", log); err != nil {
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
func (_StaticOrganization *StaticOrganizationFilterer) ParseUpdateConfiguration(log types.Log) (*StaticOrganizationUpdateConfiguration, error) {
	event := new(StaticOrganizationUpdateConfiguration)
	if err := _StaticOrganization.contract.UnpackLog(event, "UpdateConfiguration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StaticOrganizationUpdateSmartcontractIterator is returned from FilterUpdateSmartcontract and is used to iterate over the raw logs and unpacked data for UpdateSmartcontract events raised by the StaticOrganization contract.
type StaticOrganizationUpdateSmartcontractIterator struct {
	Event *StaticOrganizationUpdateSmartcontract // Event containing the contract specifics and raw log

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
func (it *StaticOrganizationUpdateSmartcontractIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StaticOrganizationUpdateSmartcontract)
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
		it.Event = new(StaticOrganizationUpdateSmartcontract)
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
func (it *StaticOrganizationUpdateSmartcontractIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StaticOrganizationUpdateSmartcontractIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StaticOrganizationUpdateSmartcontract represents a UpdateSmartcontract event raised by the StaticOrganization contract.
type StaticOrganizationUpdateSmartcontract struct {
	FileName common.Hash
	Cid      common.Hash
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUpdateSmartcontract is a free log retrieval operation binding the contract event 0xcdf23f587b224317852b246c67b694bc63fc30d22be0c20193cbe5bd69d9b7c2.
//
// Solidity: event UpdateSmartcontract(string indexed fileName, string indexed cid)
func (_StaticOrganization *StaticOrganizationFilterer) FilterUpdateSmartcontract(opts *bind.FilterOpts, fileName []string, cid []string) (*StaticOrganizationUpdateSmartcontractIterator, error) {

	var fileNameRule []interface{}
	for _, fileNameItem := range fileName {
		fileNameRule = append(fileNameRule, fileNameItem)
	}
	var cidRule []interface{}
	for _, cidItem := range cid {
		cidRule = append(cidRule, cidItem)
	}

	logs, sub, err := _StaticOrganization.contract.FilterLogs(opts, "UpdateSmartcontract", fileNameRule, cidRule)
	if err != nil {
		return nil, err
	}
	return &StaticOrganizationUpdateSmartcontractIterator{contract: _StaticOrganization.contract, event: "UpdateSmartcontract", logs: logs, sub: sub}, nil
}

// WatchUpdateSmartcontract is a free log subscription operation binding the contract event 0xcdf23f587b224317852b246c67b694bc63fc30d22be0c20193cbe5bd69d9b7c2.
//
// Solidity: event UpdateSmartcontract(string indexed fileName, string indexed cid)
func (_StaticOrganization *StaticOrganizationFilterer) WatchUpdateSmartcontract(opts *bind.WatchOpts, sink chan<- *StaticOrganizationUpdateSmartcontract, fileName []string, cid []string) (event.Subscription, error) {

	var fileNameRule []interface{}
	for _, fileNameItem := range fileName {
		fileNameRule = append(fileNameRule, fileNameItem)
	}
	var cidRule []interface{}
	for _, cidItem := range cid {
		cidRule = append(cidRule, cidItem)
	}

	logs, sub, err := _StaticOrganization.contract.WatchLogs(opts, "UpdateSmartcontract", fileNameRule, cidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StaticOrganizationUpdateSmartcontract)
				if err := _StaticOrganization.contract.UnpackLog(event, "UpdateSmartcontract", log); err != nil {
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
func (_StaticOrganization *StaticOrganizationFilterer) ParseUpdateSmartcontract(log types.Log) (*StaticOrganizationUpdateSmartcontract, error) {
	event := new(StaticOrganizationUpdateSmartcontract)
	if err := _StaticOrganization.contract.UnpackLog(event, "UpdateSmartcontract", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

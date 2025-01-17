// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package alphax

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

// AlphaXCheckInInfo is an auto generated low-level Go binding around an user-defined struct.
type AlphaXCheckInInfo struct {
	TaskId    uint32
	UserId    uint64
	Timestamp *big.Int
}

// AlphaXPredictionInfo is an auto generated low-level Go binding around an user-defined struct.
type AlphaXPredictionInfo struct {
	SignalId    uint32
	UserId      uint64
	Choice      uint8
	HasInvolved bool
}

// AlphaxMetaData contains all meta data concerning the Alphax contract.
var AlphaxMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"id\",\"type\":\"uint32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"taskId\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"userId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structAlphaX.CheckInInfo\",\"name\":\"info\",\"type\":\"tuple\"}],\"name\":\"CheckinEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"signalId\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"userId\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"choice\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"hasInvolved\",\"type\":\"bool\"}],\"indexed\":false,\"internalType\":\"structAlphaX.PredictionInfo\",\"name\":\"info\",\"type\":\"tuple\"}],\"name\":\"SignalPredictionEvent\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"day\",\"type\":\"uint256\"}],\"name\":\"checkInResult\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"taskId\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"userId\",\"type\":\"uint64\"}],\"name\":\"checkin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentDay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"signalId\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"userId\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"choice\",\"type\":\"uint8\"}],\"name\":\"signalPredict\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"signalId\",\"type\":\"uint32\"}],\"name\":\"signalPredictionResult\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// AlphaxABI is the input ABI used to generate the binding from.
// Deprecated: Use AlphaxMetaData.ABI instead.
var AlphaxABI = AlphaxMetaData.ABI

// Alphax is an auto generated Go binding around an Ethereum contract.
type Alphax struct {
	AlphaxCaller     // Read-only binding to the contract
	AlphaxTransactor // Write-only binding to the contract
	AlphaxFilterer   // Log filterer for contract events
}

// AlphaxCaller is an auto generated read-only Go binding around an Ethereum contract.
type AlphaxCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AlphaxTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AlphaxTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AlphaxFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AlphaxFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AlphaxSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AlphaxSession struct {
	Contract     *Alphax           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AlphaxCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AlphaxCallerSession struct {
	Contract *AlphaxCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// AlphaxTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AlphaxTransactorSession struct {
	Contract     *AlphaxTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AlphaxRaw is an auto generated low-level Go binding around an Ethereum contract.
type AlphaxRaw struct {
	Contract *Alphax // Generic contract binding to access the raw methods on
}

// AlphaxCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AlphaxCallerRaw struct {
	Contract *AlphaxCaller // Generic read-only contract binding to access the raw methods on
}

// AlphaxTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AlphaxTransactorRaw struct {
	Contract *AlphaxTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAlphax creates a new instance of Alphax, bound to a specific deployed contract.
func NewAlphax(address common.Address, backend bind.ContractBackend) (*Alphax, error) {
	contract, err := bindAlphax(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Alphax{AlphaxCaller: AlphaxCaller{contract: contract}, AlphaxTransactor: AlphaxTransactor{contract: contract}, AlphaxFilterer: AlphaxFilterer{contract: contract}}, nil
}

// NewAlphaxCaller creates a new read-only instance of Alphax, bound to a specific deployed contract.
func NewAlphaxCaller(address common.Address, caller bind.ContractCaller) (*AlphaxCaller, error) {
	contract, err := bindAlphax(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AlphaxCaller{contract: contract}, nil
}

// NewAlphaxTransactor creates a new write-only instance of Alphax, bound to a specific deployed contract.
func NewAlphaxTransactor(address common.Address, transactor bind.ContractTransactor) (*AlphaxTransactor, error) {
	contract, err := bindAlphax(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AlphaxTransactor{contract: contract}, nil
}

// NewAlphaxFilterer creates a new log filterer instance of Alphax, bound to a specific deployed contract.
func NewAlphaxFilterer(address common.Address, filterer bind.ContractFilterer) (*AlphaxFilterer, error) {
	contract, err := bindAlphax(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AlphaxFilterer{contract: contract}, nil
}

// bindAlphax binds a generic wrapper to an already deployed contract.
func bindAlphax(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AlphaxMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Alphax *AlphaxRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Alphax.Contract.AlphaxCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Alphax *AlphaxRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Alphax.Contract.AlphaxTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Alphax *AlphaxRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Alphax.Contract.AlphaxTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Alphax *AlphaxCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Alphax.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Alphax *AlphaxTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Alphax.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Alphax *AlphaxTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Alphax.Contract.contract.Transact(opts, method, params...)
}

// CheckInResult is a free data retrieval call binding the contract method 0x2c3ae4a7.
//
// Solidity: function checkInResult(address user, uint256 day) view returns(bool)
func (_Alphax *AlphaxCaller) CheckInResult(opts *bind.CallOpts, user common.Address, day *big.Int) (bool, error) {
	var out []interface{}
	err := _Alphax.contract.Call(opts, &out, "checkInResult", user, day)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckInResult is a free data retrieval call binding the contract method 0x2c3ae4a7.
//
// Solidity: function checkInResult(address user, uint256 day) view returns(bool)
func (_Alphax *AlphaxSession) CheckInResult(user common.Address, day *big.Int) (bool, error) {
	return _Alphax.Contract.CheckInResult(&_Alphax.CallOpts, user, day)
}

// CheckInResult is a free data retrieval call binding the contract method 0x2c3ae4a7.
//
// Solidity: function checkInResult(address user, uint256 day) view returns(bool)
func (_Alphax *AlphaxCallerSession) CheckInResult(user common.Address, day *big.Int) (bool, error) {
	return _Alphax.Contract.CheckInResult(&_Alphax.CallOpts, user, day)
}

// GetCurrentDay is a free data retrieval call binding the contract method 0x3e6968b6.
//
// Solidity: function getCurrentDay() view returns(uint256)
func (_Alphax *AlphaxCaller) GetCurrentDay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Alphax.contract.Call(opts, &out, "getCurrentDay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentDay is a free data retrieval call binding the contract method 0x3e6968b6.
//
// Solidity: function getCurrentDay() view returns(uint256)
func (_Alphax *AlphaxSession) GetCurrentDay() (*big.Int, error) {
	return _Alphax.Contract.GetCurrentDay(&_Alphax.CallOpts)
}

// GetCurrentDay is a free data retrieval call binding the contract method 0x3e6968b6.
//
// Solidity: function getCurrentDay() view returns(uint256)
func (_Alphax *AlphaxCallerSession) GetCurrentDay() (*big.Int, error) {
	return _Alphax.Contract.GetCurrentDay(&_Alphax.CallOpts)
}

// SignalPredictionResult is a free data retrieval call binding the contract method 0x703a7fa2.
//
// Solidity: function signalPredictionResult(address user, uint32 signalId) view returns(bool, uint32, uint8)
func (_Alphax *AlphaxCaller) SignalPredictionResult(opts *bind.CallOpts, user common.Address, signalId uint32) (bool, uint32, uint8, error) {
	var out []interface{}
	err := _Alphax.contract.Call(opts, &out, "signalPredictionResult", user, signalId)

	if err != nil {
		return *new(bool), *new(uint32), *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new(uint32)).(*uint32)
	out2 := *abi.ConvertType(out[2], new(uint8)).(*uint8)

	return out0, out1, out2, err

}

// SignalPredictionResult is a free data retrieval call binding the contract method 0x703a7fa2.
//
// Solidity: function signalPredictionResult(address user, uint32 signalId) view returns(bool, uint32, uint8)
func (_Alphax *AlphaxSession) SignalPredictionResult(user common.Address, signalId uint32) (bool, uint32, uint8, error) {
	return _Alphax.Contract.SignalPredictionResult(&_Alphax.CallOpts, user, signalId)
}

// SignalPredictionResult is a free data retrieval call binding the contract method 0x703a7fa2.
//
// Solidity: function signalPredictionResult(address user, uint32 signalId) view returns(bool, uint32, uint8)
func (_Alphax *AlphaxCallerSession) SignalPredictionResult(user common.Address, signalId uint32) (bool, uint32, uint8, error) {
	return _Alphax.Contract.SignalPredictionResult(&_Alphax.CallOpts, user, signalId)
}

// Checkin is a paid mutator transaction binding the contract method 0x0655e4ca.
//
// Solidity: function checkin(address user, uint32 taskId, uint64 userId) returns()
func (_Alphax *AlphaxTransactor) Checkin(opts *bind.TransactOpts, user common.Address, taskId uint32, userId uint64) (*types.Transaction, error) {
	return _Alphax.contract.Transact(opts, "checkin", user, taskId, userId)
}

// Checkin is a paid mutator transaction binding the contract method 0x0655e4ca.
//
// Solidity: function checkin(address user, uint32 taskId, uint64 userId) returns()
func (_Alphax *AlphaxSession) Checkin(user common.Address, taskId uint32, userId uint64) (*types.Transaction, error) {
	return _Alphax.Contract.Checkin(&_Alphax.TransactOpts, user, taskId, userId)
}

// Checkin is a paid mutator transaction binding the contract method 0x0655e4ca.
//
// Solidity: function checkin(address user, uint32 taskId, uint64 userId) returns()
func (_Alphax *AlphaxTransactorSession) Checkin(user common.Address, taskId uint32, userId uint64) (*types.Transaction, error) {
	return _Alphax.Contract.Checkin(&_Alphax.TransactOpts, user, taskId, userId)
}

// SignalPredict is a paid mutator transaction binding the contract method 0x591abd76.
//
// Solidity: function signalPredict(address user, uint32 signalId, uint64 userId, uint8 choice) returns()
func (_Alphax *AlphaxTransactor) SignalPredict(opts *bind.TransactOpts, user common.Address, signalId uint32, userId uint64, choice uint8) (*types.Transaction, error) {
	return _Alphax.contract.Transact(opts, "signalPredict", user, signalId, userId, choice)
}

// SignalPredict is a paid mutator transaction binding the contract method 0x591abd76.
//
// Solidity: function signalPredict(address user, uint32 signalId, uint64 userId, uint8 choice) returns()
func (_Alphax *AlphaxSession) SignalPredict(user common.Address, signalId uint32, userId uint64, choice uint8) (*types.Transaction, error) {
	return _Alphax.Contract.SignalPredict(&_Alphax.TransactOpts, user, signalId, userId, choice)
}

// SignalPredict is a paid mutator transaction binding the contract method 0x591abd76.
//
// Solidity: function signalPredict(address user, uint32 signalId, uint64 userId, uint8 choice) returns()
func (_Alphax *AlphaxTransactorSession) SignalPredict(user common.Address, signalId uint32, userId uint64, choice uint8) (*types.Transaction, error) {
	return _Alphax.Contract.SignalPredict(&_Alphax.TransactOpts, user, signalId, userId, choice)
}

// AlphaxCheckinEventIterator is returned from FilterCheckinEvent and is used to iterate over the raw logs and unpacked data for CheckinEvent events raised by the Alphax contract.
type AlphaxCheckinEventIterator struct {
	Event *AlphaxCheckinEvent // Event containing the contract specifics and raw log

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
func (it *AlphaxCheckinEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AlphaxCheckinEvent)
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
		it.Event = new(AlphaxCheckinEvent)
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
func (it *AlphaxCheckinEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AlphaxCheckinEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AlphaxCheckinEvent represents a CheckinEvent event raised by the Alphax contract.
type AlphaxCheckinEvent struct {
	User common.Address
	Info AlphaXCheckInInfo
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterCheckinEvent is a free log retrieval operation binding the contract event 0x2f384bbc453afcdcb37abef73843b9341c84045a682ee5e170a885d16b660dae.
//
// Solidity: event CheckinEvent(address indexed user, (uint32,uint64,uint256) info)
func (_Alphax *AlphaxFilterer) FilterCheckinEvent(opts *bind.FilterOpts, user []common.Address) (*AlphaxCheckinEventIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Alphax.contract.FilterLogs(opts, "CheckinEvent", userRule)
	if err != nil {
		return nil, err
	}
	return &AlphaxCheckinEventIterator{contract: _Alphax.contract, event: "CheckinEvent", logs: logs, sub: sub}, nil
}

// WatchCheckinEvent is a free log subscription operation binding the contract event 0x2f384bbc453afcdcb37abef73843b9341c84045a682ee5e170a885d16b660dae.
//
// Solidity: event CheckinEvent(address indexed user, (uint32,uint64,uint256) info)
func (_Alphax *AlphaxFilterer) WatchCheckinEvent(opts *bind.WatchOpts, sink chan<- *AlphaxCheckinEvent, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Alphax.contract.WatchLogs(opts, "CheckinEvent", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AlphaxCheckinEvent)
				if err := _Alphax.contract.UnpackLog(event, "CheckinEvent", log); err != nil {
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

// ParseCheckinEvent is a log parse operation binding the contract event 0x2f384bbc453afcdcb37abef73843b9341c84045a682ee5e170a885d16b660dae.
//
// Solidity: event CheckinEvent(address indexed user, (uint32,uint64,uint256) info)
func (_Alphax *AlphaxFilterer) ParseCheckinEvent(log types.Log) (*AlphaxCheckinEvent, error) {
	event := new(AlphaxCheckinEvent)
	if err := _Alphax.contract.UnpackLog(event, "CheckinEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AlphaxSignalPredictionEventIterator is returned from FilterSignalPredictionEvent and is used to iterate over the raw logs and unpacked data for SignalPredictionEvent events raised by the Alphax contract.
type AlphaxSignalPredictionEventIterator struct {
	Event *AlphaxSignalPredictionEvent // Event containing the contract specifics and raw log

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
func (it *AlphaxSignalPredictionEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AlphaxSignalPredictionEvent)
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
		it.Event = new(AlphaxSignalPredictionEvent)
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
func (it *AlphaxSignalPredictionEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AlphaxSignalPredictionEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AlphaxSignalPredictionEvent represents a SignalPredictionEvent event raised by the Alphax contract.
type AlphaxSignalPredictionEvent struct {
	User common.Address
	Info AlphaXPredictionInfo
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterSignalPredictionEvent is a free log retrieval operation binding the contract event 0xbe0b3fe34eb803bd35bd0e6883956b86e4bc611efcd52b146df987bad9e18ba1.
//
// Solidity: event SignalPredictionEvent(address indexed user, (uint32,uint64,uint8,bool) info)
func (_Alphax *AlphaxFilterer) FilterSignalPredictionEvent(opts *bind.FilterOpts, user []common.Address) (*AlphaxSignalPredictionEventIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Alphax.contract.FilterLogs(opts, "SignalPredictionEvent", userRule)
	if err != nil {
		return nil, err
	}
	return &AlphaxSignalPredictionEventIterator{contract: _Alphax.contract, event: "SignalPredictionEvent", logs: logs, sub: sub}, nil
}

// WatchSignalPredictionEvent is a free log subscription operation binding the contract event 0xbe0b3fe34eb803bd35bd0e6883956b86e4bc611efcd52b146df987bad9e18ba1.
//
// Solidity: event SignalPredictionEvent(address indexed user, (uint32,uint64,uint8,bool) info)
func (_Alphax *AlphaxFilterer) WatchSignalPredictionEvent(opts *bind.WatchOpts, sink chan<- *AlphaxSignalPredictionEvent, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Alphax.contract.WatchLogs(opts, "SignalPredictionEvent", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AlphaxSignalPredictionEvent)
				if err := _Alphax.contract.UnpackLog(event, "SignalPredictionEvent", log); err != nil {
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

// ParseSignalPredictionEvent is a log parse operation binding the contract event 0xbe0b3fe34eb803bd35bd0e6883956b86e4bc611efcd52b146df987bad9e18ba1.
//
// Solidity: event SignalPredictionEvent(address indexed user, (uint32,uint64,uint8,bool) info)
func (_Alphax *AlphaxFilterer) ParseSignalPredictionEvent(log types.Log) (*AlphaxSignalPredictionEvent, error) {
	event := new(AlphaxSignalPredictionEvent)
	if err := _Alphax.contract.UnpackLog(event, "SignalPredictionEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

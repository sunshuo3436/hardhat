// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package unindexed

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

// UnindexedMetaData contains all meta data concerning the Unindexed contract.
var UnindexedMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// UnindexedABI is the input ABI used to generate the binding from.
// Deprecated: Use UnindexedMetaData.ABI instead.
var UnindexedABI = UnindexedMetaData.ABI

// Unindexed is an auto generated Go binding around an Ethereum contract.
type Unindexed struct {
	UnindexedCaller     // Read-only binding to the contract
	UnindexedTransactor // Write-only binding to the contract
	UnindexedFilterer   // Log filterer for contract events
}

// UnindexedCaller is an auto generated read-only Go binding around an Ethereum contract.
type UnindexedCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UnindexedTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UnindexedTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UnindexedFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UnindexedFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UnindexedSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UnindexedSession struct {
	Contract     *Unindexed        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UnindexedCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UnindexedCallerSession struct {
	Contract *UnindexedCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// UnindexedTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UnindexedTransactorSession struct {
	Contract     *UnindexedTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// UnindexedRaw is an auto generated low-level Go binding around an Ethereum contract.
type UnindexedRaw struct {
	Contract *Unindexed // Generic contract binding to access the raw methods on
}

// UnindexedCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UnindexedCallerRaw struct {
	Contract *UnindexedCaller // Generic read-only contract binding to access the raw methods on
}

// UnindexedTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UnindexedTransactorRaw struct {
	Contract *UnindexedTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUnindexed creates a new instance of Unindexed, bound to a specific deployed contract.
func NewUnindexed(address common.Address, backend bind.ContractBackend) (*Unindexed, error) {
	contract, err := bindUnindexed(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Unindexed{UnindexedCaller: UnindexedCaller{contract: contract}, UnindexedTransactor: UnindexedTransactor{contract: contract}, UnindexedFilterer: UnindexedFilterer{contract: contract}}, nil
}

// NewUnindexedCaller creates a new read-only instance of Unindexed, bound to a specific deployed contract.
func NewUnindexedCaller(address common.Address, caller bind.ContractCaller) (*UnindexedCaller, error) {
	contract, err := bindUnindexed(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UnindexedCaller{contract: contract}, nil
}

// NewUnindexedTransactor creates a new write-only instance of Unindexed, bound to a specific deployed contract.
func NewUnindexedTransactor(address common.Address, transactor bind.ContractTransactor) (*UnindexedTransactor, error) {
	contract, err := bindUnindexed(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UnindexedTransactor{contract: contract}, nil
}

// NewUnindexedFilterer creates a new log filterer instance of Unindexed, bound to a specific deployed contract.
func NewUnindexedFilterer(address common.Address, filterer bind.ContractFilterer) (*UnindexedFilterer, error) {
	contract, err := bindUnindexed(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UnindexedFilterer{contract: contract}, nil
}

// bindUnindexed binds a generic wrapper to an already deployed contract.
func bindUnindexed(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := UnindexedMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Unindexed *UnindexedRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Unindexed.Contract.UnindexedCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Unindexed *UnindexedRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Unindexed.Contract.UnindexedTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Unindexed *UnindexedRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Unindexed.Contract.UnindexedTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Unindexed *UnindexedCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Unindexed.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Unindexed *UnindexedTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Unindexed.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Unindexed *UnindexedTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Unindexed.Contract.contract.Transact(opts, method, params...)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Unindexed *UnindexedTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Unindexed.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Unindexed *UnindexedSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Unindexed.Contract.Transfer(&_Unindexed.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Unindexed *UnindexedTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Unindexed.Contract.Transfer(&_Unindexed.TransactOpts, to, value)
}

// UnindexedTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Unindexed contract.
type UnindexedTransferIterator struct {
	Event *UnindexedTransfer // Event containing the contract specifics and raw log

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
func (it *UnindexedTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UnindexedTransfer)
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
		it.Event = new(UnindexedTransfer)
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
func (it *UnindexedTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UnindexedTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UnindexedTransfer represents a Transfer event raised by the Unindexed contract.
type UnindexedTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Unindexed *UnindexedFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*UnindexedTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Unindexed.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &UnindexedTransferIterator{contract: _Unindexed.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Unindexed *UnindexedFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *UnindexedTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Unindexed.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UnindexedTransfer)
				if err := _Unindexed.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Unindexed *UnindexedFilterer) ParseTransfer(log types.Log) (*UnindexedTransfer, error) {
	event := new(UnindexedTransfer)
	if err := _Unindexed.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package InsLogEvent

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// InsLogEventABI is the input ABI used to generate the binding from.
const InsLogEventABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"msg\",\"type\":\"string\"}],\"name\":\"AnEvent\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_acct\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_msg\",\"type\":\"string\"}],\"name\":\"IndexedEvent\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"changeOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// InsLogEvent is an auto generated Go binding around an Ethereum contract.
type InsLogEvent struct {
	InsLogEventCaller     // Read-only binding to the contract
	InsLogEventTransactor // Write-only binding to the contract
	InsLogEventFilterer   // Log filterer for contract events
}

// InsLogEventCaller is an auto generated read-only Go binding around an Ethereum contract.
type InsLogEventCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InsLogEventTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InsLogEventTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InsLogEventFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InsLogEventFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InsLogEventSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InsLogEventSession struct {
	Contract     *InsLogEvent      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// InsLogEventCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InsLogEventCallerSession struct {
	Contract *InsLogEventCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// InsLogEventTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InsLogEventTransactorSession struct {
	Contract     *InsLogEventTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// InsLogEventRaw is an auto generated low-level Go binding around an Ethereum contract.
type InsLogEventRaw struct {
	Contract *InsLogEvent // Generic contract binding to access the raw methods on
}

// InsLogEventCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InsLogEventCallerRaw struct {
	Contract *InsLogEventCaller // Generic read-only contract binding to access the raw methods on
}

// InsLogEventTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InsLogEventTransactorRaw struct {
	Contract *InsLogEventTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInsLogEvent creates a new instance of InsLogEvent, bound to a specific deployed contract.
func NewInsLogEvent(address common.Address, backend bind.ContractBackend) (*InsLogEvent, error) {
	contract, err := bindInsLogEvent(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &InsLogEvent{InsLogEventCaller: InsLogEventCaller{contract: contract}, InsLogEventTransactor: InsLogEventTransactor{contract: contract}, InsLogEventFilterer: InsLogEventFilterer{contract: contract}}, nil
}

// NewInsLogEventCaller creates a new read-only instance of InsLogEvent, bound to a specific deployed contract.
func NewInsLogEventCaller(address common.Address, caller bind.ContractCaller) (*InsLogEventCaller, error) {
	contract, err := bindInsLogEvent(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InsLogEventCaller{contract: contract}, nil
}

// NewInsLogEventTransactor creates a new write-only instance of InsLogEvent, bound to a specific deployed contract.
func NewInsLogEventTransactor(address common.Address, transactor bind.ContractTransactor) (*InsLogEventTransactor, error) {
	contract, err := bindInsLogEvent(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InsLogEventTransactor{contract: contract}, nil
}

// NewInsLogEventFilterer creates a new log filterer instance of InsLogEvent, bound to a specific deployed contract.
func NewInsLogEventFilterer(address common.Address, filterer bind.ContractFilterer) (*InsLogEventFilterer, error) {
	contract, err := bindInsLogEvent(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InsLogEventFilterer{contract: contract}, nil
}

// bindInsLogEvent binds a generic wrapper to an already deployed contract.
func bindInsLogEvent(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(InsLogEventABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InsLogEvent *InsLogEventRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _InsLogEvent.Contract.InsLogEventCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InsLogEvent *InsLogEventRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InsLogEvent.Contract.InsLogEventTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InsLogEvent *InsLogEventRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InsLogEvent.Contract.InsLogEventTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InsLogEvent *InsLogEventCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _InsLogEvent.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InsLogEvent *InsLogEventTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InsLogEvent.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InsLogEvent *InsLogEventTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InsLogEvent.Contract.contract.Transact(opts, method, params...)
}

// IndexedEvent is a paid mutator transaction binding the contract method 0xbd20b5a9.
//
// Solidity: function IndexedEvent(address _acct, string _msg) returns(bool)
func (_InsLogEvent *InsLogEventTransactor) IndexedEvent(opts *bind.TransactOpts, _acct common.Address, _msg string) (*types.Transaction, error) {
	return _InsLogEvent.contract.Transact(opts, "IndexedEvent", _acct, _msg)
}

// IndexedEvent is a paid mutator transaction binding the contract method 0xbd20b5a9.
//
// Solidity: function IndexedEvent(address _acct, string _msg) returns(bool)
func (_InsLogEvent *InsLogEventSession) IndexedEvent(_acct common.Address, _msg string) (*types.Transaction, error) {
	return _InsLogEvent.Contract.IndexedEvent(&_InsLogEvent.TransactOpts, _acct, _msg)
}

// IndexedEvent is a paid mutator transaction binding the contract method 0xbd20b5a9.
//
// Solidity: function IndexedEvent(address _acct, string _msg) returns(bool)
func (_InsLogEvent *InsLogEventTransactorSession) IndexedEvent(_acct common.Address, _msg string) (*types.Transaction, error) {
	return _InsLogEvent.Contract.IndexedEvent(&_InsLogEvent.TransactOpts, _acct, _msg)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address _newOwner) returns()
func (_InsLogEvent *InsLogEventTransactor) ChangeOwner(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _InsLogEvent.contract.Transact(opts, "changeOwner", _newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address _newOwner) returns()
func (_InsLogEvent *InsLogEventSession) ChangeOwner(_newOwner common.Address) (*types.Transaction, error) {
	return _InsLogEvent.Contract.ChangeOwner(&_InsLogEvent.TransactOpts, _newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address _newOwner) returns()
func (_InsLogEvent *InsLogEventTransactorSession) ChangeOwner(_newOwner common.Address) (*types.Transaction, error) {
	return _InsLogEvent.Contract.ChangeOwner(&_InsLogEvent.TransactOpts, _newOwner)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_InsLogEvent *InsLogEventTransactor) Kill(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InsLogEvent.contract.Transact(opts, "kill")
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_InsLogEvent *InsLogEventSession) Kill() (*types.Transaction, error) {
	return _InsLogEvent.Contract.Kill(&_InsLogEvent.TransactOpts)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_InsLogEvent *InsLogEventTransactorSession) Kill() (*types.Transaction, error) {
	return _InsLogEvent.Contract.Kill(&_InsLogEvent.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns()
func (_InsLogEvent *InsLogEventTransactor) Withdraw(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _InsLogEvent.contract.Transact(opts, "withdraw", _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns()
func (_InsLogEvent *InsLogEventSession) Withdraw(_amount *big.Int) (*types.Transaction, error) {
	return _InsLogEvent.Contract.Withdraw(&_InsLogEvent.TransactOpts, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns()
func (_InsLogEvent *InsLogEventTransactorSession) Withdraw(_amount *big.Int) (*types.Transaction, error) {
	return _InsLogEvent.Contract.Withdraw(&_InsLogEvent.TransactOpts, _amount)
}

// InsLogEventAnEventIterator is returned from FilterAnEvent and is used to iterate over the raw logs and unpacked data for AnEvent events raised by the InsLogEvent contract.
type InsLogEventAnEventIterator struct {
	Event *InsLogEventAnEvent // Event containing the contract specifics and raw log

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
func (it *InsLogEventAnEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InsLogEventAnEvent)
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
		it.Event = new(InsLogEventAnEvent)
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
func (it *InsLogEventAnEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InsLogEventAnEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InsLogEventAnEvent represents a AnEvent event raised by the InsLogEvent contract.
type InsLogEventAnEvent struct {
	Account common.Address
	Msg     string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAnEvent is a free log retrieval operation binding the contract event 0xb3b0fc93288381e6fc61b2c0bcb321880c648c3ff65e8d8e313b881415b712a7.
//
// Solidity: event AnEvent(address indexed account, string msg)
func (_InsLogEvent *InsLogEventFilterer) FilterAnEvent(opts *bind.FilterOpts, account []common.Address) (*InsLogEventAnEventIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _InsLogEvent.contract.FilterLogs(opts, "AnEvent", accountRule)
	if err != nil {
		return nil, err
	}
	return &InsLogEventAnEventIterator{contract: _InsLogEvent.contract, event: "AnEvent", logs: logs, sub: sub}, nil
}

// WatchAnEvent is a free log subscription operation binding the contract event 0xb3b0fc93288381e6fc61b2c0bcb321880c648c3ff65e8d8e313b881415b712a7.
//
// Solidity: event AnEvent(address indexed account, string msg)
func (_InsLogEvent *InsLogEventFilterer) WatchAnEvent(opts *bind.WatchOpts, sink chan<- *InsLogEventAnEvent, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _InsLogEvent.contract.WatchLogs(opts, "AnEvent", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InsLogEventAnEvent)
				if err := _InsLogEvent.contract.UnpackLog(event, "AnEvent", log); err != nil {
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

// ParseAnEvent is a log parse operation binding the contract event 0xb3b0fc93288381e6fc61b2c0bcb321880c648c3ff65e8d8e313b881415b712a7.
//
// Solidity: event AnEvent(address indexed account, string msg)
func (_InsLogEvent *InsLogEventFilterer) ParseAnEvent(log types.Log) (*InsLogEventAnEvent, error) {
	event := new(InsLogEventAnEvent)
	if err := _InsLogEvent.contract.UnpackLog(event, "AnEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

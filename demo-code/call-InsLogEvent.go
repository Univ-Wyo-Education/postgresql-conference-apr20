package main

import (
	"fmt"
	"io/ioutil"

	"github.com/Univ-Wyo-Education/S21-4010/Eth/lib/InsLogEvent"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// InsLogEventControl connection information for interface to InsLogEvent contract.
type InsLogEventControl struct {
	caller          *InsLogEvent.InsLogEventCaller
	callerOpts      *bind.CallOpts
	transactor      *InsLogEvent.InsLogEventTransactor
	transactorOpts  *bind.TransactOpts
	contract        *InsLogEvent.InsLogEvent
	contractAddress common.Address
}

// Important note on watching for Ethereum contract events in Go.
//
// In calls to an abigen generated watch, the first parameter is a filter.
// In example code, to see all events, one must pass in an empty filter.
// In geth, this doesn't work. An empty filter will result in an incorrect
// bloom filter to be selected for the Ethereum search code.
// Rather, to watch for events requires a 'nil' as the first parameter.
//
// For example:
//  	filter := nil
//  	eventSubscription, err := kg.contract.SomeContractSomeEvent(
//			filter,
//			eventChan,
//		)
//
// Will exhibit our desired behavior of selecting an empty filter.
//
// This is different from node.js/web3 code where a 'nil' is treated the same
// as an empty filter.

// InsLogEvent creates the necessary connections and configurations
// for accessing the InsLogEvent contract.
func NewInsLogEvent(gCfg *ConfigType) (*InsLogEventControl, error) {
	contractAddressHex, exists := gCfg.ContractAddr["InsLogEvent"]
	if !exists {
		return nil, fmt.Errorf("missing address for 'InsLogEvent' in ./cfg.json")
	}
	contractAddress := common.HexToAddress(contractAddressHex)

	groupTransactor, err := InsLogEvent.NewInsLogEventTransactor(contractAddress, gCfg.Client)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to instantiate a KeepRelayBeaconTranactor contract: [%v]",
			err,
		)
	}

	if gCfg.AccountKey == nil {
		key, err := DecryptKeyFile(gCfg.Account.KeyFile, gCfg.Account.KeyFilePassword)
		if err != nil {
			return nil, fmt.Errorf("failed to read and/or decrypt KeyFile: %s: [%v]", gCfg.Account.KeyFile, err)
		}
		gCfg.AccountKey = key
	}

	optsTransactor := bind.NewKeyedTransactor(gCfg.AccountKey.PrivateKey)

	groupCaller, err := InsLogEvent.NewInsLogEventCaller(contractAddress, gCfg.Client)
	if err != nil {
		return nil, fmt.Errorf("EventCaller error on InsLogEvent: [%s]", err)
	}

	optsCaller := &bind.CallOpts{From: contractAddress}

	groupContract, err := InsLogEvent.NewInsLogEvent(contractAddress, gCfg.Client)
	if err != nil {
		return nil, fmt.Errorf("LogEvent  error on InsLogEvent: [%s] at", err)
	}

	return &InsLogEventControl{
		caller:          groupCaller,
		callerOpts:      optsCaller,
		transactor:      groupTransactor,
		transactorOpts:  optsTransactor,
		contract:        groupContract,
		contractAddress: contractAddress,
	}, nil
}

// IndexedEvent Calls contract to log event
func (kg *InsLogEventControl) IndexedEvent(indexedAddr string, msg string) (*types.Transaction, error) {
	g := common.HexToAddress(indexedAddr)
	return kg.transactor.IndexedEvent(kg.transactorOpts, g, msg)
}

// DecryptKeyFile reads in a file and uses the password to decrypt it.
func DecryptKeyFile(keyFile, password string) (*keystore.Key, error) {
	data, err := ioutil.ReadFile(keyFile)
	if err != nil {
		return nil, fmt.Errorf("unable to read KeyFile %s [%v]", keyFile, err)
	}
	key, err := keystore.DecryptKey(data, password)
	if err != nil {
		return nil, fmt.Errorf("unable to decrypt %s [%v]", keyFile, err)
	}
	return key, nil
}

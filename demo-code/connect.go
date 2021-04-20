package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/ethclient"
)

func Connect(gCfg *ConfigType) (err error) {
	client, err := ethclient.Dial(gCfg.URL)
	if err != nil {
		return fmt.Errorf("Error Connecting to Ethereum Server: %s [%v]", gCfg.URL, err)
	}
	gCfg.Client = client

	LogIt, err := NewInsLogEvent(gCfg)
	if err != nil {
		return fmt.Errorf("error attaching to InsLogEvent contract: [%v]", err)
	}
	gCfg.LogIt = LogIt

	return
}

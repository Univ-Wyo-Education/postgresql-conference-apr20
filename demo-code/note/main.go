package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pschlump/godebug"
	"gitlab.com/pschlump/PureImaginationServer/ReadConfig"
)

type EthAccount struct {
	Address         string
	KeyFile         string
	KeyFilePassword string
}

type ConfigType struct {
	URL          string            `json:"URL" default:"http://127.0.0.1:8545/"` // Address of server, http://localhost:8545?
	ContractAddr map[string]string `json:"ContractAddr"`
	Account      EthAccount
	LogAddress   string `json:"LogAddress"`

	// Global Data
	Client     *ethclient.Client
	LogIt      *InsLogEventControl
	AccountKey *keystore.Key
}

var Version = flag.Bool("version", false, "Report version of code and exit")
var Cfg = flag.String("cfg", "cfg.json", "config file for this call")

var GitCommit string
var DbOn map[string]bool

func init() {
	DbOn = make(map[string]bool)
	GitCommit = "Unknown"
}

var gCfg ConfigType

func main() {
	flag.Parse() // Parse CLI arguments to this, --cfg <name>.json

	fns := flag.Args()
	if len(fns) != 0 {
		fmt.Printf("Extra arguments are not supported [%s]\n", fns)
		os.Exit(1)
	}

	if *Version {
		fmt.Printf("Version (Git Commit): %s\n", GitCommit)
		os.Exit(0)
	}

	if Cfg == nil {
		fmt.Printf("--cfg is a required parameter\n")
		os.Exit(1)
	}

	// ------------------------------------------------------------------------------
	// Read in Configuration
	// ------------------------------------------------------------------------------
	err := ReadConfig.ReadFile(*Cfg, &gCfg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to read confguration: %s error %s\n", *Cfg, err)
		os.Exit(1)
	}

	err = Connect(&gCfg)
	if err != nil {
		fmt.Printf("Unable to connect to %s: error:%s\n", gCfg.URL, err)
		os.Exit(1)
	}

	ni, err := NewInsLogEvent(&gCfg)
	if err != nil {
		fmt.Printf("Unable to instanciate the contract: %s\n", err)
		os.Exit(1)
	}

	// addr := common.HexToAddress(gCfg.LogAddress)

	tx, err := ni.IndexedEvent(gCfg.LogAddress, "log-test")
	if err != nil {
		// todo
	} else {
		fmt.Printf("Success Tx: %s\n", godebug.SVarI(tx))
	}
}

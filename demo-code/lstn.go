package main

import (
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/lib/pq"
	"gitlab.com/pschlump/PureImaginationServer/ReadConfig"
)

type PgData struct {
	Cmd  string `json:"Cmd"`
	Hash string `json:"Hash"`
}

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

	// Connect to PostgreSQL
	pgConnectString := os.Getenv("POSTGRES_CONN") // TODO - Change to use your connetion

	_, err = sql.Open("postgres", pgConnectString)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect: %s %s\n", pgConnectString, err)
		os.Exit(1)
	}

	// Main Processing
	logError := func(ev pq.ListenerEventType, err error) {
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error on event listner:, %s\n", err)
		}
	}

	err = Connect(&gCfg)
	if err != nil {
		fmt.Printf("Unable to connect to %s: error:%s\n", gCfg.URL, err)
		os.Exit(1)
	}

	// Setup Contract Call
	ethLogEvent, err := NewInsLogEvent(&gCfg)
	if err != nil {
		fmt.Printf("Unable to instanciate the contract: %s\n", err)
		os.Exit(1)
	}

	listenForEventToOccure := func(l *pq.Listener) {
		nth := 0
		period := 50
		for {
			select {
			case n := <-l.Notify:
				nth++
				// fmt.Printf("Received event on chanel=->%v<- data=->%s<-\n", n.Channel, n.Extra)
				var pd PgData
				err := json.Unmarshal([]byte(n.Extra), &pd)
				if err != nil {
					fmt.Printf("err %s data ->%s<-\n", err, pd)
				}
				if (nth % period) == 0 {
					tx, err := ethLogEvent.IndexedEvent(gCfg.LogAddress, pd.Hash)
					if err != nil {
						fmt.Printf("err %s data ->%s<-\n", err, pd.Hash)
					} else {
						fmt.Printf("Success Tx: %s\n", JsonToString(tx))
					}
				}
				fmt.Printf("%+v\n", pd)
			case <-time.After(60 * time.Second):
				nth++
				fmt.Printf("No data received after 1 min, verifying connection to database\n")
				go func() {
					l.Ping()
				}()
			}
		}
	}

	pgConn := pq.NewListener(pgConnectString, 1*time.Second, time.Minute, logError)
	err = pgConn.Listen("events")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to setup listener: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Event Listener Started \n")
	for {
		listenForEventToOccure(pgConn)
	}
}

func JsonToString(m interface{}) string {
	s, err := json.MarshalIndent(m, "", "\t")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: ", err)
		return ""
	}
	return string(s)
}

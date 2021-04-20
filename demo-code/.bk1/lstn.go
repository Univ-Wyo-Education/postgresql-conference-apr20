package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/lib/pq"
	"github.com/pschlump/radix.v2/redis"
)

type PgData struct {
	Cmd  string `json:"Cmd"`
	Hash string `json:"Hash"`
}

func main() {

	// Connect to PostgreSQL
	pgConnectString := os.Getenv("POSTGRES_CONN") // TODO - Change to use your connetion

	_, err := sql.Open("postgres", pgConnectString)
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

	// xyzzy - Connect to Geth
	// xyzzy - Setup Key
	// xyzzy - Setup Contract Call

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
					// Write to Eth-Event
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

func SetupRedis(host, port, auth string) (client *redis.Client) {
	var err error
	client, err = redis.Dial("tcp", host+":"+port)
	if err != nil {
		fmt.Printf("Redis Error: Failed to connect: %s\n", err)
		os.Exit(1)
	}

	if auth != "" {
		_, err := client.Cmd("AUTH", auth).Str()
		if err != nil {
			fmt.Printf("Redis Error: Failed to authorize: %s\n", err)
			os.Exit(1)
		}
	}
	return
}

func JsonToString(m interface{}) string {
	s, err := json.MarshalIndent(m, "", "\t")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: ", err)
		return ""
	}
	return string(s)
}

package main

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/pschlump/radix.v2/redis"
)

func main() {

	// Connect to PostgreSQL
	pgConnectString := os.Getenv("POSTGRES_CONN") // TODO - Change to use your connetion

	conn, err := sql.Open("postgres", pgConnectString)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect: %s %s\n", pgConnectString, err)
		os.Exit(1)
	}

	// Connect to Redis
	// TODO - Change to use your connetion info if necessary
	rdConnection := SetupRedis("127.0.0.1", "6379", os.Getenv("REDIS_AUTH"))

	// Basic Test Code.

	ok := SelectFromCAuthKey(conn, rdConnection, "fdf6b075-2699-f443-f4b7-3c93ef4a79e4")
	fmt.Printf("1st Select: ok=%v\n", ok)

	InsertIntoCAuthKey(conn, rdConnection, "fdf6b075-2699-f443-f4b7-3c93ef4a79e4", "e70455d8-3788-2c41-0aa4-949b857476be")

	time.Sleep(2 * time.Second)

	ok = SelectFromCAuthKey(conn, rdConnection, "fdf6b075-2699-f443-f4b7-3c93ef4a79e4")
	fmt.Printf("2nd Select: ok=%v\n", ok)
}

func InsertIntoCAuthKey(conn *sql.DB, rd *redis.Client, key, user_id string) {
	stmt := `insert into c_auth_key ( key, user_id, valid_till ) values (  $1, $2, current_timestamp + interval '1 hour' )`
	ss, err := conn.Prepare(stmt)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Invalid inset: %s error:%s\n", stmt, err)
		return
	}
	_, err = ss.Exec(key, user_id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Invalid insert exec: %s error:%s\n", stmt, err, key, user_id)
		return
	}
}

func SelectFromCAuthKey(conn *sql.DB, rd *redis.Client, data string) (found bool) {
	stmt := `select 'ok' as "OK" from c_auth_key where key = $1`
	s, err := rd.Cmd("GET", data).Str()
	if err != nil {
		// Redis failed for some reason - may want to log this.
	} else if s == "" {
		// Redis did not return data - may want to log this.
	} else {
		// Got it - do stuff to - return row.
		fmt.Printf("Found stmt=%s key=%s in Redis, data=%s\n", stmt, data, s)
		// TODO - (maybee) modify s to just pick out the row data before returning.
		//        This depends on your implemenation.
		return true
	}

	// do select to PG
	var ok string
	row := conn.QueryRow(stmt, data)
	switch err := row.Scan(&ok); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return false
	default:
		fmt.Fprintf(os.Stderr, "Select Statement Error: %s\n", err)
		return false
	}
	_ = ok
	return true
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

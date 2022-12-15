package main

import (
	"fmt"
	"github.com/aerospike/aerospike-client-go"
	"os"
	"time"
)

func ae_main() {
	as_host := os.Getenv("AS_HOST")
	client, err := aerospike.NewClient(as_host, 3000)
	if err != nil {
		fmt.Println("Client: ", err)
	} else {
		// singleRead(client)
		multipleRead(client)
	}
}
func multipleRead(client *aerospike.Client) {
	// stmt := aerospike.NewStatement("infy_cache", "endpoints")

	// // Set max records to return
	// queryPolicy := aerospike.NewQueryPolicy()
	// queryPolicy.MaxRecords = 100
	// queryPolicy.SocketTimeout = 3000

	// Execute the query
	// recordSet, err := client.Query(queryPolicy, stmt)
	// if err != nil {
	// 	fmt.Println("recordSet Err: ", err)
	// }

	// // Get the results
	// for records := range recordSet.Results() {
	// 	if records.Err == nil {
	// 		if records != nil {
	// 			// Do something
	// 			fmt.Printf("Record: %v \n", records.Record.Bins)
	// 		}
	// 	} else {
	// 		fmt.Printf("records err: %v \n", records.Err)
	// 	}

	// }
	recordCount := 0
	begin := time.Now()
	policy := aerospike.NewScanPolicy()

	recordSet, err := client.ScanAll(policy, "demo", "data")
	if err != nil {
		fmt.Println("recordSet Err: ", err)
	}
	// for res := range recordSet.Results() {
	// 	if res.Err != nil {
	// 		// handle error here
	// 		fmt.Println(res.Err)
	// 	} else {
	//      recordCount++
	// 		// process record here
	// 		fmt.Println(res.Record.Bins)
	// 	}
	// }
L:
	for {
		select {
		case rec := <-recordSet.Records:
			if rec == nil {
				break L
			}
			recordCount++

			if (recordCount % 10000) == 0 {
				fmt.Println("Records ", recordCount)
			}
		case err := <-recordSet.Errors:
			// if there was an error, stop
			if err != nil {
				fmt.Println(err)
			}
		}
	}

	end := time.Now()
	seconds := float64(end.Sub(begin)) / float64(time.Second)
	fmt.Println("====================================", seconds)
	recordSet.Close()

	// Close the connection to the server
	client.Close()

}

func singleRead(client *aerospike.Client) {
	policy := aerospike.NewPolicy()
	policy.SocketTimeout = 30000
	//key = ('infy_cache', 'endpoints', data.get('id'))
	key, err := aerospike.NewKey("infy_cache", "endpoints", 1070)
	if err != nil {
		fmt.Println("Key Err: ", err)
	}
	record, err := client.Get(nil, key)
	if err != nil {
		fmt.Println("Record Err: ", err)
	} else {
		fmt.Printf("Record: %v", record.Bins)
	}
	client.Close()
}

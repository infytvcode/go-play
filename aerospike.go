package main

import (
	"fmt"
	"github.com/aerospike/aerospike-client-go"
	"time"
)

func aeMain() {
	client, err := aerospike.NewClient("52.15.163.71", 3000)
	if err != nil {
		fmt.Println("Client: ", err)
	} else {
		// singleRead(client)
		// multipleRead(client)

		add(client, 1, "test_2")
	}
}

func add(client *aerospike.Client, count int, key_string string) {
	// Creates a key with the namespace "sandbox", set "ufodata", and user key 5001
	key, err := aerospike.NewKey("test", "counter", key_string)
	if err != nil {
		fmt.Println("add: ", err)
	}

	// Can use a previously defined write policy or create a new one
	updatePolicy := aerospike.NewWritePolicy(0, 0)
	updatePolicy.SendKey = true
	updatePolicy.RecordExistsAction = aerospike.UPDATE

	// Create bin with new value
	newPosted := aerospike.NewBin("ssai", count)

	// Update record with new bin value
	// err = client.PutBins(updatePolicy, key, newPosted)
	record, err := client.Operate(updatePolicy, key, aerospike.AddOp(newPosted), aerospike.GetOp())
	if record != nil {
    fmt.Println("add received: ", record)
    received := record.Bins[newPosted.Name]
		fmt.Println("add received: ", received)
	}
	if err != nil {
		fmt.Println("add update: ", err)
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

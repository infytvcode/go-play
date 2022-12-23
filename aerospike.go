package main

import (
	"fmt"
	"github.com/aerospike/aerospike-client-go/v6"
	"time"
)

func aeMain() {
	client, err := aerospike.NewClient("52.15.163.71", 3000)
	if err != nil {
		fmt.Println("Client: ", err)
	} else {
		// singleRead(client)
		// multipleRead(client)
		//add(client, 1, "test_2")
		//batchUpdateV6(client)
		//batchUpdate(client)
		batchRead(client)
		client.Close()
	}
}

func batchRead(client *aerospike.Client) {
	start := time.Now()
	stmt := aerospike.NewStatement("test", "counter")

	// Set max records to return
	queryPolicy := aerospike.NewQueryPolicy()
	queryPolicy.TotalTimeout = 30 * time.Second
	queryPolicy.MaxRecords = 10000

	// Execute the query
	recordSet, err := client.Query(queryPolicy, stmt)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		c := 0
		// Get the results
		for records := range recordSet.Results() {
			if records != nil {
				// Do something
				fmt.Printf("Record: %d : %v \n", c, records.Record.Bins["opportunity"])
				c++
			}
		}
	}
	recordSet.Close()
	defer fmt.Println("V6: ", time.Since(start).String())
}

func batchUpdateV6(client *aerospike.Client) {
	start := time.Now()
	batchPolicy := aerospike.NewBatchPolicy()
	batchPolicy.TotalTimeout = 30 * time.Second

	// Create batch of records
	batchRecords := []aerospike.BatchRecordIfc{}

	// Create batch of keys
	bWritePolicy := aerospike.NewBatchWritePolicy()
	bWritePolicy.SendKey = true
	bWritePolicy.RecordExistsAction = aerospike.UPDATE

	for i := 0; i < 4000; i++ {
		// Create key
		key, _ := aerospike.NewKey("test", "counter", fmt.Sprintf("infy_%d", i))
		SSAI := aerospike.NewBin("ssai", 1)
		Opportunity := aerospike.NewBin("opportunity", 1)

		// Create record
		record := aerospike.NewBatchWrite(bWritePolicy, key,
			aerospike.AddOp(SSAI), aerospike.AddOp(Opportunity),
			aerospike.GetBinOp("ssai"), aerospike.GetBinOp("opportunity"))
		// Add to batch
		batchRecords = append(batchRecords, record)
	}
	err := client.BatchOperate(batchPolicy, batchRecords)
	if err != nil {
		fmt.Println(err)
	}
	// Access the records
	//for i, records := range batchRecords {
	//	record := records.BatchRec()
	//	if record != nil {
	//		// Do something
	//		fmt.Printf("Record: %d : %v : %v \n",
	//			i, record.Record.Bins["opportunity"], record.Record.Key.String())
	//	}
	//}
	defer fmt.Println("V6: ", time.Since(start).String())
}

func batchUpdate(client *aerospike.Client) {
	//batchPolicy := aerospike.NewBatchPolicy()
	start := time.Now()
	// Create batch of keys
	keys := make([]*aerospike.Key, 400)
	updatePolicy := aerospike.NewWritePolicy(0, 0)
	updatePolicy.SendKey = true
	updatePolicy.RecordExistsAction = aerospike.UPDATE
	for i := 0; i < 400; i++ {
		keys[i], _ = aerospike.NewKey("test", "counter", fmt.Sprintf("test_%d", i))
		newPosted := aerospike.NewBin("ssai", 1)
		client.Operate(updatePolicy, keys[i], aerospike.AddOp(newPosted))
		//client.BatchOperate()
		//record, err := client.Operate(updatePolicy, keys[i], aerospike.AddOp(newPosted), aerospike.GetOp())
		//if record != nil {
		//	received := record.Bins[newPosted.Name]
		//	fmt.Println("add received: ", received)
		//}
		//if err != nil {
		//	fmt.Println("add update: ", err)
		//}
	}
	defer fmt.Println("Loop:", time.Since(start).String())
}

func add(client *aerospike.Client, count int, keyString string) {
	// Creates a key with the namespace "sandbox", set "ufodata", and user key 5001
	key, err := aerospike.NewKey("test", "counter", keyString)
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
		received := record.Bins[newPosted.Name]

		fmt.Println("add received: ", received)
	}
	if err != nil {
		fmt.Println("add update: ", err)
	}

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

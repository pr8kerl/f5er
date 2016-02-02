package main

import (
	"encoding/json"
	"github.com/pr8kerl/f5er/f5"
	"io/ioutil"
	"log"
)

type LBStack struct {
	Nodes    []json.RawMessage `json:"nodes"`
	Pools    []json.RawMessage `json:"pools"`
	Virtuals []json.RawMessage `json:"virtuals"`
}

type LBEmptyBody struct{}

func showStack() {

	stack := LBStack{}

	// read in json file
	dat, err := ioutil.ReadFile(f5Input)
	if err != nil {
		log.Fatalf("error reading input json file: %s\n", err)
	}

	// convert json to a stack struct
	err = json.Unmarshal(dat, &stack)
	if err != nil {
		log.Fatalf("error unmarshaling input json file into a stack: %s\n", err)
	}

	// show nodes
	for count, n := range stack.Nodes {

		node := f5.LBNode{}
		// convert json to a node struct
		err = json.Unmarshal(n, &node)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("\nnode[%d]: %s\n", count, node.FullPath)

		err, res := appliance.ShowNode(node.FullPath)
		if err != nil {
			log.Printf("error showing node %s : %s\n", node.FullPath, err)
		} else {
			appliance.PrintObject(&res)
		}

	}

	// show pools
	for count, p := range stack.Pools {

		pool := f5.LBPool{}
		if err := json.Unmarshal(p, &pool); err != nil {
			log.Fatal(err)
		}

		log.Printf("\npool[%d]: %s\n", count, pool.FullPath)

		err, res := appliance.ShowPool(pool.FullPath)
		if err != nil {
			log.Printf("error showing pool %s : %s\n", pool.FullPath, err)
		} else {
			appliance.PrintObject(&res)
		}

	}
	// show virtual
	for count, v := range stack.Virtuals {

		virt := f5.LBVirtual{}
		if err := json.Unmarshal(v, &virt); err != nil {
			log.Fatal(err)
		}
		log.Printf("\nvirtual[%d]: %s\n", count, virt.FullPath)

		err, res := appliance.ShowVirtual(virt.FullPath)
		if err != nil {
			log.Printf("error showing virtual %s : %s\n", virt.FullPath, err)
		} else {
			appliance.PrintObject(&res)
		}

	}

}

func addStack() {

	stack := LBStack{}
	// read in json file
	dat, err := ioutil.ReadFile(f5Input)
	if err != nil {
		log.Fatal(err)
	}

	// convert json to a stack struct
	err = json.Unmarshal(dat, &stack)
	if err != nil {
		log.Fatal(err)
	}

	err, tid := appliance.StartTransaction()
	if err != nil {
		log.Fatalf("error creating transaction: %s\n", err)
	} else {
		log.Printf("transaction %s created\n", tid)
	}

	// add nodes
	for count, n := range stack.Nodes {

		node := f5.LBNode{}
		// convert json to a node struct - make sure it is valid
		err = json.Unmarshal(n, &node)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("\nnode[%d]: %s\n", count, node.FullPath)

		err, res := appliance.AddNode(&n)
		if err != nil {
			log.Printf("error adding node %s : %s\n", node.FullPath, err)
		} else {
			appliance.PrintObject(&res)
		}

	}

	// add pools
	for count, p := range stack.Pools {

		pool := f5.LBPool{}
		// convert json to a pool struct - make sure it is valid
		if err := json.Unmarshal(p, &pool); err != nil {
			log.Fatal(err)
		}
		log.Printf("\npool[%d]: %s\n", count, pool.FullPath)

		err, res := appliance.AddPool(&p)
		if err != nil {
			log.Printf("error adding pool %s : %s\n", pool.FullPath, err)
		} else {
			appliance.PrintObject(&res)
		}

	}

	// add virtual
	for count, v := range stack.Virtuals {

		virt := f5.LBVirtual{}
		// convert json to a virtual struct - make sure it is valid
		if err := json.Unmarshal(v, &virt); err != nil {
			log.Fatal(err)
		}

		log.Printf("\nvirtual[%d]: %s\n", count, virt.FullPath)

		err, res := appliance.AddVirtual(&v)
		if err != nil {
			log.Printf("error adding virtual %s : %s\n", virt.FullPath, err)
		} else {
			appliance.PrintObject(&res)
		}

	}

	// if we made it here - commit the transaction - remove the transaction header first
	err = appliance.CommitTransaction(tid)
	if err != nil {
		log.Printf("error commiting transaction %s : %s\n", tid, err)
	} else {
		log.Printf("transaction committed : %s\n", tid)
	}

}

func updateStack() {

	stack := LBStack{}
	// read in json file
	dat, err := ioutil.ReadFile(f5Input)
	if err != nil {
		log.Fatal(err)
	}

	// convert json to a stack struct
	err = json.Unmarshal(dat, &stack)
	if err != nil {
		log.Fatal(err)
	}

	err, tid := appliance.StartTransaction()
	if err != nil {
		log.Fatalf("error creating transaction: %s\n", err)
	} else {
		log.Printf("transaction %s created\n", tid)
	}

	// nodes
	for count, n := range stack.Nodes {

		node := f5.LBNode{}
		// convert json to a node struct
		err = json.Unmarshal(n, &node)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("\nnode[%d]: %s\n", count, node.FullPath)

		err, res := appliance.UpdateNode(node.FullPath, &n)
		if err != nil {
			log.Printf("error adding virtual %s : %s\n", node.FullPath, err)
		} else {
			appliance.PrintObject(&res)
		}

	}

	// pools
	for count, p := range stack.Pools {

		pool := f5.LBPool{}
		if err := json.Unmarshal(p, &pool); err != nil {
			log.Fatal(err)
		}
		log.Printf("\npool[%d]: %s\n", count, pool.FullPath)

		err, res := appliance.UpdatePool(pool.FullPath, &p)
		if err != nil {
			log.Printf("error adding pool %s : %s\n", pool.FullPath, err)
		} else {
			appliance.PrintObject(&res)
		}

	}

	// add virtual
	for count, v := range stack.Virtuals {

		virt := f5.LBVirtual{}
		if err := json.Unmarshal(v, &virt); err != nil {
			log.Fatal(err)
		}
		log.Printf("\nvirtual[%d]: %s\n", count, virt.FullPath)

		err, res := appliance.UpdateVirtual(virt.FullPath, &v)
		if err != nil {
			log.Printf("error adding virtual %s : %s\n", virt.FullPath, err)
		} else {
			appliance.PrintObject(&res)
		}

	}

	// if we made it here - commit the transaction
	err = appliance.CommitTransaction(tid)
	if err != nil {
		log.Printf("error commiting transaction %s : %s\n", tid, err)
	} else {
		log.Printf("transaction committed : %s\n", tid)
	}

}

func deleteStack() {

	stack := LBStack{}
	// read in json file
	dat, err := ioutil.ReadFile(f5Input)
	if err != nil {
		log.Fatal(err)
	}

	// convert json to a stack struct
	err = json.Unmarshal(dat, &stack)
	if err != nil {
		log.Fatal(err)
	}

	err, tid := appliance.StartTransaction()
	if err != nil {
		log.Fatalf("error creating transaction: %s\n", err)
	} else {
		log.Printf("transaction %s created\n", tid)
	}

	// virtual
	for count, v := range stack.Virtuals {

		virt := f5.LBVirtual{}
		if err := json.Unmarshal(v, &virt); err != nil {
			log.Fatal(err)
		}
		log.Printf("\nvirtual[%d]: %s\n", count, virt.FullPath)

		err, res := appliance.DeleteVirtual(virt.FullPath)
		if err != nil {
			log.Printf("error deleting virtual %s : %s\n", virt.FullPath, err)
		} else {
			appliance.PrintObject(&res)
		}

	}

	// pools
	for count, p := range stack.Pools {

		pool := f5.LBPool{}
		if err := json.Unmarshal(p, &pool); err != nil {
			log.Fatal(err)
		}
		log.Printf("\npool[%d]: %s\n", count, pool.FullPath)

		err, res := appliance.DeletePool(pool.FullPath)
		if err != nil {
			log.Printf("error deleting pool %s : %s\n", pool.FullPath, err)
		} else {
			appliance.PrintObject(&res)
		}

	}

	// if we made it here - commit the transaction
	err = appliance.CommitTransaction(tid)
	if err != nil {
		log.Printf("error commiting transaction %s : %s\n", tid, err)
	} else {
		log.Printf("transaction committed : %s\n", tid)
	}

	// delete nodes outside of transaction - pools depend on them and won't delete otherwise
	for count, n := range stack.Nodes {

		node := f5.LBNode{}
		// convert json to a node struct
		err = json.Unmarshal(n, &node)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("\nnode[%d]: %s\n", count, node.FullPath)

		err, res := appliance.DeleteNode(node.FullPath)
		if err != nil {
			log.Printf("error deleting node %s : %s\n", node.FullPath, err)
		} else {
			appliance.PrintObject(&res)
		}
	}

}

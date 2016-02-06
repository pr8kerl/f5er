package main

import (
	"encoding/json"
	"github.com/pr8kerl/f5er/f5"
	"io/ioutil"
	"log"
)

type LBStack struct {
	ServerSsl []json.RawMessage `json:"profiles-server-ssl"`
	ClientSsl []json.RawMessage `json:"profiles-client-ssl"`
	Nodes     []json.RawMessage `json:"nodes"`
	Pools     []json.RawMessage `json:"pools"`
	Rules     []json.RawMessage `json:"rules"`
	Policies  []json.RawMessage `json:"policies"`
	Virtuals  []json.RawMessage `json:"virtuals"`
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

	// show server-ssl
	for count, n := range stack.ServerSsl {

		obj := f5.LBServerSsl{}
		// convert json to a struct
		err = json.Unmarshal(n, &obj)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("\nserver-ssl[%d]: %s\n", count, obj.FullPath)

		err, res := appliance.ShowServerSsl(obj.FullPath)
		if err != nil {
			log.Printf("error showing server-ssl %s : %s\n", obj.FullPath, err)
		} else {
			appliance.PrintObject(&res)
		}

	}

	// show client-ssl
	for count, n := range stack.ClientSsl {

		obj := f5.LBClientSsl{}
		// convert json to a struct
		err = json.Unmarshal(n, &obj)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("\nclient-ssl[%d]: %s\n", count, obj.FullPath)

		err, res := appliance.ShowClientSsl(obj.FullPath)
		if err != nil {
			log.Printf("error showing client-ssl %s : %s\n", obj.FullPath, err)
		} else {
			appliance.PrintObject(&res)
		}

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

	// show rules
	for count, n := range stack.Rules {

		obj := f5.LBRule{}
		// convert json to a struct
		err = json.Unmarshal(n, &obj)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("\nrule[%d]: %s\n", count, obj.FullPath)

		err, res := appliance.ShowRule(obj.FullPath)
		if err != nil {
			log.Printf("error showing rule %s : %s\n", obj.FullPath, err)
		} else {
			appliance.PrintObject(&res)
		}

	}

	// show policies
	for count, n := range stack.Policies {

		obj := f5.LBPolicy{}
		// convert json to an obj struct
		err = json.Unmarshal(n, &obj)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("\npolicy[%d]: %s\n", count, obj.FullPath)

		err, res := appliance.ShowPolicy(obj.FullPath)
		if err != nil {
			log.Printf("error showing policy %s : %s\n", obj.FullPath, err)
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

	// add server-ssl
	for count, n := range stack.ServerSsl {

		obj := f5.LBServerSsl{}
		// convert json to a struct - make sure it is valid
		err = json.Unmarshal(n, &obj)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("\nserver-ssl[%d]: %s\n", count, obj.FullPath)

		// use the raw json to add - only set minimal number of fields
		err, res := appliance.AddServerSsl(&n)
		if err != nil {
			log.Printf("error adding server-ssl %s : %s\n", obj.FullPath, err)
		} else {
			appliance.PrintObject(&res)
		}

	}

	// add client-ssl
	for count, n := range stack.ClientSsl {

		obj := f5.LBClientSsl{}
		// convert json to a struct - make sure it is valid
		err = json.Unmarshal(n, &obj)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("\nclient-ssl[%d]: %s\n", count, obj.FullPath)

		// use the raw json to add - only set minimal number of fields
		err, res := appliance.AddClientSsl(&n)
		if err != nil {
			log.Printf("error adding client-ssl %s : %s\n", obj.FullPath, err)
		} else {
			appliance.PrintObject(&res)
		}

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

	// add rules
	for count, n := range stack.Rules {

		obj := f5.LBRule{}
		// convert json to a struct - make sure it is valid
		err = json.Unmarshal(n, &obj)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("\nrule[%d]: %s\n", count, obj.FullPath)

		// use the raw json to add - only set minimal number of fields
		err, res := appliance.AddRule(&n)
		if err != nil {
			log.Printf("error adding rule %s : %s\n", obj.FullPath, err)
		} else {
			appliance.PrintObject(&res)
		}

	}

	// add policies
	for count, n := range stack.Policies {

		obj := f5.LBPolicy{}
		// convert json to a struct - make sure it is valid
		err = json.Unmarshal(n, &obj)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("\npolicy[%d]: %s\n", count, obj.FullPath)

		// use the raw json to add - only set minimal number of fields
		err, res := appliance.AddPolicy(&n)
		if err != nil {
			log.Printf("error adding policy %s : %s\n", obj.FullPath, err)
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

	// update server-ssl
	for count, n := range stack.ServerSsl {

		obj := f5.LBServerSsl{}
		// convert json to a struct - make sure it is valid
		err = json.Unmarshal(n, &obj)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("\nserver-ssl[%d]: %s\n", count, obj.FullPath)

		// use the raw json to add - only set minimal number of fields
		err, res := appliance.UpdateServerSsl(obj.FullPath, &n)
		if err != nil {
			log.Printf("error updating server-ssl %s : %s\n", obj.FullPath, err)
		} else {
			appliance.PrintObject(&res)
		}

	}

	// update client-ssl
	for count, n := range stack.ServerSsl {

		obj := f5.LBClientSsl{}
		// convert json to a struct - make sure it is valid
		err = json.Unmarshal(n, &obj)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("\nclient-ssl[%d]: %s\n", count, obj.FullPath)

		// use the raw json to add - only set minimal number of fields
		err, res := appliance.UpdateClientSsl(obj.FullPath, &n)
		if err != nil {
			log.Printf("error updating client-ssl %s : %s\n", obj.FullPath, err)
		} else {
			appliance.PrintObject(&res)
		}

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

	// update rules
	for count, n := range stack.Rules {

		obj := f5.LBRule{}
		// convert json to a struct - make sure it is valid
		err = json.Unmarshal(n, &obj)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("\nrule[%d]: %s\n", count, obj.FullPath)

		// use the raw json to add - only set minimal number of fields
		err, res := appliance.UpdateRule(obj.FullPath, &n)
		if err != nil {
			log.Printf("error updating rule %s : %s\n", obj.FullPath, err)
		} else {
			appliance.PrintObject(&res)
		}

	}

	// update policies
	for count, n := range stack.Policies {

		obj := f5.LBPolicy{}
		// convert json to a struct - make sure it is valid
		err = json.Unmarshal(n, &obj)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("\npolicy[%d]: %s\n", count, obj.FullPath)

		// use the raw json to add - only set minimal number of fields
		err, res := appliance.UpdatePolicy(obj.FullPath, &n)
		if err != nil {
			log.Printf("error updating policy %s : %s\n", obj.FullPath, err)
		} else {
			appliance.PrintObject(&res)
		}

	}

	// update virtual
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
		log.Printf("transaction committed : %s\n\n", tid)
	}

	// another transaction for all the other objects
	err, tid = appliance.StartTransaction()
	if err != nil {
		log.Fatalf("error creating transaction: %s\n", err)
	} else {
		log.Printf("transaction %s created\n", tid)
	}

	// delete policies
	for count, n := range stack.Policies {

		obj := f5.LBPolicy{}
		// convert json to a struct - make sure it is valid
		err = json.Unmarshal(n, &obj)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("\npolicy[%d]: %s\n", count, obj.FullPath)

		// use the raw json to add - only set minimal number of fields
		err, res := appliance.DeletePolicy(obj.FullPath)
		if err != nil {
			log.Printf("error deleting policy %s : %s\n", obj.FullPath, err)
		} else {
			appliance.PrintObject(&res)
		}

	}

	// delete rules
	for count, n := range stack.Rules {

		obj := f5.LBRule{}
		// convert json to a struct - make sure it is valid
		err = json.Unmarshal(n, &obj)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("\nrule[%d]: %s\n", count, obj.FullPath)

		// use the raw json to add - only set minimal number of fields
		err, res := appliance.DeleteRule(obj.FullPath)
		if err != nil {
			log.Printf("error deleting rule %s : %s\n", obj.FullPath, err)
		} else {
			appliance.PrintObject(&res)
		}

	}

	// delete client-ssl
	for count, n := range stack.ClientSsl {

		obj := f5.LBClientSsl{}
		// convert json to a struct - make sure it is valid
		err = json.Unmarshal(n, &obj)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("\nclient-ssl[%d]: %s\n", count, obj.FullPath)

		// use the raw json to add - only set minimal number of fields
		err, res := appliance.DeleteClientSsl(obj.FullPath)
		if err != nil {
			log.Printf("error deleting client-ssl %s : %s\n", obj.FullPath, err)
		} else {
			appliance.PrintObject(&res)
		}

	}

	// delete server-ssl
	for count, n := range stack.ServerSsl {

		obj := f5.LBServerSsl{}
		// convert json to a struct - make sure it is valid
		err = json.Unmarshal(n, &obj)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("\nserver-ssl[%d]: %s\n", count, obj.FullPath)

		// use the raw json to add - only set minimal number of fields
		err, res := appliance.DeleteServerSsl(obj.FullPath)
		if err != nil {
			log.Printf("error deleting server-ssl %s : %s\n", obj.FullPath, err)
		} else {
			appliance.PrintObject(&res)
		}

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

	// if we made it here - commit the transaction
	err = appliance.CommitTransaction(tid)
	if err != nil {
		log.Printf("error commiting transaction %s : %s\n", tid, err)
	} else {
		log.Printf("transaction committed : %s\n", tid)
	}

}

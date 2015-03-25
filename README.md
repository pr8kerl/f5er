# f5er

A golang F5 rest plaything

Supports nodes, pools, poolmembers, virtuals, nodes, policies, irules, client-ssl profiles and http monitors in full - so far.

Create, modify and delete F5 objects easily, using json input files.

A convenience entity called a **stack** can be used to act upon nodes, their pool and its virtual server as one.

Supports the REST methods GET (show), POST (create), PUT (update) and DELETE (delete).

Most commands will display the response in json as provide by the F5 device. Please note that although the response json may look similar to input json, some json object fields differ. For example, pool members within a pool are displayed within a membersReference object in a response, however members must be defined as an array within the **members** array in a pool object. Also some json object response fields are read-only and cannot be used with an input object (the object supplied in the body of a POST or PUT operation.



## credentials

F5 ip address and login credentials are stored in a json input file in the current directory.
The expected file is **f5.json**.

```
{
  "f5": "192.168.0.100",
  "username": "admin",
  "passwd": "admin"
}
```

## help

Use the help command to display hints and available subcommands

```
./f5er help show
show current state of F5 objects. Show requires an object, eg. f5er show pool

Usage: 
  f5er show [flags]
  f5er show [command]
Available Commands: 
  pool         show a pool
  poolmember   show pool members
  virtual      show a virtual server
  node         show a node
  policy       show a policy
  device       show an f5 device
  rule         show a rule
  client-ssl   show a client-ssl profile
  monitor-http show a monitor-http profile
  stack        show a stack transaction


Global Flags:
  -d, --debug=false: debug output
  -f, --f5="": IP or hostname of F5 to poke
  -h, --help=false: help for show
  -i, --input="": input json f5 configuration

Additional help topics: 
  f5er show    show F5 objects
  f5er add     add F5 objects
  f5er update  update F5 objects
  f5er delete  delete F5 objects
  f5er offline offline a pool member
  f5er online  online a pool member
  f5er help    Help about any command

Use "f5er help [command]" for more information about a command.

```

## Device

The following command will display info about the F5 device or cluster. Handy to see which is active/standby.
Only show is supported for device.

```
./f5er show device
[
	{
		"name": "bigip-1.example.com",
		"fullPath": "/Common/bigip-1.example.com",
		"failoverState": "standby",
		"managementIP": "192.168.0.100"
	},
	{
		"name": "bigip-2.example.com",
		"fullPath": "/Common/bigip-2.example.com",
		"failoverState": "active",
		"managementIP": "192.168.0.101"
	}
]
```

## Stacks

This is a convenience construct and does not exist within F5 terminology. 
It effectively allows commands to work on multiple nodes, a single pool and multiple virtual servers in one operation. It uses a REST transaction to do so.
Look at the file stack.json to see how to structure the input file.
Show, add, update and delete operations are supported.

```
./f5er help add stack
add a new stack

Usage: 
  f5er add stack [flags]

 Available Flags:
  -d, --debug=false: debug output
  -f, --f5="": IP or hostname of F5 to poke
  -h, --help=false: help for stack
  -i, --input="": input json f5 configuration


```

## Pools

Show, add, delete and update a single pool.

### Show pools

Show all pools
```
f5er show pool
```

Display a single pool in detail
```
f5er show pool /partition/poolname
```

### Add a new pool

Provide a json input file with all the new pool configuration information. You can base a new pool on the output from a current pool. 

```
f5er add pool --input=pool.json
```

### Modify an existing pool

You can modify the config of an existing pool, including the pool members.

Again, provide a json input file with the updated configuration

```
f5er update pool --input=pool.json
```


## Pool members

Pool members can be created/modified in a similar way to pools.
When pool members are created/modified, the current pool member info is always overwritten. So any new config needs to provide information for all pool members.

Additionally, pool members can be manually brought online or taken offline.

### Offline a poolmember

Provide the pool name and pool member. The following will manually mark a pool member offline. Active sessions will continue until they naturally end. This allows connection draining.
```
f5er offline poolmember --pool=/partition/poolname /partition/poolmember:portnumber
```

To take a pool member offline immediately, provide the **--now** command line option. This will prevent existing connections from continuing on the pool member.
```
f5er offline poolmember --now --pool=/partition/poolname /partition/poolmember:portnumber
```
### Bring a pool member online

The opposite to the poolmember offline command

# cross-compile 

Use [gox](https://github.com/mitchellh/gox).

* install gox
```
go get github.com/mitchellh/gox
```

* compile cross-compilation build chain for windows 32/64 bit
```
gox -build-toolchain -os="windows"
```

* create windows binaries
```
gox -os="windows"
Number of parallel builds: 4

-->     windows/386: _/home/ians/work/f5er
-->   windows/amd64: _/home/ians/work/f5er
```


# Saved F5 snippets

These bits are saved here to serve as reminders to commands that could be supported in the future.

#### cluster member status
```
curl -sk -u admin:admin -H "Content-Type: application/json" -X GET https://x.x.x.x/mgmt/tm/cm/failover-status 
{
  "kind":"tm:cm:failover-status:failover-statusstats",
  "selfLink":"https://localhost/mgmt/tm/cm/failover-status?ver=11.6.0",
  "entries":{
    "https://localhost/mgmt/tm/cm/failover-status/0":{
      "nestedStats":{
        "entries":{
          "color":{"description":"green"},
          "https://localhost/mgmt/tm/cm/failoverStatus/0/details":{
            "nestedStats":{
              "entries":{
                "https://localhost/mgmt/tm/cm/failoverStatus/0/details/0":{
                  "nestedStats":{
                    "entries":{
                      "details":{"description":"active for /Common/traffic-group-1"}
                    }
                  }
                }
              }
            }
          },
          "status":{"description":"ACTIVE"},
          "summary":{"description":"1/1 active"}
        }
      }
    }
  }
}

```

#### config sync

```
curl -sk -u admin:admin -H "Content-Type: application/json" -X POST -d '{"command":"run","utilCmdArgs":"config-sync to-group pair-group-name"}' https://x.x.x.x/mgmt/tm/cm
```

### show pool member stats

```
curl -sk -u admin:admin -H "Content-Type: application/json" https://x.x.x.x/mgmt/tm/ltm/pool/~DMZ~audmzbilltweb-sit_443_pool/members/stats
```

### show partitions

```
curl -sk -u admin:admin -H "Content-Type: application/json" https://x.x.x.x/mgmt/tm/sys/folder
```

### certificates/keys etc

```
curl -sk -u admin:admin -H "Content-Type: application/json" https://x.x.x.x/mgmt/tm/sys/crypto/
curl -sk -u admin:admin -H "Content-Type: application/json" https://x.x.x.x/mgmt/tm/sys/crypto/cert
curl -sk -u admin:admin -H "Content-Type: application/json" https://x.x.x.x/mgmt/tm/sys/crypto/key
```

### pool member status

Take pool member offline. Active sessions are no longer allowed to continue
{"state": "user-down", "session": "user-disabled"} (Member Forced Offline in GUI)

Take pool member offline, active sessions continue (drain)
{"state": "user-up", "session": "user-disabled"} (Member Disabled in GUI)

Enable a pool member
{"state": "user-up", "session": "user-enabled"}  (Member Enabled in GUI)


#### transaction info


```
POST https://192.168.25.42/mgmt/tm/transaction

{
"transId":1389812351,
"state":"STARTED",
"timeoutSeconds":30,
"kind":"tm:transactionstate",
"selfLink":"https://localhost/mgmt/tm/transaction/1389812351?ver=11.5.0"
}


GET https://192.168.25.42/mgmt/tm/transaction
GET https://192.168.25.42/mgmt/tm/transaction/<transId>


Modifying a transaction
After you create a transaction, you can populate the transaction by adding commands. Individual commands
comprise the operations that a transaction performs. Commands are added in the order they are received
but you can delete commands or change the order of the commands in the transaction.
1. To add a command to a transaction, use the POST method and specify the
X-F5-REST-Coordination-Id HTTP header with the transaction ID value from the example
(1389812351). In the example, the request creates a new pool and adds a single member to the pool.
POST https://192.168.25.42/mgmt/tm/ltm/pool
X-F5-REST-Coordination-Id:1389812351
{
"name":"tcb-xact-pool",
"members": [ {"name":"192.168.25.32:80","description":"First pool for transactions"} ]
}

The response indicates that iControl® REST added the operation to the transaction.
{
"transId":1389812351,
"state":"STARTED",
"timeoutSeconds":30,
"kind":"tm:transactionstate",
"selfLink":"https://localhost/mgmt/tm/transaction/1389813931?ver=11.5.0"
}
```

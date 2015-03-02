# f5er

A golang F5 rest plaything

## To do
Everything.  
Well not quite.

Supports pools, poolmembers and nodes in full - so far.

## credentials

F5 ip address and login credentials are stored in a json input file in the current directory.
The expected file is **f5.json**.

```
{
	"credentials": {
  				"f5": "192.168.0.100",
					"username": "admin",
					"passwd": "admin"
	}
}
```


## Pools

### Show pools

Show all pools
```
f5 show pool
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
f5er update pool --input 
```


## Pool members

Pool members can be created/modified in a similar way to pools.
When pool members are created/modified, the current pool member info is always overwritten. So any new config needs to provide information for all pool members.

Additionally, pool members can be manually brought online or taken offline.

### take a pool member offline

Provide the pool name and pool member. The following will manually mark a pool member offline. Active sessions will continue until they naturally end. This allows connection draining.
```
f5er offline poolmember --pool=/partition/poolname /partition/poolmember:portnumber
```

To take a pool member offline immediately, provide the **--now** command line option. This will prevent existing connections from continuing on the pool member.
```
f5er offline poolmember --now --pool=/partition/poolname /partition/poolmember:portnumber
```
## Bring a pool member online

## cross-compile for windows
Use [gox](https://github.com/mitchellh/gox).

* install
```
go get github.com/mitchellh/gox
```

* compile cross-compilation build chain
```
gox -build-toolchain
```

* create windows binaries
```
gox -os="windows"
Number of parallel builds: 4

-->     windows/386: _/home/ians/work/f5er
-->   windows/amd64: _/home/ians/work/f5er
```


### Pool

**https://x.x.x.x/mgmt/tm/ltm/pool/~DMZ~audmzbilltweb-sit_443_pool**

```
RawText
{"kind":"tm:ltm:pool:poolstate","name":"audmzbilltweb-sit_443_pool","partition":"DMZ","fullPath":"/DMZ/audmzbilltweb-sit_443_pool","generation":211,"selfLink":"https://localhost/mgmt/tm/ltm/pool/~DMZ~audmzbilltweb-sit_443_pool?ver=11.6.0","allowNat":"yes","allowSnat":"yes","ignorePersistedWeight":"disabled","ipTosToClient":"pass-through","ipTosToServer":"pass-through","linkQosToClient":"pass-through","linkQosToServer":"pass-through","loadBalancingMode":"round-robin","minActiveMembers":0,"minUpMembers":0,"minUpMembersAction":"failover","minUpMembersChecking":"disabled","monitor":"min 1 of { /Common/tcp }","queueDepthLimit":0,"queueOnConnectionLimit":"disabled","queueTimeLimit":0,"reselectTries":0,"serviceDownAction":"none","slowRampTime":10,"membersReference":{"link":"https://localhost/mgmt/tm/ltm/pool/~DMZ~audmzbilltweb-sit_443_pool/members?ver=11.6.0","isSubcollection":true}}
```

### Members

**https://x.x.x.x/mgmt/tm/ltm/pool/~DMZ~audmzbilltweb-sit_443_pool/members?ver=11.6.0**

```
{"kind":"tm:ltm:pool:members:memberscollectionstate","selfLink":"https://localhost/mgmt/tm/ltm/pool/~DMZ~audmzbilltweb-sit_443_pool/members?ver=11.6.0","items":[{"kind":"tm:ltm:pool:members:membersstate","name":"audmzbilltweb01-sit:443","partition":"DMZ","fullPath":"/DMZ/audmzbilltweb01-sit:443","generation":233,"selfLink":"https://localhost/mgmt/tm/ltm/pool/~DMZ~audmzbilltweb-sit_443_pool/members/~DMZ~audmzbilltweb01-sit:443?ver=11.6.0","address":"x.x.x.x%6","connectionLimit":0,"dynamicRatio":1,"ephemeral":"false","fqdn":{"autopopulate":"disabled"},"inheritProfile":"enabled","logging":"disabled","monitor":"default","priorityGroup":0,"rateLimit":"disabled","ratio":1,"session":"monitor-enabled","state":"up"},{"kind":"tm:ltm:pool:members:membersstate","name":"audmzbilltweb02-sit:443","partition":"DMZ","fullPath":"/DMZ/audmzbilltweb02-sit:443","generation":153,"selfLink":"https://localhost/mgmt/tm/ltm/pool/~DMZ~audmzbilltweb-sit_443_pool/members/~DMZ~audmzbilltweb02-sit:443?ver=11.6.0","address":"x.x.x.x%6","connectionLimit":0,"dynamicRatio":1,"ephemeral":"false","fqdn":{"autopopulate":"disabled"},"inheritProfile":"enabled","logging":"disabled","monitor":"default","priorityGroup":0,"rateLimit":"disabled","ratio":1,"session":"monitor-enabled","state":"up"}]}
```


### cluster member status
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

### config sync
```
curl -sk -u admin:admin -H "Content-Type: application/json" -X POST -d '{"command":"run","utilCmdArgs":"config-sync to-group pair-group-name"}' https://x.x.x.x/mgmt/tm/cm
```

### show interesting info about device
```
curl -sk -u admin:admin -H "Content-Type: application/json" https://x.x.x.x/mgmt/tm/cm/device
```

### show pool member stats
```
curl -sk -u admin:admin -H "Content-Type: application/json" https://x.x.x.x/mgmt/tm/ltm/pool/~DMZ~audmzbilltweb-sit_443_pool/members/stats
```

### show partitions
```
curl -sk -u admin:admin -H "Content-Type: application/json" https://x.x.x.x/mgmt/tm/sys/folder
```

### pool member status

Take pool member offline. Active sessions are no longer allowed to continue
{"state": "user-down", "session": "user-disabled"} (Member Forced Offline in GUI)

Take pool member offline, active sessions continue (drain)
{"state": "user-up", "session": "user-disabled"} (Member Disabled in GUI)

Enable a pool member
{"state": "user-up", "session": "user-enabled"}  (Member Enabled in GUI)


### transaction


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

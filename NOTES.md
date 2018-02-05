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

#### system stats

Get global interface statistics

```
/mgmt/tm/net/interface/stats
```

#### Memory stats

```
/mgmt/tm/sys/memory
```

#### CPU  stats

```
/mgmt/tm/sys/cpu
```

#### Disk stats

```
/mgmt/tm/sys/disk/logical-disk
/mgmt/tm/sys/disk/application-volume
```

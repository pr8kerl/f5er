# f5er

A golang F5 rest client

## To do
- everything

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

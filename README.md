# f5er

A golang F5 rest client

## To do
- everything


### Pool

**https://192.168.0.100/mgmt/tm/ltm/pool/~DMZ~audmzbilltweb-sit_443_pool**

```
RawText
{"kind":"tm:ltm:pool:poolstate","name":"audmzbilltweb-sit_443_pool","partition":"DMZ","fullPath":"/DMZ/audmzbilltweb-sit_443_pool","generation":211,"selfLink":"https://localhost/mgmt/tm/ltm/pool/~DMZ~audmzbilltweb-sit_443_pool?ver=11.6.0","allowNat":"yes","allowSnat":"yes","ignorePersistedWeight":"disabled","ipTosToClient":"pass-through","ipTosToServer":"pass-through","linkQosToClient":"pass-through","linkQosToServer":"pass-through","loadBalancingMode":"round-robin","minActiveMembers":0,"minUpMembers":0,"minUpMembersAction":"failover","minUpMembersChecking":"disabled","monitor":"min 1 of { /Common/tcp }","queueDepthLimit":0,"queueOnConnectionLimit":"disabled","queueTimeLimit":0,"reselectTries":0,"serviceDownAction":"none","slowRampTime":10,"membersReference":{"link":"https://localhost/mgmt/tm/ltm/pool/~DMZ~audmzbilltweb-sit_443_pool/members?ver=11.6.0","isSubcollection":true}}
```

### Members

**https://192.168.0.100/mgmt/tm/ltm/pool/~DMZ~audmzbilltweb-sit_443_pool/members?ver=11.6.0**

```
{"kind":"tm:ltm:pool:members:memberscollectionstate","selfLink":"https://localhost/mgmt/tm/ltm/pool/~DMZ~audmzbilltweb-sit_443_pool/members?ver=11.6.0","items":[{"kind":"tm:ltm:pool:members:membersstate","name":"audmzbilltweb01-sit:443","partition":"DMZ","fullPath":"/DMZ/audmzbilltweb01-sit:443","generation":233,"selfLink":"https://localhost/mgmt/tm/ltm/pool/~DMZ~audmzbilltweb-sit_443_pool/members/~DMZ~audmzbilltweb01-sit:443?ver=11.6.0","address":"10.60.61.215%6","connectionLimit":0,"dynamicRatio":1,"ephemeral":"false","fqdn":{"autopopulate":"disabled"},"inheritProfile":"enabled","logging":"disabled","monitor":"default","priorityGroup":0,"rateLimit":"disabled","ratio":1,"session":"monitor-enabled","state":"up"},{"kind":"tm:ltm:pool:members:membersstate","name":"audmzbilltweb02-sit:443","partition":"DMZ","fullPath":"/DMZ/audmzbilltweb02-sit:443","generation":153,"selfLink":"https://localhost/mgmt/tm/ltm/pool/~DMZ~audmzbilltweb-sit_443_pool/members/~DMZ~audmzbilltweb02-sit:443?ver=11.6.0","address":"10.60.61.216%6","connectionLimit":0,"dynamicRatio":1,"ephemeral":"false","fqdn":{"autopopulate":"disabled"},"inheritProfile":"enabled","logging":"disabled","monitor":"default","priorityGroup":0,"rateLimit":"disabled","ratio":1,"session":"monitor-enabled","state":"up"}]}
```

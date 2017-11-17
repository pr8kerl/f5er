# f5er

[![Build Status](https://travis-ci.org/pr8kerl/f5er.svg?branch=master)](https://travis-ci.org/pr8kerl/f5er)

An F5 rest client.

Supports nodes, pools, poolmembers, virtuals, nodes, policies, irules, client-ssl profiles and http monitors in full - so far. Some statistics retrieval.

Create, modify and delete F5 objects easily, using json input files.

A convenience entity called a **stack** can be used to act upon nodes, their pool and its virtual server as one.

Supports the REST methods GET (show), POST (create), PUT (update) and DELETE (delete).

Most commands will display the response in json as provided by the F5 device. Please note that although the response json may look similar to input json, some json object fields differ. For example, pool members within a pool are displayed within a membersReference object in a response, however members must be defined as an array within the **members** array in a pool object. Also some json object response fields are read-only and cannot be used with an input object (the object supplied in the body of a POST or PUT operation.

It can now display statistics in graphite format for virtuals, pools, nodes and rules. 
If you are a prometheus user, then also check out [bigip_exporter](https://github.com/ExpressenAB/bigip_exporter).

## Installation

Use docker and docker-compose to build.

```
docker-compose run make [linux|windows|darwin]
```

Or build manually from repo installed in your GOPATH. Requires [glide](https://glide.sh/).
```
make
```

## Configuration

**f5er** looks for configuration in the current environment, or if not found in a config file.

### Environment variables

You can use the following environment variables to specify the F5 device and credentials.
```
F5_DEVICE="192.168.0.100"
F5_USERNAME="admin"
F5_PASSWD="superSecretSquirrel"

export F5_DEVICE F5_USERNAME F5_PASSWD
```

### Config file

F5 ip address/hostname and login credentials can also be stored in a json input file in the current directory.
The expected file is called **f5.json** and it can be in the current working directory or in $HOME/.f5/.
Below is a full example of all current configurables. 

```
{
  "device": "192.168.0.100",
  "username": "admin",
  "passwd": "superSecretSquirrel",
  "stats_path_prefix": "prd.f5.bigip01",
  "stats_show_zero_values": false
}
```

The **stats_** configuration options (used for displaying statistics) are only available in a config file.

### Proxy support

**f5er** will use a proxy if the conventional proxy environment variables HTTP_PROXY or HTTPS_PROXY are set.

## Usage

Use the help command to display available commands.

```
./f5er help
A utility to manage F5 configuration objects

Usage:
  f5er [flags]
  f5er [command]

Available Commands:
  add         add F5 objects
  delete      delete F5 objects
  help        Help about any command
  offline     offline a pool member
  online      online a pool member
  run         runs a bash command on the f5
  show        show F5 objects
  stats       get F5 statistics
  update      update F5 objects
  upload      upload a file
  version     show current version

Flags:
  -d, --debug          debug output
  -f, --f5 string      IP or hostname of F5 to poke
  -i, --input string   input json f5 configuration

Use "f5er [command] --help" for more information about a command.
```

For any command or subcommand, you can get additional information by using the help flags **-h|--help**.

```
./f5er show --help
show current state of F5 objects. Show requires an object, eg. f5er show pool

Usage:
  f5er show [flags]
  f5er show [command]

Available Commands:
  cert         show a certificate
  certs        show all certificates
  client-ssl   show a client-ssl profile
  device       show an f5 device
  monitor-http show a monitor-http profile
  node         show a node
  policy       show a policy
  pool         show a pool
  poolmember   show pool members
  profile      show profiles
  rule         show a rule
  server-ssl   show a server-ssl profile
  stack        show a stack transaction
  virtual      show a virtual server

Flags:
  -h, --help   help for show

Global Flags:
  -d, --debug          debug output
  -f, --f5 string      IP or hostname of F5 to poke
  -i, --input string   input json f5 configuration

Use "f5er show [command] --help" for more information about a command.
```

## Stacks

This is a convenience construct and does not exist within F5 terminology. 
It effectively allows commands to work on multiple nodes, pools, virtual servers, rules and policies in one operation. It uses a REST transaction to do so.
Look at the file stack.json to see how to structure the input file.
Show, add, update and delete operations are supported.

```
./f5er add stack -h
add a new stack

Usage: 
  f5er add stack [flags]

 Available Flags:
  -d, --debug=false: debug output
  -f, --f5="": IP or hostname of F5 to poke
  -h, --help=false: help for stack
  -i, --input="": input json f5 configuration


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


## Statistics

Query F5 statistics for LTM virtuals, pools, nodes and rules.
Output is returned in graphite format.
To set a custom graphite path prefix, use a config file and set the configurable **stats_path_prefix**. See the conf file example above.
Only statistics that are non-zero will be returned by default. Override this by setting the configurable **stats_show_zero_values** to true.

To retrieve all stats for a single virtual server...
```
./f5er stats virtual /DMZ/virtual-prd
f5.DMZ.virtual.virtual-prd.fiveMinAvgUsageRatio 0 1465879005
f5.DMZ.virtual.virtual-prd.oneMinAvgUsageRatio 0 1465879005
f5.DMZ.virtual.virtual-prd.clientside.curConns 28 1465879005
...
```

To retrieve stats for all virtual servers in one hit, don't pass a virtual as an argument...
```
./f5er stats virtual
f5.DMZ.virtual.virtualserver.Clientside_bitsIn 160632 1465880938
f5.DMZ.virtual.virtualserver.Clientside_bitsOut 67000 1465880938
f5.DMZ.virtual.virtualserver.Clientside_maxConns 2 1465880938
f5.DMZ.virtual.virtualserver.Clientside_pktsIn 228 1465880938
f5.DMZ.virtual.virtualserver.Clientside_pktsOut 105 1465880938
...

```

## Adding TLS/SSL Certificate and Keys

```
$ ./f5er upload mysite.com.crt
Uploading file mysite.com.crt
Done

$ ./f5er upload mysite.com.key
Uploading file mysite.com.key
Done

$ ./f5er add cert mysite_com PARTITION mysite.com.crt
Name: mysite_com.crt Partition: PARTITION
Issuer: CN=DigiCert SHA2 High Assurance Server CA,OU=www.digicert.com,O=DigiCert Inc,C=US
Subject: CN=www.mysite.com,OU=IS,O=Foobar,L=Salt Lake City,ST=Utah,C=US
Strength: 2048 Curve: none Type: rsa-public
Checksum: SHA1:2148:999980b62a14054893804473bda0376f44fa51c1
Uploaded: 2017-07-31T16:30:01Z Expires May 28 12:00:00 2020 GMT

$ ./f5er add key mysite_com PARTITION mysite.com.key
Name: mysite_com.key Partition: DEV
Issuer: 
Subject: 
Strength: 0 Curve:  Type: rsa-private
Checksum: SHA1:1679:1111f23b7c0a3f53a6809c4608e6c8319a46a3a9
Uploaded: 2017-07-31T16:30:18Z Expires
```

## Running Bash Commands

```
$ ./f5er run "ls -al"
total 44
drwxr-xr-x.  5 root root 4096 Jun 30  2016 .
drwxr-xr-x. 95 root root 4096 Jun 30  2016 ..
drwxr-xr-x.  2 root root 4096 Jun 30  2016 deps
drwxr-xr-x.  2 root root 4096 Jun 30  2016 requires
lrwxrwxrwx.  1 root root   31 Jun 30  2016 run -> /etc/bigstart/scripts/restjavad
drwx------.  2 root root 4096 Jun  3 23:44 supervise
```


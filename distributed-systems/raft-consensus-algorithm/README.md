# raft-consensus-algorithm

_A simple example of distributed consensus using Hashicorp raft,
an asymmetric consensus algorithm where an elected leader issues
commands to its followers._

Documentation and reference,

* [Hashicorp Raft](https://github.com/hashicorp/raft)
  at github

[GitHub Webpage](https://jeffdecola.github.io/my-go-examples/)

## PREREQUISITES

```bash
go get -u -v github.com/hashicorp/raft
go get -u -v github.com/hashicorp/raft-boltdb
```

## SOLVING DISTRIBUTED CONSENSUS - A BRIEF HISTORY

Consensus is having all the Nodes/Servers in a distributed system
agree on a Value/State. This is more easily said than done because
how do you insure everyone does the correct things.

**Blockchains** are consensus since they run calculations to agree on
the state of things.

The **Paxos** algorithm has been wildly used in the 90's, early 2000's
but is very difficult to use.

**Raft** came along to help makes this consensus easier.

## RAFT

Raft uses asymmetric consensus, meaning there is one leader who tells the other
follower Nodes what to do.

Like otherRaft solves the failures issues such as dead Nodes
(**fail-stop**) or bad Nodes (**Byzantine Failure**).

Once a leader is chosen, we need to handle "log" (state information)
replication. Each Node/Server has a list of logs.
A committed log is when it is a majority of servers.

The client only talks to the leader and the leader in tern
tells all the followers what to do.

## HASHICORP RAFT

Their are many version of RAFT written in many languages.
This example is using Hashicorp's RAFT.
The go library gives you a simple interface to use.  So lets use it.

## LETS CODE

## RUN

```bash
go run raft-consensus-algorithm.go \
    -dataDirectory ./node1 \
    -genesis \
    -loglevel debug \
    -nodehttpport 2000 \
    -nodeip 127.0.0.1 \
    -nodetcpport 3000
```

```bash
go run raft-consensus-algorithm.go \
    -dataDirectory ./node1
    -loglevel debug \
    -networkip 127.0.0.1 \
    -networktcpport 3000 \
    -nodehttpport 2001 \
    -nodeip 127.0.0.1 \
    -nodetcpport 3001
```

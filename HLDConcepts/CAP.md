## Consistency Availability and Partition Tolerance

[Article reference](https://www.ibm.com/think/topics/cap-theorem)

### What does the Consistency stands for?
-> In distributed systems once a write succeeds in one node, every node should be having the same data across and it should behave like a single copy of data exists in the entire universe.

### What does Availability stands for?
-> Every non-failing node should be returning a non-error response, it cannot timeout or any other errors. Which basically means if the node is up and running it should be responsible for a functional response without an exception being raised.

### What does Partition Tolerance stands for?
-> The network can drop, delay or reorder an arbitary number of messages between nodes and the system must still maintain it's correctness and order.

The CAP theorem states that we often have to choose the best two out of the three worlds.

## CAP theorem NoSQL database types

NoSQL databases are ideal for distributed network applications. Unlike their vertically scalable SQL (relational) counterparts, NoSQL databases are horizontally scalable and distributed by design—they can rapidly scale across a growing network consisting of multiple interconnected nodes.


Today, NoSQL databases are classified based on the two CAP characteristics they support:

**CP database:** A CP database delivers consistency and partition tolerance at the expense of availability. When a partition occurs between any two nodes, the system has to shut down the non-consistent node (i.e., make it unavailable) until the partition is resolved.

**AP database:** An AP database delivers availability and partition tolerance at the expense of consistency. When a partition occurs, all nodes remain available but those at the wrong end of a partition might return an older version of data than others. (When the partition is resolved, the AP databases typically resync the nodes to repair all inconsistencies in the system.)

**CA database:** A CA database delivers consistency and availability across all nodes. It can’t do this if there is a partition between any two nodes in the system, however, and therefore can’t deliver fault tolerance.

### MongoDB and the CAP theorem

MongoDB is a popular NoSQL database management system that stores data as BSON (binary JSON) documents. It's frequently used for big data and real-time applications running at multiple different locations. Relative to the CAP theorem, MongoDB is a CP data store—it resolves network partitions by maintaining consistency, while compromising on availability.

[MongoDB architecture]()

MongoDB is a single-master system—each replica set can have only one primary node that receives all the write operations. All other nodes in the same replica set are secondary nodes that replicate the primary node's operation log and apply it to their own data set. By default, clients also read from the primary node, but they can also specify a read preference that allows them to read from secondary nodes.

When the primary node becomes unavailable, the secondary node with the most recent operation log will be elected as the new primary node. Once all the other secondary nodes catch up with the new master, the cluster becomes available again. As clients can't make any write requests during this interval, the data remains consistent across the entire network.

### Cassandra and the CAP theorem (AP)

Apache Cassandra is an open source NoSQL database maintained by the Apache Software Foundation. It’s a wide-column database that lets you store data on a distributed network. However, unlike MongoDB, Cassandra has a masterless architecture, and as a result, it has multiple points of failure, rather than a single one.

[Cassandra Architecture]()

Relative to the CAP theorem, Cassandra is an AP database—it delivers availability and partition tolerance but can't deliver consistency all the time. Because Cassandra doesn't have a master node, all the nodes must be available continuously. However, Cassandra provides eventual consistency by allowing clients to write to any nodes at any time and reconciling inconsistencies as quickly as possible.

As data only becomes inconsistent in the case of a network partition and inconsistencies are quickly resolved, Cassandra offers “repair” functionality to help nodes catch up with their peers. However, constant availability results in a highly performant system that might be worth the trade-off in many cases.

## The PACELC Theorem (The Real-World Extension)

The CAP Theorem only talks about what happens during a network partition. But partitions are rare (ideally). What happens during normal operations when everything is running smoothly?

It states that 
- If a partition has occurred how does the system choose between Availability and Consistency.
- Else if everything is running smooth, how the the system choose between Latency and Consistency.
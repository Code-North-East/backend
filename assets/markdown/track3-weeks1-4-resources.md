# Track 3 (Weeks 1-4): HLD Building Blocks Resources

This guide outlines essential concepts and study resources for the first 4 weeks of System Design (HLD) preparation.

## Week 1: Performance, Reliability & Networking
* **Latency vs. Throughput**: Understand the difference between response time (Latency) and volume of requests (Throughput). Why p99 matters more than the average.
* **Load Balancing**:
  * **L4 (Transport)**: Routes traffic based on IP and Port (e.g., AWS NLB). Fast, but dumb.
  * **L7 (Application)**: Routes traffic based on HTTP headers, cookies, or paths (e.g., AWS ALB, NGINX). Slower, but smart.
* **CDNs (Content Delivery Networks)**: Push vs Pull models. Using CloudFront or Fastly to cache static assets close to the user.

## Week 2: Storage Systems & Data Partitioning
* **SQL vs NoSQL**:
  * **SQL**: ACID compliant, relational data, requires schema upfront. Good for financial transactions.
  * **NoSQL**: Flexible schema, horizontally scalable.
    * *Key-Value*: Redis, DynamoDB.
    * *Document*: MongoDB.
    * *Wide-Column*: Cassandra, ScyllaDB (great for time-series/heavy writes).
    * *Graph*: Neo4j (great for recommendation engines/social networks).
* **Partitioning**:
  * **Consistent Hashing**: Understand how a hash ring works to minimize data movement when a node is added or removed.

## Week 3: Consistency Models, APIs & Caching
* **CAP Theorem**: You can only pick 2 out of 3: Consistency, Availability, Partition Tolerance. (In reality, since partitions happen, you choose between CP and AP).
* **API Design**:
  * **REST**: Resource-based, standard HTTP verbs.
  * **gRPC**: RPC-based, uses Protocol Buffers, operates over HTTP/2, highly efficient for microservice-to-microservice communication.
  * **WebSockets**: Persistent, bi-directional communication (ideal for chat apps).
* **Caching Strategies**:
  * **Cache-Aside**: Application checks cache, if miss, checks DB, updates cache.
  * **Write-Through**: Application writes to cache, cache writes to DB synchronously.

## Week 4: Asynchronous Processing & Integration
* **Message Queues vs Streams**:
  * **Queues (RabbitMQ / SQS)**: Messages are deleted once consumed. Ideal for task processing (e.g., sending emails).
  * **Streams (Kafka / Kinesis)**: Messages are appended to a log and stay there. Multiple consumers can read the same stream at different paces. Ideal for event sourcing or analytics.
* **Delivery Semantics**:
  * **At-most-once**: Fire and forget.
  * **At-least-once**: Retries on failure, requires consumers to be **idempotent**.
  * **Exactly-once**: Hardest to achieve, often requires specialized frameworks (Kafka Transactions).


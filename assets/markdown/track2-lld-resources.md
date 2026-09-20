# Track 2: LLD Architectural Hints & Resources

This document provides hints, recommended patterns, and resources for the 8 LLD builds required in Track 2.

### 1. URL Shortener
- **Key Concepts:** Base62 encoding, MD5 hashing, collision handling.
- **Hints:** Do not use sequential IDs in a distributed environment (security/predictability risks). Think about how to handle high read volume (Cache-Aside pattern).
- **Go implementation hint:** Use `crypto/md5` and a fast router like Chi.

### 2. Rate Limiter
- **Key Concepts:** Token Bucket, Leaky Bucket, Sliding Window Log.
- **Hints:** Focus on thread safety. If a map is used for per-user limits, ensure it has a `sync.Mutex` in Go or `ConcurrentHashMap` in Java.
- **Java implementation hint:** Try building this as a Spring Boot `HandlerInterceptor` or Servlet Filter.

### 3. LRU Cache
- **Key Concepts:** Hash Map + Doubly Linked List.
- **Hints:** Make sure `get` and `put` are both $O(1)$. For concurrency, consider read/write locks (`sync.RWMutex` in Go or `ReentrantReadWriteLock` in Java).

### 4. Parking Lot
- **Key Concepts:** Object-Oriented Design, State Machine.
- **Hints:** Start with the entities: `ParkingLot`, `Level`, `ParkingSpot`, `Vehicle` (Car, Truck, Motorcycle). Use the Strategy Pattern for calculating the parking fee.
- **SOLID check:** Ensure adding a new vehicle type doesn't require modifying the core `ParkingSpot` logic (Open/Closed Principle).

### 5. Notification Dispatcher
- **Key Concepts:** Strategy Pattern, Observer Pattern, Queueing.
- **Hints:** Abstract the sender into a `NotificationSender` interface. Implement `EmailSender` and `SmsSender`. Add a retry queue for failed messages.

### 6. Splitwise Ledger
- **Key Concepts:** Graphs, Graph traversal, greedy algorithms.
- **Hints:** To simplify debts, represent balances as a directed graph. Create a max-heap for creditors and a max-heap for debtors to settle amounts efficiently.

### 7. Chat System
- **Key Concepts:** WebSockets, TCP connections, Pub/Sub.
- **Hints:** Keep an in-memory registry of active WebSocket connections. When a message is sent to a group, fan it out to all active socket connections for users in that group.

### 8. Inventory Management
- **Key Concepts:** ACID transactions, Concurrency Control (Optimistic vs Pessimistic locking).
- **Hints:** Solve the double-booking problem. If two users try to buy the last item simultaneously, only one should succeed.
- **Database hint:** Use `SELECT ... FOR UPDATE` (pessimistic) or a `version` column (optimistic).


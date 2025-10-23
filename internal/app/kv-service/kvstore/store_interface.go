package kvstore

// Store defines a generic key-value storage interface.
// The current implementation uses an in-memory store (InMemoryStore),
// but this interface allows future extensions — for example,
// replacing it with Redis, Memcached, or a database-backed implementation.
type Store interface {
	Put(key string, value string)
	Get(key string) (string, bool)
	Delete(key string) bool
}

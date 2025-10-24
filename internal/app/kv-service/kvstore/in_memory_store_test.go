package kvstore

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	// Define globally as we run tests in parallel to test the concurrency.
	// Ideally we need to have test cases to test this
	store = NewInMemoryStore()
)

func TestPut(t *testing.T) {
	t.Parallel()
	store.Put("k1", "v1")

	v, ok := store.Get("k1")
	assert.True(t, ok)
	assert.Equal(t, "v1", v)
}

func TestGet(t *testing.T) {
	t.Parallel()

	// missing key
	if v, ok := store.Get("missing"); ok {
		assert.Failf(t, "expected missing key to be absent", "got %q", v)
	}

	// present key
	store.Put("k2", "v2")
	v, ok := store.Get("k2")
	assert.True(t, ok)
	assert.Equal(t, "v2", v)
}

func TestDelete(t *testing.T) {
	t.Parallel()

	// delete non-existent
	assert.False(t, store.Delete("nope"))

	// delete existing
	store.Put("k3", "v3")
	assert.True(t, store.Delete("k3"))
	if _, ok := store.Get("k3"); ok {
		assert.Fail(t, "expected key to be gone after Delete")
	}
}

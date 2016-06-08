package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLast(t *testing.T) {
	assert.Equal(t, last([]string{"a", "b", "c"}), "c")
}

// Redirect Host

// If subdomain is www redirect to root domain.
func TestRootRedirect(t *testing.T) {
	assert.Equal(t, getRedirectHost("www.foo.com"), "foo.com")
}

// Otherwise redirect to www subdomain.
func TestWwwRedirect(t *testing.T) {
	assert.Equal(t, getRedirectHost("foo.com"), "www.foo.com")
	assert.Equal(t, getRedirectHost("any.foo.com"), "www.foo.com")
	assert.Equal(t, getRedirectHost("x.y.z.foo.com"), "www.foo.com")
}

// Some registries don't allow 2th level domains.
func TestBrRegistry(t *testing.T) {
	// .com.br
	assert.Equal(t, getRedirectHost("www.foo.com.br"), "foo.com.br")
	assert.Equal(t, getRedirectHost("foo.com.br"), "www.foo.com.br")
	assert.Equal(t, getRedirectHost("any.foo.com.br"), "www.foo.com.br")
	assert.Equal(t, getRedirectHost("x.y.z.foo.com.br"), "www.foo.com.br")
}

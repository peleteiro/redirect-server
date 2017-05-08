package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// Redirect Host

// If subdomain is www redirect to root domain.
func TestRootRedirect(t *testing.T) {
	assert.Equal(t, getRedirectHost("www.foo.com"), "foo.com")
}

// Otherwise redirect to www subdomain.
func TestWwwRedirect(t *testing.T) {
	assert.Equal(t, getRedirectHostByHost("foo.com"), "www.foo.com")
	assert.Equal(t, getRedirectHostByHost("any.foo.com"), "www.foo.com")
	assert.Equal(t, getRedirectHostByHost("x.y.z.foo.com"), "www.foo.com")
}

// Some registries don't allow 2th level domains.
func TestBrRegistry(t *testing.T) {
	// .com.br
	assert.Equal(t, getRedirectHostByHost("www.foo.com.br"), "foo.com.br")
	assert.Equal(t, getRedirectHostByHost("foo.com.br"), "www.foo.com.br")
	assert.Equal(t, getRedirectHostByHost("any.foo.com.br"), "www.foo.com.br")
	assert.Equal(t, getRedirectHostByHost("x.y.z.foo.com.br"), "www.foo.com.br")
}

// Redirect by CNAME

func TestGetRedirectHostByCNAME(t *testing.T) {
	assert.Equal(t, "", GetRedirectHostByCNAME("doesnotexists.com"))
	assert.Equal(t, "www.biblebox.com", GetRedirectHostByCNAME("www.biblebox.com.foo.com."))
}

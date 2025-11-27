package scraper

import (
	"net/url"
	"testing"
)

func TestResolveURL(t *testing.T) {
	base, _ := url.Parse("http://example.com")
	resolved := resolveURL(base, "/path")
	expected := "http://example.com/path"
	if resolved != expected {
		t.Errorf("expected %s, got %s", expected, resolved)
	}
}

func TestStripQueryParams(t *testing.T) {
	stripped := stripQueryParams("http://example.com?foo=bar")
	expected := "http://example.com"
	if stripped != expected {
		t.Errorf("expected %s, got %s", expected, stripped)
	}
}

func TestIsSameDomain(t *testing.T) {
	base, _ := url.Parse("http://example.com")
	if !isSameDomain(base, "http://example.com/path") {
		t.Errorf("expected true, got false")
	}
	if isSameDomain(base, "http://another.com/path") {
		t.Errorf("expected false, got true")
	}
}

func TestIsValidLink(t *testing.T) {
	s := &scraper{}
	if !s.isValidLink("/path") {
		t.Errorf("expected true, got false")
	}
	if s.isValidLink("#") {
		t.Errorf("expected false, got true")
	}
	if s.isValidLink("http://example.com/foo.pdf") {
		t.Errorf("expected false, got true")
	}
}

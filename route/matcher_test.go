package route

import (
	"testing"

	"github.com/gobwas/glob"
)

func TestPrefixMatcher(t *testing.T) {
	tests := []struct {
		uri     string
		matches bool
		route   *Route
	}{
		{uri: "/foo", matches: true, route: &Route{Path: "/foo"}},
		{uri: "/fools", matches: true, route: &Route{Path: "/foo"}},
		{uri: "/fo", matches: false, route: &Route{Path: "/foo"}},
		{uri: "/bar", matches: false, route: &Route{Path: "/foo"}},
	}

	for _, tt := range tests {
		t.Run(tt.uri, func(t *testing.T) {
			if got, want := prefixMatcher(tt.uri, tt.route), tt.matches; got != want {
				t.Fatalf("got %v want %v", got, want)
			}
		})
	}
}

func TestGlobMatcher(t *testing.T) {
	tests := []struct {
		uri     string
		matches bool
		route   *Route
	}{
		// happy flows
		{uri: "/foo", matches: true, route: getRoute("/foo")},
		{uri: "/fool", matches: true, route: getRoute("/foo?")},
		{uri: "/fool", matches: true, route: getRoute("/foo*")},
		{uri: "/fools", matches: true, route: getRoute("/foo*")},
		{uri: "/fools", matches: true, route: getRoute("/foo*")},
		{uri: "/foo/x/bar", matches: true, route: getRoute("/foo/*/bar")},
		{uri: "/foo/x/y/z/w/bar", matches: true, route: getRoute("/foo/**")},
		{uri: "/foo/x/y/z/w/bar", matches: true, route: getRoute("/foo/**/bar")},

		// error flows
		{uri: "/fo", matches: false, route: getRoute("/foo")},
		{uri: "/fools", matches: false, route: getRoute("/foo")},
		{uri: "/fo", matches: false, route: getRoute("/foo*")},
		{uri: "/fools", matches: false, route: getRoute("/foo.*")},
		{uri: "/foo/x/y/z/w/baz", matches: false, route: getRoute("/foo/**/bar")},
	}

	for _, tt := range tests {
		t.Run(tt.uri, func(t *testing.T) {
			if got, want := globMatcher(tt.uri, tt.route), tt.matches; got != want {
				t.Fatalf("got %v want %v", got, want)
			}
		})
	}
}

func getRoute(path string) *Route {
	return &Route{Path: path, Glob: glob.MustCompile(path)}
}

package core

import "testing"

func mustVersion(t *testing.T, v string) *Version {
	t.Helper()
	version, err := createVersion(v)
	if err != nil {
		t.Fatalf("createVersion(%q) returned error: %v", v, err)
	}
	return version
}

// assertGreater fails if next is not strictly greater than current per semver
// precedence. Every increment must move forward: the reported bug was an
// increment that produced an already-existing, lower-or-equal version.
func assertGreater(t *testing.T, current, next *Version) {
	t.Helper()
	if !next.semverVersion.GreaterThan(current.semverVersion) {
		t.Errorf("expected %s to be strictly greater than %s", next, current)
	}
}

func TestIncrementPatch(t *testing.T) {
	tests := []struct {
		name    string
		current string
		want    string
	}{
		{"plain patch bump", "v28.3.1", "v28.3.2"},
		{"zero patch", "v1.0.0", "v1.0.1"},
		// Regression test for https://github.com/Wulfheart/release-ranger/issues/1:
		// a prerelease must finalise to the next patch, not to the already
		// existing v28.3.1 that it is a prerelease of.
		{"prerelease finalises to next patch", "v28.3.1-rc2000.1.0", "v28.3.2"},
		{"simple prerelease", "v1.2.3-rc.0", "v1.2.4"},
		{"prerelease with build metadata", "v1.2.3-rc.0+build.5", "v1.2.4"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			current := mustVersion(t, tt.current)
			next := current.IncrementPatch()
			if got := next.String(); got != tt.want {
				t.Errorf("IncrementPatch(%q) = %q, want %q", tt.current, got, tt.want)
			}
			assertGreater(t, current, next)
		})
	}
}

func TestIncrementPrerelease(t *testing.T) {
	tests := []struct {
		name    string
		current string
		want    string
	}{
		{"bump trailing identifier", "v28.3.1-rc2000.1.0", "v28.3.1-rc2000.1.1"},
		{"bump single identifier", "v1.2.3-rc.0", "v1.2.3-rc.1"},
		// Numeric identifiers are compared numerically, not lexically, so
		// rc.9 -> rc.10 must still sort forward (asserted below).
		{"numeric identifier past nine", "v1.2.3-rc.9", "v1.2.3-rc.10"},
		{"numeric only prerelease", "v1.2.3-1", "v1.2.3-2"},
		{"bumps last numeric identifier", "v1.2.3-alpha.2.beta", "v1.2.3-alpha.3.beta"},
		{"append when no numeric identifier", "v1.2.3-rc", "v1.2.3-rc.1"},
		{"start rc from a release", "v28.3.1", "v28.3.2-rc.0"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			current := mustVersion(t, tt.current)
			next := current.IncrementPrerelease()
			if got := next.String(); got != tt.want {
				t.Errorf("IncrementPrerelease(%q) = %q, want %q", tt.current, got, tt.want)
			}
			assertGreater(t, current, next)
		})
	}
}

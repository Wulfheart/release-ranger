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

func TestIncrementPatch(t *testing.T) {
	tests := []struct {
		name    string
		current string
		want    string
	}{
		{"plain patch bump", "v28.3.1", "v28.3.2"},
		// Regression test for https://github.com/Wulfheart/release-ranger/issues/1:
		// a prerelease must finalise to the next patch, not to the already
		// existing v28.3.1 that it is a prerelease of.
		{"prerelease finalises to next patch", "v28.3.1-rc2000.1.0", "v28.3.2"},
		{"simple prerelease", "v1.2.3-rc.0", "v1.2.4"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mustVersion(t, tt.current).IncrementPatch().String()
			if got != tt.want {
				t.Errorf("IncrementPatch(%q) = %q, want %q", tt.current, got, tt.want)
			}
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
		{"append when no numeric identifier", "v1.2.3-rc", "v1.2.3-rc.1"},
		{"start rc from a release", "v28.3.1", "v28.3.2-rc.0"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mustVersion(t, tt.current).IncrementPrerelease().String()
			if got != tt.want {
				t.Errorf("IncrementPrerelease(%q) = %q, want %q", tt.current, got, tt.want)
			}
		})
	}
}

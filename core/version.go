package core

import (
	"strconv"
	"strings"

	"github.com/Masterminds/semver"
)

type Version struct {
	semverVersion *semver.Version
}

func createVersion(version string) (*Version, error) {
	semverVersion, err := semver.NewVersion(version)

	if err != nil {
		return nil, err
	}
	return &Version{semverVersion: semverVersion}, nil
}

func (v Version) IncrementMajor() *Version {
	newVersion := v.semverVersion.IncMajor()
	return &Version{semverVersion: &newVersion}
}

func (v Version) IncrementMinor() *Version {
	newVersion := v.semverVersion.IncMinor()
	return &Version{semverVersion: &newVersion}
}

// IncrementPatch produces the next patch release.
//
// semver's IncPatch keeps the patch number when the current version carries a
// prerelease (e.g. v28.3.1-rc2000.1.0 -> v28.3.1), which collides with the
// already released version it is a prerelease of. To avoid re-emitting an
// existing tag we finalise a prerelease to the *next* patch release instead
// (v28.3.1-rc2000.1.0 -> v28.3.2).
func (v Version) IncrementPatch() *Version {
	sv := v.semverVersion
	if sv.Prerelease() != "" {
		released, _ := sv.SetPrerelease("")
		newVersion := released.IncPatch()
		return &Version{semverVersion: &newVersion}
	}
	newVersion := sv.IncPatch()
	return &Version{semverVersion: &newVersion}
}

// IncrementPrerelease produces the next release candidate.
//
// When the current version already carries a prerelease its trailing numeric
// identifier is incremented (v28.3.1-rc2000.1.0 -> v28.3.1-rc2000.1.1). When it
// does not, a fresh release candidate is started for the next patch release
// (v28.3.1 -> v28.3.2-rc.0).
func (v Version) IncrementPrerelease() *Version {
	sv := v.semverVersion
	pre := sv.Prerelease()

	if pre == "" {
		next := sv.IncPatch()
		withPre, _ := next.SetPrerelease("rc.0")
		return &Version{semverVersion: &withPre}
	}

	newVersion, _ := sv.SetPrerelease(incrementPrerelease(pre))
	return &Version{semverVersion: &newVersion}
}

// incrementPrerelease bumps the last dot-separated numeric identifier of a
// prerelease string. If no trailing numeric identifier exists a ".1" is
// appended so the result still sorts after the current version.
func incrementPrerelease(pre string) string {
	identifiers := strings.Split(pre, ".")
	for i := len(identifiers) - 1; i >= 0; i-- {
		if n, err := strconv.Atoi(identifiers[i]); err == nil {
			identifiers[i] = strconv.Itoa(n + 1)
			return strings.Join(identifiers, ".")
		}
	}
	return pre + ".1"
}

func (v Version) String() string {
	return "v" + v.semverVersion.String()
}

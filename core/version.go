package core

import "github.com/Masterminds/semver"

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

func (v Version) IncrementPatch() *Version {
	newVersion := v.semverVersion.IncPatch()
	return &Version{semverVersion: &newVersion}
}

func (v Version) String() string {
	return "v" + v.semverVersion.String()
}

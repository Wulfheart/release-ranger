package core

import (
	"os/exec"
	"strings"
)

type SortedVersionCollection []*Version

type Releaser interface {
	Retrieve() (SortedVersionCollection, error)
	Create(version *Version) error
}

type GitReleaser struct {
}

func (g GitReleaser) Retrieve() (SortedVersionCollection, error) {
	cmd := exec.Command("git", "tag", "--sort=-v:refname")

	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	versions := strings.Split(string(output), "\n")

	versionCollection := make(SortedVersionCollection, 0)

	for _, version := range versions {
		version = strings.TrimSpace(version)
		if version == "" {
			continue
		}
		sv, err := createVersion(version)
		if err != nil {
			return nil, err
		}
		versionCollection = append(versionCollection, sv)
	}
	if len(versionCollection) == 0 {
		v0, err := createVersion("v0.0.0")
		if err != nil {
			return nil, err
		}
		versionCollection = append(versionCollection, v0)
	}
	return versionCollection, nil
}

func (g GitReleaser) Create(version *Version) error {
	cmd := exec.Command("git", "tag", "-a", version.String(), "-m", version.String())
	_, err := cmd.Output()
	if err != nil {
		return err
	}
	cmd = exec.Command("git", "push", "origin", version.String())
	_, err = cmd.Output()
	return err
}

func (g GitReleaser) fetch() error {
	cmd := exec.Command("git", "fetch", "--tags")
	_, err := cmd.Output()

	return err
}

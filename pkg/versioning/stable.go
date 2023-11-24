package versioning

import (
	"github.com/Masterminds/semver/v3"
)

type stableVersionStrategy struct{}

// Stable upgrade strategy accepts all newer versions except prerelease.
func Stable() Strategy {
	return &stableVersionStrategy{}
}

func (s *stableVersionStrategy) IsAcceptable(version *semver.Version) bool {
	return len(version.Prerelease()) == 0
}

func (s *stableVersionStrategy) IsUpgrade(oldV *semver.Version, newV *semver.Version) bool {
	return s.IsAcceptable(newV) && oldV.LessThan(newV)
}

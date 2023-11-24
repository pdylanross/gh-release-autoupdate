package versioning

import (
	"github.com/Masterminds/semver/v3"
)

type alwaysUpgradeStrategy struct{}

// Always strategy accepts all versions of a higher semantic version.
func Always() Strategy {
	return &alwaysUpgradeStrategy{}
}

func (a *alwaysUpgradeStrategy) IsAcceptable(_ *semver.Version) bool {
	return true
}

func (a *alwaysUpgradeStrategy) IsUpgrade(oldV *semver.Version, newV *semver.Version) bool {
	return oldV.LessThan(newV)
}

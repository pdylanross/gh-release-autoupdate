package autoupdate

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/pdylanross/gh-release-autoupdate/autoupdate/types"
)

// Stable upgrade strategy accepts all newer versions except prerelease.
func Stable() types.VersioningStrategy {
	return &stableVersionStrategy{}
}

type stableVersionStrategy struct{}

func (s *stableVersionStrategy) IsAcceptable(version *semver.Version) bool {
	return len(version.Prerelease()) == 0
}

func (s *stableVersionStrategy) IsUpgrade(oldV *semver.Version, newV *semver.Version) bool {
	return s.IsAcceptable(newV) && oldV.LessThan(newV)
}

// Constrained upgrade strategies accept newer versions with a constraint.
func Constrained(c string) (types.VersioningStrategy, error) {
	constraint, err := semver.NewConstraint(c)
	if err != nil {
		return nil, err
	}

	return &constrainedVersionStrategy{constraint: constraint}, nil
}

// ConstrainMajor creates a constrained version strategy pinned to a major version.
func ConstrainMajor(version *semver.Version, allowPrerelease bool) types.VersioningStrategy {
	pre := ""
	if allowPrerelease {
		pre = "-0"
	}

	start, err := Constrained(fmt.Sprintf("^%d%s", version.Major(), pre))
	if err != nil {
		panic(err)
	}

	return start
}

// ConstrainMinor creates a constrained version strategy pinned to a minor version.
func ConstrainMinor(version *semver.Version, allowPrerelease bool) types.VersioningStrategy {
	pre := ""
	if allowPrerelease {
		pre = "-0"
	}

	start, err := Constrained(fmt.Sprintf("~%d.%d.%d%s", version.Major(), version.Minor(), version.Patch(), pre))
	if err != nil {
		panic(err)
	}

	return start
}

type constrainedVersionStrategy struct {
	constraint *semver.Constraints
}

func (c *constrainedVersionStrategy) IsAcceptable(version *semver.Version) bool {
	return c.constraint.Check(version)
}

func (c *constrainedVersionStrategy) IsUpgrade(oldV *semver.Version, newV *semver.Version) bool {
	return c.IsAcceptable(newV) && oldV.LessThan(newV)
}

// Always strategy accepts all versions of a higher semantic version.
func Always() types.VersioningStrategy {
	return &alwaysUpgradeStrategy{}
}

type alwaysUpgradeStrategy struct{}

func (a *alwaysUpgradeStrategy) IsAcceptable(_ *semver.Version) bool {
	return true
}

func (a *alwaysUpgradeStrategy) IsUpgrade(oldV *semver.Version, newV *semver.Version) bool {
	return oldV.LessThan(newV)
}

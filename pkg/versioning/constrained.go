package versioning

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
)

type constrainedVersionStrategy struct {
	constraint *semver.Constraints
}

// Constrained upgrade strategies accept newer versions with a constraint.
func Constrained(c string) (Strategy, error) {
	constraint, err := semver.NewConstraint(c)
	if err != nil {
		return nil, err
	}

	return &constrainedVersionStrategy{constraint: constraint}, nil
}

func ConstrainMajor(version *semver.Version, allowPrerelease bool) Strategy {
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

func ConstrainPatch(version *semver.Version, allowPrerelease bool) Strategy {
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

func (c *constrainedVersionStrategy) IsAcceptable(version *semver.Version) bool {
	return c.constraint.Check(version)
}

func (c *constrainedVersionStrategy) IsUpgrade(oldV *semver.Version, newV *semver.Version) bool {
	return c.IsAcceptable(newV) && oldV.LessThan(newV)
}

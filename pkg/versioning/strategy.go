package versioning

import "github.com/Masterminds/semver/v3"

// Strategy defines if a version is an upgrade from a prior version.
type Strategy interface {
	// IsAcceptable determines if the new version is a an accepted version (IE don't accept prerelease)
	IsAcceptable(version *semver.Version) bool
	// IsUpgrade determines if a new version is an upgrade from an old version
	IsUpgrade(oldV *semver.Version, newV *semver.Version) bool
}

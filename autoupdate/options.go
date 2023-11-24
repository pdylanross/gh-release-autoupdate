package autoupdate

import (
	"errors"
	"fmt"
	"strings"

	"github.com/pdylanross/gh-release-autoupdate/autoupdate/types"
	"github.com/pdylanross/gh-release-autoupdate/internal/gh"
)

// UpdaterOptions configures the updater with settings on how to update and metadata about the app.
type UpdaterOptions struct {
	// PackageName is the name of the app being updated
	PackageName string
	// PackageVersion is the current version of the app
	PackageVersion string

	// RepoOwner is the github owner of the package
	RepoOwner string
	// RepoName is the name of the repository
	RepoName string

	// Cache sets up the update cache. If nil, caching is disabled
	Cache *CacheOptions

	// Github sets the options for the github api
	Github *gh.GithubClientOptions

	// VersionStrategy determines how we look at versions and determine if they're valid upgrades
	VersionStrategy types.VersioningStrategy
}

// DefaultOptions creates an opinionated default set of options.
func DefaultOptions() *UpdaterOptions {
	return &UpdaterOptions{
		Cache:           DefaultCacheOptions(),
		VersionStrategy: Stable(),
	}
}

func (uo *UpdaterOptions) Validate() error {
	var errs []error
	if strings.TrimSpace(uo.PackageName) == "" {
		errs = append(errs, fmt.Errorf("PackageName must be set"))
	}

	if strings.TrimSpace(uo.PackageVersion) == "" {
		errs = append(errs, fmt.Errorf("PackageVersion must be set"))
	}

	if strings.TrimSpace(uo.RepoOwner) == "" {
		errs = append(errs, fmt.Errorf("RepoOwner must be set"))
	}
	if strings.TrimSpace(uo.RepoName) == "" {
		errs = append(errs, fmt.Errorf("RepoName must be set"))
	}

	if len(errs) != 0 {
		err := errors.Join(errs...)
		return errors.Join(errors.New("misconfigured updater options"), err)
	}

	return nil
}

func (uo *UpdaterOptions) GetGithubOpts() *gh.GithubClientOptions {
	if uo.Github == nil {
		return gh.NewOptions(uo.PackageName, uo.PackageVersion)
	}

	return uo.Github
}

func (uo *UpdaterOptions) GetStrategy() types.VersioningStrategy {
	return uo.VersionStrategy
}

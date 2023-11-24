package autoupdate

import (
	"github.com/pdylanross/gh-release-autoupdate/autoupdate/types"
)

// WithPackage sets the package name and version of the app.
func (uo *UpdaterOptions) WithPackage(name string, version string) *UpdaterOptions {
	uo.PackageName = name
	uo.PackageVersion = version
	if uo.Github != nil {
		uo.Github.WithApp(name, version)
	}
	return uo
}

// WithRepo sets up the github repo owner & name.
func (uo *UpdaterOptions) WithRepo(owner string, name string) *UpdaterOptions {
	uo.RepoOwner = owner
	uo.RepoName = name
	return uo
}

// WithVersionStrategy sets the types strategy.
func (uo *UpdaterOptions) WithVersionStrategy(start types.VersioningStrategy) *UpdaterOptions {
	uo.VersionStrategy = start
	return uo
}

// WithoutCache disables the update check cache.
func (uo *UpdaterOptions) WithoutCache() *UpdaterOptions {
	uo.Cache = nil
	return uo
}

// ConfigureCache configures the update check caching options.
func (uo *UpdaterOptions) ConfigureCache(f func(options *CacheOptions) *CacheOptions) *UpdaterOptions {
	if uo.Cache == nil {
		uo.Cache = DefaultCacheOptions()
	}

	uo.Cache = f(uo.Cache)
	return uo
}

// ConfigureGithub configures the github api settings.
func (uo *UpdaterOptions) ConfigureGithub(f func(option *types.GithubClientOptions) *types.GithubClientOptions) *UpdaterOptions {
	if uo.Github == nil {
		uo.Github = uo.defaultGithupOpts()
	}

	uo.Github = f(uo.Github)

	return uo
}

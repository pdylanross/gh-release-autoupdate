package autoupdate

import (
	"context"

	"github.com/pdylanross/gh-release-autoupdate/autoupdate/types"

	"github.com/pdylanross/gh-release-autoupdate/internal/release"

	"github.com/reugn/async"
)

// Updater is the main entrypoint to the auto update process.
type Updater struct {
	opts *UpdaterOptions
}

// NewUpdater creates a new updater with options.
func NewUpdater(opts *UpdaterOptions) (*Updater, error) {
	if err := opts.Validate(); err != nil {
		return nil, err
	}

	return &Updater{opts: opts}, nil
}

// Check for new versions.
func (u *Updater) Check(ctx context.Context) (*types.ReleaseCandidate, error) {
	checker, err := release.NewResolver(u.opts.GetGithubOpts(), u.opts.GetStrategy())
	if err != nil {
		return nil, err
	}

	return checker.Resolve(ctx, u.opts.RepoOwner, u.opts.RepoName, u.opts.PackageVersion)
}

// CheckDeferred checks for new versions in a separate goroutine.
func (u *Updater) CheckDeferred(ctx context.Context) async.Future[types.ReleaseCandidate] {
	promise := async.NewPromise[types.ReleaseCandidate]()

	go func() {
		res, err := u.Check(ctx)
		if err != nil {
			promise.Failure(err)
		} else {
			promise.Success(res)
		}
	}()

	return promise.Future()
}

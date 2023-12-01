package release

import (
	"context"
	"strings"

	"github.com/pdylanross/gh-release-autoupdate/autoupdate/types"
	"github.com/pdylanross/gh-release-autoupdate/internal/gh"

	"github.com/Masterminds/semver/v3"
	"github.com/google/go-github/v56/github"
)

type Resolver struct {
	ghClient *github.Client
	strategy types.VersioningStrategy
}

func NewResolver(ghClient *github.Client, strategy types.VersioningStrategy) (*Resolver, error) {
	return &Resolver{ghClient: ghClient, strategy: strategy}, nil
}

func (ch *Resolver) newReleaseCandidate(ctx context.Context, owner string, repo string, release *github.RepositoryRelease) (*types.ReleaseCandidate, error) {
	pager, err := gh.NewPager(ch.ghClient, func(gh *github.Client, page *github.ListOptions) ([]*github.ReleaseAsset, *github.Response, error) {
		return gh.Repositories.ListReleaseAssets(ctx, owner, repo, release.GetID(), page)
	})

	if err != nil {
		return nil, err
	}

	assets, err := pager.GetAll()
	if err != nil {
		return nil, err
	}

	rcAssets := make([]types.ReleaseCandidateAsset, len(assets))
	for _, asset := range assets {
		rcAssets = append(rcAssets, types.ReleaseCandidateAsset{
			ID:   asset.GetID(),
			Name: asset.GetName(),
		})
	}

	return &types.ReleaseCandidate{ID: release.GetID(), Name: release.GetName(), Assets: rcAssets, Owner: owner, Repository: repo}, nil
}

func (ch *Resolver) Resolve(ctx context.Context, repoOwner string, repoName string, currentVersion string) (*types.ReleaseCandidate, error) {
	currentVersionSemver, err := semver.NewVersion(currentVersion)
	if err != nil {
		return nil, err
	}

	// ASSUMPTION: the gh release API responds in descending order
	// ASSUMPTION: release versions always increment - a newer release is always a newer version
	// ASSUMPTION: we aren't interested in older versions
	// THEREFORE: we only have to walk this api until we find a suitable release candidate
	//            or start seeing versions older/equal to the current one
	//
	// thought: is there a use case for detecting if the current version is yanked
	// & if that's a special case for possibly downgrading? YAGNI for now
	pager, err := gh.NewPager(ch.ghClient, func(gh *github.Client, page *github.ListOptions) ([]*github.RepositoryRelease, *github.Response, error) {
		return gh.Repositories.ListReleases(ctx, repoOwner, repoName, page)
	})
	if err != nil {
		return nil, err
	}

	for {
		page, done, err := pager.NextPage()

		if err != nil {
			return nil, err
		}

		if resp, downgrade, err := ch.checkPage(ctx, page, repoOwner, repoName, currentVersionSemver); err != nil {
			return nil, err
		} else if resp != nil {
			return resp, err
		} else if downgrade {
			return nil, nil
		}

		if done {
			return nil, nil
		}
	}
}

func (ch *Resolver) checkPage(ctx context.Context, page []*github.RepositoryRelease, repoOwner string, repoName string, currentVersion *semver.Version) (*types.ReleaseCandidate, bool, error) {
	if err := ctx.Err(); err != nil {
		return nil, false, err
	}

	for _, rel := range page {
		if resp, downgrade, err := ch.checkReleaseItem(ctx, rel, repoOwner, repoName, currentVersion); err != nil {
			return nil, false, err
		} else if resp != nil {
			return resp, false, nil
		} else if downgrade {
			return nil, true, nil
		}
	}

	return nil, false, nil
}

func (ch *Resolver) checkReleaseItem(ctx context.Context, item *github.RepositoryRelease, repoOwner string, repoName string, currentVersion *semver.Version) (*types.ReleaseCandidate, bool, error) {
	if err := ctx.Err(); err != nil {
		return nil, false, err
	}

	itemVersionString := strings.TrimSpace(*item.Name)
	// todo: errors here shouldn't prevent looking at more releases
	itemVersion, err := semver.NewVersion(itemVersionString)
	if err != nil {
		return nil, false, err
	}

	if itemVersion.LessThan(currentVersion) || itemVersion.Equal(currentVersion) {
		return nil, true, nil
	}

	if !ch.strategy.IsAcceptable(itemVersion) {
		return nil, false, nil
	}

	if ch.strategy.IsUpgrade(currentVersion, itemVersion) {
		resp, err := ch.newReleaseCandidate(ctx, repoOwner, repoName, item)
		if err != nil {
			return nil, false, err
		}

		return resp, false, nil
	}

	return nil, false, nil
}

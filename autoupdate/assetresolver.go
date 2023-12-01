package autoupdate

import (
	"fmt"
	"regexp"
	"runtime"
	"strings"

	"github.com/pdylanross/gh-release-autoupdate/autoupdate/types"
)

// GoReleaserAssetResolver creates a ReleaseAssetResolver designed to work with GoReleaser
// based github releases
// TODO: This works well with the standard goreleaser template,
// but this needs to be able to support the full range of syntax that goreleaser can handle
// this will likely need to be completely rewritten.
func GoReleaserAssetResolver() types.ReleaseAssetResolver {
	os := runtime.GOOS
	arch := runtime.GOARCH

	return &goreleaserAssetResolver{os: os, arch: arch}
}

type goreleaserAssetResolver struct {
	os   string
	arch string
}

func (g *goreleaserAssetResolver) ResolveAsset(packageName string, candidate *types.ReleaseCandidate) *types.ReleaseCandidateAsset {
	// four capture groups
	// expecting format here of {{package name}}_{{os}}_{{arch}}{{archive format}}
	releaseNameRe, err := regexp.Compile(fmt.Sprintf("(%s)_([a-z]+)_([a-z0-9_]+)(.[a-z.]+)", packageName))
	if err != nil {
		panic(err)
	}

	for idx := range candidate.Assets {
		asset := &candidate.Assets[idx]
		assetName := strings.ToLower(asset.Name)

		matches := releaseNameRe.FindStringSubmatch(assetName)
		if len(matches) != 5 {
			continue
		}

		assetOs := matches[2]
		assetArch := matches[3]
		archive := matches[4]

		if assetOs == g.os && g.archEquals(g.arch, assetArch) && g.isUsableAssetArchive(archive) {
			return asset
		}
	}

	return nil
}

func (g *goreleaserAssetResolver) archEquals(current string, assetArch string) bool {
	if current == "amd64" {
		return assetArch == "amd64" || assetArch == "x86_64"
	} else if current == "386" {
		return assetArch == "386" || assetArch == "i386"
	} else if current == "arm" {
		return assetArch == "arm" || assetArch == "armv7"
	}

	return current == assetArch
}

func (g *goreleaserAssetResolver) isUsableAssetArchive(archive string) bool {
	return !strings.HasSuffix(archive, ".sbom")
}

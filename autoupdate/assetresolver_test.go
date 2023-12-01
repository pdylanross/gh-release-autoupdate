package autoupdate

import (
	"fmt"
	"testing"

	"github.com/pdylanross/gh-release-autoupdate/autoupdate/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGoReleaserAssetResolver(t *testing.T) {
	tests := map[string]map[string]string{
		"darwin": {
			// "386":   "goreleaser_Darwin_all.tar.gz", ?
			"amd64": "goreleaser_Darwin_x86_64.tar.gz",
			// "arm":   "goreleaser_Darwin_all.tar.gz", ?
			"arm64": "goreleaser_Darwin_arm64.tar.gz",
		},
		"linux": {
			"386":   "goreleaser_Linux_i386.tar.gz",
			"amd64": "goreleaser_Linux_x86_64.tar.gz",
			"arm":   "goreleaser_Linux_armv7.tar.gz",
			"arm64": "goreleaser_Linux_arm64.tar.gz",
		},
		"windows": {
			"386":   "goreleaser_Windows_i386.zip",
			"amd64": "goreleaser_Windows_x86_64.zip",
			//"arm": "",
			"arm64": "goreleaser_Windows_arm64.zip",
		},
	}

	for os, arches := range tests {
		for arch, expected := range arches {
			t.Run(fmt.Sprintf("CanResolveGoreleaserArtifacts/%s-%s", os, arch), func(t *testing.T) {
				resolver := goreleaserAssetResolver{os: os, arch: arch}

				rc := getTestReleaseCandidate()

				match := resolver.ResolveAsset("goreleaser", rc)

				require.NotNil(t, match)
				assert.Equal(t, expected, match.Name)
			})
		}
	}
}

func getTestReleaseCandidate() *types.ReleaseCandidate {
	return &types.ReleaseCandidate{
		Name: "goreleaser",
		Assets: []types.ReleaseCandidateAsset{
			{
				ID:   134635013,
				Name: "checksums.txt",
			},
			{
				ID:   134635017,
				Name: "checksums.txt.pem",
			},
			{
				ID:   134635016,
				Name: "checksums.txt.sig",
			},
			{
				ID:   134634984,
				Name: "goreleaser-1.22.1-1-aarch64.pkg.tar.zst",
			},
			{
				ID:   134634995,
				Name: "goreleaser-1.22.1-1-armv7h.pkg.tar.zst",
			},
			{
				ID:   134634994,
				Name: "goreleaser-1.22.1-1-i686.pkg.tar.zst",
			},
			{
				ID:   134634992,
				Name: "goreleaser-1.22.1-1-ppc64.pkg.tar.zst",
			},
			{
				ID:   134634990,
				Name: "goreleaser-1.22.1-1-x86_64.pkg.tar.zst",
			},
			{
				ID:   134634836,
				Name: "goreleaser-1.22.1-1.aarch64.rpm",
			},
			{
				ID:   134634832,
				Name: "goreleaser-1.22.1-1.armv7hl.rpm",
			},
			{
				ID:   134634834,
				Name: "goreleaser-1.22.1-1.i386.rpm",
			},
			{
				ID:   134634880,
				Name: "goreleaser-1.22.1-1.ppc64.rpm",
			},
			{
				ID:   134634839,
				Name: "goreleaser-1.22.1-1.x86_64.rpm",
			},
			{
				ID:   134634824,
				Name: "goreleaser_1.22.1_aarch64.apk",
			},
			{
				ID:   134634835,
				Name: "goreleaser_1.22.1_amd64.deb",
			},
			{
				ID:   134634827,
				Name: "goreleaser_1.22.1_arm64.deb",
			},
			{
				ID:   134634830,
				Name: "goreleaser_1.22.1_armhf.deb",
			},
			{
				ID:   134634818,
				Name: "goreleaser_1.22.1_armv7.apk",
			},
			{
				ID:   134634828,
				Name: "goreleaser_1.22.1_i386.deb",
			},
			{
				ID:   134634823,
				Name: "goreleaser_1.22.1_ppc64.apk",
			},
			{
				ID:   134634993,
				Name: "goreleaser_1.22.1_ppc64.deb",
			},
			{
				ID:   134634825,
				Name: "goreleaser_1.22.1_x86.apk",
			},
			{
				ID:   134634821,
				Name: "goreleaser_1.22.1_x86_64.apk",
			},
			{
				ID:   134634815,
				Name: "goreleaser_Darwin_all.tar.gz",
			},
			{
				ID:   134635011,
				Name: "goreleaser_Darwin_all.tar.gz.sbom",
			},
			{
				ID:   134634808,
				Name: "goreleaser_Darwin_arm64.tar.gz",
			},
			{
				ID:   134635007,
				Name: "goreleaser_Darwin_arm64.tar.gz.sbom",
			},
			{
				ID:   134634813,
				Name: "goreleaser_Darwin_x86_64.tar.gz",
			},
			{
				ID:   134635009,
				Name: "goreleaser_Darwin_x86_64.tar.gz.sbom",
			},
			{
				ID:   134634800,
				Name: "goreleaser_Linux_arm64.tar.gz",
			},
			{
				ID:   134635001,
				Name: "goreleaser_Linux_arm64.tar.gz.sbom",
			},
			{
				ID:   134634801,
				Name: "goreleaser_Linux_armv7.tar.gz",
			},
			{
				ID:   134634997,
				Name: "goreleaser_Linux_armv7.tar.gz.sbom",
			},
			{
				ID:   134634802,
				Name: "goreleaser_Linux_i386.tar.gz",
			},
			{
				ID:   134635000,
				Name: "goreleaser_Linux_i386.tar.gz.sbom",
			},
			{
				ID:   134634799,
				Name: "goreleaser_Linux_ppc64.tar.gz",
			},
			{
				ID:   134635002,
				Name: "goreleaser_Linux_ppc64.tar.gz.sbom",
			},
			{
				ID:   134635010,
				Name: "goreleaser_Linux_x86_64.tar.gz.sbom",
			},
			{
				ID:   134634809,
				Name: "goreleaser_Linux_x86_64.tar.gz",
			},
			{
				ID:   134634807,
				Name: "goreleaser_Windows_arm64.zip",
			},
			{
				ID:   134635004,
				Name: "goreleaser_Windows_arm64.zip.sbom",
			},
			{
				ID:   134634810,
				Name: "goreleaser_Windows_i386.zip",
			},
			{
				ID:   134635008,
				Name: "goreleaser_Windows_i386.zip.sbom",
			},
			{
				ID:   134634817,
				Name: "goreleaser_Windows_x86_64.zip",
			},
			{
				ID:   134635012,
				Name: "goreleaser_Windows_x86_64.zip.sbom",
			},
		},
	}
}

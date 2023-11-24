package versioning

import (
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStableVersionStrategy_IsAcceptable(t *testing.T) {
	t.Run("StandardSemverAcceptable", func(t *testing.T) {
		strategy := Stable()
		v, err := semver.NewVersion("1.0.0")

		require.Nil(t, err)

		assert.True(t, strategy.IsAcceptable(v))
	})

	t.Run("PrereleaseUnacceptable", func(t *testing.T) {
		strategy := Stable()

		v, err := semver.NewVersion("1.0.0-alpha1")

		require.Nil(t, err)

		assert.False(t, strategy.IsAcceptable(v))
	})

	t.Run("MetadataAcceptable", func(t *testing.T) {
		strategy := Stable()
		v, err := semver.NewVersion("1.0.0+stuff")

		require.Nil(t, err)

		assert.True(t, strategy.IsAcceptable(v))
	})
}

func TestStableVersionStrategy_IsUpgrade(t *testing.T) {
	t.Run("GreaterVersionIsUpgrade", func(t *testing.T) {
		strategy := Stable()
		v1, _ := semver.NewVersion("1.0.0")
		v2, _ := semver.NewVersion("1.1.0")

		assert.True(t, strategy.IsUpgrade(v1, v2))
	})

	t.Run("LesserVersionIsNotUpgrade", func(t *testing.T) {
		strategy := Stable()
		v1, _ := semver.NewVersion("1.0.0")
		v2, _ := semver.NewVersion("0.9.0")

		assert.False(t, strategy.IsUpgrade(v1, v2))
	})

	t.Run("PrereleaseIsNotUpgrade", func(t *testing.T) {
		strategy := Stable()
		v1, _ := semver.NewVersion("1.0.0")
		v2, _ := semver.NewVersion("1.1.0-alpha")

		assert.False(t, strategy.IsUpgrade(v1, v2))
	})
}

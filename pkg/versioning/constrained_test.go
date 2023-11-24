package versioning

import (
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/stretchr/testify/assert"
)

func TestConstrainedVersionStrategy_IsAcceptable(t *testing.T) {
	v099, _ := semver.NewVersion("0.9.9")
	v100, _ := semver.NewVersion("1.0.0")
	v101, _ := semver.NewVersion("1.0.1")
	v101Pre, _ := semver.NewVersion("1.0.1-pre")
	v110, _ := semver.NewVersion("1.1.0")
	v110Pre, _ := semver.NewVersion("1.1.0-pre")
	v111, _ := semver.NewVersion("1.1.1")
	v200, _ := semver.NewVersion("2.0.0")

	t.Run("MajorConstraintNoPre", func(t *testing.T) {
		strategy := ConstrainMajor(v100, false)

		assert.True(t, strategy.IsAcceptable(v100))
		assert.True(t, strategy.IsAcceptable(v110))
		assert.True(t, strategy.IsAcceptable(v111))

		assert.False(t, strategy.IsAcceptable(v110Pre))

		assert.False(t, strategy.IsAcceptable(v200))
		assert.False(t, strategy.IsAcceptable(v099))
	})

	t.Run("MajorConstraintPre", func(t *testing.T) {
		strategy := ConstrainMajor(v100, true)

		assert.True(t, strategy.IsAcceptable(v110Pre))
	})

	t.Run("PatchConstraintNoPre", func(t *testing.T) {
		strategy := ConstrainPatch(v100, false)

		assert.True(t, strategy.IsAcceptable(v100))
		assert.True(t, strategy.IsAcceptable(v101))
		assert.False(t, strategy.IsAcceptable(v110))
		assert.False(t, strategy.IsAcceptable(v099))

		assert.False(t, strategy.IsAcceptable(v101Pre))
	})

	t.Run("PatchConstraintPre", func(t *testing.T) {
		strategy := ConstrainPatch(v100, true)

		assert.True(t, strategy.IsAcceptable(v101Pre))
	})
}

func TestConstrainedVersionStrategy_IsUpgrade(t *testing.T) {
	v099, _ := semver.NewVersion("0.9.9")
	v100, _ := semver.NewVersion("1.0.0")
	v101, _ := semver.NewVersion("1.0.1")
	v101Pre, _ := semver.NewVersion("1.0.1-pre")
	v110, _ := semver.NewVersion("1.1.0")
	v110Pre, _ := semver.NewVersion("1.1.0-pre")
	v200, _ := semver.NewVersion("2.0.0")

	t.Run("MajorConstraintNoPre", func(t *testing.T) {
		strategy := ConstrainMajor(v100, false)

		assert.True(t, strategy.IsUpgrade(v100, v110))

		assert.False(t, strategy.IsUpgrade(v100, v110Pre))
		assert.False(t, strategy.IsUpgrade(v100, v200))
		assert.False(t, strategy.IsUpgrade(v100, v099))
	})

	t.Run("MajorConstraintPre", func(t *testing.T) {
		strategy := ConstrainMajor(v100, true)

		assert.True(t, strategy.IsUpgrade(v100, v110Pre))
	})

	t.Run("PatchConstraintNoPre", func(t *testing.T) {
		strategy := ConstrainPatch(v100, false)

		assert.True(t, strategy.IsUpgrade(v100, v101))

		assert.False(t, strategy.IsUpgrade(v100, v110))
		assert.False(t, strategy.IsUpgrade(v100, v101Pre))
		assert.False(t, strategy.IsUpgrade(v100, v099))
	})

	t.Run("PatchConstraintPre", func(t *testing.T) {
		strategy := ConstrainPatch(v100, true)

		assert.True(t, strategy.IsUpgrade(v100, v101Pre))
	})
}

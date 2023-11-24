package release

import (
	"context"
	"fmt"
	"testing"

	"github.com/pdylanross/gh-release-autoupdate/pkg/gh"
	"github.com/pdylanross/gh-release-autoupdate/pkg/versioning"
	"github.com/stretchr/testify/require"
)

func TestChecker_Check(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		ghOpts := gh.NewOptions("gh-release-autoupdate-tests", "0.0.0")

		start, err := versioning.Constrained("~1.2")
		require.Nil(t, err)
		c, _ := NewResolver(ghOpts, start)

		candidate, err := c.Resolve(context.Background(), "goreleaser", "goreleaser", "1.0.0")
		require.Nil(t, err)

		fmt.Println(candidate.ID, " ", candidate.Name)
	})
}

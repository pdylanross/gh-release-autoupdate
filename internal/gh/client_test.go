package gh

import (
	"testing"

	"github.com/migueleliasweb/go-github-mock/src/mock"
	"github.com/stretchr/testify/require"
)

func TestNewGithubClient(t *testing.T) {
	t.Run("CantConstructWithEmptyOptions", func(t *testing.T) {
		_, err := NewGithubClient(nil)
		require.NotNil(t, err)
	})

	t.Run("CanConstructWithMockClient", func(t *testing.T) {
		opts := &GithubClientOptions{HTTPClient: mock.NewMockedHTTPClient(), UserAgent: "test-package/v1"}
		_, err := NewGithubClient(opts)
		require.Nil(t, err)
	})
}

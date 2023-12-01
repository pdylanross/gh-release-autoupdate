package autoupdate

import (
	"context"
	"fmt"
	"math"
	"testing"
	"time"

	"github.com/pdylanross/gh-release-autoupdate/internal/gh"
	"github.com/pdylanross/gh-release-autoupdate/internal/localfs"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUpdater_Check(t *testing.T) {
	t.Run("UpdaterCanCheck", func(t *testing.T) {
		updaterOpts := DefaultOptions().
			WithPackage("goreleaser", "1.0.0").
			WithRepo("goreleaser", "goreleaser").
			WithVersionStrategy(Stable())

		updater, err := NewUpdater(updaterOpts)
		require.Nil(t, err)

		result, err := updater.Check(context.Background())
		require.Nil(t, err)

		assert.NotNil(t, result)

		asset := updater.GetAsset(result)
		assert.NotNil(t, asset)
		fmt.Println(asset)

		fs, err := localfs.NewLocalFileManager("test-stuff")
		require.Nil(t, err)
		d, err := gh.NewDownloader(context.Background(), fs, updater.ghClient, "goreleaser", "goreleaser", asset.ID)
		require.Nil(t, err)

		status, fut, err := d.StartDownload(context.Background())
		require.Nil(t, err)

		go func() {
			for s := range status.ProgressChan {
				fmt.Println(math.Round(s.GetPercent() * 100))
			}
		}()

		res, err := fut.Get(time.Hour * 2)
		require.Nil(t, err)
		fmt.Println("res: ", res)
	})
}

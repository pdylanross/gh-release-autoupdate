package gh

import (
	"context"
	"io"
	"net/http"
	"os"
	"path"
	"time"

	"github.com/google/go-github/v56/github"
	"github.com/pdylanross/gh-release-autoupdate/autoupdate/types"
	"github.com/pdylanross/gh-release-autoupdate/internal/channelutil"
	"github.com/pdylanross/gh-release-autoupdate/internal/localfs"
	"github.com/reugn/async"
)

type Downloader struct {
	ghClient  *github.Client
	localFs   *localfs.LocalFileManager
	assetInfo *github.ReleaseAsset

	repoOwner string
	repoName  string
}

type DownloadResult struct {
	Path string
}

type DownloadStatus struct {
	ProgressChan <-chan types.DownloadProgress
}

func NewDownloader(ctx context.Context, localFs *localfs.LocalFileManager, ghClient *github.Client, owner string, repo string, assetID int64) (*Downloader, error) {
	ghAsset, _, err := ghClient.Repositories.GetReleaseAsset(ctx, owner, repo, assetID)
	if err != nil {
		return nil, err
	}

	return &Downloader{
		ghClient:  ghClient,
		localFs:   localFs,
		assetInfo: ghAsset,
		repoOwner: owner,
		repoName:  repo,
	}, nil
}

func (d *Downloader) StartDownload(ctx context.Context) (*DownloadStatus, async.Future[DownloadResult], error) {
	reader, _, err := d.ghClient.Repositories.DownloadReleaseAsset(ctx, d.repoOwner, d.repoName, d.assetInfo.GetID(), http.DefaultClient)
	if err != nil {
		return nil, nil, err
	}

	tmp, err := d.localFs.GetTemp()
	if err != nil {
		return nil, nil, err
	}

	tmpFileName := path.Join(tmp, d.assetInfo.GetName())
	writer, err := os.Create(tmpFileName)
	if err != nil {
		return nil, nil, err
	}

	progressWriter := make(chan types.DownloadProgress, 10)
	progressBuffer := channelutil.NewRingbuffer(progressWriter, 10)
	progressThrottled := channelutil.NewThrottle(progressBuffer.GetReader(), time.Millisecond*10)
	status := &DownloadStatus{ProgressChan: progressThrottled.GetReader()}
	fut := async.NewPromise[DownloadResult]()

	go func() {
		tee := io.TeeReader(reader, writer)
		defer reader.Close()
		defer writer.Close()
		defer close(progressWriter)

		bytesRead := 0
		totalSize := d.assetInfo.GetSize()

		buf := make([]byte, 1024*32)
		for {
			if err := ctx.Err(); err != nil {
				return
			}

			numRead, err := tee.Read(buf)
			if err == io.EOF {
				fut.Success(&DownloadResult{Path: tmpFileName})
				return
			} else if err != nil {
				fut.Failure(err)
				return
			}

			bytesRead += numRead
			currentProgress := types.DownloadProgress{
				BytesDownloaded: bytesRead,
				BytesTotal:      totalSize,
			}
			progressWriter <- currentProgress
		}
	}()

	return status, fut.Future(), nil
}

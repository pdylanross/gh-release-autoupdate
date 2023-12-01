package types

type DownloadProgress struct {
	BytesDownloaded int
	BytesTotal      int
}

func (dp *DownloadProgress) GetPercent() float64 {
	return float64(dp.BytesDownloaded) / float64(dp.BytesTotal)
}

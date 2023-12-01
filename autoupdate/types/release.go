package types

// ReleaseCandidate tracks a suitable release candidate.
type ReleaseCandidate struct {
	ID         int64
	Name       string
	Owner      string
	Repository string
	Assets     []ReleaseCandidateAsset
}

// ReleaseCandidateAsset tracks a single file on a RC.
type ReleaseCandidateAsset struct {
	ID   int64
	Name string
}

// ReleaseAssetResolver checks the assets for a ReleaseCandidate and resolves the correct
// asset to download for the current system.
type ReleaseAssetResolver interface {
	ResolveAsset(packageName string, candidate *ReleaseCandidate) *ReleaseCandidateAsset
}

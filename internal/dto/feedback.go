package dto

type Feedback struct {
	Description string `json:"description"`
	Version     string `json:"version"`
	VersionId   uint64 `json:"versionId"`
}

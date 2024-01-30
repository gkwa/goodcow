package embedded5

import (
	"encoding/json"
	"fmt"
	"time"
)

type RepositoryInfo struct {
	BrowseURL string    `json:"browse_url"`
	CreatedAt time.Time `json:"created_at"`
	GitCommit string    `json:"git_commit"`
	GitURL    string    `json:"git_url"`
	IndexedAt time.Time `json:"indexed_at"`
	Path      string    `json:"path"`
	Release   string    `json:"release"`
	Subpath   string    `json:"subpath"`
	Version   string    `json:"version"`
}

func RunEmbedded() error {
	jsonData := `{
	    "browse_url": "https://github.com/fluent-ci-templates/shuttle-pipeline",
	    "created_at": "2024-01-29T08:13:15Z",
	    "git_commit": "0e073730b9f80be6f2c50631840b153e31c331c8",
	    "git_url": "https://github.com/fluent-ci-templates/shuttle-pipeline",
	    "indexed_at": "2024-01-29T08:14:39.002483Z",
	    "path": "github.com/fluent-ci-templates/shuttle-pipeline",
	    "release": "",
	    "version": "0e073730b9f80be6f2c50631840b153e31c331c8"
	}`

	var baseRepoInfo *RepositoryInfo

	err := json.Unmarshal([]byte(jsonData), &baseRepoInfo)
	if err != nil {
		return fmt.Errorf("error unmarshalling JSON: %s", err)
	}

	fmt.Println(baseRepoInfo)

	return nil
}

func (r *RepositoryInfo) String() string {
	b, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return fmt.Sprintf("error marshalling JSON: %v", err)
	}
	return string(b)
}
